// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package securityinsights

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures securityinsights/sentinel group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_sentinel_log_analytics_workspace_onboarding", func(r *config.Resource) {
		r.References["workspace_name"] = config.Reference{
			TerraformName: "azurerm_log_analytics_workspace",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)`,
		}
	})
}
