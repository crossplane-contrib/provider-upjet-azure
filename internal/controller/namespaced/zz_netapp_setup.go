// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-azure/v2/internal/controller/namespaced/netapp/account"
	pool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/netapp/pool"
	snapshot "github.com/upbound/provider-azure/v2/internal/controller/namespaced/netapp/snapshot"
	snapshotpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/netapp/snapshotpolicy"
	volume "github.com/upbound/provider-azure/v2/internal/controller/namespaced/netapp/volume"
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

// SetupGated_netapp creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_netapp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		pool.SetupGated,
		snapshot.SetupGated,
		snapshotpolicy.SetupGated,
		volume.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
