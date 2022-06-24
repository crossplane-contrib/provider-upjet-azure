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

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/upbound/upjet/pkg/config"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/official-providers/provider-azure/config/apimanagement"
	"github.com/upbound/official-providers/provider-azure/config/base"
	"github.com/upbound/official-providers/provider-azure/config/common"
	"github.com/upbound/official-providers/provider-azure/config/compute"
	"github.com/upbound/official-providers/provider-azure/config/cosmosdb"
	"github.com/upbound/official-providers/provider-azure/config/datashare"
	"github.com/upbound/official-providers/provider-azure/config/devices"
	"github.com/upbound/official-providers/provider-azure/config/eventhub"
	"github.com/upbound/official-providers/provider-azure/config/ip"
	"github.com/upbound/official-providers/provider-azure/config/keyvault"
	"github.com/upbound/official-providers/provider-azure/config/kubernetes"
	"github.com/upbound/official-providers/provider-azure/config/loganalytics"
	"github.com/upbound/official-providers/provider-azure/config/logic"
	"github.com/upbound/official-providers/provider-azure/config/management"
	"github.com/upbound/official-providers/provider-azure/config/mariadb"
	"github.com/upbound/official-providers/provider-azure/config/monitor"
	"github.com/upbound/official-providers/provider-azure/config/network"
	"github.com/upbound/official-providers/provider-azure/config/notificationhubs"
	"github.com/upbound/official-providers/provider-azure/config/postgresql"
	"github.com/upbound/official-providers/provider-azure/config/redis"
	"github.com/upbound/official-providers/provider-azure/config/resource"
	"github.com/upbound/official-providers/provider-azure/config/security"
	"github.com/upbound/official-providers/provider-azure/config/sql"
	"github.com/upbound/official-providers/provider-azure/config/storage"
	"github.com/upbound/official-providers/provider-azure/config/storagesync"
	"github.com/upbound/official-providers/provider-azure/config/subnet"
)

const (
	resourcePrefix = "azure"
	modulePath     = "github.com/upbound/official-providers/provider-azure"
)

//go:embed schema.json
var providerSchema string

var includedResources = []string{
	// "azurerm_.+",
	"azurerm_virtual_.+",
	"azurerm_kubernetes_.+",
	"azurerm_postgresql_.+",
	"azurerm_cosmosdb_.+",
	"azurerm_redis_.+",
	"azurerm_resource_group",
	"azurerm_subnet",
	"azurerm_storage_account$",
	"azurerm_storage_container$",
	"azurerm_storage_blob$",
	"azurerm_sql_server",
	"azurerm_mssql_server$",
	"azurerm_mssql_server_transparent_data_encryption$",
	"azurerm_lb.*",
	"azurerm_local_network_gateway$",
	"azurerm_log_analytics_workspace",
	"azurerm_iothub.*",
	"azurerm_monitor_metric_alert",
	"azurerm_network_connection_monitor$",
	"azurerm_network_ddos_protection_plan$",
	"azurerm_network_watcher$",
	"azurerm_application_security_group$",
	"azurerm_network_interface_application_security_group_association$",
	"azurerm_nat_gateway.*",
	"azurerm_key_vault.*",
	"azurerm_eventhub_namespace$",
	"azurerm_eventhub$",
	"azurerm_eventhub_consumer_group$",
	"azurerm_eventhub_authorization_rule$",
	"azurerm_network_interface$",
	"azurerm_mariadb_.+",
	"azurerm_public_ip.*",
	"azurerm_disk_encryption_set$",
	"azurerm_(windows|linux)_virtual_.+",
}

// These resources cannot be generated because of their suffixes colliding with
// kubebuilder-accepted type suffixes.
var skipList = []string{
	"azurerm_mssql_server_extended_auditing_policy",
	// group prefix collision
	"azurerm_api_management_group",
	"azurerm_api_management_product_group",
	"azurerm_dedicated_host_group",
	"azurerm_storage_sync_group",
	"azurerm_virtual_desktop_application_group",
	// associated with non-generated
	"azurerm_virtual_desktop_workspace_application_group_association",
	// generated name too long
	"azurerm_network_interface_application_gateway_backend_address_pool_association",
	"azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection",
	// cannot generate a unique name
	"azurerm_route",
	"azurerm_network_security_rule",
	// deprecated
	"azurerm_virtual_machine_scale_set",
	"azurerm_virtual_machine_configuration_policy_assignment",
	"azurerm_virtual_machine",
	"azurerm_virtual_machine_extension",
	"azurerm_virtual_machine_data_disk_attachment",
	"azurerm_virtual_machine_scale_set_extension",
	"azurerm_sql_server",
	// irrelevant
	"azurerm_virtual_desktop_application",
	"azurerm_virtual_desktop_host_pool",
	"azurerm_virtual_desktop_workspace",
	// other upjet issues
	"azurerm_container_registry_task",
	"azurerm_dashboard",
	// doc not found in Terraform Azurerm provider
	"azurerm_virtual_network_dns_servers",
	// unsupported sensitive field type
	"azurerm_security_center_automation",
	"azurerm_data_factory_trigger_tumbling_window",
	"azurerm_storage_share_file",
	"azurerm_sql_virtual_network_rule",
	"azurerm_virtual_desktop_workspace",
	"azurerm_data_lake_analytics_account",
	"azurerm_virtual_hub",
	"azurerm_log_analytics_storage_insights",
	"azurerm_virtual_hub_bgp_connection",
	"azurerm_automation_dsc_configuration",
	"azurerm_monitor_log_profile",
	"azurerm_machine_learning_inference_cluster",
	"azurerm_sql_failover_group",
	"azurerm_logic_app_integration_account_certificate",
	"azurerm_container_group",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, "config/provider-metadata.yaml",
		tjconfig.WithShortName("azure"),
		tjconfig.WithRootGroup("azure.upbound.io"),
		tjconfig.WithIncludeList(includedResources),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithDefaultResourceFn(defaultResource(defaultVersion(), externalNameConfig(), groupOverrides())),
	)
	// external-name configuration for all resources
	for name := range pc.Resources {
		// external-name configuration
		pc.AddResourceConfigurator(name, func(r *tjconfig.Resource) {
			r.ExternalName = tjconfig.IdentifierFromProvider
		})
	}

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		network.Configure,
		ip.Configure,
		management.Configure,
		redis.Configure,
		resource.Configure,
		kubernetes.Configure,
		postgresql.Configure,
		cosmosdb.Configure,
		sql.Configure,
		subnet.Configure,
		storage.Configure,
		loganalytics.Configure,
		devices.Configure,
		monitor.Configure,
		apimanagement.Configure,
		logic.Configure,
		security.Configure,
		base.Configure,
		datashare.Configure,
		notificationhubs.Configure,
		storagesync.Configure,
		keyvault.Configure,
		eventhub.Configure,
		mariadb.Configure,
		compute.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()

	// This function runs after the special configurations were applied. However, if some references were configured in
	// the configuration files, the reference generator does not override them.
	for _, r := range pc.Resources {
		if err := common.AddCommonReferences(r); err != nil {
			panic(err)
		}
	}

	return pc
}

func defaultResource(opts ...tjconfig.ResourceOption) tjconfig.DefaultResourceFn {
	return func(name string, terraformResource *schema.Resource, orgOpts ...tjconfig.ResourceOption) *tjconfig.Resource {
		return tjconfig.DefaultResource(name, terraformResource, append(orgOpts, opts...)...)
	}
}
