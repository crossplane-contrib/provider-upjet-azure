/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/maps/account"
	creator "github.com/upbound/provider-azure/internal/controller/maps/creator"
)

// Setup_maps creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_maps(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		creator.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
