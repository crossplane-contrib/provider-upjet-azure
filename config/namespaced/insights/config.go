// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package insights

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/upbound/provider-azure/v2/apis/namespaced/rconfig"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures insights group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_monitor_metric_alert", func(r *config.Resource) {
		r.References["scopes"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["action.action_group_id"] = config.Reference{
			TerraformName: "azurerm_monitor_action_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_monitor_private_link_scoped_service", func(r *config.Resource) {
		r.References["scope_name"] = config.Reference{
			TerraformName: "azurerm_monitor_private_link_scope",
		}
		r.References["linked_resource_id"] = config.Reference{
			TerraformName: "azurerm_application_insights",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_activity_log_alert", func(r *config.Resource) {
		r.References["scopes"] = config.Reference{
			TerraformName: "azurerm_resource_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_scheduled_query_rules_alert", func(r *config.Resource) {
		r.References["action.action_group"] = config.Reference{
			TerraformName: "azurerm_monitor_action_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_scheduled_query_rules_alert_v2", func(r *config.Resource) {
		r.References["scopes"] = config.Reference{
			TerraformName: "azurerm_application_insights",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) { //nolint:gocyclo
			if state == nil || state.Empty() || diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			if config == nil {
				return nil, errors.New("resource config cannot be nil")
			}

			duratioDiffKeys := []string{
				"evaluation_frequency",
				"window_duration",
			}
			for _, key := range duratioDiffKeys {
				rcValue, ok := config.Get(key)
				if !ok {
					continue
				}
				rcString, ok := rcValue.(string)
				if !ok {
					continue
				}
				stateValue, ok := state.Attributes[key]
				if !ok {
					continue
				}
				if rcString == "PT60M" && stateValue == "PT1H" {
					delete(diff.Attributes, key)
				}
				if rcString == "PT24H" && stateValue == "P1D" {
					delete(diff.Attributes, key)
				}
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_diagnostic_setting", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			// log is deprecated in favour of the enabled_log property and will be removed in version 4.0 of the AzureRM Provider.
			IgnoredFields: []string{"log", "enabled_log"},
		}
		// target_resource_id references several different types, terraform example auto configures this to only keyvault type
		delete(r.References, "target_resource_id")
	})
}
