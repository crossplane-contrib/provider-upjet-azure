// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/cosmosdb/account"
	cassandracluster "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/provider-azure/internal/controller/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/provider-azure/internal/controller/cosmosdb/gremlingraph"
	mongocollection "github.com/upbound/provider-azure/internal/controller/cosmosdb/mongocollection"
	mongodatabase "github.com/upbound/provider-azure/internal/controller/cosmosdb/mongodatabase"
	mongoroledefinition "github.com/upbound/provider-azure/internal/controller/cosmosdb/mongoroledefinition"
	mongouserdefinition "github.com/upbound/provider-azure/internal/controller/cosmosdb/mongouserdefinition"
	sqlcontainer "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlcontainer"
	sqldatabase "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqldatabase"
	sqldedicatedgateway "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqldedicatedgateway"
	sqlfunction "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqltrigger"
	table "github.com/upbound/provider-azure/internal/controller/cosmosdb/table"
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
