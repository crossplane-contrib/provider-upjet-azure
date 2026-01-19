// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	_ "embed"

	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

const (
	resourcePrefix = "azure"
	modulePath     = "github.com/upbound/provider-azure/v2"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte

	// oldSingletonListAPIs is a newline-delimited list of Terraform resource
	// names with converted singleton list APIs with at least CRD API version
	// containing the old singleton list API. This is to prevent the API
	// conversion for the newly added resources whose CRD APIs will already
	// use embedded objects instead of the singleton lists and thus, will
	// not possess a CRD API version with the singleton list. Thus, for
	// the newly added resources (resources added after the singleton lists
	// have been converted), we do not need the CRD API conversion
	// functions that convert between singleton lists and embedded objects,
	// but we need only the Terraform conversion functions.
	// This list is immutable and represents the set of resources with the
	// already generated CRD API versions with now converted singleton lists.
	// Because new resources should never have singleton lists in their
	// generated APIs, there should be no need to add them to this list.
	// However, bugs might result in exceptions in the future.
	// Please see:
	// https://github.com/crossplane-contrib/provider-upjet-azure/pull/733
	// for more context on singleton list to embedded object conversions.
	//go:embed old-singleton-list-apis.txt
	oldSingletonListAPIs string
)

// These resources cannot be generated because of their suffixes colliding with
// kubebuilder-accepted type suffixes.
var skipList = []string{
	"azurerm_mssql_server_extended_auditing_policy",
	// group prefix collision
	"azurerm_dedicated_host_group",
	"azurerm_storage_disks_pool",
	"azurerm_storage_sync_group",
	"azurerm_storage_sync_cloud_endpoint", // depends on azurerm_storage_sync_group
	"azurerm_virtual_desktop_application_group",
	// associated with non-generated
	"azurerm_virtual_desktop_workspace_application_group_association",
	// generated name too long
	"azurerm_network_interface_application_gateway_backend_address_pool_association",
	"azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection",
	// deprecated
	"^azurerm_virtual_machine_scale_set$", // to match exactly azurerm_virtual_machine_scale_set not everything which starts from azurerm_virtual_machine_scale_set
	"azurerm_virtual_machine_configuration_policy_assignment",
	"azurerm_virtual_machine_scale_set_extension",
	"azurerm_sql_server",
	"azurerm_video_analyzer",
	"azurerm_video_analyzer_edge_module",
	"azurerm_sql_managed_database",
	"azurerm_sql_managed_instance",
	"azurerm_sql_managed_instance_active_directory_administrator",
	"azurerm_sql_managed_instance_failover_group",
	"azurerm_sql_active_directory_administrator",
	"azurerm_sql_database",
	"azurerm_sql_elasticpool",
	"azurerm_sql_firewall_rule",
	"azurerm_site_recovery_replicated_vm", // depends on azurerm_virtual_machine
	// irrelevant
	"azurerm_virtual_desktop_application",
	"azurerm_virtual_desktop_host_pool",
	"azurerm_virtual_desktop_workspace",
	"azurerm_virtual_desktop_scaling_plan",                // depends on azurerm_virtual_desktop_host_pool
	"azurerm_virtual_desktop_host_pool_registration_info", // depends on azurerm_virtual_desktop_host_pool
	// other upjet issues
	"azurerm_container_registry_task",
	// unsupported sensitive field type
	"azurerm_security_center_automation",
	"azurerm_data_factory_trigger_tumbling_window",
	"azurerm_storage_share_file",
	"azurerm_sql_virtual_network_rule",
	"azurerm_virtual_desktop_workspace",
	"azurerm_data_lake_analytics_account",
	"azurerm_log_analytics_storage_insights",
	"azurerm_virtual_hub_bgp_connection",
	"azurerm_automation_dsc_configuration",
	"azurerm_automation_dsc_nodeconfiguration", // depends on azurerm_automation_dsc_configuration
	"azurerm_monitor_log_profile",
	"azurerm_machine_learning_inference_cluster",
	"azurerm_sql_failover_group",
	"azurerm_logic_app_integration_account_certificate",
	"azurerm_container_group",
	// Azure are officially halting the preview of Azure Disk Pools, and it will not be made generally available.
	"azurerm_disk_pool_iscsi_target",
	// DEPRECATED: azurerm_service_plan should be used instead
	"azurerm_app_service_source_control_token",
	// DEPRECATED:
	"azurerm_cdn_frontdoor_route_disable_link_to_default_domain",
}

// workaround for the TF Azure v3.57.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// TerraformPluginSDKResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
