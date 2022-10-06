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
}
