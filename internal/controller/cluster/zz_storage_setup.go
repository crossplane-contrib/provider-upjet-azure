// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/cluster/storage/account"
	accountlocaluser "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountlocaluser"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/cluster/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/cluster/storage/container"
	containerimmutabilitypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/containerimmutabilitypolicy"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2filesystem"
	datalakegen2path "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2path"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/cluster/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/cluster/storage/objectreplication"
	queue "github.com/upbound/provider-azure/internal/controller/cluster/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/cluster/storage/share"
	sharedirectory "github.com/upbound/provider-azure/internal/controller/cluster/storage/sharedirectory"
	table "github.com/upbound/provider-azure/internal/controller/cluster/storage/table"
	tableentity "github.com/upbound/provider-azure/internal/controller/cluster/storage/tableentity"
)

// Setup_storage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		accountlocaluser.Setup,
		accountnetworkrules.Setup,
		blob.Setup,
		blobinventorypolicy.Setup,
		container.Setup,
		containerimmutabilitypolicy.Setup,
		datalakegen2filesystem.Setup,
		datalakegen2path.Setup,
		encryptionscope.Setup,
		managementpolicy.Setup,
		objectreplication.Setup,
		queue.Setup,
		share.Setup,
		sharedirectory.Setup,
		table.Setup,
		tableentity.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_storage creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_storage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		accountlocaluser.SetupGated,
		accountnetworkrules.SetupGated,
		blob.SetupGated,
		blobinventorypolicy.SetupGated,
		container.SetupGated,
		containerimmutabilitypolicy.SetupGated,
		datalakegen2filesystem.SetupGated,
		datalakegen2path.SetupGated,
		encryptionscope.SetupGated,
		managementpolicy.SetupGated,
		objectreplication.SetupGated,
		queue.SetupGated,
		share.SetupGated,
		sharedirectory.SetupGated,
		table.SetupGated,
		tableentity.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
