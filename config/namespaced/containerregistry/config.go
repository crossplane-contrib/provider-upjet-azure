// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package containerregistry

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/apis/namespaced/rconfig"
	"github.com/upbound/provider-azure/config/namespaced/common"
)

// Configure configures containerregistry group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_container_registry", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"encryption"},
		}
	})

	p.AddResourceConfigurator("azurerm_container_registry_token", func(r *config.Resource) {
		r.References["scope_map_id"] = config.Reference{
			TerraformName: "azurerm_container_registry_scope_map",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_container_connected_registry", func(r *config.Resource) {
		r.References["container_registry_id"] = config.Reference{
			TerraformName: "azurerm_container_registry",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["sync_token_id"] = config.Reference{
			TerraformName: "azurerm_container_registry_token",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "container_registry_id")
	})

	p.AddResourceConfigurator("azurerm_container_registry_webhook", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "registry_name")
	})

	p.AddResourceConfigurator("azurerm_container_registry_credential_set", func(r *config.Resource) {
		r.References["container_registry_id"] = config.Reference{
			TerraformName: "azurerm_container_registry",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["authentication_credentials.username_secret_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_secret",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("versionless_id",true)`,
		}
		r.References["authentication_credentials.password_secret_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_secret",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("versionless_id",true)`,
		}
	})
}
