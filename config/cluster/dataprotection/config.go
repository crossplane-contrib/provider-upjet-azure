// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dataprotection

import (
	"github.com/upbound/provider-azure/apis/cluster/rconfig"

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
}
