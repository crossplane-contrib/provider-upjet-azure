/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	federatedidentitycredential "github.com/upbound/provider-azure/internal/controller/managedidentity/federatedidentitycredential"
	userassignedidentity "github.com/upbound/provider-azure/internal/controller/managedidentity/userassignedidentity"
)

// Setup_managedidentity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managedidentity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		federatedidentitycredential.Setup,
		userassignedidentity.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
