// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	loganalyticsdataexportrule "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticsquerypack "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticsquerypack"
	loganalyticsquerypackquery "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticsquerypackquery"
	loganalyticssavedsearch "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/loganalyticssavedsearch"
	workspace "github.com/upbound/provider-azure/v2/internal/controller/cluster/operationalinsights/workspace"
)

// Setup_operationalinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_operationalinsights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loganalyticsdataexportrule.Setup,
		loganalyticsdatasourcewindowsevent.Setup,
		loganalyticsdatasourcewindowsperformancecounter.Setup,
		loganalyticslinkedservice.Setup,
		loganalyticslinkedstorageaccount.Setup,
		loganalyticsquerypack.Setup,
		loganalyticsquerypackquery.Setup,
		loganalyticssavedsearch.Setup,
		workspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_operationalinsights creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_operationalinsights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loganalyticsdataexportrule.SetupGated,
		loganalyticsdatasourcewindowsevent.SetupGated,
		loganalyticsdatasourcewindowsperformancecounter.SetupGated,
		loganalyticslinkedservice.SetupGated,
		loganalyticslinkedstorageaccount.SetupGated,
		loganalyticsquerypack.SetupGated,
		loganalyticsquerypackquery.SetupGated,
		loganalyticssavedsearch.SetupGated,
		workspace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
