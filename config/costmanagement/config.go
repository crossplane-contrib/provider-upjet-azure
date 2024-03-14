// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package costmanagement

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures costmanagement group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_subscription_cost_management_export", func(r *config.Resource) {
		r.References["subscription_id"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/azure/v1beta1.Subscription",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
