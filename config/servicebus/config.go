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
}
