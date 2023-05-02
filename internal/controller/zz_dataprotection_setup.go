/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancedisk"
	backupinstancepostgresql "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancepostgresql"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicydisk"
	backupvault "github.com/upbound/provider-azure/internal/controller/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/internal/controller/dataprotection/resourceguard"
)

// Setup_dataprotection creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataprotection(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupinstanceblobstorage.Setup,
		backupinstancedisk.Setup,
		backupinstancepostgresql.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydisk.Setup,
		backupvault.Setup,
		resourceguard.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
