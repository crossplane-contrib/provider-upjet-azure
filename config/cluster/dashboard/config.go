// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dashboard

import (
	"github.com/upbound/provider-azure/v2/apis/cluster/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures dashboard group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_dashboard_grafana", func(r *config.Resource) {
		r.References["azure_monitor_workspace_integrations.resource_id"] = config.Reference{
			TerraformName: "azurerm_monitor_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
