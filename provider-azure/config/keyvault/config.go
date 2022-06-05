/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keyvault

import (
	"context"
	"fmt"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
	"github.com/upbound/official-providers/provider-azure/config/common"
)

// getURLIDFn returns a GetIDFn that can generate Azure vault secret IDs.
// An example identifier is as follows:
// https://example.vault.azure.net/secrets/example-secret/c0ffee5f4d45440cb60c28672887f832
// https://example.vault.azure.net/keys/example-key/fdf067c93bbb4b22bff4d8b7a9a56217
func getURLIDFn(resourceType string) config.GetIDFn {
	return func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		keyVaultID, err := common.ParseNameFromIDField(parameters, "key_vault_id")
		if err != nil {
			return "", err
		}
		version, err := common.GetAttributeValue(parameters, "version")
		if err != nil { // then persistent ID is not yet available, thus do not return an error
			return "", nil
		}
		name, err := common.GetAttributeValue(parameters, "name")
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("https://%s.vault.azure.net/%s/%s/%s",
			keyVaultID, resourceType, name, version), nil
	}
}

// getIssuerURLIDFn is similar to getURLIDFn for vault certificate issuers and
// does not include a version component.
// Example:
// https://key-vault-name.vault.azure.net/certificates/issuers/example
// https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01
func getIssuerURLIDFn(resourceType string) config.GetIDFn {
	return func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		keyVaultID, err := common.ParseNameFromIDField(parameters, "key_vault_id")
		if err != nil {
			return "", err
		}
		name, err := common.GetAttributeValue(parameters, "name")
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("https://%s.vault.azure.net/%s/%s",
			keyVaultID, resourceType, name), nil
	}
}

// TODO(aru): Adopt this function for azurerm_key_vault_access_policy
// example ID: /subscriptions/038f2b7c-3265-43b8-8624-c9ad5da610a8/resourceGroups/alper/providers/Microsoft.KeyVault/vaults/examplekeyvault-alper/objectId/a8090ca4-b5a5-4594-aa6d-362aa682f168
// can also contain application ID as the last component
/*func getAccessPolicyID() config.GetIDFn {
	return func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		id, err := common.GetAttributeValue(parameters, "key_vault_id")
		if err != nil {
			return "", err
		}
		id = fmt.Sprintf("%s/objectId/%s", id, parameters["object_id"])
		appID, ok := parameters["application_id"]
		if ok && len(appID.(string)) > 0 {
			id = fmt.Sprintf("%s/application_id/%s", id, appID)
		}
		return id, nil
	}
}*/

// Configure configures keyvault group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_key_vault", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		delete(r.TerraformResource.Schema, "access_policy") // we have the keyvault.AccessPolicy instead
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/vaults/vault1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.KeyVault", "vaults", "name")
	})

	p.AddResourceConfigurator("azurerm_key_vault_secret", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetResourceNameFromIDURLFn(2)
		// https://example.vault.azure.net/secrets/example-secret/c0ffee5f4d45440cb60c28672887f832
		r.ExternalName.GetIDFn = getURLIDFn("secrets")
		r.References = config.References{
			"key_vault_id": config.Reference{
				Type:      "Vault",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_key", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetResourceNameFromIDURLFn(2)
		// https://example-keyvault.vault.azure.net/keys/example/fdf067c93bbb4b22bff4d8b7a9a56217
		r.ExternalName.GetIDFn = getURLIDFn("keys")
		r.References = config.References{
			"key_vault_id": config.Reference{
				Type:      "Vault",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_certificate", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetResourceNameFromIDURLFn(2)
		// https://example-keyvault.vault.azure.net/certificates/example/fdf067c93bbb4b22bff4d8b7a9a56217
		r.ExternalName.GetIDFn = getURLIDFn("certificates")
		r.References = config.References{
			"key_vault_id": config.Reference{
				Type:      "Vault",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_certificate_issuer", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetResourceNameFromIDURLFn(1)
		// https://key-vault-name.vault.azure.net/certificates/issuers/example
		r.ExternalName.GetIDFn = getIssuerURLIDFn("certificates/issuers")
		r.References = config.References{
			"key_vault_id": config.Reference{
				Type:      "Vault",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_managed_storage_account", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetResourceNameFromIDURLFn(1)
		// https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01
		r.ExternalName.GetIDFn = getIssuerURLIDFn("storage")
		r.References = config.References{
			"key_vault_id": config.Reference{
				Type:      "Vault",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"storage_account_id": config.Reference{
				Type:      rconfig.StorageAccountReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_managed_hardware_security_module", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.KeyVault/managedHSMs/hsm1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.KeyVault", "managedHSMs", "name")
	})

	p.AddResourceConfigurator("azurerm_key_vault_access_policy", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		// TODO(aru): Adopt this function for azurerm_key_vault_access_policy
		/* r.ExternalName = config.NameAsIdentifier
		// /subscriptions/038f2b7c-3265-43b8-8624-c9ad5da610a8/resourceGroups/alper/providers/Microsoft.KeyVault/vaults/examplekeyvault-alper/objectId/a8090ca4-b5a5-4594-aa6d-362aa682f168
		r.ExternalName.GetIDFn = getAccessPolicyID()
		r.ExternalName.SetIdentifierArgumentFn = config.NopSetIdentifierArgument */
		r.ExternalName = config.IdentifierFromProvider
		r.References = config.References{
			"key_vault_id": config.Reference{
				Type:      "Vault",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_managed_storage_account_sas_token_definition", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetResourceNameFromIDURLFn(1)
		// https://example-keyvault.vault.azure.net/storage/exampleStorageAcc01/sas/exampleSasDefinition01
		r.ExternalName.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			managedStorageAccountID, err := common.GetAttributeValue(parameters, "managed_storage_account_id")
			if err != nil {
				return "", err
			}
			name, err := common.GetAttributeValue(parameters, "name")
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("%s/sas/%s", managedStorageAccountID, name), nil
		}
		r.References = config.References{
			"managed_storage_account_id": config.Reference{
				Type:      "ManagedStorageAccount",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})
}
