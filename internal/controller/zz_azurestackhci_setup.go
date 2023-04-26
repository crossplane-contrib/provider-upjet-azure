/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/upbound/provider-azure/internal/controller/azurestackhci/cluster"
)

// Setup_azurestackhci creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_azurestackhci(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
