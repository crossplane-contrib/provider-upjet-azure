// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	containerapp "github.com/upbound/provider-azure/internal/controller/containerapp/containerapp"
	containerjob "github.com/upbound/provider-azure/internal/controller/containerapp/containerjob"
	customdomain "github.com/upbound/provider-azure/internal/controller/containerapp/customdomain"
	environment "github.com/upbound/provider-azure/internal/controller/containerapp/environment"
	environmentcertificate "github.com/upbound/provider-azure/internal/controller/containerapp/environmentcertificate"
	environmentcustomdomain "github.com/upbound/provider-azure/internal/controller/containerapp/environmentcustomdomain"
	environmentdaprcomponent "github.com/upbound/provider-azure/internal/controller/containerapp/environmentdaprcomponent"
	environmentstorage "github.com/upbound/provider-azure/internal/controller/containerapp/environmentstorage"
)

// Setup_containerapp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerapp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerapp.Setup,
		containerjob.Setup,
		customdomain.Setup,
		environment.Setup,
		environmentcertificate.Setup,
		environmentcustomdomain.Setup,
		environmentdaprcomponent.Setup,
		environmentstorage.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
