// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/namespaced/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/namespaced/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/namespaced/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/namespaced/resources/subscriptiontemplatedeployment"
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

// SetupGated_resources creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_resources(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcedeploymentscriptazurecli.SetupGated,
		resourcedeploymentscriptazurepowershell.SetupGated,
		resourcegrouptemplatedeployment.SetupGated,
		subscriptiontemplatedeployment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
