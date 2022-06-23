/*
Copyright 2022 Upbound Inc.
*/

package compute

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_disk_encryption_set", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("azurerm_linux_virtual_machine", func(r *config.Resource) {
		r.UseAsync = true
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
