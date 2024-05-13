// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package appplatform

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures appplatform group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_spring_cloud_api_portal", func(r *config.Resource) {
		r.References["gateway_ids"] = config.Reference{
			TerraformName: "azurerm_spring_cloud_gateway",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
