// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package desktopvirtualization

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/apis/cluster/rconfig"
)

const group = "desktopvirtualization"

// Configure configures desktopvirtualization group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_virtual_desktop_host_pool", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "VirtualDesktopHostPool"
	})
	p.AddResourceConfigurator("azurerm_virtual_desktop_host_pool_registration_info", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "VirtualDesktopHostPoolRegistrationInfo"
		r.References["hostpool_id"] = config.Reference{
			TerraformName: "azurerm_virtual_desktop_host_pool",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_virtual_desktop_application_group", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "VirtualDesktopApplicationGroup"
		r.References["host_pool_id"] = config.Reference{
			TerraformName: "azurerm_virtual_desktop_host_pool",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_virtual_desktop_workspace_application_group_association", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "VirtualDesktopWorkspaceAppplicationGroupAssociation"
		r.References["workspace_id"] = config.Reference{
			TerraformName: "azurerm_virtual_desktop_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["application_group_id"] = config.Reference{
			TerraformName: "azurerm_virtual_desktop_application_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
