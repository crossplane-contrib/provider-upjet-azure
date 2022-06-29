package insights

import "github.com/upbound/upjet/pkg/config"

// Configure configures insights group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_monitor_metric_alert", func(r *config.Resource) {
		r.UseAsync = false
	})
}
