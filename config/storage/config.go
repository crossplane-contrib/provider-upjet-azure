// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package storage

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
	"github.com/upbound/provider-azure/config/common"
)

// Configure configures storage group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_storage_blob", func(r *config.Resource) {
		r.References["storage_account_name"] = config.Reference{
			TerraformName: "azurerm_storage_account",
		}
		r.References["storage_container_name"] = config.Reference{
			TerraformName: "azurerm_storage_container",
		}
	})

	p.AddResourceConfigurator("azurerm_storage_container", func(r *config.Resource) {
		r.References["storage_account_name"] = config.Reference{
			TerraformName: "azurerm_storage_account",
		}
	})

	p.AddResourceConfigurator("azurerm_storage_data_lake_gen2_filesystem", func(r *config.Resource) {
		r.References["storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_storage_encryption_scope", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "storage_account_id")
	})

	p.AddResourceConfigurator("azurerm_storage_management_policy", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "storage_account_id")
	})

	p.AddResourceConfigurator("azurerm_storage_account_network_rules", func(r *config.Resource) {
		r.OverrideFieldNames["PrivateLinkAccessParameters"] = "AccountNetworkRulesPrivateLinkAccessParameters"
		r.OverrideFieldNames["PrivateLinkAccessInitParameters"] = "AccountNetworkRulesPrivateLinkAccessInitParameters"
		r.OverrideFieldNames["PrivateLinkAccessObservation"] = "AccountNetworkRulesPrivateLinkAccessObservation"
	})

	p.AddResourceConfigurator("azurerm_storage_object_replication", func(r *config.Resource) {
		r.OverrideFieldNames["RulesParameters"] = "ObjectReplicationRulesParameters"
		r.OverrideFieldNames["RulesInitParameters"] = "ObjectReplicationRulesInitParameters"
		r.OverrideFieldNames["RulesObservation"] = "ObjectReplicationRulesObservation"
	})
}
