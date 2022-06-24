/*
Copyright 2022 Upbound Inc.
*/

package compute

import (
	"context"
	"fmt"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/config/common"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_disk_encryption_set", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/group1/providers/Microsoft.Compute/diskEncryptionSets/encryptionSet1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"diskEncryptionSets", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_linux_virtual_machine", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachines", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_linux_virtual_machine_scale_set", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachineScaleSets", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_windows_virtual_machine", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachines", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_windows_virtual_machine_scale_set", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachineScaleSets", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_availability_set", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"availabilitySets", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_disk_access", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/diskAccesses/diskAccess1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"diskAccesses", "name",
		)
	})

	/* Note on testing:
	* - create a storage account
	* - upload a text file with *.vhd extension
	* - provide as blob URI
	 */
	p.AddResourceConfigurator("azurerm_image", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/images/image1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"images", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_managed_disk", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/disks/manageddisk1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"disks", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_orchestrated_virtual_machine_scale_set", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachineScaleSets", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_proximity_placement_group", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/proximityPlacementGroups/example-ppg
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"proximityPlacementGroups", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_shared_image_gallery", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"galleries", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_snapshot", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/snapshots/snapshot1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"snapshots", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_marketplace_agreement", func(r *config.Resource) {
		r.UseAsync = true
		// /subscriptions/000-000/providers/Microsoft.MarketplaceOrdering/agreements/publisher1/offers/offer1/plans/plan1
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("azurerm_dedicated_host", func(r *config.Resource) {
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/hostGroups/group1/hosts/host1
		r.ExternalName.GetIDFn = func(_ context.Context, name string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			return fmt.Sprintf("%s/hosts/%s", parameters["dedicated_host_group_id"], name), nil
		}
	})

}
