// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package orbital

import (
	"github.com/upbound/provider-azure/v2/apis/namespaced/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures orbital group
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("azurerm_orbital_contact", func(r *config.Resource) {
		r.References["spacecraft_id"] = config.Reference{
			TerraformName: "azurerm_orbital_spacecraft",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["contact_profile_id"] = config.Reference{
			TerraformName: "azurerm_orbital_contact_profile",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
