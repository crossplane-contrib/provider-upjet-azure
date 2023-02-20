package appplatform

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures appplatform group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_spring_cloud_api_portal", func(r *config.Resource) {
		r.References["gateway_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudGateway",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
