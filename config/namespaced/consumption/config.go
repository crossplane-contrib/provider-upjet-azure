// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package consumption

import (
	"github.com/upbound/provider-azure/apis/namespaced/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures consumption group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_consumption_budget_subscription", func(r *config.Resource) {
		r.References["notification.contact_groups"] = config.Reference{
			TerraformName: "azurerm_monitor_action_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
