/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	resourcepolicyremediation "github.com/upbound/provider-azure/internal/controller/policyinsights/resourcepolicyremediation"
	subscriptionpolicyremediation "github.com/upbound/provider-azure/internal/controller/policyinsights/subscriptionpolicyremediation"
)

// Setup_policyinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_policyinsights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcepolicyremediation.Setup,
		subscriptionpolicyremediation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
