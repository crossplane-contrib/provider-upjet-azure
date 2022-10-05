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

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures virtual group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_ip_group", func(r *config.Resource) {
		r.UseAsync = false
		r.Kind = "IPGroup"
	})

	p.AddResourceConfigurator("azurerm_network_interface", func(r *config.Resource) {
		r.UseAsync = false
		r.Kind = "NetworkInterface"
	})

	p.AddResourceConfigurator("azurerm_lb", func(r *config.Resource) {
		r.Kind = "LoadBalancer"

		r.References["frontend_ip_configuration.public_ip_address_id"] = config.Reference{
			Type:      "PublicIP",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_backend_address_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerBackendAddressPool"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
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
	})

	p.AddResourceConfigurator("azurerm_lb_nat_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatPool"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_nat_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatRule"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
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
	})

	p.AddResourceConfigurator("azurerm_lb_probe", func(r *config.Resource) {
		r.Kind = "LoadBalancerProbe"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerRule"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_nat_gateway_public_ip_association", func(r *config.Resource) {
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
		r.UseAsync = false
		r.Kind = "Watcher"
	})

	p.AddResourceConfigurator("azurerm_network_watcher_flow_log", func(r *config.Resource) {
		r.UseAsync = false
		r.References["network_watcher_name"] = config.Reference{
			Type: "Watcher",
		}
		r.References["network_security_group_id"] = config.Reference{
			Type:      "SecurityGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_account_id"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/storage/v1beta1.Account",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		delete(r.References, "traffic_analytics.workspace_region")
	})

	p.AddResourceConfigurator("azurerm_network_connection_monitor", func(r *config.Resource) {
		r.UseAsync = false
		r.Kind = "ConnectionMonitor"
		r.References["network_watcher_id"] = config.Reference{
			Type:      "Watcher",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_ddos_protection_plan", func(r *config.Resource) {
		r.UseAsync = false
		r.Kind = "DDoSProtectionPlan"
	})

	p.AddResourceConfigurator("azurerm_network_security_rule", func(r *config.Resource) {
		r.UseAsync = false
		r.References["network_security_group_name"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_application_security_group_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceApplicationSecurityGroupAssociation"
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
	})

	p.AddResourceConfigurator("azurerm_network_profile", func(r *config.Resource) {
		r.UseAsync = false
		r.References["container_network_interface.ip_configuration.subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_a_record", func(r *config.Resource) {
		r.UseAsync = false
		r.References["zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_aaaa_record", func(r *config.Resource) {
		r.UseAsync = false
		r.References["zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_cname_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_mx_record", func(r *config.Resource) {
		r.UseAsync = false
		r.References["zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_ptr_record", func(r *config.Resource) {
		r.UseAsync = false
		r.References["zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_srv_record", func(r *config.Resource) {
		r.UseAsync = false
		r.References["zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_txt_record", func(r *config.Resource) {
		r.UseAsync = false
		r.References["zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_zone_virtual_network_link", func(r *config.Resource) {
		r.References["private_dns_zone_name"] = config.Reference{
			Type: "PrivateDNSZone",
		}
		r.References["virtual_network_id"] = config.Reference{
			Type:      "VirtualNetwork",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_dns_a_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_aaaa_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_caa_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_cname_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_mx_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_ns_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_ptr_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_srv_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_txt_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			Type: "DNSZone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_link_service", func(r *config.Resource) {
		r.References["nat_ip_configuration.subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_private_endpoint", func(r *config.Resource) {
		r.References["private_dns_zone_group.private_dns_zone_ids"] = config.Reference{
			Type:      "PrivateDNSZone",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		delete(r.References, "private_service_connection.private_connection_resource_id")
	})

	p.AddResourceConfigurator("azurerm_network_packet_capture", func(r *config.Resource) {
		r.References["network_watcher_name"] = config.Reference{
			Type: "Watcher",
		}
		r.References["storage_location.storage_account_id"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/storage/v1beta1.Account",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_point_to_site_vpn_gateway", func(r *config.Resource) {
		r.References["virtual_hub_id"] = config.Reference{
			Type:      "VirtualHub",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["vpn_server_configuration_id"] = config.Reference{
			Type:      "VPNServerConfiguration",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_circuit_authorization", func(r *config.Resource) {
		r.References["express_route_circuit_name"] = config.Reference{
			Type: "ExpressRouteCircuit",
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_circuit_peering", func(r *config.Resource) {
		r.References["express_route_circuit_name"] = config.Reference{
			Type: "ExpressRouteCircuit",
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_circuit_connection", func(r *config.Resource) {
		r.References["peering_id"] = config.Reference{
			Type:      "ExpressRouteCircuitPeering",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["peer_peering_id"] = config.Reference{
			Type:      "ExpressRouteCircuitPeering",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_gateway", func(r *config.Resource) {
		r.References["virtual_hub_id"] = config.Reference{
			Type:      "VirtualHub",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_connection", func(r *config.Resource) {
		r.References["express_route_gateway_id"] = config.Reference{
			Type:      "ExpressRouteGateway",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["express_route_circuit_peering_id"] = config.Reference{
			Type:      "ExpressRouteCircuitPeering",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_firewall", func(r *config.Resource) {
		r.References["ip_configuration.subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["ip_configuration.public_ip_address_id"] = config.Reference{
			Type:      "PublicIP",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_application_rule_collection", func(r *config.Resource) {
		r.References["azure_firewall_name"] = config.Reference{
			Type: "Firewall",
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_nat_rule_collection", func(r *config.Resource) {
		r.References["azure_firewall_name"] = config.Reference{
			Type: "Firewall",
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_network_rule_collection", func(r *config.Resource) {
		r.References["azure_firewall_name"] = config.Reference{
			Type: "Firewall",
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_policy_rule_collection_group", func(r *config.Resource) {
		r.References["firewall_policy_id"] = config.Reference{
			Type:      "FirewallPolicy",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_frontdoor_custom_https_configuration", func(r *config.Resource) {
		r.References["custom_https_configuration.azure_key_vault_certificate_vault_id"] = config.Reference{
			Type:      rconfig.VaultKeyReferencePath,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_frontdoor_rules_engine", func(r *config.Resource) {
		r.References["frontdoor_name"] = config.Reference{
			Type: "FrontDoor",
		}
	})

	p.AddResourceConfigurator("azurerm_application_gateway", func(r *config.Resource) {
		r.References["frontend_ip_configuration.public_ip_address_id"] = config.Reference{
			Type:      "PublicIP",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
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
	})

	p.AddResourceConfigurator("azurerm_virtual_hub", func(r *config.Resource) {
		r.References["virtual_wan_id"] = config.Reference{
			Type:      "VirtualWAN",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_frontdoor", func(r *config.Resource) {
		r.UseAsync = false
		r.Kind = "FrontDoor"
	})

	p.AddResourceConfigurator("azurerm_subnet", func(r *config.Resource) {
		r.Kind = "Subnet"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"address_prefix"},
		}
		r.References["virtual_network_name"] = config.Reference{
			Type: "VirtualNetwork",
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_nat_gateway_association", func(r *config.Resource) {
		r.Kind = "SubnetNATGatewayAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_network_security_group_association", func(r *config.Resource) {
		r.Kind = "SubnetNetworkSecurityGroupAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["network_security_group_id"] = config.Reference{
			Type:      "SecurityGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_service_endpoint_storage_policy", func(r *config.Resource) {
		r.Kind = "SubnetServiceEndpointStoragePolicy"
	})

	p.AddResourceConfigurator("azurerm_subnet_route_table_association", func(r *config.Resource) {
		r.Kind = "SubnetRouteTableAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["route_table_id"] = config.Reference{
			Type:      "RouteTable",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_application_security_group", func(r *config.Resource) {
		r.UseAsync = false
	})

	p.AddResourceConfigurator("azurerm_private_dns_zone", func(r *config.Resource) {
		r.UseAsync = false
	})

	p.AddResourceConfigurator("azurerm_public_ip", func(r *config.Resource) {
		r.UseAsync = false
	})

	p.AddResourceConfigurator("azurerm_public_ip_prefix", func(r *config.Resource) {
		r.UseAsync = false
	})

	p.AddResourceConfigurator("azurerm_network_security_group", func(r *config.Resource) {
		r.UseAsync = false
	})

	p.AddResourceConfigurator("azurerm_virtual_network", func(r *config.Resource) {
		r.UseAsync = false
	})
}
