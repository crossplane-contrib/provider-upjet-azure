// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cognitiveservices

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/apis/namespaced/rconfig"
)

const group = "cognitiveservices"

// Configure configures cognitiveservices group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_cognitive_account", func(r *config.Resource) {
		r.ShortGroup = group
		r.References["storage.storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage.identity_client_id"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("client_id",true)`,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_ai_services", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AIServices"
		r.References["storage.storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage.identity_client_id"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("client_id",true)`,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_cognitive_account_rai_blocklist", func(r *config.Resource) {
		r.ShortGroup = group
		r.References["cognitive_account_id"] = config.Reference{
			TerraformName: "azurerm_cognitive_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_cognitive_account_rai_policy", func(r *config.Resource) {
		r.ShortGroup = group
		r.References["cognitive_account_id"] = config.Reference{
			TerraformName: "azurerm_cognitive_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
