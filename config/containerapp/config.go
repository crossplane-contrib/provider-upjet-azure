// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package containerapp

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures kubernetes group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_container_app", func(r *config.Resource) {
		r.ShortGroup = "containerapp"
	})

	p.AddResourceConfigurator("azurerm_container_app_environment", func(r *config.Resource) {
		r.ShortGroup = "containerapp"
		r.Kind = "Environment"
	})
}
