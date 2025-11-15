// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/json"
	"strings"
	"sync/atomic"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/upbound/provider-azure/v2/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-azure/v2/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal Azure credentials as JSON"
	errSubscriptionIDNotSet = "subscription ID must be set in ProviderConfig when credential source is InjectedIdentity, OIDCTokenFile or Upbound"
	errTenantIDNotSet       = "tenant ID must be set in ProviderConfig when credential source is InjectedIdentity, OIDCTokenFile or Upbound"
	errClientIDNotSet       = "client ID must be set in ProviderConfig when credential source is OIDCTokenFile or Upbound"
	// Azure service principal credentials file JSON keys
	keyAzureSubscriptionID = "subscriptionId"
	keyAzureClientID       = "clientId"
	keyAzureClientSecret   = "clientSecret"
	keyAzureClientCert     = "clientCertificate"
	keyAzureClientCertPass = "clientCertificatePassword"
	keyAzureTenantID       = "tenantId"
	// Terraform Provider configuration block keys
	keyTerraformFeatures        = "features"
	keySkipProviderRegistration = "skip_provider_registration"
	keyUseMSI                   = "use_msi"
	keyClientID                 = "client_id"
	keySubscriptionID           = "subscription_id"
	keyTenantID                 = "tenant_id"
	keyMSIEndpoint              = "msi_endpoint"
	keyClientSecret             = "client_secret"
	keyClientCert               = "client_certificate"
	keyClientCertPassword       = "client_certificate_password"
	keyEnvironment              = "environment"
	keyOidcTokenFilePath        = "oidc_token_file_path"
	keyUseOIDC                  = "use_oidc"
	// Default OidcTokenFilePath
	defaultOidcTokenFilePath = "/var/run/secrets/azure/tokens/azure-identity-token"
)

var (
	credentialsSourceUserAssignedManagedIdentity   xpv1.CredentialsSource = "UserAssignedManagedIdentity"
	credentialsSourceSystemAssignedManagedIdentity xpv1.CredentialsSource = "SystemAssignedManagedIdentity"
	credentialsSourceOIDCTokenFile                 xpv1.CredentialsSource = "OIDCTokenFile"
	credentialsSourceUpbound                       xpv1.CredentialsSource = "Upbound"

	upboundProviderIdentityTokenFile = "/var/run/secrets/upbound.io/provider/token"

	// Round-robin counter for service principal selection
	servicePrincipalCounter uint64
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, err
		}

		ps.Configuration = map[string]interface{}{
			keyTerraformFeatures: map[string]interface{}{},
			// Terraform AzureRM provider tries to register all resource providers
			// in Azure just in case if the provider of the resource you're
			// trying to create is not registered and the returned error is
			// ambiguous. However, this requires service principal to have provider
			// registration permissions which are irrelevant in most contexts.
			// For details, see https://github.com/upbound/provider-azure/v2/issues/104
			keySkipProviderRegistration: true,
		}

		switch pcSpec.Credentials.Source { //nolint:exhaustive
		case credentialsSourceSystemAssignedManagedIdentity, credentialsSourceUserAssignedManagedIdentity:
			err = msiAuth(pcSpec, &ps)
		case credentialsSourceOIDCTokenFile:
			err = oidcAuth(pcSpec, &ps)
		case credentialsSourceUpbound:
			err = upboundAuth(pcSpec, &ps)
		default:
			err = spAuth(ctx, pcSpec, &ps, client)
		}
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to prepare terraform.Setup")
		}

		return ps, errors.Wrap(configureNoForkAzureClient(ctx, &ps, *tfProvider), "failed to configure the no-fork Azure client")
	}
}

func configureNoForkAzureClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}

func spAuth(ctx context.Context, pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup, client client.Client) error {
	data, err := xpresource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return errors.Wrap(err, errExtractCredentials)
	}
	data = []byte(strings.TrimSpace(string(data)))

	// Try to unmarshal as array of service principals first
	var servicePrincipals []map[string]string
	if err := json.Unmarshal(data, &servicePrincipals); err == nil && len(servicePrincipals) > 0 {
		// Round-robin selection
		index := atomic.AddUint64(&servicePrincipalCounter, 1) % uint64(len(servicePrincipals))
		azureCreds := servicePrincipals[index]
		return configureSpCredentials(azureCreds, pcSpec, ps)
	}

	// Fallback to single service principal format
	azureCreds := map[string]string{}
	if err := json.Unmarshal(data, &azureCreds); err != nil {
		return errors.Wrap(err, errUnmarshalCredentials)
	}
	return configureSpCredentials(azureCreds, pcSpec, ps)
}

func configureSpCredentials(azureCreds map[string]string, pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	ps.Configuration[keySubscriptionID] = azureCreds[keyAzureSubscriptionID]
	ps.Configuration[keyTenantID] = azureCreds[keyAzureTenantID]
	ps.Configuration[keyClientID] = azureCreds[keyAzureClientID]
	ps.Configuration[keyClientSecret] = azureCreds[keyAzureClientSecret]
	if clientCert, ok := azureCreds[keyAzureClientCert]; ok {
		ps.Configuration[keyClientCert] = clientCert
		if clientCertPass, passwordOk := azureCreds[keyAzureClientCertPass]; passwordOk {
			ps.Configuration[keyClientCertPassword] = clientCertPass
		}
	}
	if pcSpec.SubscriptionID != nil {
		ps.Configuration[keySubscriptionID] = *pcSpec.SubscriptionID
	}
	if pcSpec.TenantID != nil {
		ps.Configuration[keyTenantID] = *pcSpec.TenantID
	}
	if pcSpec.ClientID != nil {
		ps.Configuration[keyClientID] = *pcSpec.ClientID
	}
	if pcSpec.Environment != nil {
		ps.Configuration[keyEnvironment] = *pcSpec.Environment
	}
	return nil
}

func msiAuth(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	if pcSpec.SubscriptionID == nil || len(*pcSpec.SubscriptionID) == 0 {
		return errors.New(errSubscriptionIDNotSet)
	}
	if pcSpec.TenantID == nil || len(*pcSpec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	ps.Configuration[keySubscriptionID] = *pcSpec.SubscriptionID
	ps.Configuration[keyTenantID] = *pcSpec.TenantID
	ps.Configuration[keyUseMSI] = "true"
	if pcSpec.MSIEndpoint != nil {
		ps.Configuration[keyMSIEndpoint] = *pcSpec.MSIEndpoint
	}
	if pcSpec.ClientID != nil {
		ps.Configuration[keyClientID] = *pcSpec.ClientID
	}
	if pcSpec.Environment != nil {
		ps.Configuration[keyEnvironment] = *pcSpec.Environment
	}
	return nil
}

func oidcAuth(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	if pcSpec.SubscriptionID == nil || len(*pcSpec.SubscriptionID) == 0 {
		return errors.New(errSubscriptionIDNotSet)
	}
	if pcSpec.TenantID == nil || len(*pcSpec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	if pcSpec.ClientID == nil || len(*pcSpec.ClientID) == 0 {
		return errors.New(errClientIDNotSet)
	}
	// OIDC Token File Path defaults to a projected-volume path mounted in the pod running in the AKS cluster, when workload identity is enabled on the pod.
	ps.Configuration[keyOidcTokenFilePath] = defaultOidcTokenFilePath
	if pcSpec.OidcTokenFilePath != nil {
		ps.Configuration[keyOidcTokenFilePath] = *pcSpec.OidcTokenFilePath
	}
	ps.Configuration[keySubscriptionID] = *pcSpec.SubscriptionID
	ps.Configuration[keyTenantID] = *pcSpec.TenantID
	ps.Configuration[keyClientID] = *pcSpec.ClientID
	ps.Configuration[keyUseOIDC] = "true"
	return nil

}

func upboundAuth(pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup) error {
	if pcSpec.SubscriptionID == nil || len(*pcSpec.SubscriptionID) == 0 {
		return errors.New(errSubscriptionIDNotSet)
	}
	if pcSpec.TenantID == nil || len(*pcSpec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	if pcSpec.ClientID == nil || len(*pcSpec.ClientID) == 0 {
		return errors.New(errClientIDNotSet)
	}
	ps.Configuration[keyOidcTokenFilePath] = upboundProviderIdentityTokenFile
	ps.Configuration[keySubscriptionID] = *pcSpec.SubscriptionID
	ps.Configuration[keyTenantID] = *pcSpec.TenantID
	ps.Configuration[keyClientID] = *pcSpec.ClientID
	ps.Configuration[keyUseOIDC] = "true"
	return nil

}

func legacyToModernProviderConfigSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	// TODO(erhan): this is hacky and potentially lossy, generate or manually implement
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func enrichLocalSecretRefs(pc *namespacedv1beta1.ProviderConfig, mg xpresource.Managed) {
	if pc != nil && pc.Spec.Credentials.SecretRef != nil {
		pc.Spec.Credentials.SecretRef.Namespace = mg.GetNamespace()
	}
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg xpresource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case xpresource.LegacyManaged:
		return resolveProviderConfigLegacy(ctx, crClient, managed)
	case xpresource.ModernManaged:
		return resolveProviderConfigModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed")
	}
}

func resolveProviderConfigLegacy(ctx context.Context, client client.Client, mg xpresource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := xpresource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return legacyToModernProviderConfigSpec(pc)
}

func resolveProviderConfigModern(ctx context.Context, crClient client.Client, mg xpresource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrapf(err, "referenced provider config kind %q is invalid for %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}
	pcObj, ok := pcRuntimeObj.(xpresource.ProviderConfig)
	if !ok {
		return nil, errors.Errorf("referenced provider config kind %q is not a provider config type %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		enrichLocalSecretRefs(pc, mg)
		pcSpec = pc.Spec
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		// TODO(erhan)
		return nil, errors.New("unknown")
	}
	t := xpresource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
