// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package operationalinsights

import (
	"github.com/upbound/provider-azure/apis/namespaced/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures operationalinsights group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_log_analytics_workspace", func(r *config.Resource) {
		r.Kind = "Workspace"
	})

	p.AddResourceConfigurator("azurerm_log_analytics_linked_storage_account", func(r *config.Resource) {
		r.References["storage_account_ids"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
