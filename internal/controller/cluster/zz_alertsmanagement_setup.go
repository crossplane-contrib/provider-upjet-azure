// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	monitoralertprocessingruleactiongroup "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoralertprocessingruleactiongroup"
	monitoralertprocessingrulesuppression "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoralertprocessingrulesuppression"
	monitorsmartdetectoralertrule "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitorsmartdetectoralertrule"
)

// Setup_alertsmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_alertsmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoralertprocessingruleactiongroup.Setup,
		monitoralertprocessingrulesuppression.Setup,
		monitorsmartdetectoralertrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_alertsmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_alertsmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoralertprocessingruleactiongroup.SetupGated,
		monitoralertprocessingrulesuppression.SetupGated,
		monitorsmartdetectoralertrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
