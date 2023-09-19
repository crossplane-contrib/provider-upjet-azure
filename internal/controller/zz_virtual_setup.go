/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	desktophostpool "github.com/upbound/provider-azure/internal/controller/virtual/desktophostpool"
)

// Setup_virtual creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_virtual(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		desktophostpool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
