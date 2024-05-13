// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package logic

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures logic group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_integration_service_environment", func(r *config.Resource) {
		r.Kind = "IntegrationServiceEnvironment"

		r.References["virtual_network_subnet_ids"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
