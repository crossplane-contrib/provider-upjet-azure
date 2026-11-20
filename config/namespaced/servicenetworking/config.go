// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package servicenetworking

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/apis/namespaced/rconfig"
)

const group = "servicenetworking"

// Configure configures servicenetworking group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_application_load_balancer", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "ApplicationLoadBalancer"
	})
	p.AddResourceConfigurator("azurerm_application_load_balancer_frontend", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "ApplicationLoadBalancerFrontend"
		r.References["application_load_balancer_id"] = config.Reference{
			TerraformName: "azurerm_application_load_balancer",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_application_load_balancer_security_policy", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "ApplicationLoadBalancerSecurityPolicy"
		r.References["application_load_balancer_id"] = config.Reference{
			TerraformName: "azurerm_application_load_balancer",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["web_application_firewall_policy_id"] = config.Reference{
			TerraformName: "azurerm_web_application_firewall_policy",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_application_load_balancer_subnet_association", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "ApplicationLoadBalancerSubnetAssociation"
		r.References["application_load_balancer_id"] = config.Reference{
			TerraformName: "azurerm_application_load_balancer",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
