// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	loganalyticsdataexportrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticsquerypack "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticsquerypack"
	loganalyticsquerypackquery "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticsquerypackquery"
	loganalyticssavedsearch "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/loganalyticssavedsearch"
	workspace "github.com/upbound/provider-azure/v2/internal/controller/namespaced/operationalinsights/workspace"
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

// SetupWebhookWithManager_operationalinsights registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_operationalinsights(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		loganalyticsdataexportrule.SetupWebhookWithManager,
		loganalyticsdatasourcewindowsevent.SetupWebhookWithManager,
		loganalyticsdatasourcewindowsperformancecounter.SetupWebhookWithManager,
		loganalyticslinkedservice.SetupWebhookWithManager,
		loganalyticslinkedstorageaccount.SetupWebhookWithManager,
		loganalyticsquerypack.SetupWebhookWithManager,
		loganalyticsquerypackquery.SetupWebhookWithManager,
		loganalyticssavedsearch.SetupWebhookWithManager,
		workspace.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
