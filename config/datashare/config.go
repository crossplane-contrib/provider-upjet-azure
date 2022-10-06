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

package datashare

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures datashare group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_data_share_account", func(r *config.Resource) {
		r.Kind = "Account"
	})

	p.AddResourceConfigurator("azurerm_data_share", func(r *config.Resource) {
		r.Kind = "DataShare"
		r.References["account_id"] = config.Reference{
			Type:      "Account",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_blob_storage", func(r *config.Resource) {
		r.References["data_share_id"] = config.Reference{
			Type:      "DataShare",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["container_name"] = config.Reference{
			Type: rconfig.APISPackagePath + "/storage/v1beta1.Container",
		}
		r.References["storage_account.name"] = config.Reference{
			Type: rconfig.APISPackagePath + "/storage/v1beta1.Account",
		}
		r.References["storage_account.resource_group_name"] = config.Reference{
			Type: rconfig.APISPackagePath + "/azure/v1beta1.ResourceGroup",
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_data_lake_gen2", func(r *config.Resource) {
		r.References["share_id"] = config.Reference{
			Type:      "DataShare",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_account_id"] = config.Reference{
			Type:      rconfig.StorageAccountReferencePath,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["file_system_name"] = config.Reference{
			Type: rconfig.APISPackagePath + "/storage/v1beta1.DataLakeGen2FileSystem",
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_kusto_cluster", func(r *config.Resource) {
		r.References["share_id"] = config.Reference{
			Type:      "DataShare",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["kusto_cluster_id"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/kusto/v1beta1.Cluster",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_share_dataset_kusto_database", func(r *config.Resource) {
		r.References["share_id"] = config.Reference{
			Type:      "DataShare",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["kusto_database_id"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/kusto/v1beta1.Database",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
