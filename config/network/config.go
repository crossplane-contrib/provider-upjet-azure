// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package network

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures virtual group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_ip_group", func(r *config.Resource) {
		r.Kind = "IPGroup"
	})

	p.AddResourceConfigurator("azurerm_network_interface", func(r *config.Resource) {
		r.Kind = "NetworkInterface"

		r.References["ip_configuration.public_ip_address_id"] = config.Reference{
			TerraformName: "azurerm_public_ip",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}

		r.OverrideFieldNames["IPConfigurationInitParameters"] = "NetworkInterfaceIPConfigurationInitParameters"
		r.OverrideFieldNames["IPConfigurationParameters"] = "NetworkInterfaceIPConfigurationParameters"
		r.OverrideFieldNames["IPConfigurationObservation"] = "NetworkInterfaceIPConfigurationObservation"
	})

	p.AddResourceConfigurator("azurerm_lb", func(r *config.Resource) {
		r.Kind = "LoadBalancer"

		r.References["frontend_ip_configuration.public_ip_address_id"] = config.Reference{
			TerraformName: "azurerm_public_ip",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.OverrideFieldNames["FrontendIPConfigurationParameters"] = "LoadBalancerFrontendIPConfigurationParameters"
		r.OverrideFieldNames["FrontendIPConfigurationInitParameters"] = "LoadBalancerFrontendIPConfigurationInitParameters"
		r.OverrideFieldNames["FrontendIPConfigurationObservation"] = "LoadBalancerFrontendIPConfigurationObservation"
	})

	p.AddResourceConfigurator("azurerm_lb_backend_address_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerBackendAddressPool"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "azurerm_lb",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_backend_address_pool_address", func(r *config.Resource) {
		r.Kind = "LoadBalancerBackendAddressPoolAddress"
		r.References["backend_address_pool_id"] = config.Reference{
			TerraformName: "azurerm_lb_backend_address_pool",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["virtual_network_id"] = config.Reference{
			TerraformName: "azurerm_virtual_network",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_nat_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatPool"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "azurerm_lb",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_nat_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatRule"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "azurerm_lb",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_outbound_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerOutboundRule"
		r.References["backend_address_pool_id"] = config.Reference{
			TerraformName: "azurerm_lb_backend_address_pool",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "azurerm_lb",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}

		r.OverrideFieldNames["FrontendIPConfigurationInitParameters"] = "LoadBalancerOutboundRuleFrontendIPConfigurationInitParameters"
		r.OverrideFieldNames["FrontendIPConfigurationParameters"] = "LoadBalancerOutboundRuleFrontendIPConfigurationParameters"
		r.OverrideFieldNames["FrontendIPConfigurationObservation"] = "LoadBalancerOutboundRuleFrontendIPConfigurationObservation"
	})

	p.AddResourceConfigurator("azurerm_lb_probe", func(r *config.Resource) {
		r.Kind = "LoadBalancerProbe"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "azurerm_lb",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_lb_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerRule"
		r.References["loadbalancer_id"] = config.Reference{
			TerraformName: "azurerm_lb",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_nat_gateway_public_ip_association", func(r *config.Resource) {
		r.References["nat_gateway_id"] = config.Reference{
			TerraformName: "azurerm_nat_gateway",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["public_ip_address_id"] = config.Reference{
			TerraformName: "azurerm_public_ip",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_nat_gateway_public_ip_prefix_association", func(r *config.Resource) {
		r.References["nat_gateway_id"] = config.Reference{
			TerraformName: "azurerm_nat_gateway",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["public_ip_prefix_id"] = config.Reference{
			TerraformName: "azurerm_public_ip_prefix",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_watcher", func(r *config.Resource) {
		r.Kind = "Watcher"
	})

	p.AddResourceConfigurator("azurerm_network_watcher_flow_log", func(r *config.Resource) {
		r.References["network_watcher_name"] = config.Reference{
			TerraformName: "azurerm_network_watcher",
		}
		r.References["network_security_group_id"] = config.Reference{
			TerraformName: "azurerm_network_security_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		delete(r.References, "traffic_analytics.workspace_region")
	})

	p.AddResourceConfigurator("azurerm_network_connection_monitor", func(r *config.Resource) {
		r.Kind = "ConnectionMonitor"
		r.References["network_watcher_id"] = config.Reference{
			TerraformName: "azurerm_network_watcher",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_ddos_protection_plan", func(r *config.Resource) {
		r.Kind = "DDoSProtectionPlan"
	})

	p.AddResourceConfigurator("azurerm_network_security_rule", func(r *config.Resource) {
		r.References["network_security_group_name"] = config.Reference{
			TerraformName: "azurerm_network_security_group",
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_application_security_group_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceApplicationSecurityGroupAssociation"
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "azurerm_network_interface",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["application_security_group_id"] = config.Reference{
			TerraformName: "azurerm_application_security_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_backend_address_pool_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceBackendAddressPoolAssociation"
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "azurerm_network_interface",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["backend_address_pool_id"] = config.Reference{
			TerraformName: "azurerm_lb_backend_address_pool",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_nat_rule_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceNatRuleAssociation"
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "azurerm_network_interface",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["nat_rule_id"] = config.Reference{
			TerraformName: "azurerm_lb_nat_rule",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_security_group_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceSecurityGroupAssociation"
		r.References["network_interface_id"] = config.Reference{
			TerraformName: "azurerm_network_interface",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["network_security_group_id"] = config.Reference{
			TerraformName: "azurerm_network_security_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
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
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_virtual_network_peering", func(r *config.Resource) {
		r.Kind = "VirtualNetworkPeering"
		r.References["virtual_network_name"] = config.Reference{
			TerraformName: "azurerm_virtual_network",
		}
		r.References["remote_virtual_network_id"] = config.Reference{
			TerraformName: "azurerm_virtual_network",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_profile", func(r *config.Resource) {
		r.References["container_network_interface.ip_configuration.subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_a_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_aaaa_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_cname_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_mx_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_ptr_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_srv_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_txt_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_dns_zone_virtual_network_link", func(r *config.Resource) {
		r.References["private_dns_zone_name"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
		}
		r.References["virtual_network_id"] = config.Reference{
			TerraformName: "azurerm_virtual_network",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_dns_a_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_aaaa_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_caa_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_cname_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_mx_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_ns_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_ptr_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_srv_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_dns_txt_record", func(r *config.Resource) {
		r.References["zone_name"] = config.Reference{
			TerraformName: "azurerm_dns_zone",
		}
	})

	p.AddResourceConfigurator("azurerm_private_link_service", func(r *config.Resource) {
		r.References["nat_ip_configuration.subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_private_endpoint", func(r *config.Resource) {
		r.References["private_dns_zone_group.private_dns_zone_ids"] = config.Reference{
			TerraformName: "azurerm_private_dns_zone",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		delete(r.References, "private_service_connection.private_connection_resource_id")
	})

	p.AddResourceConfigurator("azurerm_network_packet_capture", func(r *config.Resource) {
		r.References["network_watcher_name"] = config.Reference{
			TerraformName: "azurerm_network_watcher",
		}
		r.References["storage_location.storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_point_to_site_vpn_gateway", func(r *config.Resource) {
		r.References["virtual_hub_id"] = config.Reference{
			TerraformName: "azurerm_virtual_hub",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["vpn_server_configuration_id"] = config.Reference{
			TerraformName: "azurerm_vpn_server_configuration",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_circuit_authorization", func(r *config.Resource) {
		r.References["express_route_circuit_name"] = config.Reference{
			TerraformName: "azurerm_express_route_circuit",
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_circuit_peering", func(r *config.Resource) {
		r.References["express_route_circuit_name"] = config.Reference{
			TerraformName: "azurerm_express_route_circuit",
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_circuit_connection", func(r *config.Resource) {
		r.References["peering_id"] = config.Reference{
			TerraformName: "azurerm_express_route_circuit_peering",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["peer_peering_id"] = config.Reference{
			TerraformName: "azurerm_express_route_circuit_peering",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_gateway", func(r *config.Resource) {
		r.References["virtual_hub_id"] = config.Reference{
			TerraformName: "azurerm_virtual_hub",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_express_route_connection", func(r *config.Resource) {
		r.References["express_route_gateway_id"] = config.Reference{
			TerraformName: "azurerm_express_route_gateway",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["express_route_circuit_peering_id"] = config.Reference{
			TerraformName: "azurerm_express_route_circuit_peering",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_firewall", func(r *config.Resource) {
		r.References["ip_configuration.subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["ip_configuration.public_ip_address_id"] = config.Reference{
			TerraformName: "azurerm_public_ip",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_application_rule_collection", func(r *config.Resource) {
		r.References["azure_firewall_name"] = config.Reference{
			TerraformName: "azurerm_firewall",
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_nat_rule_collection", func(r *config.Resource) {
		r.References["azure_firewall_name"] = config.Reference{
			TerraformName: "azurerm_firewall",
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_network_rule_collection", func(r *config.Resource) {
		r.References["azure_firewall_name"] = config.Reference{
			TerraformName: "azurerm_firewall",
		}
	})

	p.AddResourceConfigurator("azurerm_firewall_policy_rule_collection_group", func(r *config.Resource) {
		r.References["firewall_policy_id"] = config.Reference{
			TerraformName: "azurerm_firewall_policy",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_frontdoor_custom_https_configuration", func(r *config.Resource) {
		r.References["custom_https_configuration.azure_key_vault_certificate_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_key",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_frontdoor_rules_engine", func(r *config.Resource) {
		r.References["frontdoor_name"] = config.Reference{
			TerraformName: "azurerm_frontdoor",
		}
	})

	p.AddResourceConfigurator("azurerm_application_gateway", func(r *config.Resource) {
		r.References["frontend_ip_configuration.public_ip_address_id"] = config.Reference{
			TerraformName: "azurerm_public_ip",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
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
			TerraformName: "azurerm_virtual_network_gateway",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["peer_virtual_network_gateway_id"] = config.Reference{
			TerraformName: "azurerm_virtual_network_gateway",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
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
			TerraformName: "azurerm_virtual_wan",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.OverrideFieldNames["VirtualHubParameters"] = "VirtualHubParameters_2"
		r.OverrideFieldNames["VirtualHubInitParameters"] = "VirtualHubInitParameters_2"
		r.OverrideFieldNames["VirtualHubObservation"] = "VirtualHubObservation_2"

		r.OverrideFieldNames["RouteInitParameters"] = "VirtualHubRouteInitParameters"
		r.OverrideFieldNames["RouteParameters"] = "VirtualHubRouteParameters"
		r.OverrideFieldNames["RouteObservation"] = "VirtualHubRouteObservation"

	})

	p.AddResourceConfigurator("azurerm_frontdoor", func(r *config.Resource) {
		r.Kind = "FrontDoor"
	})

	p.AddResourceConfigurator("azurerm_subnet", func(r *config.Resource) {
		r.Kind = "Subnet"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"address_prefix"},
		}
		r.References["virtual_network_name"] = config.Reference{
			TerraformName: "azurerm_virtual_network",
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_nat_gateway_association", func(r *config.Resource) {
		r.Kind = "SubnetNATGatewayAssociation"
		r.References["subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_network_security_group_association", func(r *config.Resource) {
		r.Kind = "SubnetNetworkSecurityGroupAssociation"
		r.References["subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["network_security_group_id"] = config.Reference{
			TerraformName: "azurerm_network_security_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_vpn_gateway", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"bgp_settings"},
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_service_endpoint_storage_policy", func(r *config.Resource) {
		r.Kind = "SubnetServiceEndpointStoragePolicy"
	})

	p.AddResourceConfigurator("azurerm_route_table", func(r *config.Resource) {
		r.OverrideFieldNames["RouteInitParameters"] = "RouteTableRouteInitParameters"
		r.OverrideFieldNames["RouteParameters"] = "RouteTableRouteParameters"
		r.OverrideFieldNames["RouteObservation"] = "RouteTableRouteObservation"
	})

	p.AddResourceConfigurator("azurerm_virtual_hub_route_table", func(r *config.Resource) {
		r.OverrideFieldNames["RouteInitParameters"] = "VirtualHubRouteTableRouteInitParameters"
		r.OverrideFieldNames["RouteParameters"] = "VirtualHubRouteTableRouteParameters"
		r.OverrideFieldNames["RouteObservation"] = "VirtualHubRouteTableRouteObservation"
	})

	p.AddResourceConfigurator("azurerm_subnet_route_table_association", func(r *config.Resource) {
		r.Kind = "SubnetRouteTableAssociation"
		r.References["subnet_id"] = config.Reference{
			TerraformName: "azurerm_subnet",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["route_table_id"] = config.Reference{
			TerraformName: "azurerm_route_table",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_virtual_hub_connection", func(r *config.Resource) {
		r.References["routing.associated_route_table_id"] = config.Reference{
			TerraformName: "azurerm_virtual_hub_route_table",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_manager_static_member", func(r *config.Resource) {
		r.References["target_virtual_network_id"] = config.Reference{
			TerraformName: "azurerm_virtual_network",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_manager_management_group_connection", func(r *config.Resource) {
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "network_manager_id")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("azurerm_frontdoor_firewall_policy", func(r *config.Resource) {
		r.OverrideFieldNames["ExclusionParameters"] = "ManagedRuleExclusionParameters"
		r.OverrideFieldNames["ExclusionInitParameters"] = "ManagedRuleExclusionInitParameters"
		r.OverrideFieldNames["ExclusionObservation"] = "ManagedRuleExclusionObservation"
	})

	p.AddResourceConfigurator("azurerm_route", func(r *config.Resource) {
		r.OverrideFieldNames["RouteParameters"] = "RouteParameters_2"
		r.OverrideFieldNames["RouteInitParameters"] = "RouteInitParameters_2"
		r.OverrideFieldNames["RouteObservation"] = "RouteObservation_2"
	})

	p.AddResourceConfigurator("azurerm_route_map", func(r *config.Resource) {
		r.OverrideFieldNames["ActionParameters"] = "RuleActionParameters"
		r.OverrideFieldNames["ActionInitParameters"] = "RuleActionInitParameters"
		r.OverrideFieldNames["ActionObservation"] = "RuleActionObservation"
	})

	p.AddResourceConfigurator("azurerm_traffic_manager_azure_endpoint", func(r *config.Resource) {
		r.OverrideFieldNames["SubnetParameters"] = "TrafficManagerAzureEndpointSubnetParameters"
		r.OverrideFieldNames["SubnetInitParameters"] = "TrafficManagerAzureEndpointSubnetInitParameters"
		r.OverrideFieldNames["SubnetObservation"] = "TrafficManagerAzureEndpointSubnetObservation"
	})

	p.AddResourceConfigurator("azurerm_traffic_manager_external_endpoint", func(r *config.Resource) {
		r.OverrideFieldNames["SubnetInitParameters"] = "TrafficManagerExternalEndpointSubnetInitParameters"
		r.OverrideFieldNames["SubnetParameters"] = "TrafficManagerExternalEndpointSubnetParameters"
		r.OverrideFieldNames["SubnetObservation"] = "TrafficManagerExternalEndpointSubnetObservation"
	})

	p.AddResourceConfigurator("azurerm_traffic_manager_nested_endpoint", func(r *config.Resource) {
		r.OverrideFieldNames["SubnetInitParameters"] = "TrafficManagerNestedEndpointSubnetInitParameters"
		r.OverrideFieldNames["SubnetParameters"] = "TrafficManagerNestedEndpointSubnetParameters"
		r.OverrideFieldNames["SubnetObservation"] = "TrafficManagerNestedEndpointSubnetObservation"
	})
}
