// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/storage/account"
	accountlocaluser "github.com/upbound/provider-azure/internal/controller/storage/accountlocaluser"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/storage/container"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/storage/datalakegen2filesystem"
	datalakegen2path "github.com/upbound/provider-azure/internal/controller/storage/datalakegen2path"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/storage/objectreplication"
	queue "github.com/upbound/provider-azure/internal/controller/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/storage/share"
	sharedirectory "github.com/upbound/provider-azure/internal/controller/storage/sharedirectory"
	table "github.com/upbound/provider-azure/internal/controller/storage/table"
	tableentity "github.com/upbound/provider-azure/internal/controller/storage/tableentity"
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
