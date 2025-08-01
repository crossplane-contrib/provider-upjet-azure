// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appactiveslot "github.com/upbound/provider-azure/internal/controller/cluster/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/web/apphybridconnection"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/cluster/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/functionapp"
	functionappactiveslot "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappactiveslot"
	functionappfunction "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/web/functionapphybridconnection"
	functionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappslot"
	linuxfunctionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/internal/controller/cluster/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/cluster/web/sourcecontroltoken"
	staticsite "github.com/upbound/provider-azure/internal/controller/cluster/web/staticsite"
	windowsfunctionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/internal/controller/cluster/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/windowswebappslot"
)

// Setup_web creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_web(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appactiveslot.Setup,
		apphybridconnection.Setup,
		appserviceplan.Setup,
		functionapp.Setup,
		functionappactiveslot.Setup,
		functionappfunction.Setup,
		functionapphybridconnection.Setup,
		functionappslot.Setup,
		linuxfunctionapp.Setup,
		linuxfunctionappslot.Setup,
		linuxwebapp.Setup,
		linuxwebappslot.Setup,
		serviceplan.Setup,
		sourcecontroltoken.Setup,
		staticsite.Setup,
		windowsfunctionapp.Setup,
		windowsfunctionappslot.Setup,
		windowswebapp.Setup,
		windowswebappslot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
