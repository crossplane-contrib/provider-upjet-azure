/*
Copyright 2022 Upbound Inc.
*/

package compute

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_linux_virtual_machine", func(r *config.Resource) {
		r.References["network_interface_ids"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/network/v1beta1.NetworkInterface",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"platform_fault_domain",
				"virtual_machine_scale_set_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_linux_virtual_machine_scale_set", func(r *config.Resource) {
		// In version 3.38.0 the `scale_in_policy` parameter was removed, and replaced by `scale_in`
		config.MoveToStatus(r.TerraformResource, "scale_in_policy")
	})
	p.AddResourceConfigurator("azurerm_windows_virtual_machine", func(r *config.Resource) {
		r.References["network_interface_ids"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/network/v1beta1.NetworkInterface",
			Extractor: rconfig.ExtractResourceIDFuncPath,
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
	})
	/* Note on testing:
	* - create a storage account
	* - upload a text file with *.vhd extension
	* - provide as blob URI
	 */
}
