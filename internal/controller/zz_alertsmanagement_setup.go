// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	monitoractionruleactiongroup "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoractionrulesuppression"
	monitoralertprocessingruleactiongroup "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoralertprocessingruleactiongroup"
	monitoralertprocessingrulesuppression "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoralertprocessingrulesuppression"
	monitorsmartdetectoralertrule "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitorsmartdetectoralertrule"
)

// Setup_alertsmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_alertsmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoractionruleactiongroup.Setup,
		monitoractionrulesuppression.Setup,
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
