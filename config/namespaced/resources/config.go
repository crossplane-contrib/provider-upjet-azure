// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package resources

import (
	"github.com/upbound/provider-azure/apis/namespaced/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures resources group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_resource_deployment_script_azure_cli", func(r *config.Resource) {
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.Path = "resourcedeploymentscriptazureclicli"
	})
	p.AddResourceConfigurator("azurerm_resource_deployment_script_azure_power_shell", func(r *config.Resource) {
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
