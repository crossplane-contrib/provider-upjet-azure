// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package netapp

import (
	"github.com/upbound/provider-azure/v2/apis/cluster/rconfig"
	"github.com/upbound/provider-azure/v2/config/cluster/common"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures netapp group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_netapp_pool", func(r *config.Resource) {
		r.References["account_name"] = config.Reference{
			TerraformName: "azurerm_netapp_account",
		}
	})

	p.AddResourceConfigurator("azurerm_netapp_snapshot_policy", func(r *config.Resource) {
		r.References["account_name"] = config.Reference{
			TerraformName: "azurerm_netapp_account",
		}
	})

	p.AddResourceConfigurator("azurerm_netapp_snapshot", func(r *config.Resource) {
		r.References["account_name"] = config.Reference{
			TerraformName: "azurerm_netapp_account",
		}
		r.References["pool_name"] = config.Reference{
			TerraformName: "azurerm_netapp_pool",
		}
		r.References["volume_name"] = config.Reference{
			TerraformName: "azurerm_netapp_volume",
		}
	})

	p.AddResourceConfigurator("azurerm_netapp_volume", func(r *config.Resource) {
		r.Version = "v1beta3"
		r.PreviousVersions = []string{"v1beta2"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta2")
		r.References["account_name"] = config.Reference{
			TerraformName: "azurerm_netapp_account",
		}
		r.References["pool_name"] = config.Reference{
			TerraformName: "azurerm_netapp_pool",
		}
		r.References["data_protection_snapshot_policy.snapshot_policy_id"] = config.Reference{
			TerraformName: "azurerm_netapp_snapshot_policy",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["data_protection_replication.remote_volume_resource_id"] = config.Reference{
			TerraformName: "azurerm_netapp_volume",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["create_from_snapshot_resource_id"] = config.Reference{
			TerraformName: "azurerm_netapp_snapshot",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		delete(r.References, "data_protection_replication.remote_volume_location")
	})
}
