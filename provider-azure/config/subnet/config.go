/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package subnet

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
)

const groupNetwork = "network"

// Configure configures subnet group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_subnet", func(r *config.Resource) {
		r.Kind = "Subnet"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"address_prefix"},
		}
		r.References["virtual_network_name"] = config.Reference{
			Type: "VirtualNetwork",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_subnet_nat_gateway_association", func(r *config.Resource) {
		r.Kind = "SubnetNATGatewayAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_subnet_network_security_group_association", func(r *config.Resource) {
		r.Kind = "SubnetNetworkSecurityGroupAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_subnet_service_endpoint_storage_policy", func(r *config.Resource) {
		r.Kind = "SubnetServiceEndpointStoragePolicy"
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_subnet_route_table_association", func(r *config.Resource) {
		r.Kind = "SubnetRouteTableAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})
}
