// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package compute

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-azure/apis/cluster/rconfig"
)

const group = "compute"

// Configure configures compute group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_linux_virtual_machine", func(r *config.Resource) {
		r.References["network_interface_ids"] = config.Reference{
			TerraformName: "azurerm_network_interface",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"platform_fault_domain",
				"virtual_machine_scale_set_id"},
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "termination_notification.#")
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator("azurerm_linux_virtual_machine_scale_set", func(r *config.Resource) {
		// In version 3.38.0 the `scale_in_policy` parameter was removed, and replaced by `scale_in`
		config.MoveToStatus(r.TerraformResource, "scale_in_policy")
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "termination_notification.#")
				delete(diff.Attributes, "terminate_notification.#")
				delete(diff.Attributes, "gallery_application.#")
				delete(diff.Attributes, "gallery_applications.#")
				delete(diff.Attributes, "spot_restore.#")
			}
			return diff, nil
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"instances"},
		}
	})
	p.AddResourceConfigurator("azurerm_windows_virtual_machine", func(r *config.Resource) {
		r.References["network_interface_ids"] = config.Reference{
			TerraformName: "azurerm_network_interface",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"platform_fault_domain",
				"virtual_machine_scale_set_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_windows_virtual_machine_scale_set", func(r *config.Resource) {
		// In version 3.21.0 the `scale_in_policy` parameter was removed, and replaced by `scale_in`
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"scale_in_policy"},
		}
		r.MetaResource.ArgumentDocs["scale_in_policy"] = "Deprecated: scaleInPolicy will be removed in favour of the scaleIn code block."
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "termination_notification.#")
				delete(diff.Attributes, "terminate_notification.#")
				delete(diff.Attributes, "gallery_application.#")
				delete(diff.Attributes, "gallery_applications.#")
				delete(diff.Attributes, "spot_restore.#")
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator("azurerm_virtual_machine_extension", func(r *config.Resource) {
		r.Kind = "VirtualMachineExtension"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azurerm_virtual_machine_run_command", func(r *config.Resource) {
		r.References["virtual_machine_id"] = config.Reference{
			TerraformName: "azurerm_linux_virtual_machine",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_virtual_machine_data_disk_attachment", func(r *config.Resource) {
		r.References["virtual_machine_id"] = config.Reference{
			TerraformName: "azurerm_linux_virtual_machine",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_orchestrated_virtual_machine_scale_set", func(r *config.Resource) {
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil && diff.Attributes != nil {
				delete(diff.Attributes, "termination_notification.#")
			}
			return diff, nil
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"instances"},
		}
	})
	p.AddResourceConfigurator("azurerm_virtual_machine_scale_set_standby_pool", func(r *config.Resource) {
		r.ShortGroup = group
		r.References["attached_virtual_machine_scale_set_id"] = config.Reference{
			TerraformName: "azurerm_orchestrated_virtual_machine_scale_set",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
	/* Note on testing:
	* - create a storage account
	* - upload a text file with *.vhd extension
	* - provide as blob URI
	 */
}
