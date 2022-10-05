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

package storage

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures storage group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_storage_blob", func(r *config.Resource) {
		r.References["storage_account_name"] = config.Reference{
			Type: "Account",
		}
		r.References["storage_container_name"] = config.Reference{
			Type: "Container",
		}
	})

	p.AddResourceConfigurator("azurerm_storage_container", func(r *config.Resource) {
		r.References["storage_account_name"] = config.Reference{
			Type: "Account",
		}
	})

	p.AddResourceConfigurator("azurerm_storage_data_lake_gen2_filesystem", func(r *config.Resource) {
		r.References["storage_account_id"] = config.Reference{
			Type:      "Account",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
