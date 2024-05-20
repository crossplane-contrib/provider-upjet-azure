// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package synapse

import "github.com/crossplane/upjet/pkg/config"

// Configure configures synapse group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_synapse_workspace", func(r *config.Resource) {
		r.TerraformResource.Schema["aad_admin"].MaxItems = 1
		r.TerraformResource.Schema["sql_aad_admin"].MaxItems = 1

		r.AddSingletonListConversion("aad_admin", "aad_admin")
		r.AddSingletonListConversion("sql_aad_admin", "sql_aad_admin")
	})
}
