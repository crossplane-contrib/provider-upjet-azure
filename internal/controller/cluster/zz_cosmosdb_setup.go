// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/account"
	cassandracluster "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/gremlingraph"
	mongocluster "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongocluster"
	mongocollection "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongocollection"
	mongodatabase "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongodatabase"
	mongoroledefinition "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongoroledefinition"
	mongouserdefinition "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/mongouserdefinition"
	sqlcontainer "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlcontainer"
	sqldatabase "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqldatabase"
	sqldedicatedgateway "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqldedicatedgateway"
	sqlfunction "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/sqltrigger"
	table "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/table"
)

// Setup_cosmosdb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cosmosdb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		cassandracluster.Setup,
		cassandradatacenter.Setup,
		cassandrakeyspace.Setup,
		cassandratable.Setup,
		gremlindatabase.Setup,
		gremlingraph.Setup,
		mongocluster.Setup,
		mongocollection.Setup,
		mongodatabase.Setup,
		mongoroledefinition.Setup,
		mongouserdefinition.Setup,
		sqlcontainer.Setup,
		sqldatabase.Setup,
		sqldedicatedgateway.Setup,
		sqlfunction.Setup,
		sqlroleassignment.Setup,
		sqlroledefinition.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		table.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cosmosdb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cosmosdb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		cassandracluster.SetupGated,
		cassandradatacenter.SetupGated,
		cassandrakeyspace.SetupGated,
		cassandratable.SetupGated,
		gremlindatabase.SetupGated,
		gremlingraph.SetupGated,
		mongocluster.SetupGated,
		mongocollection.SetupGated,
		mongodatabase.SetupGated,
		mongoroledefinition.SetupGated,
		mongouserdefinition.SetupGated,
		sqlcontainer.SetupGated,
		sqldatabase.SetupGated,
		sqldedicatedgateway.SetupGated,
		sqlfunction.SetupGated,
		sqlroleassignment.SetupGated,
		sqlroledefinition.SetupGated,
		sqlstoredprocedure.SetupGated,
		sqltrigger.SetupGated,
		table.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
