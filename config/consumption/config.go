package consumption

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures consumption group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_consumption_budget_subscription", func(r *config.Resource) {
		r.References["notification.contact_groups"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/insights/v1beta1.MonitorActionGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
