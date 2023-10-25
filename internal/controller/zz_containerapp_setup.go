/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	containerapp "github.com/upbound/provider-azure/internal/controller/containerapp/containerapp"
	environment "github.com/upbound/provider-azure/internal/controller/containerapp/environment"
)

// Setup_containerapp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerapp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerapp.Setup,
		environment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
