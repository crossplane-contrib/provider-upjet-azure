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
	"github.com/upbound/official-providers/provider-azure/config/common"
)

const groupNetwork = "network"

// Configure configures virtual group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_network_interface", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.Kind = "NetworkInterface"
		r.ShortGroup = groupNetwork
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "networkInterfaces", "name")
	})

	p.AddResourceConfigurator("azurerm_lb", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.Kind = "LoadBalancer"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "loadBalancers", "name")
	})

	p.AddResourceConfigurator("azurerm_virtual_network", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.Kind = "VirtualNetwork"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
		}
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
		r.Version = common.VersionV1Alpha2
		r.Kind = "VirtualNetworkGateway"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"ip_configuration.subnet_id": config.Reference{
				Type:      "Subnet",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
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
		r.Version = common.VersionV1Alpha2
		r.Kind = "VirtualNetworkPeering"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"virtual_network_name": config.Reference{
				Type: "VirtualNetwork",
			},
			"remote_virtual_network_id": config.Reference{
				Type:      "VirtualNetwork",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
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
		r.Version = common.VersionV1Alpha2
		r.Kind = "VirtualNetworkGatewayConnection"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"virtual_network_gateway_id": config.Reference{
				Type:      "VirtualNetworkGateway",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"peer_virtual_network_gateway_id": config.Reference{
				Type:      "VirtualNetworkGateway",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
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
		r.Version = common.VersionV1Alpha2
		r.Kind = "VirtualWAN"
		r.ShortGroup = groupNetwork
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
		}
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
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPAddresses/myPublicIpAddress1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Network", "publicIPAddresses", "name")
	})
}
