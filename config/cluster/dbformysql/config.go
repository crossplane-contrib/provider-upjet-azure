// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dbformysql

import (
	"github.com/upbound/provider-azure/apis/cluster/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures dbformysql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_mysql_flexible_server", func(r *config.Resource) {
		r.References["private_dns_zone_id"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mysql_flexible_database", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_mysql_flexible_server",
		}
	})

	p.AddResourceConfigurator("azurerm_mysql_flexible_server_configuration", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_mysql_flexible_server",
		}
	})

	p.AddResourceConfigurator("azurerm_mysql_flexible_server_firewall_rule", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_mysql_flexible_server",
		}
	})
}
