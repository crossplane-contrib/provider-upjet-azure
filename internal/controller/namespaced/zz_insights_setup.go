// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	applicationinsights "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitordatacollectionruleassociation"
	monitordiagnosticsetting "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitordiagnosticsetting"
	monitormetricalert "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitorscheduledqueryruleslog"
	monitorworkspace "github.com/upbound/provider-azure/v2/internal/controller/namespaced/insights/monitorworkspace"
)

// Setup_insights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_insights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationinsights.Setup,
		applicationinsightsanalyticsitem.Setup,
		applicationinsightsapikey.Setup,
		applicationinsightssmartdetectionrule.Setup,
		applicationinsightsstandardwebtest.Setup,
		applicationinsightswebtest.Setup,
		applicationinsightsworkbook.Setup,
		applicationinsightsworkbooktemplate.Setup,
		monitoractiongroup.Setup,
		monitoractivitylogalert.Setup,
		monitorautoscalesetting.Setup,
		monitordatacollectionendpoint.Setup,
		monitordatacollectionrule.Setup,
		monitordatacollectionruleassociation.Setup,
		monitordiagnosticsetting.Setup,
		monitormetricalert.Setup,
		monitorprivatelinkscope.Setup,
		monitorprivatelinkscopedservice.Setup,
		monitorscheduledqueryrulesalert.Setup,
		monitorscheduledqueryrulesalertv2.Setup,
		monitorscheduledqueryruleslog.Setup,
		monitorworkspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_insights creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_insights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationinsights.SetupGated,
		applicationinsightsanalyticsitem.SetupGated,
		applicationinsightsapikey.SetupGated,
		applicationinsightssmartdetectionrule.SetupGated,
		applicationinsightsstandardwebtest.SetupGated,
		applicationinsightswebtest.SetupGated,
		applicationinsightsworkbook.SetupGated,
		applicationinsightsworkbooktemplate.SetupGated,
		monitoractiongroup.SetupGated,
		monitoractivitylogalert.SetupGated,
		monitorautoscalesetting.SetupGated,
		monitordatacollectionendpoint.SetupGated,
		monitordatacollectionrule.SetupGated,
		monitordatacollectionruleassociation.SetupGated,
		monitordiagnosticsetting.SetupGated,
		monitormetricalert.SetupGated,
		monitorprivatelinkscope.SetupGated,
		monitorprivatelinkscopedservice.SetupGated,
		monitorscheduledqueryrulesalert.SetupGated,
		monitorscheduledqueryrulesalertv2.SetupGated,
		monitorscheduledqueryruleslog.SetupGated,
		monitorworkspace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
