/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	iothubdeviceupdateaccount "github.com/upbound/provider-azure/internal/controller/deviceupdate/iothubdeviceupdateaccount"
	iothubdeviceupdateinstance "github.com/upbound/provider-azure/internal/controller/deviceupdate/iothubdeviceupdateinstance"
)

// Setup_deviceupdate creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_deviceupdate(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		iothubdeviceupdateaccount.Setup,
		iothubdeviceupdateinstance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
