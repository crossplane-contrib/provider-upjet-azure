// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dataprotection

import (
	"github.com/upbound/provider-azure/v2/apis/namespaced/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures dataprotection group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_blob_storage", func(r *config.Resource) {
		r.References["vault_id"] = config.Reference{
			TerraformName: "azurerm_data_protection_backup_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_data_lake_storage", func(r *config.Resource) {
		r.References["data_protection_backup_vault_id"] = config.Reference{
			TerraformName: "azurerm_data_protection_backup_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_protection_backup_instance_data_lake_storage", func(r *config.Resource) {
		r.References["data_protection_backup_vault_id"] = config.Reference{
			TerraformName: "azurerm_data_protection_backup_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["backup_policy_data_lake_storage_id"] = config.Reference{
			TerraformName: "azurerm_data_protection_backup_policy_data_lake_storage",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_disk", func(r *config.Resource) {
		r.References["vault_id"] = config.Reference{
			TerraformName: "azurerm_data_protection_backup_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_postgresql", func(r *config.Resource) {
		r.References["vault_name"] = config.Reference{
			TerraformName: "azurerm_data_protection_backup_vault",
		}
	})

	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_postgresql_flexible_server", func(r *config.Resource) {
		r.References["vault_id"] = config.Reference{
			TerraformName: "azurerm_data_protection_backup_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
