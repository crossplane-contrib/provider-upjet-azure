// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package web

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/apis/cluster/rconfig"
)

// Configure configures web group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_linux_web_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_function_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_function_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_linux_function_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_linux_function_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_linux_web_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_windows_function_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_windows_function_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_windows_web_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_windows_web_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_function_app_hybrid_connection", func(r *config.Resource) {
		r.References["function_app_id"] = config.Reference{
			TerraformName: "azurerm_windows_function_app",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_function_app_flex_consumption", func(r *config.Resource) {
		r.References["service_plan_id"] = config.Reference{
			TerraformName: "azurerm_service_plan",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_user_assigned_identity_id"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["virtual_network_subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["site_config.ip_restriction.virtual_network_subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["site_config.scm_ip_restriction.virtual_network_subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["site_config.container_registry_managed_identity_client_id"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["site_config.api_management_api_id"] = config.Reference{
			TerraformName: "azurerm_api_management_api",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_container_endpoint"] = config.Reference{
			TerraformName: "azurerm_storage_container",
			Extractor:     rconfig.ExtractAccountContainerEndpointFuncPath,
		}
	})
}
