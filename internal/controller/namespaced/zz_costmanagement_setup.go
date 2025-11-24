// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	costanomalyalert "github.com/upbound/provider-azure/v2/internal/controller/namespaced/costmanagement/costanomalyalert"
	resourcegroupcostmanagementexport "github.com/upbound/provider-azure/v2/internal/controller/namespaced/costmanagement/resourcegroupcostmanagementexport"
	subscriptioncostmanagementexport "github.com/upbound/provider-azure/v2/internal/controller/namespaced/costmanagement/subscriptioncostmanagementexport"
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

// SetupGated_costmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_costmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		costanomalyalert.SetupGated,
		resourcegroupcostmanagementexport.SetupGated,
		subscriptioncostmanagementexport.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
