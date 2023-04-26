/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	healthbot "github.com/upbound/provider-azure/internal/controller/healthbot/healthbot"
)

// Setup_healthbot creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_healthbot(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		healthbot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
