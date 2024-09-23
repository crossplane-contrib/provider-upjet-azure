// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancedisk"
	backupinstancekubernetescluster "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancekubernetescluster"
	backupinstancepostgresql "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancepostgresql"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicydisk"
	backuppolicykubernetescluster "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicykubernetescluster"
	backuppolicypostgresql "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicypostgresql"
	backupvault "github.com/upbound/provider-azure/internal/controller/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/internal/controller/dataprotection/resourceguard"
)

// Setup_dataprotection creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataprotection(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupinstanceblobstorage.Setup,
		backupinstancedisk.Setup,
		backupinstancekubernetescluster.Setup,
		backupinstancepostgresql.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydisk.Setup,
		backuppolicykubernetescluster.Setup,
		backuppolicypostgresql.Setup,
		backupvault.Setup,
		resourceguard.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
