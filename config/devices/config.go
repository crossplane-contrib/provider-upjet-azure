// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package devices

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures iothub group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_iothub", func(r *config.Resource) {
		// Note(ezgidemirel): Following fields are not marked as "sensitive" in Terraform cli schema output.
		// We need to configure them explicitly to store in connectionDetails secret.
		r.TerraformResource.Schema["endpoint"].Elem.(*schema.Resource).
			Schema["connection_string"].Sensitive = true
		r.TerraformResource.Schema["shared_access_policy"].Elem.(*schema.Resource).
			Schema["primary_key"].Sensitive = true
		r.TerraformResource.Schema["shared_access_policy"].Elem.(*schema.Resource).
			Schema["secondary_key"].Sensitive = true

		// Mutually exclusive with
		// azurerm_iothub_endpoint_*
		// azurerm_iothub_route
		// azurerm_iothub_enrichment
		// azurerm_iothub_fallback_route
		config.MoveToStatus(r.TerraformResource, "endpoint", "route", "enrichment", "fallback_route")

		r.Kind = "IOTHub"
		r.References["endpoint.container_name"] = config.Reference{
			TerraformName: "azurerm_storage_container",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_consumer_group", func(r *config.Resource) {
		r.References["iothub_name"] = config.Reference{
			TerraformName: "azurerm_iothub",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_dps_certificate", func(r *config.Resource) {
		r.References["iot_dps_name"] = config.Reference{
			TerraformName: "azurerm_iothub_dps",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_dps_shared_access_policy", func(r *config.Resource) {
		r.References["iot_dps_name"] = config.Reference{
			TerraformName: "azurerm_iothub_dps",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_shared_access_policy", func(r *config.Resource) {
		r.References["iot_dps_name"] = config.Reference{
			TerraformName: "azurerm_iothub_dps",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_endpoint_storage_container", func(r *config.Resource) {
		r.References["container_name"] = config.Reference{
			TerraformName: "azurerm_storage_container",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_fallback_route", func(r *config.Resource) {
		r.References["iothub_name"] = config.Reference{
			TerraformName: "azurerm_iothub",
		}
		r.References["endpoint_names"] = config.Reference{
			TerraformName: "azurerm_iothub_endpoint_storage_container",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_dps", func(r *config.Resource) {
		r.Path = "iothubdps"
	})
}
