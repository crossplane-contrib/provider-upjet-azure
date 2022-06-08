/*
Copyright 2022 Upbound Inc.
*/

package compute

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/config/common"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_disk_encryption_set", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/group1/providers/Microsoft.Compute/diskEncryptionSets/encryptionSet1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"diskEncryptionSets", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_linux_virtual_machine", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachines", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_linux_virtual_machine_scale_set", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachineScaleSets", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_windows_virtual_machine", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/machine1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachines", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_windows_virtual_machine_scale_set", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Compute",
			"virtualMachineScaleSets", "name",
		)
	})

}
