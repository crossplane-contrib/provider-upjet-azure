/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	managedapplicationdefinition "github.com/upbound/provider-azure/internal/controller/solutions/managedapplicationdefinition"
)

// Setup_solutions creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_solutions(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managedapplicationdefinition.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
