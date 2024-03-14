// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eventhub

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_eventhub_namespace", func(r *config.Resource) {
		r.Kind = "EventHubNamespace"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"network_rulesets"},
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			Type: "EventHubNamespace",
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub_consumer_group", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			Type: "EventHubNamespace",
		}
		r.References["eventhub_name"] = config.Reference{
			Type: "EventHub",
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub_authorization_rule", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			Type: "EventHubNamespace",
		}
		r.References["eventhub_name"] = config.Reference{
			Type: "EventHub",
		}
	})
}
