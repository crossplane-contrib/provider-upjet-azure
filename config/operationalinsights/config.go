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

package operationalinsights

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures operationalinsights group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_log_analytics_workspace", func(r *config.Resource) {
		r.Kind = "Workspace"
	})

	p.AddResourceConfigurator("azurerm_log_analytics_linked_storage_account", func(r *config.Resource) {
		r.References["storage_account_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/storage/v1beta1.Account",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
