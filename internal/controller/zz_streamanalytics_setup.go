// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/upbound/provider-azure/internal/controller/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/streamanalytics/functionjavascriptuda"
	job "github.com/upbound/provider-azure/internal/controller/streamanalytics/job"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/streamanalytics/managedprivateendpoint"
	outputblob "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputblob"
	outputeventhub "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputeventhub"
	outputfunction "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputfunction"
	outputmssql "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputmssql"
	outputpowerbi "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputpowerbi"
	outputservicebusqueue "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputservicebustopic"
	outputsynapse "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputsynapse"
	outputtable "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputtable"
	referenceinputblob "github.com/upbound/provider-azure/internal/controller/streamanalytics/referenceinputblob"
	referenceinputmssql "github.com/upbound/provider-azure/internal/controller/streamanalytics/referenceinputmssql"
	streaminputblob "github.com/upbound/provider-azure/internal/controller/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/upbound/provider-azure/internal/controller/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/upbound/provider-azure/internal/controller/streamanalytics/streaminputiothub"
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
