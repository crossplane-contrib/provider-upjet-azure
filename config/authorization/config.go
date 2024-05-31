// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package authorization

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/config/common"
)

// Configure configures authorization group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_role_assignment", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"role_definition_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_resource_group_policy_assignment", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "resource_group_id")
	})
	p.AddResourceConfigurator("azurerm_pim_active_role_assignment", func(r *config.Resource) {
		r.PreviousVersions = nil
		r.Version = "v1beta1"
		r.Conversions = nil
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}
	})
	p.AddResourceConfigurator("azurerm_pim_eligible_role_assignment", func(r *config.Resource) {
		r.PreviousVersions = nil
		r.Version = "v1beta1"
		r.Conversions = nil
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}
	})
}
