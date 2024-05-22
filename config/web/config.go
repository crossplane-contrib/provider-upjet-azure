// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package web

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures web group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_linux_web_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_function_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_function_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_linux_function_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_linux_function_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_linux_web_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_windows_function_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_windows_function_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_windows_web_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_windows_web_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}

		sc := r.TerraformResource.Schema["site_config"].Elem.(*schema.Resource)
		sc.Schema["ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1
		sc.Schema["scm_ip_restriction"].Elem.(*schema.Resource).
			Schema["headers"].MaxItems = 1

		r.AddSingletonListConversion("site_config[*].ip_restriction[*].headers", "siteConfig[*].ipRestriction[*].headers")
		r.AddSingletonListConversion("site_config[*].scm_ip_restriction[*].headers", "siteConfig[*].scmIpRestriction[*].headers")
	})
	p.AddResourceConfigurator("azurerm_function_app_hybrid_connection", func(r *config.Resource) {
		r.References["function_app_id"] = config.Reference{
			TerraformName: "azurerm_windows_function_app",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
