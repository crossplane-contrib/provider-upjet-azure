// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appactiveslot "github.com/upbound/provider-azure/internal/controller/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/internal/controller/web/apphybridconnection"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/web/functionapp"
	functionappactiveslot "github.com/upbound/provider-azure/internal/controller/web/functionappactiveslot"
	functionappfunction "github.com/upbound/provider-azure/internal/controller/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/internal/controller/web/functionapphybridconnection"
	functionappslot "github.com/upbound/provider-azure/internal/controller/web/functionappslot"
	linuxfunctionapp "github.com/upbound/provider-azure/internal/controller/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/internal/controller/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/internal/controller/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/internal/controller/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/web/sourcecontroltoken"
	staticsite "github.com/upbound/provider-azure/internal/controller/web/staticsite"
	windowsfunctionapp "github.com/upbound/provider-azure/internal/controller/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/internal/controller/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/internal/controller/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/internal/controller/web/windowswebappslot"
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
