// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	applicationinsights "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionruleassociation"
	monitordiagnosticsetting "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordiagnosticsetting"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryruleslog"
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
