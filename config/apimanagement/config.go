// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apimanagement

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
	"github.com/upbound/provider-azure/config/common"
)

// Configure configures apimanagement group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_api_management", func(r *config.Resource) {
		r.Kind = "Management"
		// Mutually exclusive with azurerm_api_management_custom_domain
		config.MoveToStatus(r.TerraformResource, "hostname_configuration")
	})
	p.AddResourceConfigurator("azurerm_api_management_api_operation", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			Type: "API",
		}
		r.References["api_management_name"] = config.Reference{
			Type: "Management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_api_policy", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			Type: "API",
		}
		r.References["api_management_name"] = config.Reference{
			Type: "Management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_api_schema", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			Type: "API",
		}
		r.References["api_management_name"] = config.Reference{
			Type: "Management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_product_api", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			Type: "API",
		}
		r.References["product_id"] = config.Reference{
			Type: "Product",
		}
		r.References["api_management_name"] = config.Reference{
			Type: "Management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_product_policy", func(r *config.Resource) {
		r.References["product_id"] = config.Reference{
			Type: "Product",
		}
		r.References["api_management_name"] = config.Reference{
			Type: "Management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_subscription", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			Type:      "User",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["product_id"] = config.Reference{
			Type:      "Product",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["api_management_name"] = config.Reference{
			Type: "Management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_authorization_server", func(r *config.Resource) {
		r.References["api_management_name"] = config.Reference{
			Type: "Management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_gateway_api", func(r *config.Resource) {
		r.References["gateway_id"] = config.Reference{
			Type:      "Gateway",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["api_id"] = config.Reference{
			Type:      "API",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_gateway", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "api_management_id")
	})
}
