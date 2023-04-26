/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	attacheddatabaseconfiguration "github.com/upbound/provider-azure/internal/controller/kusto/attacheddatabaseconfiguration"
	cluster "github.com/upbound/provider-azure/internal/controller/kusto/cluster"
	clustermanagedprivateendpoint "github.com/upbound/provider-azure/internal/controller/kusto/clustermanagedprivateendpoint"
	clusterprincipalassignment "github.com/upbound/provider-azure/internal/controller/kusto/clusterprincipalassignment"
	database "github.com/upbound/provider-azure/internal/controller/kusto/database"
	databaseprincipalassignment "github.com/upbound/provider-azure/internal/controller/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/upbound/provider-azure/internal/controller/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/upbound/provider-azure/internal/controller/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/upbound/provider-azure/internal/controller/kusto/iothubdataconnection"
)

// Setup_kusto creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kusto(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attacheddatabaseconfiguration.Setup,
		cluster.Setup,
		clustermanagedprivateendpoint.Setup,
		clusterprincipalassignment.Setup,
		database.Setup,
		databaseprincipalassignment.Setup,
		eventgriddataconnection.Setup,
		eventhubdataconnection.Setup,
		iothubdataconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
