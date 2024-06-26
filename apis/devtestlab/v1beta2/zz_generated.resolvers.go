// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *GlobalVMShutdownSchedule) ResolveReferences( // ResolveReferences of this GlobalVMShutdownSchedule.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("compute.azure.upbound.io", "v1beta2", "LinuxVirtualMachine", "LinuxVirtualMachineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualMachineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VirtualMachineIDRef,
			Selector:     mg.Spec.ForProvider.VirtualMachineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualMachineID")
	}
	mg.Spec.ForProvider.VirtualMachineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualMachineIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.azure.upbound.io", "v1beta2", "LinuxVirtualMachine", "LinuxVirtualMachineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VirtualMachineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.VirtualMachineIDRef,
			Selector:     mg.Spec.InitProvider.VirtualMachineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VirtualMachineID")
	}
	mg.Spec.InitProvider.VirtualMachineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VirtualMachineIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LinuxVirtualMachine.
func (mg *LinuxVirtualMachine) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta1", "Lab", "LabList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LabNameRef,
			Selector:     mg.Spec.ForProvider.LabNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabName")
	}
	mg.Spec.ForProvider.LabName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabSubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LabSubnetNameRef,
			Selector:     mg.Spec.ForProvider.LabSubnetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabSubnetName")
	}
	mg.Spec.ForProvider.LabSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabSubnetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta2", "VirtualNetwork", "VirtualNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabVirtualNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.LabVirtualNetworkIDRef,
			Selector:     mg.Spec.ForProvider.LabVirtualNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabVirtualNetworkID")
	}
	mg.Spec.ForProvider.LabVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabVirtualNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta1", "Lab", "LabList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LabName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LabNameRef,
			Selector:     mg.Spec.InitProvider.LabNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LabName")
	}
	mg.Spec.InitProvider.LabName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LabNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LabSubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LabSubnetNameRef,
			Selector:     mg.Spec.InitProvider.LabSubnetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LabSubnetName")
	}
	mg.Spec.InitProvider.LabSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LabSubnetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta2", "VirtualNetwork", "VirtualNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LabVirtualNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.LabVirtualNetworkIDRef,
			Selector:     mg.Spec.InitProvider.LabVirtualNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LabVirtualNetworkID")
	}
	mg.Spec.InitProvider.LabVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LabVirtualNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Schedule.
func (mg *Schedule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta1", "Lab", "LabList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LabNameRef,
			Selector:     mg.Spec.ForProvider.LabNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabName")
	}
	mg.Spec.ForProvider.LabName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetwork.
func (mg *VirtualNetwork) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta1", "Lab", "LabList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LabNameRef,
			Selector:     mg.Spec.ForProvider.LabNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabName")
	}
	mg.Spec.ForProvider.LabName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta1", "Lab", "LabList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LabName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LabNameRef,
			Selector:     mg.Spec.InitProvider.LabNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LabName")
	}
	mg.Spec.InitProvider.LabName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LabNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WindowsVirtualMachine.
func (mg *WindowsVirtualMachine) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta1", "Lab", "LabList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LabNameRef,
			Selector:     mg.Spec.ForProvider.LabNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabName")
	}
	mg.Spec.ForProvider.LabName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabSubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LabSubnetNameRef,
			Selector:     mg.Spec.ForProvider.LabSubnetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabSubnetName")
	}
	mg.Spec.ForProvider.LabSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabSubnetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta2", "VirtualNetwork", "VirtualNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LabVirtualNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.LabVirtualNetworkIDRef,
			Selector:     mg.Spec.ForProvider.LabVirtualNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LabVirtualNetworkID")
	}
	mg.Spec.ForProvider.LabVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LabVirtualNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta1", "Lab", "LabList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LabName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LabNameRef,
			Selector:     mg.Spec.InitProvider.LabNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LabName")
	}
	mg.Spec.InitProvider.LabName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LabNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LabSubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LabSubnetNameRef,
			Selector:     mg.Spec.InitProvider.LabSubnetNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LabSubnetName")
	}
	mg.Spec.InitProvider.LabSubnetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LabSubnetNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devtestlab.azure.upbound.io", "v1beta2", "VirtualNetwork", "VirtualNetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LabVirtualNetworkID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.LabVirtualNetworkIDRef,
			Selector:     mg.Spec.InitProvider.LabVirtualNetworkIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LabVirtualNetworkID")
	}
	mg.Spec.InitProvider.LabVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LabVirtualNetworkIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
