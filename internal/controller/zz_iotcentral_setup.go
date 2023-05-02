/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/upbound/provider-azure/internal/controller/iotcentral/application"
	applicationnetworkruleset "github.com/upbound/provider-azure/internal/controller/iotcentral/applicationnetworkruleset"
)

// Setup_iotcentral creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iotcentral(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationnetworkruleset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
