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

package network

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
)

// Configure configures virtual group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_ip_group", func(r *config.Resource) {
		r.Kind = "IPGroup"
	})

	p.AddResourceConfigurator("azurerm_network_interface", func(r *config.Resource) {
		r.Kind = "NetworkInterface"
	})

	p.AddResourceConfigurator("azurerm_lb", func(r *config.Resource) {
		r.Kind = "LoadBalancer"

		r.References["frontend_ip_configuration.public_ip_address_id"] = config.Reference{
			Type:      "PublicIP",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_lb_backend_address_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerBackendAddressPool"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_lb_backend_address_pool_address", func(r *config.Resource) {
		r.Kind = "LoadBalancerBackendAddressPoolAddress"
		r.References["backend_address_pool_id"] = config.Reference{
			Type:      "LoadBalancerBackendAddressPool",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["virtual_network_id"] = config.Reference{
			Type:      "VirtualNetwork",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_lb_nat_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatPool"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_lb_nat_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatRule"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_lb_outbound_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerOutboundRule"
		r.References["backend_address_pool_id"] = config.Reference{
			Type:      "LoadBalancerBackendAddressPool",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_lb_probe", func(r *config.Resource) {
		r.Kind = "LoadBalancerProbe"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_lb_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerRule"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_local_network_gateway", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_nat_gateway", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_nat_gateway_public_ip_association", func(r *config.Resource) {
		r.UseAsync = true
		r.References["nat_gateway_id"] = config.Reference{
			Type:      "NATGateway",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["public_ip_address_id"] = config.Reference{
			Type:      "PublicIP",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_nat_gateway_public_ip_prefix_association", func(r *config.Resource) {
		r.UseAsync = true
		r.References["nat_gateway_id"] = config.Reference{
			Type:      "NATGateway",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["public_ip_prefix_id"] = config.Reference{
			Type:      "PublicIPPrefix",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_watcher", func(r *config.Resource) {
		r.Kind = "Watcher"
	})

	p.AddResourceConfigurator("azurerm_network_connection_monitor", func(r *config.Resource) {
		r.Kind = "ConnectionMonitor"
		r.References["network_watcher_id"] = config.Reference{
			Type:      "Watcher",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_ddos_protection_plan", func(r *config.Resource) {
		r.Kind = "DDoSProtectionPlan"
	})

	p.AddResourceConfigurator("azurerm_network_security_rule", func(r *config.Resource) {
		r.References["network_security_group_name"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_application_security_group_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceApplicationSecurityGroupAssociation"
		r.UseAsync = true
		r.References["network_interface_id"] = config.Reference{
			Type:      "NetworkInterface",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["application_security_group_id"] = config.Reference{
			Type:      "ApplicationSecurityGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_backend_address_pool_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceBackendAddressPoolAssociation"
		r.UseAsync = true
		r.References["network_interface_id"] = config.Reference{
			Type:      "NetworkInterface",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["backend_address_pool_id"] = config.Reference{
			Type:      "LoadBalancerBackendAddressPool",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_nat_rule_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceNatRuleAssociation"
		r.UseAsync = true
		r.References["network_interface_id"] = config.Reference{
			Type:      "NetworkInterface",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["nat_rule_id"] = config.Reference{
			Type:      "LoadBalancerNatRule",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_security_group_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceSecurityGroupAssociation"
		r.UseAsync = true
		r.References["network_interface_id"] = config.Reference{
			Type:      "NetworkInterface",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["network_security_group_id"] = config.Reference{
			Type:      "SecurityGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_virtual_network", func(r *config.Resource) {
		r.Kind = "VirtualNetwork"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"subnet"},
		}
		config.MoveToStatus(r.TerraformResource, "subnet")
	})

	p.AddResourceConfigurator("azurerm_virtual_network_gateway", func(r *config.Resource) {
		r.Kind = "VirtualNetworkGateway"
		r.References["ip_configuration.subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_virtual_network_peering", func(r *config.Resource) {
		r.Kind = "VirtualNetworkPeering"
		r.References["virtual_network_name"] = config.Reference{
			Type: "VirtualNetwork",
		}
		r.References["remote_virtual_network_id"] = config.Reference{
			Type:      "VirtualNetwork",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	/*p.AddResourceConfigurator("azurerm_virtual_desktop_application", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
	})*/

	/*p.AddResourceConfigurator("azurerm_virtual_desktop_host_pool", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
	})*/

	p.AddResourceConfigurator("azurerm_virtual_network_gateway_connection", func(r *config.Resource) {
		r.Kind = "VirtualNetworkGatewayConnection"
		r.References["virtual_network_gateway_id"] = config.Reference{
			Type:      "VirtualNetworkGateway",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["peer_virtual_network_gateway_id"] = config.Reference{
			Type:      "VirtualNetworkGateway",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	/*p.AddResourceConfigurator("azurerm_virtual_desktop_workspace", func(r *config.Resource) {
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type:      rconfig.APISPackagePath + "/azure/v1alpha1.ResourceGroup",
			},
		}
		r.UseAsync = true
	})*/

	p.AddResourceConfigurator("azurerm_virtual_wan", func(r *config.Resource) {
		r.Kind = "VirtualWAN"
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_frontdoor", func(r *config.Resource) {
		r.Kind = "FrontDoor"
	})

	p.AddResourceConfigurator("azurerm_network_packet_capture", func(r *config.Resource) {
		r.Kind = "NetworkPacketCapture"
	})

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
