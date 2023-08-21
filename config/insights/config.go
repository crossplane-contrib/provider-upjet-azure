package insights

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures insights group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_monitor_metric_alert", func(r *config.Resource) {
		r.References["scopes"] = config.Reference{
			Type:      rconfig.StorageAccountReferencePath,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["action.action_group_id"] = config.Reference{
			Type:      "MonitorActionGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_monitor_private_link_scoped_service", func(r *config.Resource) {
		r.References["scope_name"] = config.Reference{
			Type: "MonitorPrivateLinkScope",
		}
		r.References["linked_resource_id"] = config.Reference{
			Type:      "ApplicationInsights",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_activity_log_alert", func(r *config.Resource) {
		r.References["scopes"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_scheduled_query_rules_alert", func(r *config.Resource) {
		r.References["action.action_group"] = config.Reference{
			Type:      "MonitorActionGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_scheduled_query_rules_alert_v2", func(r *config.Resource) {
		r.References["scopes"] = config.Reference{
			Type:      "ApplicationInsights",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_monitor_diagnostic_setting", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			// log is deprecated in favour of the enabled_log property and will be removed in version 4.0 of the AzureRM Provider.
			IgnoredFields: []string{"log", "enabled_log"},
		}
	})
}
