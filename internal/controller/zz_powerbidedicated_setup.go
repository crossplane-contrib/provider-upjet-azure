/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	powerbiembedded "github.com/upbound/provider-azure/internal/controller/powerbidedicated/powerbiembedded"
)

// Setup_powerbidedicated creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_powerbidedicated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		powerbiembedded.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
