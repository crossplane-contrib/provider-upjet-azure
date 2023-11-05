// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/netapp/account"
	pool "github.com/upbound/provider-azure/internal/controller/netapp/pool"
	snapshot "github.com/upbound/provider-azure/internal/controller/netapp/snapshot"
	snapshotpolicy "github.com/upbound/provider-azure/internal/controller/netapp/snapshotpolicy"
	volume "github.com/upbound/provider-azure/internal/controller/netapp/volume"
)

// Setup_netapp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_netapp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		pool.Setup,
		snapshot.Setup,
		snapshotpolicy.Setup,
		volume.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
