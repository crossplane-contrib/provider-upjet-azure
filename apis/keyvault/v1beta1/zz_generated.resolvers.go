// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	rconfig "github.com/upbound/provider-azure/apis/rconfig"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AccessPolicy) ResolveReferences( // ResolveReferences of this AccessPolicy.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Certificate.
func (mg *Certificate) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CertificateContacts.
func (mg *CertificateContacts) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CertificateIssuer.
func (mg *CertificateIssuer) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Key.
func (mg *Key) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
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

// ResolveReferences of this ManagedStorageAccount.
func (mg *ManagedStorageAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccountID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.StorageAccountIDRef,
			Selector:     mg.Spec.InitProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccountID")
	}
	mg.Spec.InitProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ManagedStorageAccountSASTokenDefinition.
func (mg *ManagedStorageAccountSASTokenDefinition) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "ManagedStorageAccount", "ManagedStorageAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedStorageAccountID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ManagedStorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.ManagedStorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagedStorageAccountID")
	}
	mg.Spec.ForProvider.ManagedStorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManagedStorageAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Secret.
func (mg *Secret) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vault.
func (mg *Vault) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
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
