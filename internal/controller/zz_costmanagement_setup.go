/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	costanomalyalert "github.com/upbound/provider-azure/internal/controller/costmanagement/costanomalyalert"
	resourcegroupcostmanagementexport "github.com/upbound/provider-azure/internal/controller/costmanagement/resourcegroupcostmanagementexport"
	subscriptioncostmanagementexport "github.com/upbound/provider-azure/internal/controller/costmanagement/subscriptioncostmanagementexport"
)

// Setup_costmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_costmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		costanomalyalert.Setup,
		resourcegroupcostmanagementexport.Setup,
		subscriptioncostmanagementexport.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
