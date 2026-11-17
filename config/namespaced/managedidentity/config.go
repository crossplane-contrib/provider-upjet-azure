// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package managedidentity

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v3/apis/namespaced/rconfig"
)

// Configure configures managedidentity group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_federated_identity_credential", func(r *config.Resource) {
		r.References["user_assigned_identity_id"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
