/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	managementapitag "github.com/upbound/official-providers/provider-azure/internal/controller/api/managementapitag"
	managementnotificationrecipientuser "github.com/upbound/official-providers/provider-azure/internal/controller/api/managementnotificationrecipientuser"
	api "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/api"
	apidiagnostic "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apidiagnostic"
	apioperation "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apioperation"
	apioperationpolicy "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apioperationpolicy"
	apioperationtag "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apioperationtag"
	apipolicy "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apipolicy"
	apirelease "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apirelease"
	apischema "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apischema"
	apiversionset "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/apiversionset"
	authorizationserver "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/authorizationserver"
	backend "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/backend"
	certificate "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/certificate"
	customdomain "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/customdomain"
	diagnostic "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/diagnostic"
	emailtemplate "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/emailtemplate"
	gateway "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/gateway"
	gatewayapi "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/gatewayapi"
	identityprovideraad "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/identityprovideraad"
	identityprovideraadb2c "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/identityprovideraadb2c"
	identityproviderfacebook "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/identityprovidertwitter"
	logger "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/logger"
	management "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/management"
	namedvalue "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/namedvalue"
	notificationrecipientemail "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/notificationrecipientemail"
	openidconnectprovider "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/openidconnectprovider"
	policy "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/policy"
	product "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/product"
	productapi "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/productapi"
	productpolicy "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/productpolicy"
	rediscache "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/rediscache"
	subscription "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/subscription"
	tag "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/tag"
	user "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/user"
	managementgrouppolicyassignment "github.com/upbound/official-providers/provider-azure/internal/controller/authorization/managementgrouppolicyassignment"
	resourcegrouppolicyassignment "github.com/upbound/official-providers/provider-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	subscriptionpolicyassignment "github.com/upbound/official-providers/provider-azure/internal/controller/authorization/subscriptionpolicyassignment"
	resourcegroup "github.com/upbound/official-providers/provider-azure/internal/controller/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/official-providers/provider-azure/internal/controller/azure/resourceproviderregistration"
	subscriptionazure "github.com/upbound/official-providers/provider-azure/internal/controller/azure/subscription"
	rediscachecache "github.com/upbound/official-providers/provider-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/official-providers/provider-azure/internal/controller/cache/redislinkedserver"
	availabilityset "github.com/upbound/official-providers/provider-azure/internal/controller/compute/availabilityset"
	dedicatedhost "github.com/upbound/official-providers/provider-azure/internal/controller/compute/dedicatedhost"
	diskaccess "github.com/upbound/official-providers/provider-azure/internal/controller/compute/diskaccess"
	diskencryptionset "github.com/upbound/official-providers/provider-azure/internal/controller/compute/diskencryptionset"
	image "github.com/upbound/official-providers/provider-azure/internal/controller/compute/image"
	linuxvirtualmachine "github.com/upbound/official-providers/provider-azure/internal/controller/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/official-providers/provider-azure/internal/controller/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/official-providers/provider-azure/internal/controller/compute/manageddisk"
	orchestratedvirtualmachinescaleset "github.com/upbound/official-providers/provider-azure/internal/controller/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/official-providers/provider-azure/internal/controller/compute/proximityplacementgroup"
	sharedimagegallery "github.com/upbound/official-providers/provider-azure/internal/controller/compute/sharedimagegallery"
	snapshot "github.com/upbound/official-providers/provider-azure/internal/controller/compute/snapshot"
	windowsvirtualmachine "github.com/upbound/official-providers/provider-azure/internal/controller/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/official-providers/provider-azure/internal/controller/compute/windowsvirtualmachinescaleset"
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
	accountdatashare "github.com/upbound/official-providers/provider-azure/internal/controller/datashare/account"
	datasetblobstorage "github.com/upbound/official-providers/provider-azure/internal/controller/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/official-providers/provider-azure/internal/controller/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/official-providers/provider-azure/internal/controller/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/official-providers/provider-azure/internal/controller/datashare/datasetkustodatabase"
	datashare "github.com/upbound/official-providers/provider-azure/internal/controller/datashare/datashare"
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
	cluster "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/cluster"
	consumergroup "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/consumergroup"
	eventhub "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/namespaceauthorizationrule"
	namespacecustomermanagedkey "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/namespacecustomermanagedkey"
	namespacedisasterrecoveryconfig "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/namespacedisasterrecoveryconfig"
	monitormetricalert "github.com/upbound/official-providers/provider-azure/internal/controller/insights/monitormetricalert"
	certificateiothub "github.com/upbound/official-providers/provider-azure/internal/controller/iothub/certificate"
	accesspolicy "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/accesspolicy"
	certificatekeyvault "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/certificate"
	certificateissuer "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/certificateissuer"
	key "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/secret"
	vault "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/vault"
	integrationserviceenvironment "github.com/upbound/official-providers/provider-azure/internal/controller/logic/integrationserviceenvironment"
	disksastoken "github.com/upbound/official-providers/provider-azure/internal/controller/managed/disksastoken"
	grouppolicyexemption "github.com/upbound/official-providers/provider-azure/internal/controller/management/grouppolicyexemption"
	grouppolicyremediation "github.com/upbound/official-providers/provider-azure/internal/controller/management/grouppolicyremediation"
	managementgroup "github.com/upbound/official-providers/provider-azure/internal/controller/management/managementgroup"
	marketplaceagreement "github.com/upbound/official-providers/provider-azure/internal/controller/marketplaceordering/marketplaceagreement"
	serverdnsalias "github.com/upbound/official-providers/provider-azure/internal/controller/mssql/serverdnsalias"
	ipgroup "github.com/upbound/official-providers/provider-azure/internal/controller/network/ipgroup"
	subnet "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetserviceendpointstoragepolicy"
	virtualnetwork "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkpeering"
	authorizationrulenotificationhubs "github.com/upbound/official-providers/provider-azure/internal/controller/notificationhubs/authorizationrule"
	namespace "github.com/upbound/official-providers/provider-azure/internal/controller/notificationhubs/namespace"
	notificationhub "github.com/upbound/official-providers/provider-azure/internal/controller/notificationhubs/notificationhub"
	workspace "github.com/upbound/official-providers/provider-azure/internal/controller/operationalinsights/workspace"
	providerconfig "github.com/upbound/official-providers/provider-azure/internal/controller/providerconfig"
	groupcostmanagementexport "github.com/upbound/official-providers/provider-azure/internal/controller/resource/groupcostmanagementexport"
	grouppolicyexemptionresource "github.com/upbound/official-providers/provider-azure/internal/controller/resource/grouppolicyexemption"
	grouppolicyremediationresource "github.com/upbound/official-providers/provider-azure/internal/controller/resource/grouppolicyremediation"
	managementgroupsubscriptionassociation "github.com/upbound/official-providers/provider-azure/internal/controller/resources/managementgroupsubscriptionassociation"
	managementgrouptemplatedeployment "github.com/upbound/official-providers/provider-azure/internal/controller/resources/managementgrouptemplatedeployment"
	resourcegrouptemplatedeployment "github.com/upbound/official-providers/provider-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/official-providers/provider-azure/internal/controller/resources/subscriptiontemplatedeployment"
	advancedthreatprotection "github.com/upbound/official-providers/provider-azure/internal/controller/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/official-providers/provider-azure/internal/controller/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/official-providers/provider-azure/internal/controller/security/iotsecuritysolution"
	mssqlserver "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlserver"
	mssqlserversecurityalertpolicy "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlservervulnerabilityassessment"
	accountstorage "github.com/upbound/official-providers/provider-azure/internal/controller/storage/account"
	accountcustomermanagedkey "github.com/upbound/official-providers/provider-azure/internal/controller/storage/accountcustomermanagedkey"
	accountnetworkrules "github.com/upbound/official-providers/provider-azure/internal/controller/storage/accountnetworkrules"
	blob "github.com/upbound/official-providers/provider-azure/internal/controller/storage/blob"
	blobinventorypolicy "github.com/upbound/official-providers/provider-azure/internal/controller/storage/blobinventorypolicy"
	container "github.com/upbound/official-providers/provider-azure/internal/controller/storage/container"
	cloudendpoint "github.com/upbound/official-providers/provider-azure/internal/controller/storagesync/cloudendpoint"
	storagesync "github.com/upbound/official-providers/provider-azure/internal/controller/storagesync/storagesync"
	costmanagementexport "github.com/upbound/official-providers/provider-azure/internal/controller/subscription/costmanagementexport"
	policyexemption "github.com/upbound/official-providers/provider-azure/internal/controller/subscription/policyexemption"
	policyremediation "github.com/upbound/official-providers/provider-azure/internal/controller/subscription/policyremediation"
	networkgatewaynatrule "github.com/upbound/official-providers/provider-azure/internal/controller/virtual/networkgatewaynatrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementapitag.Setup,
		managementnotificationrecipientuser.Setup,
		api.Setup,
		apidiagnostic.Setup,
		apioperation.Setup,
		apioperationpolicy.Setup,
		apioperationtag.Setup,
		apipolicy.Setup,
		apirelease.Setup,
		apischema.Setup,
		apiversionset.Setup,
		authorizationserver.Setup,
		backend.Setup,
		certificate.Setup,
		customdomain.Setup,
		diagnostic.Setup,
		emailtemplate.Setup,
		gateway.Setup,
		gatewayapi.Setup,
		identityprovideraad.Setup,
		identityprovideraadb2c.Setup,
		identityproviderfacebook.Setup,
		identityprovidergoogle.Setup,
		identityprovidermicrosoft.Setup,
		identityprovidertwitter.Setup,
		logger.Setup,
		management.Setup,
		namedvalue.Setup,
		notificationrecipientemail.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		product.Setup,
		productapi.Setup,
		productpolicy.Setup,
		rediscache.Setup,
		subscription.Setup,
		tag.Setup,
		user.Setup,
		managementgrouppolicyassignment.Setup,
		resourcegrouppolicyassignment.Setup,
		subscriptionpolicyassignment.Setup,
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		subscriptionazure.Setup,
		rediscachecache.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
		availabilityset.Setup,
		dedicatedhost.Setup,
		diskaccess.Setup,
		diskencryptionset.Setup,
		image.Setup,
		linuxvirtualmachine.Setup,
		linuxvirtualmachinescaleset.Setup,
		manageddisk.Setup,
		orchestratedvirtualmachinescaleset.Setup,
		proximityplacementgroup.Setup,
		sharedimagegallery.Setup,
		snapshot.Setup,
		windowsvirtualmachine.Setup,
		windowsvirtualmachinescaleset.Setup,
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
		accountdatashare.Setup,
		datasetblobstorage.Setup,
		datasetdatalakegen2.Setup,
		datasetkustocluster.Setup,
		datasetkustodatabase.Setup,
		datashare.Setup,
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
		cluster.Setup,
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		namespaceauthorizationrule.Setup,
		namespacecustomermanagedkey.Setup,
		namespacedisasterrecoveryconfig.Setup,
		monitormetricalert.Setup,
		certificateiothub.Setup,
		accesspolicy.Setup,
		certificatekeyvault.Setup,
		certificateissuer.Setup,
		key.Setup,
		managedhardwaresecuritymodule.Setup,
		managedstorageaccount.Setup,
		managedstorageaccountsastokendefinition.Setup,
		secret.Setup,
		vault.Setup,
		integrationserviceenvironment.Setup,
		disksastoken.Setup,
		grouppolicyexemption.Setup,
		grouppolicyremediation.Setup,
		managementgroup.Setup,
		marketplaceagreement.Setup,
		serverdnsalias.Setup,
		ipgroup.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		virtualnetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		authorizationrulenotificationhubs.Setup,
		namespace.Setup,
		notificationhub.Setup,
		workspace.Setup,
		providerconfig.Setup,
		groupcostmanagementexport.Setup,
		grouppolicyexemptionresource.Setup,
		grouppolicyremediationresource.Setup,
		managementgroupsubscriptionassociation.Setup,
		managementgrouptemplatedeployment.Setup,
		resourcegrouptemplatedeployment.Setup,
		subscriptiontemplatedeployment.Setup,
		advancedthreatprotection.Setup,
		iotsecuritydevicegroup.Setup,
		iotsecuritysolution.Setup,
		mssqlserver.Setup,
		mssqlserversecurityalertpolicy.Setup,
		mssqlservertransparentdataencryption.Setup,
		mssqlservervulnerabilityassessment.Setup,
		accountstorage.Setup,
		accountcustomermanagedkey.Setup,
		accountnetworkrules.Setup,
		blob.Setup,
		blobinventorypolicy.Setup,
		container.Setup,
		cloudendpoint.Setup,
		storagesync.Setup,
		costmanagementexport.Setup,
		policyexemption.Setup,
		policyremediation.Setup,
		networkgatewaynatrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
