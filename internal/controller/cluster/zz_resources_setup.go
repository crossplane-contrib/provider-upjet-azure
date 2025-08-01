// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/subscriptiontemplatedeployment"
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
