package clients

import (
	"context"
	"encoding/json"
	"strings"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/upbound/provider-azure/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal Azure credentials as JSON"
	errSubscriptionIDNotSet = "subscription ID must be set in ProviderConfig when credential source is InjectedIdentity"
	errTenantIDNotSet       = "tenant ID must be set in ProviderConfig when credential source is InjectedIdentity"
	// Azure service principal credentials file JSON keys
	keyAzureSubscriptionID = "subscriptionId"
	keyAzureClientID       = "clientId"
	keyAzureClientSecret   = "clientSecret"
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
	keyEnvironment              = "environment"
)

var (
	credentialsSourceUserAssignedManagedIdentity   xpv1.CredentialsSource = "UserAssignedManagedIdentity"
	credentialsSourceSystemAssignedManagedIdentity xpv1.CredentialsSource = "SystemAssignedManagedIdentity"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := xpresource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		ps.Configuration = map[string]interface{}{
			keyTerraformFeatures: struct{}{},
			// Terraform AzureRM provider tries to register all resource providers
			// in Azure just in case if the provider of the resource you're
			// trying to create is not registered and the returned error is
			// ambiguous. However, this requires service principal to have provider
			// registration permissions which are irrelevant in most contexts.
			// For details, see https://github.com/upbound/provider-azure/issues/104
			keySkipProviderRegistration: true,
		}

		var err error
		switch pc.Spec.Credentials.Source { //nolint:exhaustive
		case credentialsSourceSystemAssignedManagedIdentity, credentialsSourceUserAssignedManagedIdentity:
			err = msiAuth(pc, &ps)
		default:
			err = spAuth(ctx, pc, &ps, client)
		}
		return ps, err
	}
}

func spAuth(ctx context.Context, pc *v1beta1.ProviderConfig, ps *terraform.Setup, client client.Client) error {
	data, err := xpresource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return errors.Wrap(err, errExtractCredentials)
	}
	data = []byte(strings.TrimSpace(string(data)))
	azureCreds := map[string]string{}
	if err := json.Unmarshal(data, &azureCreds); err != nil {
		return errors.Wrap(err, errUnmarshalCredentials)
	}
	// set credentials configuration
	ps.Configuration[keySubscriptionID] = azureCreds[keyAzureSubscriptionID]
	ps.Configuration[keyTenantID] = azureCreds[keyAzureTenantID]
	ps.Configuration[keyClientID] = azureCreds[keyAzureClientID]
	ps.Configuration[keyClientSecret] = azureCreds[keyAzureClientSecret]
	if pc.Spec.Environment != nil {
		ps.Configuration[keyEnvironment] = *pc.Spec.Environment
	}
	return nil
}

func msiAuth(pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	if pc.Spec.SubscriptionID == nil || len(*pc.Spec.SubscriptionID) == 0 {
		return errors.New(errSubscriptionIDNotSet)
	}
	if pc.Spec.TenantID == nil || len(*pc.Spec.TenantID) == 0 {
		return errors.New(errTenantIDNotSet)
	}
	ps.Configuration[keySubscriptionID] = *pc.Spec.SubscriptionID
	ps.Configuration[keyTenantID] = *pc.Spec.TenantID
	ps.Configuration[keyUseMSI] = "true"
	if pc.Spec.MSIEndpoint != nil {
		ps.Configuration[keyMSIEndpoint] = *pc.Spec.MSIEndpoint
	}
	if pc.Spec.ClientID != nil {
		ps.Configuration[keyClientID] = *pc.Spec.ClientID
	}
	if pc.Spec.Environment != nil {
		ps.Configuration[keyEnvironment] = *pc.Spec.Environment
	}
	return nil
}
