// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	_ "embed"

	"github.com/crossplane/upjet/pkg/config"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/xpprovider"
	"github.com/pkg/errors"

	"github.com/upbound/provider-azure/config/common"
	"github.com/upbound/provider-azure/hack"
)

const (
	resourcePrefix = "azure"
	modulePath     = "github.com/upbound/provider-azure"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

// These resources cannot be generated because of their suffixes colliding with
// kubebuilder-accepted type suffixes.
var skipList = []string{
	"azurerm_mssql_server_extended_auditing_policy",
	// group prefix collision
	"azurerm_api_management_group",
	"azurerm_api_management_product_group",
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
	"azurerm_virtual_machine_scale_set",
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

// GetProvider returns provider configuration
func GetProvider(ctx context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p, err = xpprovider.GetProviderSchema(ctx)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		ujconfig.WithShortName("azure"),
		ujconfig.WithRootGroup("azure.upbound.io"),
		ujconfig.WithIncludeList(CLIReconciledResourceList()),
		ujconfig.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		ujconfig.WithSkipList(skipList),
		ujconfig.WithDefaultResourceOptions(ResourceConfigurator()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithTerraformProvider(p),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	bumpVersionsWithEmbeddedLists(pc)
	// "azure" group contains resources that actually do not have a specific
	// group, e.g. ResourceGroup with APIVersion "azure.upbound.io/v1beta1".
	// We need to include the controllers for this group into the base packages
	// list to get their controllers packaged together with the config package
	// controllers (provider family config package).
	for _, c := range []string{"internal/controller/azure/resourcegroup", "internal/controller/azure/resourceproviderregistration", "internal/controller/azure/subscription"} {
		pc.BasePackages.ControllerMap[c] = "config"
	}

	// API group overrides from Terraform import statements
	for _, r := range pc.Resources {
		groupKindOverride(r)
	}

	// add custom config functions
	for _, configure := range ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()

	// This function runs after the special configurations were applied. However, if some references were configured in
	// the configuration files, the reference generator does not override them.
	for _, r := range pc.Resources {
		delete(r.References, "resource_group_name")
		if err := common.AddCommonReferences(r); err != nil {
			panic(err)
		}
	}
	return pc, nil
}

func bumpVersionsWithEmbeddedLists(pc *ujconfig.Provider) {
	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}
		r.Version = "v1beta2"
		// we would like to set the storage version to v1beta1 to facilitate
		// downgrades.
		r.SetCRDStorageVersion("v1beta1")
		r.Conversions = []conversion.Conversion{
			conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, []string{"spec.forProvider", "spec.initProvider", "status.atProvider"}, r.CRDListConversionPaths()...),
			conversion.NewSingletonListConversion("v1beta1", "v1beta2", r.CRDListConversionPaths(), conversion.ToEmbeddedObject),
			conversion.NewSingletonListConversion("v1beta2", "v1beta1", r.CRDListConversionPaths(), conversion.ToSingletonList)}
		pc.Resources[name] = r
	}
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

// Configure configures the specified Provider.
type Configure func(provider *config.Provider)

// Configurator is a registry for provider Configs.
type Configurator []Configure

// AddConfig adds a Config to the Configurator registry.
func (c *Configurator) AddConfig(conf Configure) {
	*c = append(*c, conf)
}

// ProviderConfiguration is a global registry to be used by
// the resource providers to register their Config functions.
var ProviderConfiguration = Configurator{}
