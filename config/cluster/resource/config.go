// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package resource

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_resource_group_template_deployment", func(r *config.Resource) {
		r.Kind = "ResourceGroupTemplateDeployment"
	})

	p.AddResourceConfigurator("azurerm_resource_group_policy_assignment", func(r *config.Resource) {
		r.Kind = "ResourceGroupPolicyAssignment"
	})
}
