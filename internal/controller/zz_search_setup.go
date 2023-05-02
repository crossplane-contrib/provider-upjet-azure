/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	service "github.com/upbound/provider-azure/internal/controller/search/service"
	sharedprivatelinkservice "github.com/upbound/provider-azure/internal/controller/search/sharedprivatelinkservice"
)

// Setup_search creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_search(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		service.Setup,
		sharedprivatelinkservice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
