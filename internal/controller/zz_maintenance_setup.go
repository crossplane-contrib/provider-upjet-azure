// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	maintenanceassignmentdedicatedhost "github.com/upbound/provider-azure/internal/controller/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceconfiguration "github.com/upbound/provider-azure/internal/controller/maintenance/maintenanceconfiguration"
)

// Setup_maintenance creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_maintenance(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		maintenanceassignmentdedicatedhost.Setup,
		maintenanceassignmentvirtualmachine.Setup,
		maintenanceconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
