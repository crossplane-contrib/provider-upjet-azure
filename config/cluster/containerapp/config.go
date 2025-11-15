// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package containerapp

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/provider-azure/v2/apis/cluster/rconfig"
)

const group = "containerapp"

// Configure configures containerapp group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_container_app", func(r *config.Resource) {
		r.ShortGroup = group
		r.TerraformResource.Schema["secret"].Elem.(*schema.Resource).
			Schema["name"].Sensitive = true
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"secret"},
		}
	})

	p.AddResourceConfigurator("azurerm_container_app_environment", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "Environment"
	})

	p.AddResourceConfigurator("azurerm_container_app_custom_domain", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "CustomDomain"
	})

	p.AddResourceConfigurator("azurerm_container_app_environment_certificate", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "EnvironmentCertificate"
	})

	p.AddResourceConfigurator("azurerm_container_app_environment_custom_domain", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "EnvironmentCustomDomain"
	})

	p.AddResourceConfigurator("azurerm_container_app_environment_dapr_component", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "EnvironmentDaprComponent"
	})

	p.AddResourceConfigurator("azurerm_container_app_environment_storage", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "EnvironmentStorage"
	})

	p.AddResourceConfigurator("azurerm_container_app_job", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "ContainerJob"
		r.References["container_app_environment_id"] = config.Reference{
			TerraformName: "azurerm_container_app_environment",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["secret.key_vault_secret_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_secret",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
