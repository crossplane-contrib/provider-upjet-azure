// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package synapse

import "github.com/crossplane/upjet/pkg/config"

// Configure configures synapse group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_synapse_workspace", func(r *config.Resource) {
	})
}
