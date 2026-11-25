// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package securityinsights

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/config/namespaced/common"
)

// Configure configures securityinsights/sentinel group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_sentinel_log_analytics_workspace_onboarding", func(r *config.Resource) {
		r.References["workspace_name"] = config.Reference{
			TerraformName: "azurerm_log_analytics_workspace",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",true)`,
		}
	})
	p.AddResourceConfigurator("azurerm_sentinel_alert_rule_fusion", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
	})
}
