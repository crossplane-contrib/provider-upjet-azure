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
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures keyvault group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_key_vault", func(r *config.Resource) {
		// Mutually exclusive with azurerm_key_vault_access_policy
		config.MoveToStatus(r.TerraformResource, "access_policy")
	})

	p.AddResourceConfigurator("azurerm_key_vault_secret", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			Type:      "Vault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_key", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			Type:      "Vault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_certificate", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			Type:      "Vault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_certificate_issuer", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			Type:      "Vault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_managed_storage_account", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			Type:      "Vault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_account_id"] = config.Reference{
			Type:      rconfig.StorageAccountReferencePath,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_access_policy", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			Type:      "Vault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_managed_storage_account_sas_token_definition", func(r *config.Resource) {
		r.References["managed_storage_account_id"] = config.Reference{
			Type:      "ManagedStorageAccount",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
