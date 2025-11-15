// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/functionjavascriptuda"
	job "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/job"
	managedprivateendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/managedprivateendpoint"
	outputblob "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputblob"
	outputeventhub "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputeventhub"
	outputfunction "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputfunction"
	outputmssql "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputmssql"
	outputpowerbi "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputpowerbi"
	outputservicebusqueue "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputservicebustopic"
	outputsynapse "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputsynapse"
	outputtable "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/outputtable"
	referenceinputblob "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/referenceinputblob"
	referenceinputmssql "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/referenceinputmssql"
	streaminputblob "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/upbound/provider-azure/v2/internal/controller/namespaced/streamanalytics/streaminputiothub"
)

// Setup_streamanalytics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_streamanalytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		functionjavascriptuda.Setup,
		job.Setup,
		managedprivateendpoint.Setup,
		outputblob.Setup,
		outputeventhub.Setup,
		outputfunction.Setup,
		outputmssql.Setup,
		outputpowerbi.Setup,
		outputservicebusqueue.Setup,
		outputservicebustopic.Setup,
		outputsynapse.Setup,
		outputtable.Setup,
		referenceinputblob.Setup,
		referenceinputmssql.Setup,
		streaminputblob.Setup,
		streaminputeventhub.Setup,
		streaminputiothub.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_streamanalytics creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_streamanalytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.SetupGated,
		functionjavascriptuda.SetupGated,
		job.SetupGated,
		managedprivateendpoint.SetupGated,
		outputblob.SetupGated,
		outputeventhub.SetupGated,
		outputfunction.SetupGated,
		outputmssql.SetupGated,
		outputpowerbi.SetupGated,
		outputservicebusqueue.SetupGated,
		outputservicebustopic.SetupGated,
		outputsynapse.SetupGated,
		outputtable.SetupGated,
		referenceinputblob.SetupGated,
		referenceinputmssql.SetupGated,
		streaminputblob.SetupGated,
		streaminputeventhub.SetupGated,
		streaminputiothub.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
