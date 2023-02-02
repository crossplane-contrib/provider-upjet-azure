/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	monitoractionruleactiongroup "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoractionrulesuppression"
	monitorsmartdetectoralertrule "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitorsmartdetectoralertrule"
	server "github.com/upbound/provider-azure/internal/controller/analysisservices/server"
	api "github.com/upbound/provider-azure/internal/controller/apimanagement/api"
	apidiagnostic "github.com/upbound/provider-azure/internal/controller/apimanagement/apidiagnostic"
	apioperation "github.com/upbound/provider-azure/internal/controller/apimanagement/apioperation"
	apioperationpolicy "github.com/upbound/provider-azure/internal/controller/apimanagement/apioperationpolicy"
	apioperationtag "github.com/upbound/provider-azure/internal/controller/apimanagement/apioperationtag"
	apipolicy "github.com/upbound/provider-azure/internal/controller/apimanagement/apipolicy"
	apirelease "github.com/upbound/provider-azure/internal/controller/apimanagement/apirelease"
	apischema "github.com/upbound/provider-azure/internal/controller/apimanagement/apischema"
	apitag "github.com/upbound/provider-azure/internal/controller/apimanagement/apitag"
	apiversionset "github.com/upbound/provider-azure/internal/controller/apimanagement/apiversionset"
	authorizationserver "github.com/upbound/provider-azure/internal/controller/apimanagement/authorizationserver"
	backend "github.com/upbound/provider-azure/internal/controller/apimanagement/backend"
	certificate "github.com/upbound/provider-azure/internal/controller/apimanagement/certificate"
	diagnostic "github.com/upbound/provider-azure/internal/controller/apimanagement/diagnostic"
	emailtemplate "github.com/upbound/provider-azure/internal/controller/apimanagement/emailtemplate"
	gateway "github.com/upbound/provider-azure/internal/controller/apimanagement/gateway"
	identityprovideraad "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovideraad"
	identityproviderfacebook "github.com/upbound/provider-azure/internal/controller/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovidertwitter"
	logger "github.com/upbound/provider-azure/internal/controller/apimanagement/logger"
	management "github.com/upbound/provider-azure/internal/controller/apimanagement/management"
	namedvalue "github.com/upbound/provider-azure/internal/controller/apimanagement/namedvalue"
	notificationrecipientemail "github.com/upbound/provider-azure/internal/controller/apimanagement/notificationrecipientemail"
	notificationrecipientuser "github.com/upbound/provider-azure/internal/controller/apimanagement/notificationrecipientuser"
	openidconnectprovider "github.com/upbound/provider-azure/internal/controller/apimanagement/openidconnectprovider"
	policy "github.com/upbound/provider-azure/internal/controller/apimanagement/policy"
	product "github.com/upbound/provider-azure/internal/controller/apimanagement/product"
	productapi "github.com/upbound/provider-azure/internal/controller/apimanagement/productapi"
	productpolicy "github.com/upbound/provider-azure/internal/controller/apimanagement/productpolicy"
	rediscache "github.com/upbound/provider-azure/internal/controller/apimanagement/rediscache"
	subscription "github.com/upbound/provider-azure/internal/controller/apimanagement/subscription"
	tag "github.com/upbound/provider-azure/internal/controller/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/apimanagement/user"
	springcloudactivedeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudactivedeployment"
	springcloudapp "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappredisassociation"
	springcloudcertificate "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcertificate"
	springcloudcustomdomain "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcustomdomain"
	springcloudjavadeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudservice"
	springcloudstorage "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudstorage"
	provider "github.com/upbound/provider-azure/internal/controller/attestation/provider"
	managementlock "github.com/upbound/provider-azure/internal/controller/authorization/managementlock"
	policydefinition "github.com/upbound/provider-azure/internal/controller/authorization/policydefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	resourcepolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/resourcepolicyexemption"
	roleassignment "github.com/upbound/provider-azure/internal/controller/authorization/roleassignment"
	roledefinition "github.com/upbound/provider-azure/internal/controller/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/subscriptionpolicyexemption"
	account "github.com/upbound/provider-azure/internal/controller/automation/account"
	credential "github.com/upbound/provider-azure/internal/controller/automation/credential"
	module "github.com/upbound/provider-azure/internal/controller/automation/module"
	variablebool "github.com/upbound/provider-azure/internal/controller/automation/variablebool"
	variableint "github.com/upbound/provider-azure/internal/controller/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/automation/variablestring"
	resourcegroup "github.com/upbound/provider-azure/internal/controller/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/azure/resourceproviderregistration"
	subscriptionazure "github.com/upbound/provider-azure/internal/controller/azure/subscription"
	cluster "github.com/upbound/provider-azure/internal/controller/azurestackhci/cluster"
	botchannelalexa "github.com/upbound/provider-azure/internal/controller/botservice/botchannelalexa"
	botchanneldirectline "github.com/upbound/provider-azure/internal/controller/botservice/botchanneldirectline"
	botchannelline "github.com/upbound/provider-azure/internal/controller/botservice/botchannelline"
	botchannelmsteams "github.com/upbound/provider-azure/internal/controller/botservice/botchannelmsteams"
	botchannelslack "github.com/upbound/provider-azure/internal/controller/botservice/botchannelslack"
	botchannelsms "github.com/upbound/provider-azure/internal/controller/botservice/botchannelsms"
	botchannelsregistration "github.com/upbound/provider-azure/internal/controller/botservice/botchannelsregistration"
	botchannelwebchat "github.com/upbound/provider-azure/internal/controller/botservice/botchannelwebchat"
	botconnection "github.com/upbound/provider-azure/internal/controller/botservice/botconnection"
	botwebapp "github.com/upbound/provider-azure/internal/controller/botservice/botwebapp"
	rediscachecache "github.com/upbound/provider-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "github.com/upbound/provider-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/internal/controller/cache/redislinkedserver"
	endpoint "github.com/upbound/provider-azure/internal/controller/cdn/endpoint"
	frontdoorcustomdomain "github.com/upbound/provider-azure/internal/controller/cdn/frontdoorcustomdomain"
	frontdoorcustomdomainassociation "github.com/upbound/provider-azure/internal/controller/cdn/frontdoorcustomdomainassociation"
	frontdoorendpoint "github.com/upbound/provider-azure/internal/controller/cdn/frontdoorendpoint"
	frontdoororigin "github.com/upbound/provider-azure/internal/controller/cdn/frontdoororigin"
	frontdoororigingroup "github.com/upbound/provider-azure/internal/controller/cdn/frontdoororigingroup"
	frontdoorprofile "github.com/upbound/provider-azure/internal/controller/cdn/frontdoorprofile"
	frontdoorroute "github.com/upbound/provider-azure/internal/controller/cdn/frontdoorroute"
	frontdoorrule "github.com/upbound/provider-azure/internal/controller/cdn/frontdoorrule"
	frontdoorruleset "github.com/upbound/provider-azure/internal/controller/cdn/frontdoorruleset"
	profile "github.com/upbound/provider-azure/internal/controller/cdn/profile"
	accountcognitiveservices "github.com/upbound/provider-azure/internal/controller/cognitiveservices/account"
	service "github.com/upbound/provider-azure/internal/controller/communication/service"
	availabilityset "github.com/upbound/provider-azure/internal/controller/compute/availabilityset"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/compute/diskencryptionset"
	image "github.com/upbound/provider-azure/internal/controller/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/compute/manageddisk"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/compute/proximityplacementgroup"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/compute/snapshot"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/windowsvirtualmachinescaleset"
	ledger "github.com/upbound/provider-azure/internal/controller/confidentialledger/ledger"
	budgetmanagementgroup "github.com/upbound/provider-azure/internal/controller/consumption/budgetmanagementgroup"
	budgetresourcegroup "github.com/upbound/provider-azure/internal/controller/consumption/budgetresourcegroup"
	budgetsubscription "github.com/upbound/provider-azure/internal/controller/consumption/budgetsubscription"
	agentpool "github.com/upbound/provider-azure/internal/controller/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/provider-azure/internal/controller/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/provider-azure/internal/controller/containerregistry/registry"
	scopemap "github.com/upbound/provider-azure/internal/controller/containerregistry/scopemap"
	token "github.com/upbound/provider-azure/internal/controller/containerregistry/token"
	webhook "github.com/upbound/provider-azure/internal/controller/containerregistry/webhook"
	kubernetescluster "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetescluster"
	kubernetesclusternodepool "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetesclusternodepool"
	kubernetesfleetmanager "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetesfleetmanager"
	accountcosmosdb "github.com/upbound/provider-azure/internal/controller/cosmosdb/account"
	cassandracluster "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/provider-azure/internal/controller/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/provider-azure/internal/controller/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/provider-azure/internal/controller/cosmosdb/gremlingraph"
	mongocollection "github.com/upbound/provider-azure/internal/controller/cosmosdb/mongocollection"
	mongodatabase "github.com/upbound/provider-azure/internal/controller/cosmosdb/mongodatabase"
	sqlcontainer "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlcontainer"
	sqldatabase "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqldatabase"
	sqlfunction "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqltrigger"
	table "github.com/upbound/provider-azure/internal/controller/cosmosdb/table"
	resourcegroupcostmanagementexport "github.com/upbound/provider-azure/internal/controller/costmanagement/resourcegroupcostmanagementexport"
	subscriptioncostmanagementexport "github.com/upbound/provider-azure/internal/controller/costmanagement/subscriptioncostmanagementexport"
	device "github.com/upbound/provider-azure/internal/controller/databoxedge/device"
	workspace "github.com/upbound/provider-azure/internal/controller/databricks/workspace"
	workspacecustomermanagedkey "github.com/upbound/provider-azure/internal/controller/databricks/workspacecustomermanagedkey"
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
	factory "github.com/upbound/provider-azure/internal/controller/datafactory/factory"
	integrationruntimeazure "github.com/upbound/provider-azure/internal/controller/datafactory/integrationruntimeazure"
	integrationruntimeazuressis "github.com/upbound/provider-azure/internal/controller/datafactory/integrationruntimeazuressis"
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
	databasemigrationproject "github.com/upbound/provider-azure/internal/controller/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/upbound/provider-azure/internal/controller/datamigration/databasemigrationservice"
	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancedisk"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicydisk"
	backupvault "github.com/upbound/provider-azure/internal/controller/dataprotection/backupvault"
	accountdatashare "github.com/upbound/provider-azure/internal/controller/datashare/account"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/datashare/datashare"
	configuration "github.com/upbound/provider-azure/internal/controller/dbformariadb/configuration"
	database "github.com/upbound/provider-azure/internal/controller/dbformariadb/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/dbformariadb/firewallrule"
	serverdbformariadb "github.com/upbound/provider-azure/internal/controller/dbformariadb/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/dbformariadb/virtualnetworkrule"
	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/dbformysql/activedirectoryadministrator"
	configurationdbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/configuration"
	databasedbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/database"
	firewallruledbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/firewallrule"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverfirewallrule"
	serverdbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/server"
	virtualnetworkruledbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/virtualnetworkrule"
	activedirectoryadministratordbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/configuration"
	databasedbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/database"
	firewallruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfigurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/server"
	serverkey "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/serverkey"
	virtualnetworkruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
	iothub "github.com/upbound/provider-azure/internal/controller/devices/iothub"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/devices/iothubsharedaccesspolicy"
	iothubdeviceupdateaccount "github.com/upbound/provider-azure/internal/controller/deviceupdate/iothubdeviceupdateaccount"
	iothubdeviceupdateinstance "github.com/upbound/provider-azure/internal/controller/deviceupdate/iothubdeviceupdateinstance"
	globalvmshutdownschedule "github.com/upbound/provider-azure/internal/controller/devtestlab/globalvmshutdownschedule"
	lab "github.com/upbound/provider-azure/internal/controller/devtestlab/lab"
	linuxvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/devtestlab/linuxvirtualmachine"
	policydevtestlab "github.com/upbound/provider-azure/internal/controller/devtestlab/policy"
	schedule "github.com/upbound/provider-azure/internal/controller/devtestlab/schedule"
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/devtestlab/virtualnetwork"
	windowsvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/devtestlab/windowsvirtualmachine"
	instance "github.com/upbound/provider-azure/internal/controller/digitaltwins/instance"
	cloudelasticsearch "github.com/upbound/provider-azure/internal/controller/elastic/cloudelasticsearch"
	domain "github.com/upbound/provider-azure/internal/controller/eventgrid/domain"
	domaintopic "github.com/upbound/provider-azure/internal/controller/eventgrid/domaintopic"
	eventsubscription "github.com/upbound/provider-azure/internal/controller/eventgrid/eventsubscription"
	systemtopic "github.com/upbound/provider-azure/internal/controller/eventgrid/systemtopic"
	topic "github.com/upbound/provider-azure/internal/controller/eventgrid/topic"
	authorizationrule "github.com/upbound/provider-azure/internal/controller/eventhub/authorizationrule"
	consumergroup "github.com/upbound/provider-azure/internal/controller/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/eventhub/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/eventhub/namespacedisasterrecoveryconfig"
	healthcareservice "github.com/upbound/provider-azure/internal/controller/healthcareapis/healthcareservice"
	applicationinsights "github.com/upbound/provider-azure/internal/controller/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightssmartdetectionrule"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/insights/monitorautoscalesetting"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryruleslog"
	application "github.com/upbound/provider-azure/internal/controller/iotcentral/application"
	applicationnetworkruleset "github.com/upbound/provider-azure/internal/controller/iotcentral/applicationnetworkruleset"
	accesspolicy "github.com/upbound/provider-azure/internal/controller/keyvault/accesspolicy"
	certificatekeyvault "github.com/upbound/provider-azure/internal/controller/keyvault/certificate"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/keyvault/secret"
	vault "github.com/upbound/provider-azure/internal/controller/keyvault/vault"
	attacheddatabaseconfiguration "github.com/upbound/provider-azure/internal/controller/kusto/attacheddatabaseconfiguration"
	clusterkusto "github.com/upbound/provider-azure/internal/controller/kusto/cluster"
	clustermanagedprivateendpoint "github.com/upbound/provider-azure/internal/controller/kusto/clustermanagedprivateendpoint"
	clusterprincipalassignment "github.com/upbound/provider-azure/internal/controller/kusto/clusterprincipalassignment"
	databasekusto "github.com/upbound/provider-azure/internal/controller/kusto/database"
	databaseprincipalassignment "github.com/upbound/provider-azure/internal/controller/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/upbound/provider-azure/internal/controller/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/upbound/provider-azure/internal/controller/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/upbound/provider-azure/internal/controller/kusto/iothubdataconnection"
	appactioncustom "github.com/upbound/provider-azure/internal/controller/logic/appactioncustom"
	appactionhttp "github.com/upbound/provider-azure/internal/controller/logic/appactionhttp"
	appintegrationaccount "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccount"
	appintegrationaccountbatchconfiguration "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountschema "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccountsession"
	apptriggercustom "github.com/upbound/provider-azure/internal/controller/logic/apptriggercustom"
	apptriggerhttprequest "github.com/upbound/provider-azure/internal/controller/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/upbound/provider-azure/internal/controller/logic/apptriggerrecurrence"
	appworkflow "github.com/upbound/provider-azure/internal/controller/logic/appworkflow"
	integrationserviceenvironment "github.com/upbound/provider-azure/internal/controller/logic/integrationserviceenvironment"
	userassignedidentity "github.com/upbound/provider-azure/internal/controller/managedidentity/userassignedidentity"
	managementgroup "github.com/upbound/provider-azure/internal/controller/management/managementgroup"
	accountmaps "github.com/upbound/provider-azure/internal/controller/maps/account"
	creator "github.com/upbound/provider-azure/internal/controller/maps/creator"
	marketplaceagreement "github.com/upbound/provider-azure/internal/controller/marketplaceordering/marketplaceagreement"
	asset "github.com/upbound/provider-azure/internal/controller/media/asset"
	assetfilter "github.com/upbound/provider-azure/internal/controller/media/assetfilter"
	contentkeypolicy "github.com/upbound/provider-azure/internal/controller/media/contentkeypolicy"
	liveevent "github.com/upbound/provider-azure/internal/controller/media/liveevent"
	liveeventoutput "github.com/upbound/provider-azure/internal/controller/media/liveeventoutput"
	servicesaccount "github.com/upbound/provider-azure/internal/controller/media/servicesaccount"
	streamingendpoint "github.com/upbound/provider-azure/internal/controller/media/streamingendpoint"
	streaminglocator "github.com/upbound/provider-azure/internal/controller/media/streaminglocator"
	streamingpolicy "github.com/upbound/provider-azure/internal/controller/media/streamingpolicy"
	transform "github.com/upbound/provider-azure/internal/controller/media/transform"
	spatialanchorsaccount "github.com/upbound/provider-azure/internal/controller/mixedreality/spatialanchorsaccount"
	accountnetapp "github.com/upbound/provider-azure/internal/controller/netapp/account"
	pool "github.com/upbound/provider-azure/internal/controller/netapp/pool"
	snapshotnetapp "github.com/upbound/provider-azure/internal/controller/netapp/snapshot"
	snapshotpolicy "github.com/upbound/provider-azure/internal/controller/netapp/snapshotpolicy"
	volume "github.com/upbound/provider-azure/internal/controller/netapp/volume"
	applicationgateway "github.com/upbound/provider-azure/internal/controller/network/applicationgateway"
	applicationsecuritygroup "github.com/upbound/provider-azure/internal/controller/network/applicationsecuritygroup"
	connectionmonitor "github.com/upbound/provider-azure/internal/controller/network/connectionmonitor"
	ddosprotectionplan "github.com/upbound/provider-azure/internal/controller/network/ddosprotectionplan"
	dnsaaaarecord "github.com/upbound/provider-azure/internal/controller/network/dnsaaaarecord"
	dnsarecord "github.com/upbound/provider-azure/internal/controller/network/dnsarecord"
	dnscaarecord "github.com/upbound/provider-azure/internal/controller/network/dnscaarecord"
	dnscnamerecord "github.com/upbound/provider-azure/internal/controller/network/dnscnamerecord"
	dnsmxrecord "github.com/upbound/provider-azure/internal/controller/network/dnsmxrecord"
	dnsnsrecord "github.com/upbound/provider-azure/internal/controller/network/dnsnsrecord"
	dnsptrrecord "github.com/upbound/provider-azure/internal/controller/network/dnsptrrecord"
	dnssrvrecord "github.com/upbound/provider-azure/internal/controller/network/dnssrvrecord"
	dnstxtrecord "github.com/upbound/provider-azure/internal/controller/network/dnstxtrecord"
	dnszone "github.com/upbound/provider-azure/internal/controller/network/dnszone"
	expressroutecircuit "github.com/upbound/provider-azure/internal/controller/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/upbound/provider-azure/internal/controller/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/upbound/provider-azure/internal/controller/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/upbound/provider-azure/internal/controller/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/upbound/provider-azure/internal/controller/network/expressrouteconnection"
	expressroutegateway "github.com/upbound/provider-azure/internal/controller/network/expressroutegateway"
	expressrouteport "github.com/upbound/provider-azure/internal/controller/network/expressrouteport"
	firewall "github.com/upbound/provider-azure/internal/controller/network/firewall"
	firewallapplicationrulecollection "github.com/upbound/provider-azure/internal/controller/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/upbound/provider-azure/internal/controller/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/upbound/provider-azure/internal/controller/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/upbound/provider-azure/internal/controller/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/upbound/provider-azure/internal/controller/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/upbound/provider-azure/internal/controller/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/upbound/provider-azure/internal/controller/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicy "github.com/upbound/provider-azure/internal/controller/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/upbound/provider-azure/internal/controller/network/frontdoorrulesengine"
	ipgroup "github.com/upbound/provider-azure/internal/controller/network/ipgroup"
	loadbalancer "github.com/upbound/provider-azure/internal/controller/network/loadbalancer"
	loadbalancerbackendaddresspool "github.com/upbound/provider-azure/internal/controller/network/loadbalancerbackendaddresspool"
	loadbalancerbackendaddresspooladdress "github.com/upbound/provider-azure/internal/controller/network/loadbalancerbackendaddresspooladdress"
	loadbalancernatpool "github.com/upbound/provider-azure/internal/controller/network/loadbalancernatpool"
	loadbalancernatrule "github.com/upbound/provider-azure/internal/controller/network/loadbalancernatrule"
	loadbalanceroutboundrule "github.com/upbound/provider-azure/internal/controller/network/loadbalanceroutboundrule"
	loadbalancerprobe "github.com/upbound/provider-azure/internal/controller/network/loadbalancerprobe"
	loadbalancerrule "github.com/upbound/provider-azure/internal/controller/network/loadbalancerrule"
	localnetworkgateway "github.com/upbound/provider-azure/internal/controller/network/localnetworkgateway"
	natgateway "github.com/upbound/provider-azure/internal/controller/network/natgateway"
	natgatewaypublicipassociation "github.com/upbound/provider-azure/internal/controller/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/upbound/provider-azure/internal/controller/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/upbound/provider-azure/internal/controller/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/upbound/provider-azure/internal/controller/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/upbound/provider-azure/internal/controller/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/network/networkinterfacesecuritygroupassociation"
	packetcapture "github.com/upbound/provider-azure/internal/controller/network/packetcapture"
	pointtositevpngateway "github.com/upbound/provider-azure/internal/controller/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/upbound/provider-azure/internal/controller/network/privatednsaaaarecord"
	privatednsarecord "github.com/upbound/provider-azure/internal/controller/network/privatednsarecord"
	privatednscnamerecord "github.com/upbound/provider-azure/internal/controller/network/privatednscnamerecord"
	privatednsmxrecord "github.com/upbound/provider-azure/internal/controller/network/privatednsmxrecord"
	privatednsptrrecord "github.com/upbound/provider-azure/internal/controller/network/privatednsptrrecord"
	privatednssrvrecord "github.com/upbound/provider-azure/internal/controller/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/provider-azure/internal/controller/network/privatednstxtrecord"
	privatednszone "github.com/upbound/provider-azure/internal/controller/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/provider-azure/internal/controller/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/provider-azure/internal/controller/network/privateendpoint"
	privatelinkservice "github.com/upbound/provider-azure/internal/controller/network/privatelinkservice"
	profilenetwork "github.com/upbound/provider-azure/internal/controller/network/profile"
	publicip "github.com/upbound/provider-azure/internal/controller/network/publicip"
	publicipprefix "github.com/upbound/provider-azure/internal/controller/network/publicipprefix"
	routetable "github.com/upbound/provider-azure/internal/controller/network/routetable"
	securitygroup "github.com/upbound/provider-azure/internal/controller/network/securitygroup"
	securityrule "github.com/upbound/provider-azure/internal/controller/network/securityrule"
	subnet "github.com/upbound/provider-azure/internal/controller/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/provider-azure/internal/controller/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/provider-azure/internal/controller/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/provider-azure/internal/controller/network/subnetserviceendpointstoragepolicy"
	trafficmanagerprofile "github.com/upbound/provider-azure/internal/controller/network/trafficmanagerprofile"
	virtualhub "github.com/upbound/provider-azure/internal/controller/network/virtualhub"
	virtualnetworknetwork "github.com/upbound/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/internal/controller/network/virtualwan"
	vpngateway "github.com/upbound/provider-azure/internal/controller/network/vpngateway"
	vpngatewayconnection "github.com/upbound/provider-azure/internal/controller/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/upbound/provider-azure/internal/controller/network/vpnserverconfiguration"
	vpnsite "github.com/upbound/provider-azure/internal/controller/network/vpnsite"
	watcher "github.com/upbound/provider-azure/internal/controller/network/watcher"
	watcherflowlog "github.com/upbound/provider-azure/internal/controller/network/watcherflowlog"
	webapplicationfirewallpolicy "github.com/upbound/provider-azure/internal/controller/network/webapplicationfirewallpolicy"
	authorizationrulenotificationhubs "github.com/upbound/provider-azure/internal/controller/notificationhubs/authorizationrule"
	notificationhub "github.com/upbound/provider-azure/internal/controller/notificationhubs/notificationhub"
	notificationhubnamespace "github.com/upbound/provider-azure/internal/controller/notificationhubs/notificationhubnamespace"
	loganalyticsdataexportrule "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticssavedsearch "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticssavedsearch"
	workspaceoperationalinsights "github.com/upbound/provider-azure/internal/controller/operationalinsights/workspace"
	loganalyticssolution "github.com/upbound/provider-azure/internal/controller/operationsmanagement/loganalyticssolution"
	subscriptionpolicyremediation "github.com/upbound/provider-azure/internal/controller/policyinsights/subscriptionpolicyremediation"
	dashboard "github.com/upbound/provider-azure/internal/controller/portal/dashboard"
	powerbiembedded "github.com/upbound/provider-azure/internal/controller/powerbidedicated/powerbiembedded"
	providerconfig "github.com/upbound/provider-azure/internal/controller/providerconfig"
	accountpurview "github.com/upbound/provider-azure/internal/controller/purview/account"
	backupcontainerstorageaccount "github.com/upbound/provider-azure/internal/controller/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/upbound/provider-azure/internal/controller/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/upbound/provider-azure/internal/controller/recoveryservices/backuppolicyvm"
	backuppolicyvmworkload "github.com/upbound/provider-azure/internal/controller/recoveryservices/backuppolicyvmworkload"
	backupprotectedfileshare "github.com/upbound/provider-azure/internal/controller/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/upbound/provider-azure/internal/controller/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryfabric"
	vaultrecoveryservices "github.com/upbound/provider-azure/internal/controller/recoveryservices/vault"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/resources/subscriptiontemplatedeployment"
	servicesearch "github.com/upbound/provider-azure/internal/controller/search/service"
	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/security/iotsecuritysolution"
	securitycenterassessment "github.com/upbound/provider-azure/internal/controller/security/securitycenterassessment"
	securitycenterassessmentpolicy "github.com/upbound/provider-azure/internal/controller/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/upbound/provider-azure/internal/controller/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/upbound/provider-azure/internal/controller/security/securitycentercontact"
	securitycenterservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/security/securitycenterservervulnerabilityassessment"
	securitycenterservervulnerabilityassessmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/security/securitycenterservervulnerabilityassessmentvirtualmachine"
	securitycentersetting "github.com/upbound/provider-azure/internal/controller/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/upbound/provider-azure/internal/controller/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/upbound/provider-azure/internal/controller/security/securitycenterworkspace"
	sentinelautomationrule "github.com/upbound/provider-azure/internal/controller/securityinsights/sentinelautomationrule"
	sentinelwatchlist "github.com/upbound/provider-azure/internal/controller/securityinsights/sentinelwatchlist"
	namespaceauthorizationruleservicebus "github.com/upbound/provider-azure/internal/controller/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfigservicebus "github.com/upbound/provider-azure/internal/controller/servicebus/namespacedisasterrecoveryconfig"
	namespacenetworkruleset "github.com/upbound/provider-azure/internal/controller/servicebus/namespacenetworkruleset"
	queue "github.com/upbound/provider-azure/internal/controller/servicebus/queue"
	queueauthorizationrule "github.com/upbound/provider-azure/internal/controller/servicebus/queueauthorizationrule"
	servicebusnamespace "github.com/upbound/provider-azure/internal/controller/servicebus/servicebusnamespace"
	subscriptionservicebus "github.com/upbound/provider-azure/internal/controller/servicebus/subscription"
	subscriptionrule "github.com/upbound/provider-azure/internal/controller/servicebus/subscriptionrule"
	topicservicebus "github.com/upbound/provider-azure/internal/controller/servicebus/topic"
	topicauthorizationrule "github.com/upbound/provider-azure/internal/controller/servicebus/topicauthorizationrule"
	clusterservicefabric "github.com/upbound/provider-azure/internal/controller/servicefabric/cluster"
	managedcluster "github.com/upbound/provider-azure/internal/controller/servicefabric/managedcluster"
	networkacl "github.com/upbound/provider-azure/internal/controller/signalrservice/networkacl"
	servicesignalrservice "github.com/upbound/provider-azure/internal/controller/signalrservice/service"
	managedapplicationdefinition "github.com/upbound/provider-azure/internal/controller/solutions/managedapplicationdefinition"
	activedirectoryadministratorsql "github.com/upbound/provider-azure/internal/controller/sql/activedirectoryadministrator"
	databasesql "github.com/upbound/provider-azure/internal/controller/sql/database"
	elasticpool "github.com/upbound/provider-azure/internal/controller/sql/elasticpool"
	firewallrulesql "github.com/upbound/provider-azure/internal/controller/sql/firewallrule"
	mssqldatabase "github.com/upbound/provider-azure/internal/controller/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/sql/mssqldatabaseextendedauditingpolicy"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/sql/mssqlfirewallrule"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/sql/mssqlserverdnsalias"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/sql/mssqlservertransparentdataencryption"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/sql/mssqlvirtualnetworkrule"
	accountstorage "github.com/upbound/provider-azure/internal/controller/storage/account"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/storage/container"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/storage/datalakegen2filesystem"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/storage/objectreplication"
	queuestorage "github.com/upbound/provider-azure/internal/controller/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/storage/share"
	tablestorage "github.com/upbound/provider-azure/internal/controller/storage/table"
	hpccache "github.com/upbound/provider-azure/internal/controller/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccachenfstarget"
	diskpool "github.com/upbound/provider-azure/internal/controller/storagepool/diskpool"
	storagesync "github.com/upbound/provider-azure/internal/controller/storagesync/storagesync"
	clusterstreamanalytics "github.com/upbound/provider-azure/internal/controller/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/streamanalytics/functionjavascriptuda"
	job "github.com/upbound/provider-azure/internal/controller/streamanalytics/job"
	managedprivateendpointstreamanalytics "github.com/upbound/provider-azure/internal/controller/streamanalytics/managedprivateendpoint"
	outputblob "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputblob"
	outputfunction "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputfunction"
	outputmssql "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputmssql"
	outputservicebusqueue "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputservicebustopic"
	outputsynapse "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputsynapse"
	referenceinputblob "github.com/upbound/provider-azure/internal/controller/streamanalytics/referenceinputblob"
	streaminputblob "github.com/upbound/provider-azure/internal/controller/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/upbound/provider-azure/internal/controller/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/upbound/provider-azure/internal/controller/streamanalytics/streaminputiothub"
	firewallrulesynapse "github.com/upbound/provider-azure/internal/controller/synapse/firewallrule"
	integrationruntimeazuresynapse "github.com/upbound/provider-azure/internal/controller/synapse/integrationruntimeazure"
	integrationruntimeselfhostedsynapse "github.com/upbound/provider-azure/internal/controller/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/internal/controller/synapse/linkedservice"
	managedprivateendpointsynapse "github.com/upbound/provider-azure/internal/controller/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/internal/controller/synapse/privatelinkhub"
	sparkpool "github.com/upbound/provider-azure/internal/controller/synapse/sparkpool"
	sqlpool "github.com/upbound/provider-azure/internal/controller/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolworkloadclassifier "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolworkloadclassifier"
	sqlpoolworkloadgroup "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolworkloadgroup"
	workspacesynapse "github.com/upbound/provider-azure/internal/controller/synapse/workspace"
	workspaceaadadmin "github.com/upbound/provider-azure/internal/controller/synapse/workspaceaadadmin"
	workspaceextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/synapse/workspaceextendedauditingpolicy"
	workspacesecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/synapse/workspacesecurityalertpolicy"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/web/functionapp"
	functionappslot "github.com/upbound/provider-azure/internal/controller/web/functionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/web/linuxwebapp"
	serviceplan "github.com/upbound/provider-azure/internal/controller/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/web/sourcecontroltoken"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoractionruleactiongroup.Setup,
		monitoractionrulesuppression.Setup,
		monitorsmartdetectoralertrule.Setup,
		server.Setup,
		api.Setup,
		apidiagnostic.Setup,
		apioperation.Setup,
		apioperationpolicy.Setup,
		apioperationtag.Setup,
		apipolicy.Setup,
		apirelease.Setup,
		apischema.Setup,
		apitag.Setup,
		apiversionset.Setup,
		authorizationserver.Setup,
		backend.Setup,
		certificate.Setup,
		diagnostic.Setup,
		emailtemplate.Setup,
		gateway.Setup,
		identityprovideraad.Setup,
		identityproviderfacebook.Setup,
		identityprovidergoogle.Setup,
		identityprovidermicrosoft.Setup,
		identityprovidertwitter.Setup,
		logger.Setup,
		management.Setup,
		namedvalue.Setup,
		notificationrecipientemail.Setup,
		notificationrecipientuser.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		product.Setup,
		productapi.Setup,
		productpolicy.Setup,
		rediscache.Setup,
		subscription.Setup,
		tag.Setup,
		user.Setup,
		springcloudactivedeployment.Setup,
		springcloudapp.Setup,
		springcloudappcosmosdbassociation.Setup,
		springcloudappmysqlassociation.Setup,
		springcloudappredisassociation.Setup,
		springcloudcertificate.Setup,
		springcloudcustomdomain.Setup,
		springcloudjavadeployment.Setup,
		springcloudservice.Setup,
		springcloudstorage.Setup,
		provider.Setup,
		managementlock.Setup,
		policydefinition.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcepolicyassignment.Setup,
		resourcepolicyexemption.Setup,
		roleassignment.Setup,
		roledefinition.Setup,
		subscriptionpolicyassignment.Setup,
		subscriptionpolicyexemption.Setup,
		account.Setup,
		credential.Setup,
		module.Setup,
		variablebool.Setup,
		variableint.Setup,
		variablestring.Setup,
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		subscriptionazure.Setup,
		cluster.Setup,
		botchannelalexa.Setup,
		botchanneldirectline.Setup,
		botchannelline.Setup,
		botchannelmsteams.Setup,
		botchannelslack.Setup,
		botchannelsms.Setup,
		botchannelsregistration.Setup,
		botchannelwebchat.Setup,
		botconnection.Setup,
		botwebapp.Setup,
		rediscachecache.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
		endpoint.Setup,
		frontdoorcustomdomain.Setup,
		frontdoorcustomdomainassociation.Setup,
		frontdoorendpoint.Setup,
		frontdoororigin.Setup,
		frontdoororigingroup.Setup,
		frontdoorprofile.Setup,
		frontdoorroute.Setup,
		frontdoorrule.Setup,
		frontdoorruleset.Setup,
		profile.Setup,
		accountcognitiveservices.Setup,
		service.Setup,
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
		ledger.Setup,
		budgetmanagementgroup.Setup,
		budgetresourcegroup.Setup,
		budgetsubscription.Setup,
		agentpool.Setup,
		containerconnectedregistry.Setup,
		registry.Setup,
		scopemap.Setup,
		token.Setup,
		webhook.Setup,
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		kubernetesfleetmanager.Setup,
		accountcosmosdb.Setup,
		cassandracluster.Setup,
		cassandradatacenter.Setup,
		cassandrakeyspace.Setup,
		cassandratable.Setup,
		gremlindatabase.Setup,
		gremlingraph.Setup,
		mongocollection.Setup,
		mongodatabase.Setup,
		sqlcontainer.Setup,
		sqldatabase.Setup,
		sqlfunction.Setup,
		sqlroleassignment.Setup,
		sqlroledefinition.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		table.Setup,
		resourcegroupcostmanagementexport.Setup,
		subscriptioncostmanagementexport.Setup,
		device.Setup,
		workspace.Setup,
		workspacecustomermanagedkey.Setup,
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
		databasemigrationproject.Setup,
		databasemigrationservice.Setup,
		backupinstanceblobstorage.Setup,
		backupinstancedisk.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydisk.Setup,
		backupvault.Setup,
		accountdatashare.Setup,
		datasetblobstorage.Setup,
		datasetdatalakegen2.Setup,
		datasetkustocluster.Setup,
		datasetkustodatabase.Setup,
		datashare.Setup,
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		serverdbformariadb.Setup,
		virtualnetworkrule.Setup,
		activedirectoryadministrator.Setup,
		configurationdbformysql.Setup,
		databasedbformysql.Setup,
		firewallruledbformysql.Setup,
		flexibledatabase.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverfirewallrule.Setup,
		serverdbformysql.Setup,
		virtualnetworkruledbformysql.Setup,
		activedirectoryadministratordbforpostgresql.Setup,
		configurationdbforpostgresql.Setup,
		databasedbforpostgresql.Setup,
		firewallruledbforpostgresql.Setup,
		flexibleserverdbforpostgresql.Setup,
		flexibleserverconfigurationdbforpostgresql.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallruledbforpostgresql.Setup,
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
		iothubdeviceupdateaccount.Setup,
		iothubdeviceupdateinstance.Setup,
		globalvmshutdownschedule.Setup,
		lab.Setup,
		linuxvirtualmachinedevtestlab.Setup,
		policydevtestlab.Setup,
		schedule.Setup,
		virtualnetwork.Setup,
		windowsvirtualmachinedevtestlab.Setup,
		instance.Setup,
		cloudelasticsearch.Setup,
		domain.Setup,
		domaintopic.Setup,
		eventsubscription.Setup,
		systemtopic.Setup,
		topic.Setup,
		authorizationrule.Setup,
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		namespaceauthorizationrule.Setup,
		namespacedisasterrecoveryconfig.Setup,
		healthcareservice.Setup,
		applicationinsights.Setup,
		applicationinsightsanalyticsitem.Setup,
		applicationinsightsapikey.Setup,
		applicationinsightssmartdetectionrule.Setup,
		monitoractiongroup.Setup,
		monitoractivitylogalert.Setup,
		monitorautoscalesetting.Setup,
		monitormetricalert.Setup,
		monitorprivatelinkscope.Setup,
		monitorprivatelinkscopedservice.Setup,
		monitorscheduledqueryrulesalert.Setup,
		monitorscheduledqueryruleslog.Setup,
		application.Setup,
		applicationnetworkruleset.Setup,
		accesspolicy.Setup,
		certificatekeyvault.Setup,
		certificateissuer.Setup,
		key.Setup,
		managedhardwaresecuritymodule.Setup,
		managedstorageaccount.Setup,
		managedstorageaccountsastokendefinition.Setup,
		secret.Setup,
		vault.Setup,
		attacheddatabaseconfiguration.Setup,
		clusterkusto.Setup,
		clustermanagedprivateendpoint.Setup,
		clusterprincipalassignment.Setup,
		databasekusto.Setup,
		databaseprincipalassignment.Setup,
		eventgriddataconnection.Setup,
		eventhubdataconnection.Setup,
		iothubdataconnection.Setup,
		appactioncustom.Setup,
		appactionhttp.Setup,
		appintegrationaccount.Setup,
		appintegrationaccountbatchconfiguration.Setup,
		appintegrationaccountschema.Setup,
		appintegrationaccountsession.Setup,
		apptriggercustom.Setup,
		apptriggerhttprequest.Setup,
		apptriggerrecurrence.Setup,
		appworkflow.Setup,
		integrationserviceenvironment.Setup,
		userassignedidentity.Setup,
		managementgroup.Setup,
		accountmaps.Setup,
		creator.Setup,
		marketplaceagreement.Setup,
		asset.Setup,
		assetfilter.Setup,
		contentkeypolicy.Setup,
		liveevent.Setup,
		liveeventoutput.Setup,
		servicesaccount.Setup,
		streamingendpoint.Setup,
		streaminglocator.Setup,
		streamingpolicy.Setup,
		transform.Setup,
		spatialanchorsaccount.Setup,
		accountnetapp.Setup,
		pool.Setup,
		snapshotnetapp.Setup,
		snapshotpolicy.Setup,
		volume.Setup,
		applicationgateway.Setup,
		applicationsecuritygroup.Setup,
		connectionmonitor.Setup,
		ddosprotectionplan.Setup,
		dnsaaaarecord.Setup,
		dnsarecord.Setup,
		dnscaarecord.Setup,
		dnscnamerecord.Setup,
		dnsmxrecord.Setup,
		dnsnsrecord.Setup,
		dnsptrrecord.Setup,
		dnssrvrecord.Setup,
		dnstxtrecord.Setup,
		dnszone.Setup,
		expressroutecircuit.Setup,
		expressroutecircuitauthorization.Setup,
		expressroutecircuitconnection.Setup,
		expressroutecircuitpeering.Setup,
		expressrouteconnection.Setup,
		expressroutegateway.Setup,
		expressrouteport.Setup,
		firewall.Setup,
		firewallapplicationrulecollection.Setup,
		firewallnatrulecollection.Setup,
		firewallnetworkrulecollection.Setup,
		firewallpolicy.Setup,
		firewallpolicyrulecollectiongroup.Setup,
		frontdoor.Setup,
		frontdoorcustomhttpsconfiguration.Setup,
		frontdoorfirewallpolicy.Setup,
		frontdoorrulesengine.Setup,
		ipgroup.Setup,
		loadbalancer.Setup,
		loadbalancerbackendaddresspool.Setup,
		loadbalancerbackendaddresspooladdress.Setup,
		loadbalancernatpool.Setup,
		loadbalancernatrule.Setup,
		loadbalanceroutboundrule.Setup,
		loadbalancerprobe.Setup,
		loadbalancerrule.Setup,
		localnetworkgateway.Setup,
		natgateway.Setup,
		natgatewaypublicipassociation.Setup,
		natgatewaypublicipprefixassociation.Setup,
		networkinterface.Setup,
		networkinterfaceapplicationsecuritygroupassociation.Setup,
		networkinterfacebackendaddresspoolassociation.Setup,
		networkinterfacenatruleassociation.Setup,
		networkinterfacesecuritygroupassociation.Setup,
		packetcapture.Setup,
		pointtositevpngateway.Setup,
		privatednsaaaarecord.Setup,
		privatednsarecord.Setup,
		privatednscnamerecord.Setup,
		privatednsmxrecord.Setup,
		privatednsptrrecord.Setup,
		privatednssrvrecord.Setup,
		privatednstxtrecord.Setup,
		privatednszone.Setup,
		privatednszonevirtualnetworklink.Setup,
		privateendpoint.Setup,
		privatelinkservice.Setup,
		profilenetwork.Setup,
		publicip.Setup,
		publicipprefix.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securityrule.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		trafficmanagerprofile.Setup,
		virtualhub.Setup,
		virtualnetworknetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
		vpngateway.Setup,
		vpngatewayconnection.Setup,
		vpnserverconfiguration.Setup,
		vpnsite.Setup,
		watcher.Setup,
		watcherflowlog.Setup,
		webapplicationfirewallpolicy.Setup,
		authorizationrulenotificationhubs.Setup,
		notificationhub.Setup,
		notificationhubnamespace.Setup,
		loganalyticsdataexportrule.Setup,
		loganalyticsdatasourcewindowsevent.Setup,
		loganalyticsdatasourcewindowsperformancecounter.Setup,
		loganalyticslinkedservice.Setup,
		loganalyticslinkedstorageaccount.Setup,
		loganalyticssavedsearch.Setup,
		workspaceoperationalinsights.Setup,
		loganalyticssolution.Setup,
		subscriptionpolicyremediation.Setup,
		dashboard.Setup,
		powerbiembedded.Setup,
		providerconfig.Setup,
		accountpurview.Setup,
		backupcontainerstorageaccount.Setup,
		backuppolicyfileshare.Setup,
		backuppolicyvm.Setup,
		backuppolicyvmworkload.Setup,
		backupprotectedfileshare.Setup,
		backupprotectedvm.Setup,
		siterecoveryfabric.Setup,
		vaultrecoveryservices.Setup,
		resourcegrouptemplatedeployment.Setup,
		subscriptiontemplatedeployment.Setup,
		servicesearch.Setup,
		advancedthreatprotection.Setup,
		iotsecuritydevicegroup.Setup,
		iotsecuritysolution.Setup,
		securitycenterassessment.Setup,
		securitycenterassessmentpolicy.Setup,
		securitycenterautoprovisioning.Setup,
		securitycentercontact.Setup,
		securitycenterservervulnerabilityassessment.Setup,
		securitycenterservervulnerabilityassessmentvirtualmachine.Setup,
		securitycentersetting.Setup,
		securitycentersubscriptionpricing.Setup,
		securitycenterworkspace.Setup,
		sentinelautomationrule.Setup,
		sentinelwatchlist.Setup,
		namespaceauthorizationruleservicebus.Setup,
		namespacedisasterrecoveryconfigservicebus.Setup,
		namespacenetworkruleset.Setup,
		queue.Setup,
		queueauthorizationrule.Setup,
		servicebusnamespace.Setup,
		subscriptionservicebus.Setup,
		subscriptionrule.Setup,
		topicservicebus.Setup,
		topicauthorizationrule.Setup,
		clusterservicefabric.Setup,
		managedcluster.Setup,
		networkacl.Setup,
		servicesignalrservice.Setup,
		managedapplicationdefinition.Setup,
		activedirectoryadministratorsql.Setup,
		databasesql.Setup,
		elasticpool.Setup,
		firewallrulesql.Setup,
		mssqldatabase.Setup,
		mssqldatabaseextendedauditingpolicy.Setup,
		mssqlfailovergroup.Setup,
		mssqlfirewallrule.Setup,
		mssqlmanageddatabase.Setup,
		mssqlmanagedinstance.Setup,
		mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		mssqlmanagedinstancefailovergroup.Setup,
		mssqlmanagedinstancevulnerabilityassessment.Setup,
		mssqloutboundfirewallrule.Setup,
		mssqlserver.Setup,
		mssqlserverdnsalias.Setup,
		mssqlserversecurityalertpolicy.Setup,
		mssqlservertransparentdataencryption.Setup,
		mssqlvirtualnetworkrule.Setup,
		accountstorage.Setup,
		accountnetworkrules.Setup,
		blob.Setup,
		blobinventorypolicy.Setup,
		container.Setup,
		datalakegen2filesystem.Setup,
		encryptionscope.Setup,
		managementpolicy.Setup,
		objectreplication.Setup,
		queuestorage.Setup,
		share.Setup,
		tablestorage.Setup,
		hpccache.Setup,
		hpccacheaccesspolicy.Setup,
		hpccacheblobnfstarget.Setup,
		hpccacheblobtarget.Setup,
		hpccachenfstarget.Setup,
		diskpool.Setup,
		storagesync.Setup,
		clusterstreamanalytics.Setup,
		functionjavascriptuda.Setup,
		job.Setup,
		managedprivateendpointstreamanalytics.Setup,
		outputblob.Setup,
		outputfunction.Setup,
		outputmssql.Setup,
		outputservicebusqueue.Setup,
		outputservicebustopic.Setup,
		outputsynapse.Setup,
		referenceinputblob.Setup,
		streaminputblob.Setup,
		streaminputeventhub.Setup,
		streaminputiothub.Setup,
		firewallrulesynapse.Setup,
		integrationruntimeazuresynapse.Setup,
		integrationruntimeselfhostedsynapse.Setup,
		linkedservice.Setup,
		managedprivateendpointsynapse.Setup,
		privatelinkhub.Setup,
		sparkpool.Setup,
		sqlpool.Setup,
		sqlpoolextendedauditingpolicy.Setup,
		sqlpoolsecurityalertpolicy.Setup,
		sqlpoolworkloadclassifier.Setup,
		sqlpoolworkloadgroup.Setup,
		workspacesynapse.Setup,
		workspaceaadadmin.Setup,
		workspaceextendedauditingpolicy.Setup,
		workspacesecurityalertpolicy.Setup,
		appserviceplan.Setup,
		functionapp.Setup,
		functionappslot.Setup,
		linuxwebapp.Setup,
		serviceplan.Setup,
		sourcecontroltoken.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
