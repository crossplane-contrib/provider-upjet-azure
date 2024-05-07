// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package servicebus

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures security group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_servicebus_namespace", func(r *config.Resource) {
		// Issues on using Namespace as Kind in ServiceBus
		// https://github.com/crossplane/terrajet/issues/234
		// https://github.com/kubernetes/kubernetes/pull/108382
		r.Kind = "ServiceBusNamespace"
	})

	p.AddResourceConfigurator("azurerm_servicebus_namespace_network_rule_set", func(r *config.Resource) {
		r.OverrideFieldNames["NetworkRulesParameters"] = "NamespaceNetworkRuleSetNetworkRulesParameters"
		r.OverrideFieldNames["NetworkRulesInitParameters"] = "NamespaceNetworkRuleSetNetworkRulesInitParameters"
		r.OverrideFieldNames["NetworkRulesObservation"] = "NamespaceNetworkRuleSetNetworkRulesObservation"
	})
}
