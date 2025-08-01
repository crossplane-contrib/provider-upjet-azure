// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package base

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the base group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_subscription", func(r *config.Resource) {
		r.ShortGroup = ""
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"billing_scope_id", "subscription_id"},
		}
	})

	p.AddResourceConfigurator("azurerm_resource_provider_registration", func(r *config.Resource) {
		r.ShortGroup = ""
	})

	p.AddResourceConfigurator("azurerm_resource_group", func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "ResourceGroup"
		r.ShortGroup = ""
	})
}
