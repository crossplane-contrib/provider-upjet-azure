// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dbforpostgresql

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures dbforpostgresql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_postgresql_flexible_server", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"zone"},
		}
	})
}
