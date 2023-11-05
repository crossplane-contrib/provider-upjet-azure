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
