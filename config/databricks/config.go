/*
Copyright 2023 The Crossplane Authors.

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

package databricks

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures databricks group
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("azurerm_databricks_workspace_customer_managed_key", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			TerraformName: "azurerm_databricks_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["key_vault_key_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_key",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
