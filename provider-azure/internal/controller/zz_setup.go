/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	resourcegrouppolicyassignment "github.com/upbound/official-providers/provider-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	resourcegroup "github.com/upbound/official-providers/provider-azure/internal/controller/azure/resourcegroup"
	rediscache "github.com/upbound/official-providers/provider-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redislinkedserver"
	diskencryptionset "github.com/upbound/official-providers/provider-azure/internal/controller/compute/diskencryptionset"
	kubernetescluster "github.com/upbound/official-providers/provider-azure/internal/controller/containerservice/kubernetescluster"
	kubernetesclusternodepool "github.com/upbound/official-providers/provider-azure/internal/controller/containerservice/kubernetesclusternodepool"
	account "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/account"
	cassandracluster "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/gremlingraph"
	mongocollection "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/mongocollection"
	mongodatabase "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/mongodatabase"
	notebookworkspace "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/notebookworkspace"
	sqlcontainer "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/sqlcontainer"
	sqldatabase "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/sqldatabase"
	sqlfunction "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/sqltrigger"
	table "github.com/upbound/official-providers/provider-azure/internal/controller/cosmosdb/table"
	configuration "github.com/upbound/official-providers/provider-azure/internal/controller/dbformariadb/configuration"
	database "github.com/upbound/official-providers/provider-azure/internal/controller/dbformariadb/database"
	firewallrule "github.com/upbound/official-providers/provider-azure/internal/controller/dbformariadb/firewallrule"
	server "github.com/upbound/official-providers/provider-azure/internal/controller/dbformariadb/server"
	virtualnetworkrule "github.com/upbound/official-providers/provider-azure/internal/controller/dbformariadb/virtualnetworkrule"
	activedirectoryadministrator "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/configuration"
	databasedbforpostgresql "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/database"
	firewallruledbforpostgresql "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserver "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/server"
	serverkey "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/serverkey"
	virtualnetworkruledbforpostgresql "github.com/upbound/official-providers/provider-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
	iothub "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothub"
	iothubconsumergroup "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/official-providers/provider-azure/internal/controller/devices/iothubsharedaccesspolicy"
	authorizationrule "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/authorizationrule"
	consumergroup "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/consumergroup"
	eventhub "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/eventhubnamespace"
	monitormetricalert "github.com/upbound/official-providers/provider-azure/internal/controller/insights/monitormetricalert"
	certificate "github.com/upbound/official-providers/provider-azure/internal/controller/iothub/certificate"
	accesspolicy "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/accesspolicy"
	certificatekeyvault "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/certificate"
	certificateissuer "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/certificateissuer"
	key "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/secret"
	vault "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/vault"
	workspace "github.com/upbound/official-providers/provider-azure/internal/controller/loganalytics/workspace"
	loadbalancer "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancer"
	networkinterface "github.com/upbound/official-providers/provider-azure/internal/controller/network/networkinterface"
	publicip "github.com/upbound/official-providers/provider-azure/internal/controller/network/publicip"
	subnet "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetserviceendpointstoragepolicy"
	virtualnetwork "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualwan"
	providerconfig "github.com/upbound/official-providers/provider-azure/internal/controller/providerconfig"
	groupcostmanagementexport "github.com/upbound/official-providers/provider-azure/internal/controller/resource/groupcostmanagementexport"
	grouppolicyexemption "github.com/upbound/official-providers/provider-azure/internal/controller/resource/grouppolicyexemption"
	grouppolicyremediation "github.com/upbound/official-providers/provider-azure/internal/controller/resource/grouppolicyremediation"
	resourcegrouptemplatedeployment "github.com/upbound/official-providers/provider-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	mssqlserver "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlserver"
	mssqlservertransparentdataencryption "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlservertransparentdataencryption"
	accountstorage "github.com/upbound/official-providers/provider-azure/internal/controller/storage/account"
	blob "github.com/upbound/official-providers/provider-azure/internal/controller/storage/blob"
	container "github.com/upbound/official-providers/provider-azure/internal/controller/storage/container"
	desktopscalingplan "github.com/upbound/official-providers/provider-azure/internal/controller/virtual/desktopscalingplan"
	networkgatewaynatrule "github.com/upbound/official-providers/provider-azure/internal/controller/virtual/networkgatewaynatrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcegrouppolicyassignment.Setup,
		resourcegroup.Setup,
		rediscache.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
		diskencryptionset.Setup,
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		account.Setup,
		cassandracluster.Setup,
		cassandradatacenter.Setup,
		cassandrakeyspace.Setup,
		cassandratable.Setup,
		gremlindatabase.Setup,
		gremlingraph.Setup,
		mongocollection.Setup,
		mongodatabase.Setup,
		notebookworkspace.Setup,
		sqlcontainer.Setup,
		sqldatabase.Setup,
		sqlfunction.Setup,
		sqlroleassignment.Setup,
		sqlroledefinition.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		table.Setup,
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
		activedirectoryadministrator.Setup,
		configurationdbforpostgresql.Setup,
		databasedbforpostgresql.Setup,
		firewallruledbforpostgresql.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		serverdbforpostgresql.Setup,
		serverkey.Setup,
		virtualnetworkruledbforpostgresql.Setup,
		iothub.Setup,
		iothubconsumergroup.Setup,
		iothubdps.Setup,
		iothubdpscertificate.Setup,
		iothubdpssharedaccesspolicy.Setup,
		iothubendpointeventhub.Setup,
		iothubendpointservicebusqueue.Setup,
		iothubendpointservicebustopic.Setup,
		iothubendpointstoragecontainer.Setup,
		iothubenrichment.Setup,
		iothubfallbackroute.Setup,
		iothubroute.Setup,
		iothubsharedaccesspolicy.Setup,
		authorizationrule.Setup,
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		monitormetricalert.Setup,
		certificate.Setup,
		accesspolicy.Setup,
		certificatekeyvault.Setup,
		certificateissuer.Setup,
		key.Setup,
		managedhardwaresecuritymodule.Setup,
		managedstorageaccount.Setup,
		managedstorageaccountsastokendefinition.Setup,
		secret.Setup,
		vault.Setup,
		workspace.Setup,
		loadbalancer.Setup,
		networkinterface.Setup,
		publicip.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		virtualnetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
		providerconfig.Setup,
		groupcostmanagementexport.Setup,
		grouppolicyexemption.Setup,
		grouppolicyremediation.Setup,
		resourcegrouptemplatedeployment.Setup,
		mssqlserver.Setup,
		mssqlservertransparentdataencryption.Setup,
		accountstorage.Setup,
		blob.Setup,
		container.Setup,
		desktopscalingplan.Setup,
		networkgatewaynatrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
