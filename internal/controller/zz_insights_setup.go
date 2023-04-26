/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	applicationinsights "github.com/upbound/provider-azure/internal/controller/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/internal/controller/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/internal/controller/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/internal/controller/insights/monitordatacollectionruleassociation"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryruleslog"
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
