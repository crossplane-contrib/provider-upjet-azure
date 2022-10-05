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

package devices

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
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
			Type: rconfig.ContainerReferencePath,
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_consumer_group", func(r *config.Resource) {
		r.References["iothub_name"] = config.Reference{
			Type: "IOTHub",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_dps_certificate", func(r *config.Resource) {
		r.References["iot_dps_name"] = config.Reference{
			Type: "IOTHubDPS",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_dps_shared_access_policy", func(r *config.Resource) {
		r.References["iot_dps_name"] = config.Reference{
			Type: "IOTHubDPS",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_shared_access_policy", func(r *config.Resource) {
		r.References["iot_dps_name"] = config.Reference{
			Type: "IOTHubDPS",
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_endpoint_storage_container", func(r *config.Resource) {
		r.References["container_name"] = config.Reference{
			Type: rconfig.ContainerReferencePath,
		}
	})

	p.AddResourceConfigurator("azurerm_iothub_fallback_route", func(r *config.Resource) {
		r.References["iothub_name"] = config.Reference{
			Type: "IOTHub",
		}
		r.References["endpoint_names"] = config.Reference{
			Type: "IOTHubEndpointStorageContainer",
		}
	})
}
