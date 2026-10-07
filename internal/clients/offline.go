// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfazureclient "github.com/hashicorp/terraform-provider-azurerm/xpprovider"
	"sigs.k8s.io/controller-runtime/pkg/client"

	namespacedv1beta1 "github.com/upbound/provider-azure/v2/apis/namespaced/v1beta1"
)

const (
	// Terraform Provider configuration key for the Azure CLI authenticator,
	// which defaults to true and shells out to the `az` binary.
	keyUseCLI = "use_cli"
	// Terraform Provider configuration key for the provider's (4.x,
	// soon-to-be-deprecated) top-level enhanced-validation block.
	keyEnhancedValidation = "enhanced_validation"
	// Within an enhanced_validation block, whether to validate resource
	// provider arguments against the list of supported providers. This is
	// deliberately disabled for offline diffs: EnsureRegistered - the same
	// eager ARM call skip_provider_registration is meant to avoid - only
	// skips populating its resource-provider cache when *both*
	// skip_provider_registration is set *and* this is false. On this
	// provider's current (4.x) branch it defaults to true (FivePointOh() is
	// gated behind ARM_FIVEPOINTZERO_BETA, which is never set here), so
	// skip_provider_registration alone does not suppress the call; see
	// offlineConfiguration's own comment for how this was found.
	keyEnhancedValidationResourceProviders = "resource_providers"
	// Within an enhanced_validation block, whether to validate the
	// location argument against Azure's metadata service. Same defaulting
	// problem as resource_providers above, different eager call: left
	// enabled, Configure fetches the supported-locations list from Azure's
	// metadata endpoint.
	keyEnhancedValidationLocations = "locations"

	// Placeholder identities used when configuring the Terraform provider for
	// offline diffs. They are only ever observed locally: nothing derived from
	// them leaves the process, because the token they end up in is minted by
	// offlineTokenClient rather than by Azure.
	offlineSubscriptionID = "00000000-0000-0000-0000-000000000000"
	offlineTenantID       = "00000000-0000-0000-0000-000000000001"
	offlineClientID       = "00000000-0000-0000-0000-000000000002"
	offlineObjectID       = "00000000-0000-0000-0000-000000000003"
	offlineClientSecret   = "offline-diff-placeholder"

	// errRegisterRequestMiddleware is returned when the egress guard cannot be
	// installed. Offline diffs rely on it, so failing to install it is fatal
	// rather than a degradation.
	errRegisterRequestMiddleware = "cannot register the offline request middleware: provider meta is not an AzureRM client"

	// fmtErrOfflineRequestBlocked reports an Azure API call that the offline
	// diff server refused to send. The request's query string is deliberately
	// left out: it can carry values taken from the resource's own fields.
	fmtErrOfflineRequestBlocked = "offline diff server blocked an outbound Azure API request to %s (%s %s): computing this diff requires calling Azure, which is not possible offline."

	// offlineTokenLifetime is what the offline token claims for expires_in.
	// The value is immaterial: when the cached token lapses, the Azure SDK
	// simply asks offlineTokenClient for another one.
	offlineTokenLifetime = "3600"

	valUnknown = "<unknown>"
)

// EnableOfflineAuthentication replaces the HTTP clients that the Azure SDK uses
// to obtain access tokens with an in-process stub, so that configuring the
// AzureRM Terraform provider neither requires credentials nor reaches Azure.
//
// This is for the offline diff server, which computes the difference between a
// desired and an actual managed resource state and never reconciles the
// external resource. It must not be called in a provider that reconciles, as it
// disables authentication for the whole process: both auth.Client and
// auth.MetadataClient are package-level variables in the Azure SDK.
//
// It is safe to call more than once; only the first call takes effect.
func EnableOfflineAuthentication() {
	offlineAuthOnce.Do(func() {
		auth.Client = offlineTokenClient{}
		auth.MetadataClient = offlineTokenClient{}
	})
}

var offlineAuthOnce sync.Once

// OfflineTerraformSetupBuilder returns a Terraform setup that configures the
// AzureRM Terraform provider without credentials and without contacting Azure.
// The returned terraform.Setup carries a fully built *clients.Client as its
// Meta, which the Terraform provider's CustomizeDiff functions require, but one
// holding an access token that no Azure API will accept.
//
// Use it in place of TerraformSetupBuilder in the diff server only. The
// ProviderConfig is still resolved, so that configuration which affects the
// diff - the cloud environment above all - is honoured; only its credentials
// are ignored.
func OfflineTerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	EnableOfflineAuthentication()
	return func(ctx context.Context, client client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, err
		}

		return configureOffline(ctx, tfProvider, pcSpec)
	}
}

// configureOffline configures the Terraform provider from pcSpec without
// credentials, and installs the egress guard on the resulting client.
func configureOffline(ctx context.Context, tfProvider *schema.Provider, pcSpec *namespacedv1beta1.ProviderConfigSpec) (terraform.Setup, error) {
	ps := terraform.Setup{
		Configuration: offlineConfiguration(pcSpec),
	}
	if err := configureNoForkAzureClient(ctx, &ps, *tfProvider); err != nil {
		return terraform.Setup{}, errors.Wrap(err, "failed to configure the no-fork Azure client")
	}
	// The stubbed HTTP clients cover token acquisition only. Azure's ARM and
	// Microsoft Graph APIs are reached through separate clients, so a
	// CustomizeDiff function that calls either - see denyOutboundRequest -
	// would otherwise retry against an unreachable endpoint until it timed
	// out.
	//
	// The guard can only go on here, after Configure: it is registered on the
	// clients held by the provider's Meta, which Configure is what creates.
	// That leaves Configure itself unguarded, and it does contain two eager
	// calls, each closed off differently:
	//
	//   - buildClient's resource-provider registration (EnsureRegistered) is
	//     closed by configuration: skip_provider_registration and
	//     enhanced_validation.resource_providers=false together, both set
	//     unconditionally in offlineConfiguration (see its own comment for
	//     why both are required - confirmed by a genuine e2e run that still
	//     reached the network with only one of them set).
	//
	//   - clients.Build's oid-discovery fallback (NewResourceManagerAccount,
	//     reached only when the access token carries no oid claim) is closed
	//     by the token the offline auth client mints always carrying one, so
	//     the fallback is never entered. This is a *single* layer of
	//     protection, not two: unlike an equivalent gap this architecture was
	//     compared against elsewhere, where a missing deadline on the
	//     request's context additionally blocked the call even if it had been
	//     reached, azurerm's internal/clients/graph helpers synthesize their
	//     own 5-minute deadline when the incoming context has none - so if
	//     offlineAccessToken ever stopped setting oid, this call would really
	//     reach Microsoft Graph, through a client that bypasses this guard
	//     entirely (it is constructed ad hoc in clients/graph, never added to
	//     the client.armClients set AppendRequestMiddleware iterates). See
	//     offline_egress_test.go for the test that pins this asymmetry.
	if !tfazureclient.RegisterRequestMiddleware(ps.Meta, denyOutboundRequest) {
		return terraform.Setup{}, errors.New(errRegisterRequestMiddleware)
	}
	return ps, nil
}

// denyOutboundRequest refuses every Azure API request, failing the diff
// immediately instead of letting it retry against an API that will not accept
// the offline access token. The Azure SDK returns a request middleware's error
// to the caller without sending the request.
func denyOutboundRequest(req *http.Request) (*http.Request, error) {
	host := valUnknown
	method := valUnknown
	path := valUnknown
	if req != nil {
		method = req.Method
		if req.URL != nil {
			host = req.URL.Host
			path = req.URL.Path
		}
	}
	return nil, diffserver.NewDiffComputationNotSupportedError(errors.Errorf(fmtErrOfflineRequestBlocked, host, method, path))
}

// offlineConfiguration returns the Terraform provider configuration to use for
// offline diffs. It never reads the ProviderConfig's credentials, which are
// assumed to be unavailable.
func offlineConfiguration(pcSpec *namespacedv1beta1.ProviderConfigSpec) map[string]any {
	cfg := map[string]any{
		keyTerraformFeatures: []interface{}{map[string]interface{}{}}, // one empty features block
		keySubscriptionID:    offlineSubscriptionID,
		keyTenantID:          offlineTenantID,
		keyClientID:          offlineClientID,
		// A non-empty client secret makes the Azure SDK pick the client
		// credentials authorizer, whose token request offlineTokenClient
		// answers locally. Were all credentials empty, the SDK would fall
		// through to the Azure CLI authorizer and shell out to `az`.
		keyClientSecret: offlineClientSecret,
		keyUseCLI:       false,
		// Terraform AzureRM provider would otherwise try to register every
		// resource provider in the subscription during Configure, an eager
		// ARM call that an offline, credential-less diff cannot make either.
		// This is already unconditional in TerraformSetupBuilder for an
		// unrelated reason (see its own comment); kept unconditional here too
		// so the behaviour does not depend on which builder is in use.
		//
		// It is not sufficient by itself, though: confirmed by a genuine e2e
		// run that still reached management.azure.com with only this key set.
		// buildClient's EnsureRegistered only skips populating its
		// resource-provider cache when the required-providers set is empty
		// *and* enhanced validation for resource providers is disabled - and
		// that defaults to enabled on this provider's current branch, as does
		// enhanced validation for locations, an independent eager call to
		// Azure's metadata service found the same way. Both have to be turned
		// off alongside skip_provider_registration, not instead of it.
		keySkipProviderRegistration: true,
		keyEnhancedValidation: []interface{}{
			map[string]interface{}{
				keyEnhancedValidationResourceProviders: false,
				keyEnhancedValidationLocations:         false,
			},
		},
	}

	if pcSpec == nil {
		return cfg
	}
	if pcSpec.SubscriptionID != nil && *pcSpec.SubscriptionID != "" {
		cfg[keySubscriptionID] = *pcSpec.SubscriptionID
	}
	if pcSpec.TenantID != nil && *pcSpec.TenantID != "" {
		cfg[keyTenantID] = *pcSpec.TenantID
	}
	if pcSpec.ClientID != nil && *pcSpec.ClientID != "" {
		cfg[keyClientID] = *pcSpec.ClientID
	}
	if pcSpec.Environment != nil && *pcSpec.Environment != "" {
		cfg[keyEnvironment] = *pcSpec.Environment
	}
	return cfg
}

// offlineTokenClient answers every access token request locally. It satisfies
// auth.HTTPClient, the interface behind the Azure SDK's auth.Client and
// auth.MetadataClient, and covers every authentication method the provider
// supports, because they all obtain their token over HTTP and parse the same
// response shape.
type offlineTokenClient struct{}

func (offlineTokenClient) Do(req *http.Request) (*http.Response, error) {
	token, err := offlineAccessToken()
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(map[string]any{
		"access_token": token,
		"token_type":   "Bearer",
		// The Azure SDK accepts expires_in as either a string or a number.
		"expires_in": offlineTokenLifetime,
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot marshal the offline token response")
	}

	return &http.Response{
		Status:        "200 OK",
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        http.Header{"Content-Type": []string{"application/json"}},
		Body:          io.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Request:       req,
	}, nil
}

// offlineAccessToken mints an unsigned JWT carrying the claims that configuring
// the provider depends on. The Azure SDK decodes the token's payload to read
// them and does not verify its signature. The `oid` claim matters most: were
// it absent, NewResourceManagerAccount would query Microsoft Graph to discover
// the object ID of the authenticated principal, which offline it cannot do.
func offlineAccessToken() (string, error) {
	payload, err := json.Marshal(map[string]any{
		"aud":   "https://graph.microsoft.com",
		"iss":   "https://sts.windows.net/" + offlineTenantID + "/",
		"oid":   offlineObjectID,
		"tid":   offlineTenantID,
		"appid": offlineClientID,
		"scp":   "",
	})
	if err != nil {
		return "", errors.Wrap(err, "cannot marshal the offline token claims")
	}

	enc := base64.RawURLEncoding.EncodeToString
	return enc([]byte(`{"alg":"none","typ":"JWT"}`)) + "." + enc(payload) + ".offline", nil
}
