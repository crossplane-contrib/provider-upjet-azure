// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
