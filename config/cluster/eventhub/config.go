// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eventhub

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/apis/cluster/rconfig"
	"github.com/upbound/provider-azure/v2/config/cluster/common"
)

// Configure configures resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_eventhub_namespace", func(r *config.Resource) {
		r.Kind = "EventHubNamespace"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"network_rulesets"},
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub", func(r *config.Resource) {
		r.Version = "v1beta3"
		r.PreviousVersions = []string{"v1beta2"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta2")
		r.References["namespace_id"] = config.Reference{
			TerraformName: "azurerm_eventhub_namespace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub_consumer_group", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			TerraformName: "azurerm_eventhub_namespace",
		}
		r.References["eventhub_name"] = config.Reference{
			TerraformName: "azurerm_eventhub",
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub_authorization_rule", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			TerraformName: "azurerm_eventhub_namespace",
		}
		r.References["eventhub_name"] = config.Reference{
			TerraformName: "azurerm_eventhub",
		}
	})
}
