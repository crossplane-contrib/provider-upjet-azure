// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package operationalinsights

import (
	"github.com/upbound/provider-azure/v2/apis/namespaced/rconfig"
	"github.com/upbound/provider-azure/v2/config/namespaced/common"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures operationalinsights group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_log_analytics_workspace", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
		r.Kind = "Workspace"
	})

	p.AddResourceConfigurator("azurerm_log_analytics_linked_storage_account", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
		r.References["storage_account_ids"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
