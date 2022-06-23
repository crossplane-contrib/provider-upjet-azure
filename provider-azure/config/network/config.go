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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
	"github.com/upbound/official-providers/provider-azure/config/common"
)

// getParameterBasedIDFn returns a GetIDFn that returns load balancer based ID FQDNs
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/parameter/parameter1/resourceType/resourceType1
// TODO(ytsarev): deperecate it when similar upjet function is available https://github.com/upbound/upjet/pull/22
func getParameterBasedIDFn(parameter string, resourceType string) config.GetIDFn {
	return func(_ context.Context, name string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		parameterID, ok := parameters[parameter]
		if !ok {
			return "", errors.Errorf(common.ErrFmtNoAttribute, parameter)
		}
		parameterIDStr, ok := parameterID.(string)
		if !ok {
			return "", errors.Errorf(common.ErrFmtUnexpectedType, parameter)
		}
		return fmt.Sprintf("%s/%s/%s", parameterIDStr, resourceType, name), nil
	}
}

// Configure configures virtual group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_ip_group", func(r *config.Resource) {
		r.Kind = "IPGroup"
	})
	p.AddResourceConfigurator("azurerm_network_interface", func(r *config.Resource) {
		r.Kind = "NetworkInterface"
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "networkInterfaces", "name")
	})

	p.AddResourceConfigurator("azurerm_lb", func(r *config.Resource) {
		r.Kind = "LoadBalancer"

		r.References["frontend_ip_configuration.public_ip_address_id"] = config.Reference{
			Type:      "PublicIP",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "loadBalancers", "name")
	})

	p.AddResourceConfigurator("azurerm_lb_backend_address_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerBackendAddressPool"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("loadbalancer_id", "backendAddressPools")
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
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/loadBalancer1/backendAddressPools/backendAddressPool1/addresses/address1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("backend_address_pool_id", "addresses")
	})

	p.AddResourceConfigurator("azurerm_lb_nat_pool", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatPool"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatPools/pool1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("loadbalancer_id", "inboundNatPools")
	})

	p.AddResourceConfigurator("azurerm_lb_nat_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerNatRule"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("loadbalancer_id", "inboundNatRules")
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
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/outboundRules/rule1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("loadbalancer_id", "outboundRules")
	})

	p.AddResourceConfigurator("azurerm_lb_probe", func(r *config.Resource) {
		r.Kind = "LoadBalancerProbe"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/probes/probe1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("loadbalancer_id", "probes")
	})

	p.AddResourceConfigurator("azurerm_lb_rule", func(r *config.Resource) {
		r.Kind = "LoadBalancerRule"
		r.References["loadbalancer_id"] = config.Reference{
			Type:      "LoadBalancer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/rule1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("loadbalancer_id", "loadBalancingRules")
	})

	p.AddResourceConfigurator("azurerm_local_network_gateway", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/localNetworkGateways/lng1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "localNetworkGateways", "name")
	})

	p.AddResourceConfigurator("azurerm_nat_gateway", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/natGateways/gateway1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "natGateways", "name")
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
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/natGateways/gateway1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPAddresses/myPublicIpAddress1
		r.ExternalName = config.IdentifierFromProvider
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
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/natGateways/gateway1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPPrefixes/myPublicIpPrefix1
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("azurerm_network_watcher", func(r *config.Resource) {
		r.Kind = "Watcher"
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"networkWatchers", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_network_connection_monitor", func(r *config.Resource) {
		r.Kind = "ConnectionMonitor"
		r.References["network_watcher_id"] = config.Reference{
			Type:      "Watcher",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/networkWatchers/watcher1/connectionMonitors/connectionMonitor1
		r.ExternalName.GetIDFn = getParameterBasedIDFn("network_watcher_id", "connectionMonitors")
	})

	p.AddResourceConfigurator("azurerm_network_ddos_protection_plan", func(r *config.Resource) {
		r.Kind = "DDoSProtectionPlan"
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/ddosProtectionPlans/testddospplan
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"ddosProtectionPlans", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_application_security_group", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationSecurityGroups/securitygroup1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"applicationSecurityGroups", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_network_security_group", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"networkSecurityGroups", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_network_security_rule", func(r *config.Resource) {
		r.References["network_security_group_name"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup/securityRules/rule1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"networkSecurityGroups", "network_security_group_name",
			"securityRules", "name",
		)
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
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationSecurityGroups/securityGroup1
		r.ExternalName = config.IdentifierFromProvider
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
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1
		r.ExternalName = config.IdentifierFromProvider
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
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
		r.ExternalName = config.IdentifierFromProvider
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
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("azurerm_virtual_network", func(r *config.Resource) {
		r.Kind = "VirtualNetwork"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"subnet"},
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualNetworks", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_virtual_network_gateway", func(r *config.Resource) {
		r.Kind = "VirtualNetworkGateway"
		r.References["ip_configuration.subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/virtualNetworkGateways/myGateway1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualNetworkGateways", "name",
		)
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
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/virtualNetworkPeerings/myvnet1peering
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualNetworks", "virtual_network_name",
			"virtualNetworkPeerings", "name",
		)
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
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/connections/myConnection1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"connections", "name",
		)
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
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualWans/testvwan
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network",
			"virtualWans", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_frontdoor", func(r *config.Resource) {
		r.Kind = "FrontDoor"
	})

	p.AddResourceConfigurator("azurerm_network_packet_capture", func(r *config.Resource) {
		r.Kind = "NetworkPacketCapture"
	})

	p.AddResourceConfigurator("azurerm_public_ip", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPAddresses/myPublicIpAddress1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "publicIPAddresses", "name")
	})

	p.AddResourceConfigurator("azurerm_public_ip_prefix", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPPrefixes/myPublicIpPrefix1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "publicIPPrefixes", "name")
	})
}
