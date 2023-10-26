/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cdn

import (
	"github.com/crossplane/upjet/pkg/config"
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
			Type:      "github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorOrigin",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["cdn_frontdoor_rule_set_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorRuleSet",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["cdn_frontdoor_custom_domain_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorCustomDomain",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
	p.AddResourceConfigurator("azurerm_cdn_frontdoor_custom_domain_association", func(r *config.Resource) {
		r.References["cdn_frontdoor_route_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorRoute",
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
