// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	rconfig "github.com/upbound/provider-azure/apis/rconfig"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Account.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *Account) ResolveReferences(ctx context.Context, c client.Reader) error {
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

// ResolveReferences of this DataSetBlobStorage.
func (mg *DataSetBlobStorage) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ContainerNameRef,
			Selector:     mg.Spec.ForProvider.ContainerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ContainerName")
	}
	mg.Spec.ForProvider.ContainerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ContainerNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("datashare.azure.upbound.io", "v1beta1", "DataShare", "DataShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataShareID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DataShareIDRef,
			Selector:     mg.Spec.ForProvider.DataShareIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DataShareID")
	}
	mg.Spec.ForProvider.DataShareID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DataShareIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageAccount); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccount[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.StorageAccount[i3].NameRef,
				Selector:     mg.Spec.ForProvider.StorageAccount[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccount[i3].Name")
		}
		mg.Spec.ForProvider.StorageAccount[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageAccount[i3].NameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageAccount); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupNameRef,
				Selector:     mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupName")
		}
		mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageAccount[i3].ResourceGroupNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ContainerName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ContainerNameRef,
			Selector:     mg.Spec.InitProvider.ContainerNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ContainerName")
	}
	mg.Spec.InitProvider.ContainerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ContainerNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageAccount); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccount[i3].Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.StorageAccount[i3].NameRef,
				Selector:     mg.Spec.InitProvider.StorageAccount[i3].NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccount[i3].Name")
		}
		mg.Spec.InitProvider.StorageAccount[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageAccount[i3].NameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageAccount); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccount[i3].ResourceGroupName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.StorageAccount[i3].ResourceGroupNameRef,
				Selector:     mg.Spec.InitProvider.StorageAccount[i3].ResourceGroupNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccount[i3].ResourceGroupName")
		}
		mg.Spec.InitProvider.StorageAccount[i3].ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageAccount[i3].ResourceGroupNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DataSetDataLakeGen2.
func (mg *DataSetDataLakeGen2) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "DataLakeGen2FileSystem", "DataLakeGen2FileSystemList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.FileSystemNameRef,
			Selector:     mg.Spec.ForProvider.FileSystemNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemName")
	}
	mg.Spec.ForProvider.FileSystemName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FileSystemNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("datashare.azure.upbound.io", "v1beta2", "DataShare", "DataShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ShareID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ShareIDRef,
			Selector:     mg.Spec.ForProvider.ShareIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ShareID")
	}
	mg.Spec.ForProvider.ShareID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ShareIDRef = rsp.ResolvedReference
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
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "DataLakeGen2FileSystem", "DataLakeGen2FileSystemList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FileSystemName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.FileSystemNameRef,
			Selector:     mg.Spec.InitProvider.FileSystemNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FileSystemName")
	}
	mg.Spec.InitProvider.FileSystemName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FileSystemNameRef = rsp.ResolvedReference
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

// ResolveReferences of this DataSetKustoCluster.
func (mg *DataSetKustoCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KustoClusterID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KustoClusterIDRef,
			Selector:     mg.Spec.ForProvider.KustoClusterIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KustoClusterID")
	}
	mg.Spec.ForProvider.KustoClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KustoClusterIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("datashare.azure.upbound.io", "v1beta2", "DataShare", "DataShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ShareID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ShareIDRef,
			Selector:     mg.Spec.ForProvider.ShareIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ShareID")
	}
	mg.Spec.ForProvider.ShareID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ShareIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KustoClusterID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KustoClusterIDRef,
			Selector:     mg.Spec.InitProvider.KustoClusterIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KustoClusterID")
	}
	mg.Spec.InitProvider.KustoClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KustoClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DataSetKustoDatabase.
func (mg *DataSetKustoDatabase) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta1", "Database", "DatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KustoDatabaseID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KustoDatabaseIDRef,
			Selector:     mg.Spec.ForProvider.KustoDatabaseIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KustoDatabaseID")
	}
	mg.Spec.ForProvider.KustoDatabaseID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KustoDatabaseIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("datashare.azure.upbound.io", "v1beta2", "DataShare", "DataShareList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ShareID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ShareIDRef,
			Selector:     mg.Spec.ForProvider.ShareIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ShareID")
	}
	mg.Spec.ForProvider.ShareID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ShareIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kusto.azure.upbound.io", "v1beta1", "Database", "DatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KustoDatabaseID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KustoDatabaseIDRef,
			Selector:     mg.Spec.InitProvider.KustoDatabaseIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KustoDatabaseID")
	}
	mg.Spec.InitProvider.KustoDatabaseID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KustoDatabaseIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DataShare.
func (mg *DataShare) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("datashare.azure.upbound.io", "v1beta1", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AccountIDRef,
			Selector:     mg.Spec.ForProvider.AccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	return nil
}
