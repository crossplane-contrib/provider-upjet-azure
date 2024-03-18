// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	backupcontainerstorageaccount "github.com/upbound/provider-azure/internal/controller/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/upbound/provider-azure/internal/controller/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/upbound/provider-azure/internal/controller/recoveryservices/backuppolicyvm"
	backuppolicyvmworkload "github.com/upbound/provider-azure/internal/controller/recoveryservices/backuppolicyvmworkload"
	backupprotectedfileshare "github.com/upbound/provider-azure/internal/controller/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/upbound/provider-azure/internal/controller/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicationpolicy "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryreplicationpolicy"
	vault "github.com/upbound/provider-azure/internal/controller/recoveryservices/vault"
)

// Setup_recoveryservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_recoveryservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupcontainerstorageaccount.Setup,
		backuppolicyfileshare.Setup,
		backuppolicyvm.Setup,
		backuppolicyvmworkload.Setup,
		backupprotectedfileshare.Setup,
		backupprotectedvm.Setup,
		siterecoveryfabric.Setup,
		siterecoverynetworkmapping.Setup,
		siterecoveryprotectioncontainer.Setup,
		siterecoveryprotectioncontainermapping.Setup,
		siterecoveryreplicationpolicy.Setup,
		vault.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
