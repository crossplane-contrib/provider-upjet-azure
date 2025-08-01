// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	customdataset "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/customdataset"
	dataflow "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/dataflow"
	datasetazureblob "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetazureblob"
	datasetbinary "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetbinary"
	datasetcosmosdbsqlapi "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetcosmosdbsqlapi"
	datasetdelimitedtext "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetdelimitedtext"
	datasethttp "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasethttp"
	datasetjson "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetjson"
	datasetmysql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetmysql"
	datasetparquet "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetparquet"
	datasetpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetpostgresql"
	datasetsnowflake "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetsnowflake"
	datasetsqlservertable "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/datasetsqlservertable"
	factory "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/factory"
	integrationruntimeazure "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimeazure"
	integrationruntimeazuressis "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimeazuressis"
	integrationruntimeselfhosted "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/integrationruntimeselfhosted"
	linkedcustomservice "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedcustomservice"
	linkedserviceazureblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazureblobstorage"
	linkedserviceazuredatabricks "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuredatabricks"
	linkedserviceazurefilestorage "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazurefilestorage"
	linkedserviceazurefunction "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazurefunction"
	linkedserviceazuresearch "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuresearch"
	linkedserviceazuresqldatabase "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuresqldatabase"
	linkedserviceazuretablestorage "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceazuretablestorage"
	linkedservicecosmosdb "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicecosmosdb"
	linkedservicecosmosdbmongoapi "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicecosmosdbmongoapi"
	linkedservicedatalakestoragegen2 "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicedatalakestoragegen2"
	linkedservicekeyvault "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicekeyvault"
	linkedservicekusto "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicekusto"
	linkedservicemysql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicemysql"
	linkedserviceodata "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceodata"
	linkedserviceodbc "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceodbc"
	linkedservicepostgresql "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicepostgresql"
	linkedservicesftp "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesftp"
	linkedservicesnowflake "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesnowflake"
	linkedservicesqlserver "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesqlserver"
	linkedservicesynapse "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedservicesynapse"
	linkedserviceweb "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/linkedserviceweb"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/managedprivateendpoint"
	pipeline "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/pipeline"
	triggerblobevent "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggerblobevent"
	triggercustomevent "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggercustomevent"
	triggerschedule "github.com/upbound/provider-azure/internal/controller/cluster/datafactory/triggerschedule"
)

// Setup_datafactory creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datafactory(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customdataset.Setup,
		dataflow.Setup,
		datasetazureblob.Setup,
		datasetbinary.Setup,
		datasetcosmosdbsqlapi.Setup,
		datasetdelimitedtext.Setup,
		datasethttp.Setup,
		datasetjson.Setup,
		datasetmysql.Setup,
		datasetparquet.Setup,
		datasetpostgresql.Setup,
		datasetsnowflake.Setup,
		datasetsqlservertable.Setup,
		factory.Setup,
		integrationruntimeazure.Setup,
		integrationruntimeazuressis.Setup,
		integrationruntimeselfhosted.Setup,
		linkedcustomservice.Setup,
		linkedserviceazureblobstorage.Setup,
		linkedserviceazuredatabricks.Setup,
		linkedserviceazurefilestorage.Setup,
		linkedserviceazurefunction.Setup,
		linkedserviceazuresearch.Setup,
		linkedserviceazuresqldatabase.Setup,
		linkedserviceazuretablestorage.Setup,
		linkedservicecosmosdb.Setup,
		linkedservicecosmosdbmongoapi.Setup,
		linkedservicedatalakestoragegen2.Setup,
		linkedservicekeyvault.Setup,
		linkedservicekusto.Setup,
		linkedservicemysql.Setup,
		linkedserviceodata.Setup,
		linkedserviceodbc.Setup,
		linkedservicepostgresql.Setup,
		linkedservicesftp.Setup,
		linkedservicesnowflake.Setup,
		linkedservicesqlserver.Setup,
		linkedservicesynapse.Setup,
		linkedserviceweb.Setup,
		managedprivateendpoint.Setup,
		pipeline.Setup,
		triggerblobevent.Setup,
		triggercustomevent.Setup,
		triggerschedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datafactory creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datafactory(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customdataset.SetupGated,
		dataflow.SetupGated,
		datasetazureblob.SetupGated,
		datasetbinary.SetupGated,
		datasetcosmosdbsqlapi.SetupGated,
		datasetdelimitedtext.SetupGated,
		datasethttp.SetupGated,
		datasetjson.SetupGated,
		datasetmysql.SetupGated,
		datasetparquet.SetupGated,
		datasetpostgresql.SetupGated,
		datasetsnowflake.SetupGated,
		datasetsqlservertable.SetupGated,
		factory.SetupGated,
		integrationruntimeazure.SetupGated,
		integrationruntimeazuressis.SetupGated,
		integrationruntimeselfhosted.SetupGated,
		linkedcustomservice.SetupGated,
		linkedserviceazureblobstorage.SetupGated,
		linkedserviceazuredatabricks.SetupGated,
		linkedserviceazurefilestorage.SetupGated,
		linkedserviceazurefunction.SetupGated,
		linkedserviceazuresearch.SetupGated,
		linkedserviceazuresqldatabase.SetupGated,
		linkedserviceazuretablestorage.SetupGated,
		linkedservicecosmosdb.SetupGated,
		linkedservicecosmosdbmongoapi.SetupGated,
		linkedservicedatalakestoragegen2.SetupGated,
		linkedservicekeyvault.SetupGated,
		linkedservicekusto.SetupGated,
		linkedservicemysql.SetupGated,
		linkedserviceodata.SetupGated,
		linkedserviceodbc.SetupGated,
		linkedservicepostgresql.SetupGated,
		linkedservicesftp.SetupGated,
		linkedservicesnowflake.SetupGated,
		linkedservicesqlserver.SetupGated,
		linkedservicesynapse.SetupGated,
		linkedserviceweb.SetupGated,
		managedprivateendpoint.SetupGated,
		pipeline.SetupGated,
		triggerblobevent.SetupGated,
		triggercustomevent.SetupGated,
		triggerschedule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
