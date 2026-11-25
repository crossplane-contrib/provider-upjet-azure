// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kusto

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/apis/namespaced/rconfig"
	"github.com/upbound/provider-azure/v2/config/namespaced/common"
)

// Configure configures kusto group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_kusto_database", func(r *config.Resource) {
		r.References["cluster_name"] = config.Reference{
			TerraformName: "azurerm_kusto_cluster",
		}
	})

	p.AddResourceConfigurator("azurerm_kusto_attached_database_configuration", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
		r.References["cluster_id"] = config.Reference{
			TerraformName: "azurerm_kusto_cluster",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_kusto_cluster", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
	})

	p.AddResourceConfigurator("azurerm_kusto_eventgrid_data_connection", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
	})
}
