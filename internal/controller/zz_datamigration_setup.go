/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	databasemigrationproject "github.com/upbound/provider-azure/internal/controller/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/upbound/provider-azure/internal/controller/datamigration/databasemigrationservice"
)

// Setup_datamigration creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datamigration(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		databasemigrationproject.Setup,
		databasemigrationservice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
