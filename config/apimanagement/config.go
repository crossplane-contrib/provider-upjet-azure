/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apimanagement

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures apimanagement group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_api_management", func(r *config.Resource) {
		r.Kind = "Management"
		r.UseAsync = false
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
}
