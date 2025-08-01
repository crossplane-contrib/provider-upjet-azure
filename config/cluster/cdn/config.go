// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cdn

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures cdn group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_cdn_frontdoor_origin", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"health_probes_enabled"},
		}
	})
	p.AddResourceConfigurator("azurerm_cdn_frontdoor_route", func(r *config.Resource) {
		r.References["cdn_frontdoor_origin_ids"] = config.Reference{
			TerraformName: "azurerm_cdn_frontdoor_origin",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
		r.References["cdn_frontdoor_rule_set_ids"] = config.Reference{
			TerraformName: "azurerm_cdn_frontdoor_rule_set",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
		r.References["cdn_frontdoor_custom_domain_ids"] = config.Reference{
			TerraformName: "azurerm_cdn_frontdoor_custom_domain",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("azurerm_cdn_frontdoor_custom_domain_association", func(r *config.Resource) {
		r.References["cdn_frontdoor_route_ids"] = config.Reference{
			TerraformName: "azurerm_cdn_frontdoor_route",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("azurerm_cdn_frontdoor_firewall_policy", func(r *config.Resource) {
		r.ShortGroup = "cdn"
		r.Kind = "FrontdoorFirewallPolicy"
	})
}
