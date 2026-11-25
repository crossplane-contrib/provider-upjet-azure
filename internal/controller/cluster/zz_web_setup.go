// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appactiveslot "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/apphybridconnection"
	appservicecertificatebinding "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/appservicecertificatebinding"
	appservicecustomhostnamebinding "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/appservicecustomhostnamebinding"
	appservicemanagedcertificate "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/appservicemanagedcertificate"
	functionappactiveslot "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/functionappactiveslot"
	functionappflexconsumption "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/functionappflexconsumption"
	functionappfunction "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/functionapphybridconnection"
	linuxfunctionapp "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/sourcecontroltoken"
	windowsfunctionapp "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/v2/internal/controller/cluster/web/windowswebappslot"
)

// Setup_web creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_web(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appactiveslot.Setup,
		apphybridconnection.Setup,
		appservicecertificatebinding.Setup,
		appservicecustomhostnamebinding.Setup,
		appservicemanagedcertificate.Setup,
		functionappactiveslot.Setup,
		functionappflexconsumption.Setup,
		functionappfunction.Setup,
		functionapphybridconnection.Setup,
		linuxfunctionapp.Setup,
		linuxfunctionappslot.Setup,
		linuxwebapp.Setup,
		linuxwebappslot.Setup,
		serviceplan.Setup,
		sourcecontroltoken.Setup,
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

// SetupGated_web creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_web(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appactiveslot.SetupGated,
		apphybridconnection.SetupGated,
		appservicecertificatebinding.SetupGated,
		appservicecustomhostnamebinding.SetupGated,
		appservicemanagedcertificate.SetupGated,
		functionappactiveslot.SetupGated,
		functionappflexconsumption.SetupGated,
		functionappfunction.SetupGated,
		functionapphybridconnection.SetupGated,
		linuxfunctionapp.SetupGated,
		linuxfunctionappslot.SetupGated,
		linuxwebapp.SetupGated,
		linuxwebappslot.SetupGated,
		serviceplan.SetupGated,
		sourcecontroltoken.SetupGated,
		windowsfunctionapp.SetupGated,
		windowsfunctionappslot.SetupGated,
		windowswebapp.SetupGated,
		windowswebappslot.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_web registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_web(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		appactiveslot.SetupWebhookWithManager,
		apphybridconnection.SetupWebhookWithManager,
		appservicecertificatebinding.SetupWebhookWithManager,
		appservicecustomhostnamebinding.SetupWebhookWithManager,
		appservicemanagedcertificate.SetupWebhookWithManager,
		functionappactiveslot.SetupWebhookWithManager,
		functionappflexconsumption.SetupWebhookWithManager,
		functionappfunction.SetupWebhookWithManager,
		functionapphybridconnection.SetupWebhookWithManager,
		linuxfunctionapp.SetupWebhookWithManager,
		linuxfunctionappslot.SetupWebhookWithManager,
		linuxwebapp.SetupWebhookWithManager,
		linuxwebappslot.SetupWebhookWithManager,
		serviceplan.SetupWebhookWithManager,
		sourcecontroltoken.SetupWebhookWithManager,
		windowsfunctionapp.SetupWebhookWithManager,
		windowsfunctionappslot.SetupWebhookWithManager,
		windowswebapp.SetupWebhookWithManager,
		windowswebappslot.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
