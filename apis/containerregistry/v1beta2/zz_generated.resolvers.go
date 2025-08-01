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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Registry.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *Registry) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption.IdentityClientID),
				Extract:      resource.ExtractParamPath("client_id", true),
				Reference:    mg.Spec.ForProvider.Encryption.IdentityClientIDRef,
				Selector:     mg.Spec.ForProvider.Encryption.IdentityClientIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption.IdentityClientID")
		}
		mg.Spec.ForProvider.Encryption.IdentityClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption.IdentityClientIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.Identity != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Identity.IdentityIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.ForProvider.Identity.IdentityIdsRefs,
				Selector:      mg.Spec.ForProvider.Identity.IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Identity.IdentityIds")
		}
		mg.Spec.ForProvider.Identity.IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Identity.IdentityIdsRefs = mrsp.ResolvedReferences

	}
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

	if mg.Spec.InitProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption.IdentityClientID),
				Extract:      resource.ExtractParamPath("client_id", true),
				Reference:    mg.Spec.InitProvider.Encryption.IdentityClientIDRef,
				Selector:     mg.Spec.InitProvider.Encryption.IdentityClientIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption.IdentityClientID")
		}
		mg.Spec.InitProvider.Encryption.IdentityClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption.IdentityClientIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Identity != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Identity.IdentityIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.InitProvider.Identity.IdentityIdsRefs,
				Selector:      mg.Spec.InitProvider.Identity.IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Identity.IdentityIds")
		}
		mg.Spec.InitProvider.Identity.IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Identity.IdentityIdsRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this TokenPassword.
func (mg *TokenPassword) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("containerregistry.azure.upbound.io", "v1beta1", "Token", "TokenList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerRegistryTokenID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ContainerRegistryTokenIDRef,
			Selector:     mg.Spec.ForProvider.ContainerRegistryTokenIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerRegistryTokenID")
	}
	mg.Spec.ForProvider.ContainerRegistryTokenID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerRegistryTokenIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("containerregistry.azure.upbound.io", "v1beta1", "Token", "TokenList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ContainerRegistryTokenID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ContainerRegistryTokenIDRef,
			Selector:     mg.Spec.InitProvider.ContainerRegistryTokenIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ContainerRegistryTokenID")
	}
	mg.Spec.InitProvider.ContainerRegistryTokenID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ContainerRegistryTokenIDRef = rsp.ResolvedReference

	return nil
}
