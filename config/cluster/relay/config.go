// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package relay

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures relay group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_relay_namespace", func(r *config.Resource) {
		r.Kind = "EventRelayNamespace"
	})
}
