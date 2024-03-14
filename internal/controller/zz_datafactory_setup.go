// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	customdataset "github.com/upbound/provider-azure/internal/controller/datafactory/customdataset"
	dataflow "github.com/upbound/provider-azure/internal/controller/datafactory/dataflow"
	datasetazureblob "github.com/upbound/provider-azure/internal/controller/datafactory/datasetazureblob"
	datasetbinary "github.com/upbound/provider-azure/internal/controller/datafactory/datasetbinary"
	datasetcosmosdbsqlapi "github.com/upbound/provider-azure/internal/controller/datafactory/datasetcosmosdbsqlapi"
	datasetdelimitedtext "github.com/upbound/provider-azure/internal/controller/datafactory/datasetdelimitedtext"
	datasethttp "github.com/upbound/provider-azure/internal/controller/datafactory/datasethttp"
	datasetjson "github.com/upbound/provider-azure/internal/controller/datafactory/datasetjson"
	datasetmysql "github.com/upbound/provider-azure/internal/controller/datafactory/datasetmysql"
	datasetparquet "github.com/upbound/provider-azure/internal/controller/datafactory/datasetparquet"
	datasetpostgresql "github.com/upbound/provider-azure/internal/controller/datafactory/datasetpostgresql"
	datasetsnowflake "github.com/upbound/provider-azure/internal/controller/datafactory/datasetsnowflake"
	datasetsqlservertable "github.com/upbound/provider-azure/internal/controller/datafactory/datasetsqlservertable"
	factory "github.com/upbound/provider-azure/internal/controller/datafactory/factory"
	integrationruntimeazure "github.com/upbound/provider-azure/internal/controller/datafactory/integrationruntimeazure"
	integrationruntimeazuressis "github.com/upbound/provider-azure/internal/controller/datafactory/integrationruntimeazuressis"
	integrationruntimemanaged "github.com/upbound/provider-azure/internal/controller/datafactory/integrationruntimemanaged"
	integrationruntimeselfhosted "github.com/upbound/provider-azure/internal/controller/datafactory/integrationruntimeselfhosted"
	linkedcustomservice "github.com/upbound/provider-azure/internal/controller/datafactory/linkedcustomservice"
	linkedserviceazureblobstorage "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceazureblobstorage"
	linkedserviceazuredatabricks "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceazuredatabricks"
	linkedserviceazurefilestorage "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceazurefilestorage"
	linkedserviceazurefunction "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceazurefunction"
	linkedserviceazuresearch "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceazuresearch"
	linkedserviceazuresqldatabase "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceazuresqldatabase"
	linkedserviceazuretablestorage "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceazuretablestorage"
	linkedservicecosmosdb "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicecosmosdb"
	linkedservicecosmosdbmongoapi "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicecosmosdbmongoapi"
	linkedservicedatalakestoragegen2 "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicedatalakestoragegen2"
	linkedservicekeyvault "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicekeyvault"
	linkedservicekusto "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicekusto"
	linkedservicemysql "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicemysql"
	linkedserviceodata "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceodata"
	linkedserviceodbc "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceodbc"
	linkedservicepostgresql "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicepostgresql"
	linkedservicesftp "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicesftp"
	linkedservicesnowflake "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicesnowflake"
	linkedservicesqlserver "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicesqlserver"
	linkedservicesynapse "github.com/upbound/provider-azure/internal/controller/datafactory/linkedservicesynapse"
	linkedserviceweb "github.com/upbound/provider-azure/internal/controller/datafactory/linkedserviceweb"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/datafactory/managedprivateendpoint"
	pipeline "github.com/upbound/provider-azure/internal/controller/datafactory/pipeline"
	triggerblobevent "github.com/upbound/provider-azure/internal/controller/datafactory/triggerblobevent"
	triggercustomevent "github.com/upbound/provider-azure/internal/controller/datafactory/triggercustomevent"
	triggerschedule "github.com/upbound/provider-azure/internal/controller/datafactory/triggerschedule"
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
		integrationruntimemanaged.Setup,
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
