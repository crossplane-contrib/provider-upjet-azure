/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

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
