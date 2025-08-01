// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	containerapp "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/containerapp"
	customdomain "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/customdomain"
	environment "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environment"
	environmentcertificate "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentcertificate"
	environmentcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentcustomdomain"
	environmentdaprcomponent "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentdaprcomponent"
	environmentstorage "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentstorage"
)

// Setup_containerapp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerapp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerapp.Setup,
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
