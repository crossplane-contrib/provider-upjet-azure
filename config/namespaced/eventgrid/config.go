// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eventgrid

import (
	"github.com/upbound/provider-azure/apis/namespaced/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures eventgrid group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_eventgrid_namespace", func(r *config.Resource) {
		r.Kind = "EventGridNamespace"
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	// Uncomment when terraform provider is updated to 4.38.0 which introduces azurerm_eventgrid_partner_registration
	// p.AddResourceConfigurator("azurerm_eventgrid_partner_configuration", func(r *config.Resource) {
	// 	r.References["partner_authorization.partner_registration_id"] = config.Reference{
	// 		TerraformName: "azurerm_eventgrid_partner_registration",
	// 		Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("partner_registration_id",true)`,
	// 	}
	// })
}
