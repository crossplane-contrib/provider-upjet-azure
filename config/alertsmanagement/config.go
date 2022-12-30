package alertsmanagement

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure alertsmanagement group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_monitor_smart_detector_alert_rule", func(r *config.Resource) {
		r.References["action_group.ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/insights/v1beta1.MonitorActionGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["scope_resource_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
