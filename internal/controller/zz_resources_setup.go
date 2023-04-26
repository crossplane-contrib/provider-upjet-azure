/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/resources/subscriptiontemplatedeployment"
)

// Setup_resources creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_resources(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcedeploymentscriptazurecli.Setup,
		resourcedeploymentscriptazurepowershell.Setup,
		resourcegrouptemplatedeployment.Setup,
		subscriptiontemplatedeployment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
