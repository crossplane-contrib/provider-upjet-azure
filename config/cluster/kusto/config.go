// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kusto

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures kusto group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_kusto_database", func(r *config.Resource) {
		r.References["cluster_name"] = config.Reference{
			TerraformName: "azurerm_kusto_cluster",
		}
	})
}
