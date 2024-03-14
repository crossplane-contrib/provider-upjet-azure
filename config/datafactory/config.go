// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package datafactory

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures datafactory group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_data_factory_trigger_schedule", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"pipeline_name", "pipeline"},
		}
	})

	p.AddResourceConfigurator("azurerm_data_factory_linked_service_kusto", func(r *config.Resource) {
		r.Path = "linkedservicekustoes"
	})

	p.AddResourceConfigurator("azurerm_data_factory_integration_runtime_azure_ssis", func(r *config.Resource) {
		r.Path = "integrationruntimeazuressis"
	})
}
