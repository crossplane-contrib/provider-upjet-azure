/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	resourcegroup "github.com/upbound/provider-azure/internal/controller/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/azure/resourceproviderregistration"
	subscription "github.com/upbound/provider-azure/internal/controller/azure/subscription"
	providerconfig "github.com/upbound/provider-azure/internal/controller/providerconfig"
)

// Setup_config creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_config(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		subscription.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
