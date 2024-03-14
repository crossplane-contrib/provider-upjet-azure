// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package netapp

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures netapp group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_netapp_pool", func(r *config.Resource) {
		r.References["account_name"] = config.Reference{
			Type: "Account",
		}
	})

	p.AddResourceConfigurator("azurerm_netapp_snapshot_policy", func(r *config.Resource) {
		r.References["account_name"] = config.Reference{
			Type: "Account",
		}
	})

	p.AddResourceConfigurator("azurerm_netapp_snapshot", func(r *config.Resource) {
		r.References["account_name"] = config.Reference{
			Type: "Account",
		}
		r.References["pool_name"] = config.Reference{
			Type: "Pool",
		}
		r.References["volume_name"] = config.Reference{
			Type: "Volume",
		}
	})

	p.AddResourceConfigurator("azurerm_netapp_volume", func(r *config.Resource) {
		r.References["account_name"] = config.Reference{
			Type: "Account",
		}
		r.References["pool_name"] = config.Reference{
			Type: "Pool",
		}
		r.References["data_protection_snapshot_policy.snapshot_policy_id"] = config.Reference{
			Type:      "SnapshotPolicy",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["data_protection_replication.remote_volume_resource_id"] = config.Reference{
			Type:      "Volume",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["create_from_snapshot_resource_id"] = config.Reference{
			Type:      "Snapshot",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		delete(r.References, "data_protection_replication.remote_volume_location")
	})
}
