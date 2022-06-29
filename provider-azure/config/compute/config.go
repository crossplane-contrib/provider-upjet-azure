/*
Copyright 2022 Upbound Inc.
*/

package compute

import (
	"github.com/upbound/official-providers/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_disk_encryption_set", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_linux_virtual_machine", func(r *config.Resource) {
		r.UseAsync = true
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
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_windows_virtual_machine", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_windows_virtual_machine_scale_set", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_availability_set", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_disk_access", func(r *config.Resource) {
		r.UseAsync = true
	})

	/* Note on testing:
	* - create a storage account
	* - upload a text file with *.vhd extension
	* - provide as blob URI
	 */
	p.AddResourceConfigurator("azurerm_image", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_managed_disk", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_orchestrated_virtual_machine_scale_set", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_proximity_placement_group", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_shared_image_gallery", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_snapshot", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_marketplace_agreement", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_dedicated_host", func(r *config.Resource) {
		r.UseAsync = true
	})
}
