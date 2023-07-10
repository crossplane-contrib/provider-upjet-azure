/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	monitoractionruleactiongroup "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoractionruleactiongroup"
	monitoractionrulesuppression "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoractionrulesuppression"
	monitoralertprocessingruleactiongroup "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoralertprocessingruleactiongroup"
	monitoralertprocessingrulesuppression "github.com/upbound/provider-azure/internal/controller/alertsmanagement/monitoralertprocessingrulesuppression"
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
	gatewayapi "github.com/upbound/provider-azure/internal/controller/apimanagement/gatewayapi"
	globalschema "github.com/upbound/provider-azure/internal/controller/apimanagement/globalschema"
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
	producttag "github.com/upbound/provider-azure/internal/controller/apimanagement/producttag"
	rediscache "github.com/upbound/provider-azure/internal/controller/apimanagement/rediscache"
	subscription "github.com/upbound/provider-azure/internal/controller/apimanagement/subscription"
	tag "github.com/upbound/provider-azure/internal/controller/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/apimanagement/user"
	configuration "github.com/upbound/provider-azure/internal/controller/appconfiguration/configuration"
	springcloudaccelerator "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudaccelerator"
	springcloudactivedeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudactivedeployment"
	springcloudapiportal "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudapiportal"
	springcloudapiportalcustomdomain "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudapiportalcustomdomain"
	springcloudapp "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappredisassociation"
	springcloudbuilddeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudbuilddeployment"
	springcloudbuilder "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudbuilder"
	springcloudbuildpackbinding "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudbuildpackbinding"
	springcloudcertificate "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcertificate"
	springcloudconfigurationservice "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudconfigurationservice"
	springcloudcontainerdeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcontainerdeployment"
	springcloudcustomdomain "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcustomdomain"
	springcloudcustomizedaccelerator "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcustomizedaccelerator"
	springclouddevtoolportal "github.com/upbound/provider-azure/internal/controller/appplatform/springclouddevtoolportal"
	springcloudgateway "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudgateway"
	springcloudgatewaycustomdomain "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudgatewaycustomdomain"
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
	connection "github.com/upbound/provider-azure/internal/controller/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/internal/controller/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/internal/controller/automation/connectiontype"
	credential "github.com/upbound/provider-azure/internal/controller/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/internal/controller/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/internal/controller/automation/module"
	runbook "github.com/upbound/provider-azure/internal/controller/automation/runbook"
	schedule "github.com/upbound/provider-azure/internal/controller/automation/schedule"
	variablebool "github.com/upbound/provider-azure/internal/controller/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/internal/controller/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/internal/controller/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/automation/variablestring"
	webhook "github.com/upbound/provider-azure/internal/controller/automation/webhook"
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
	appservicecertificateorder "github.com/upbound/provider-azure/internal/controller/certificateregistration/appservicecertificateorder"
	accountcognitiveservices "github.com/upbound/provider-azure/internal/controller/cognitiveservices/account"
	service "github.com/upbound/provider-azure/internal/controller/communication/service"
	availabilityset "github.com/upbound/provider-azure/internal/controller/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/compute/snapshot"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/compute/sshpublickey"
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
	tokenpassword "github.com/upbound/provider-azure/internal/controller/containerregistry/tokenpassword"
	webhookcontainerregistry "github.com/upbound/provider-azure/internal/controller/containerregistry/webhook"
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
	sqldedicatedgateway "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqldedicatedgateway"
	sqlfunction "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/provider-azure/internal/controller/cosmosdb/sqltrigger"
	table "github.com/upbound/provider-azure/internal/controller/cosmosdb/table"
	costanomalyalert "github.com/upbound/provider-azure/internal/controller/costmanagement/costanomalyalert"
	resourcegroupcostmanagementexport "github.com/upbound/provider-azure/internal/controller/costmanagement/resourcegroupcostmanagementexport"
	subscriptioncostmanagementexport "github.com/upbound/provider-azure/internal/controller/costmanagement/subscriptioncostmanagementexport"
	customprovider "github.com/upbound/provider-azure/internal/controller/customproviders/customprovider"
	device "github.com/upbound/provider-azure/internal/controller/databoxedge/device"
	accessconnector "github.com/upbound/provider-azure/internal/controller/databricks/accessconnector"
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
	databasemigrationproject "github.com/upbound/provider-azure/internal/controller/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/upbound/provider-azure/internal/controller/datamigration/databasemigrationservice"
	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancedisk"
	backupinstancepostgresql "github.com/upbound/provider-azure/internal/controller/dataprotection/backupinstancepostgresql"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicydisk"
	backuppolicypostgresql "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicypostgresql"
	backupvault "github.com/upbound/provider-azure/internal/controller/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/internal/controller/dataprotection/resourceguard"
	accountdatashare "github.com/upbound/provider-azure/internal/controller/datashare/account"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/datashare/datashare"
	configurationdbformariadb "github.com/upbound/provider-azure/internal/controller/dbformariadb/configuration"
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
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/devices/iothubcertificate"
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
	scheduledevtestlab "github.com/upbound/provider-azure/internal/controller/devtestlab/schedule"
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
	namespaceschemagroup "github.com/upbound/provider-azure/internal/controller/eventhub/namespaceschemagroup"
	serverfluidrelay "github.com/upbound/provider-azure/internal/controller/fluidrelay/server"
	policyvirtualmachineconfigurationassignment "github.com/upbound/provider-azure/internal/controller/guestconfiguration/policyvirtualmachineconfigurationassignment"
	hadoopcluster "github.com/upbound/provider-azure/internal/controller/hdinsight/hadoopcluster"
	hbasecluster "github.com/upbound/provider-azure/internal/controller/hdinsight/hbasecluster"
	interactivequerycluster "github.com/upbound/provider-azure/internal/controller/hdinsight/interactivequerycluster"
	kafkacluster "github.com/upbound/provider-azure/internal/controller/hdinsight/kafkacluster"
	sparkcluster "github.com/upbound/provider-azure/internal/controller/hdinsight/sparkcluster"
	healthbot "github.com/upbound/provider-azure/internal/controller/healthbot/healthbot"
	healthcaredicomservice "github.com/upbound/provider-azure/internal/controller/healthcareapis/healthcaredicomservice"
	healthcarefhirservice "github.com/upbound/provider-azure/internal/controller/healthcareapis/healthcarefhirservice"
	healthcaremedtechservice "github.com/upbound/provider-azure/internal/controller/healthcareapis/healthcaremedtechservice"
	healthcaremedtechservicefhirdestination "github.com/upbound/provider-azure/internal/controller/healthcareapis/healthcaremedtechservicefhirdestination"
	healthcareservice "github.com/upbound/provider-azure/internal/controller/healthcareapis/healthcareservice"
	healthcareworkspace "github.com/upbound/provider-azure/internal/controller/healthcareapis/healthcareworkspace"
	applicationinsights "github.com/upbound/provider-azure/internal/controller/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/internal/controller/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/internal/controller/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/internal/controller/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/internal/controller/insights/monitordatacollectionruleassociation"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/insights/monitorscheduledqueryruleslog"
	application "github.com/upbound/provider-azure/internal/controller/iotcentral/application"
	applicationnetworkruleset "github.com/upbound/provider-azure/internal/controller/iotcentral/applicationnetworkruleset"
	accesspolicy "github.com/upbound/provider-azure/internal/controller/keyvault/accesspolicy"
	certificatekeyvault "github.com/upbound/provider-azure/internal/controller/keyvault/certificate"
	certificatecontacts "github.com/upbound/provider-azure/internal/controller/keyvault/certificatecontacts"
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
	labservicelab "github.com/upbound/provider-azure/internal/controller/labservices/labservicelab"
	labserviceplan "github.com/upbound/provider-azure/internal/controller/labservices/labserviceplan"
	appactioncustom "github.com/upbound/provider-azure/internal/controller/logic/appactioncustom"
	appactionhttp "github.com/upbound/provider-azure/internal/controller/logic/appactionhttp"
	appintegrationaccount "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccount"
	appintegrationaccountbatchconfiguration "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountpartner "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccountpartner"
	appintegrationaccountschema "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/upbound/provider-azure/internal/controller/logic/appintegrationaccountsession"
	apptriggercustom "github.com/upbound/provider-azure/internal/controller/logic/apptriggercustom"
	apptriggerhttprequest "github.com/upbound/provider-azure/internal/controller/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/upbound/provider-azure/internal/controller/logic/apptriggerrecurrence"
	appworkflow "github.com/upbound/provider-azure/internal/controller/logic/appworkflow"
	integrationserviceenvironment "github.com/upbound/provider-azure/internal/controller/logic/integrationserviceenvironment"
	monitor "github.com/upbound/provider-azure/internal/controller/logz/monitor"
	subaccount "github.com/upbound/provider-azure/internal/controller/logz/subaccount"
	subaccounttagrule "github.com/upbound/provider-azure/internal/controller/logz/subaccounttagrule"
	tagrule "github.com/upbound/provider-azure/internal/controller/logz/tagrule"
	computecluster "github.com/upbound/provider-azure/internal/controller/machinelearningservices/computecluster"
	computeinstance "github.com/upbound/provider-azure/internal/controller/machinelearningservices/computeinstance"
	synapsespark "github.com/upbound/provider-azure/internal/controller/machinelearningservices/synapsespark"
	workspacemachinelearningservices "github.com/upbound/provider-azure/internal/controller/machinelearningservices/workspace"
	maintenanceassignmentdedicatedhost "github.com/upbound/provider-azure/internal/controller/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceconfiguration "github.com/upbound/provider-azure/internal/controller/maintenance/maintenanceconfiguration"
	federatedidentitycredential "github.com/upbound/provider-azure/internal/controller/managedidentity/federatedidentitycredential"
	userassignedidentity "github.com/upbound/provider-azure/internal/controller/managedidentity/userassignedidentity"
	managementgroup "github.com/upbound/provider-azure/internal/controller/management/managementgroup"
	managementgroupsubscriptionassociation "github.com/upbound/provider-azure/internal/controller/management/managementgroupsubscriptionassociation"
	accountmaps "github.com/upbound/provider-azure/internal/controller/maps/account"
	creator "github.com/upbound/provider-azure/internal/controller/maps/creator"
	marketplaceagreement "github.com/upbound/provider-azure/internal/controller/marketplaceordering/marketplaceagreement"
	asset "github.com/upbound/provider-azure/internal/controller/media/asset"
	assetfilter "github.com/upbound/provider-azure/internal/controller/media/assetfilter"
	contentkeypolicy "github.com/upbound/provider-azure/internal/controller/media/contentkeypolicy"
	job "github.com/upbound/provider-azure/internal/controller/media/job"
	liveevent "github.com/upbound/provider-azure/internal/controller/media/liveevent"
	liveeventoutput "github.com/upbound/provider-azure/internal/controller/media/liveeventoutput"
	servicesaccount "github.com/upbound/provider-azure/internal/controller/media/servicesaccount"
	servicesaccountfilter "github.com/upbound/provider-azure/internal/controller/media/servicesaccountfilter"
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
	manager "github.com/upbound/provider-azure/internal/controller/network/manager"
	managermanagementgroupconnection "github.com/upbound/provider-azure/internal/controller/network/managermanagementgroupconnection"
	managernetworkgroup "github.com/upbound/provider-azure/internal/controller/network/managernetworkgroup"
	managerstaticmember "github.com/upbound/provider-azure/internal/controller/network/managerstaticmember"
	managersubscriptionconnection "github.com/upbound/provider-azure/internal/controller/network/managersubscriptionconnection"
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
	privatednsresolver "github.com/upbound/provider-azure/internal/controller/network/privatednsresolver"
	privatednssrvrecord "github.com/upbound/provider-azure/internal/controller/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/provider-azure/internal/controller/network/privatednstxtrecord"
	privatednszone "github.com/upbound/provider-azure/internal/controller/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/provider-azure/internal/controller/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/provider-azure/internal/controller/network/privateendpoint"
	privateendpointapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/network/privateendpointapplicationsecuritygroupassociation"
	privatelinkservice "github.com/upbound/provider-azure/internal/controller/network/privatelinkservice"
	profilenetwork "github.com/upbound/provider-azure/internal/controller/network/profile"
	publicip "github.com/upbound/provider-azure/internal/controller/network/publicip"
	publicipprefix "github.com/upbound/provider-azure/internal/controller/network/publicipprefix"
	route "github.com/upbound/provider-azure/internal/controller/network/route"
	routefilter "github.com/upbound/provider-azure/internal/controller/network/routefilter"
	routemap "github.com/upbound/provider-azure/internal/controller/network/routemap"
	routeserver "github.com/upbound/provider-azure/internal/controller/network/routeserver"
	routeserverbgpconnection "github.com/upbound/provider-azure/internal/controller/network/routeserverbgpconnection"
	routetable "github.com/upbound/provider-azure/internal/controller/network/routetable"
	securitygroup "github.com/upbound/provider-azure/internal/controller/network/securitygroup"
	securityrule "github.com/upbound/provider-azure/internal/controller/network/securityrule"
	subnet "github.com/upbound/provider-azure/internal/controller/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/provider-azure/internal/controller/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/provider-azure/internal/controller/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/provider-azure/internal/controller/network/subnetserviceendpointstoragepolicy"
	trafficmanagerazureendpoint "github.com/upbound/provider-azure/internal/controller/network/trafficmanagerazureendpoint"
	trafficmanagerexternalendpoint "github.com/upbound/provider-azure/internal/controller/network/trafficmanagerexternalendpoint"
	trafficmanagernestedendpoint "github.com/upbound/provider-azure/internal/controller/network/trafficmanagernestedendpoint"
	trafficmanagerprofile "github.com/upbound/provider-azure/internal/controller/network/trafficmanagerprofile"
	virtualhub "github.com/upbound/provider-azure/internal/controller/network/virtualhub"
	virtualhubconnection "github.com/upbound/provider-azure/internal/controller/network/virtualhubconnection"
	virtualhubip "github.com/upbound/provider-azure/internal/controller/network/virtualhubip"
	virtualhubroutetable "github.com/upbound/provider-azure/internal/controller/network/virtualhubroutetable"
	virtualhubroutetableroute "github.com/upbound/provider-azure/internal/controller/network/virtualhubroutetableroute"
	virtualhubsecuritypartnerprovider "github.com/upbound/provider-azure/internal/controller/network/virtualhubsecuritypartnerprovider"
	virtualnetworknetwork "github.com/upbound/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/internal/controller/network/virtualwan"
	vpngateway "github.com/upbound/provider-azure/internal/controller/network/vpngateway"
	vpngatewayconnection "github.com/upbound/provider-azure/internal/controller/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/upbound/provider-azure/internal/controller/network/vpnserverconfiguration"
	vpnserverconfigurationpolicygroup "github.com/upbound/provider-azure/internal/controller/network/vpnserverconfigurationpolicygroup"
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
	loganalyticsquerypack "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticsquerypack"
	loganalyticsquerypackquery "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticsquerypackquery"
	loganalyticssavedsearch "github.com/upbound/provider-azure/internal/controller/operationalinsights/loganalyticssavedsearch"
	workspaceoperationalinsights "github.com/upbound/provider-azure/internal/controller/operationalinsights/workspace"
	loganalyticssolution "github.com/upbound/provider-azure/internal/controller/operationsmanagement/loganalyticssolution"
	contactprofile "github.com/upbound/provider-azure/internal/controller/orbital/contactprofile"
	spacecraft "github.com/upbound/provider-azure/internal/controller/orbital/spacecraft"
	resourcepolicyremediation "github.com/upbound/provider-azure/internal/controller/policyinsights/resourcepolicyremediation"
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
	siterecoverynetworkmapping "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicationpolicy "github.com/upbound/provider-azure/internal/controller/recoveryservices/siterecoveryreplicationpolicy"
	vaultrecoveryservices "github.com/upbound/provider-azure/internal/controller/recoveryservices/vault"
	eventrelaynamespace "github.com/upbound/provider-azure/internal/controller/relay/eventrelaynamespace"
	hybridconnection "github.com/upbound/provider-azure/internal/controller/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/upbound/provider-azure/internal/controller/relay/hybridconnectionauthorizationrule"
	namespaceauthorizationrulerelay "github.com/upbound/provider-azure/internal/controller/relay/namespaceauthorizationrule"
	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/resources/subscriptiontemplatedeployment"
	servicesearch "github.com/upbound/provider-azure/internal/controller/search/service"
	sharedprivatelinkservice "github.com/upbound/provider-azure/internal/controller/search/sharedprivatelinkservice"
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
	sentinelalertrulefusion "github.com/upbound/provider-azure/internal/controller/securityinsights/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/upbound/provider-azure/internal/controller/securityinsights/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/upbound/provider-azure/internal/controller/securityinsights/sentinelalertrulemssecurityincident"
	sentinelautomationrule "github.com/upbound/provider-azure/internal/controller/securityinsights/sentinelautomationrule"
	sentineldataconnectoriot "github.com/upbound/provider-azure/internal/controller/securityinsights/sentineldataconnectoriot"
	sentinelloganalyticsworkspaceonboarding "github.com/upbound/provider-azure/internal/controller/securityinsights/sentinelloganalyticsworkspaceonboarding"
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
	springcloudconnection "github.com/upbound/provider-azure/internal/controller/servicelinker/springcloudconnection"
	networkacl "github.com/upbound/provider-azure/internal/controller/signalrservice/networkacl"
	servicesignalrservice "github.com/upbound/provider-azure/internal/controller/signalrservice/service"
	signalrsharedprivatelinkresource "github.com/upbound/provider-azure/internal/controller/signalrservice/signalrsharedprivatelinkresource"
	webpubsub "github.com/upbound/provider-azure/internal/controller/signalrservice/webpubsub"
	webpubsubhub "github.com/upbound/provider-azure/internal/controller/signalrservice/webpubsubhub"
	webpubsubnetworkacl "github.com/upbound/provider-azure/internal/controller/signalrservice/webpubsubnetworkacl"
	managedapplicationdefinition "github.com/upbound/provider-azure/internal/controller/solutions/managedapplicationdefinition"
	cloudapplicationliveview "github.com/upbound/provider-azure/internal/controller/spring/cloudapplicationliveview"
	mssqldatabase "github.com/upbound/provider-azure/internal/controller/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/upbound/provider-azure/internal/controller/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/upbound/provider-azure/internal/controller/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/sql/mssqlfirewallrule"
	mssqljobagent "github.com/upbound/provider-azure/internal/controller/sql/mssqljobagent"
	mssqljobcredential "github.com/upbound/provider-azure/internal/controller/sql/mssqljobcredential"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "github.com/upbound/provider-azure/internal/controller/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/sql/mssqlvirtualnetworkrule"
	accountstorage "github.com/upbound/provider-azure/internal/controller/storage/account"
	accountlocaluser "github.com/upbound/provider-azure/internal/controller/storage/accountlocaluser"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/storage/container"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/storage/datalakegen2filesystem"
	datalakegen2path "github.com/upbound/provider-azure/internal/controller/storage/datalakegen2path"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/storage/objectreplication"
	queuestorage "github.com/upbound/provider-azure/internal/controller/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/storage/share"
	sharedirectory "github.com/upbound/provider-azure/internal/controller/storage/sharedirectory"
	tablestorage "github.com/upbound/provider-azure/internal/controller/storage/table"
	tableentity "github.com/upbound/provider-azure/internal/controller/storage/tableentity"
	hpccache "github.com/upbound/provider-azure/internal/controller/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccachenfstarget"
	diskpool "github.com/upbound/provider-azure/internal/controller/storagepool/diskpool"
	storagesync "github.com/upbound/provider-azure/internal/controller/storagesync/storagesync"
	clusterstreamanalytics "github.com/upbound/provider-azure/internal/controller/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/streamanalytics/functionjavascriptuda"
	jobstreamanalytics "github.com/upbound/provider-azure/internal/controller/streamanalytics/job"
	managedprivateendpointstreamanalytics "github.com/upbound/provider-azure/internal/controller/streamanalytics/managedprivateendpoint"
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
	firewallrulesynapse "github.com/upbound/provider-azure/internal/controller/synapse/firewallrule"
	integrationruntimeazuresynapse "github.com/upbound/provider-azure/internal/controller/synapse/integrationruntimeazure"
	integrationruntimeselfhostedsynapse "github.com/upbound/provider-azure/internal/controller/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/internal/controller/synapse/linkedservice"
	managedprivateendpointsynapse "github.com/upbound/provider-azure/internal/controller/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/internal/controller/synapse/privatelinkhub"
	roleassignmentsynapse "github.com/upbound/provider-azure/internal/controller/synapse/roleassignment"
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
	workspacesqlaadadmin "github.com/upbound/provider-azure/internal/controller/synapse/workspacesqlaadadmin"
	workspacevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/synapse/workspacevulnerabilityassessment"
	eventsourceeventhub "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/eventsourceeventhub"
	eventsourceiothub "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/eventsourceiothub"
	gen2environment "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/gen2environment"
	referencedataset "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/referencedataset"
	standardenvironment "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/standardenvironment"
	appactiveslot "github.com/upbound/provider-azure/internal/controller/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/internal/controller/web/apphybridconnection"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/web/functionapp"
	functionappactiveslot "github.com/upbound/provider-azure/internal/controller/web/functionappactiveslot"
	functionappfunction "github.com/upbound/provider-azure/internal/controller/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/internal/controller/web/functionapphybridconnection"
	functionappslot "github.com/upbound/provider-azure/internal/controller/web/functionappslot"
	linuxfunctionapp "github.com/upbound/provider-azure/internal/controller/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/internal/controller/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/internal/controller/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/internal/controller/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/web/sourcecontroltoken"
	staticsite "github.com/upbound/provider-azure/internal/controller/web/staticsite"
	windowsfunctionapp "github.com/upbound/provider-azure/internal/controller/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/internal/controller/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/internal/controller/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/internal/controller/web/windowswebappslot"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoractionruleactiongroup.Setup,
		monitoractionrulesuppression.Setup,
		monitoralertprocessingruleactiongroup.Setup,
		monitoralertprocessingrulesuppression.Setup,
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
		gatewayapi.Setup,
		globalschema.Setup,
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
		producttag.Setup,
		rediscache.Setup,
		subscription.Setup,
		tag.Setup,
		user.Setup,
		configuration.Setup,
		springcloudaccelerator.Setup,
		springcloudactivedeployment.Setup,
		springcloudapiportal.Setup,
		springcloudapiportalcustomdomain.Setup,
		springcloudapp.Setup,
		springcloudappcosmosdbassociation.Setup,
		springcloudappmysqlassociation.Setup,
		springcloudappredisassociation.Setup,
		springcloudbuilddeployment.Setup,
		springcloudbuilder.Setup,
		springcloudbuildpackbinding.Setup,
		springcloudcertificate.Setup,
		springcloudconfigurationservice.Setup,
		springcloudcontainerdeployment.Setup,
		springcloudcustomdomain.Setup,
		springcloudcustomizedaccelerator.Setup,
		springclouddevtoolportal.Setup,
		springcloudgateway.Setup,
		springcloudgatewaycustomdomain.Setup,
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
		connection.Setup,
		connectionclassiccertificate.Setup,
		connectiontype.Setup,
		credential.Setup,
		hybridrunbookworkergroup.Setup,
		module.Setup,
		runbook.Setup,
		schedule.Setup,
		variablebool.Setup,
		variabledatetime.Setup,
		variableint.Setup,
		variablestring.Setup,
		webhook.Setup,
		resourcegroup.Setup,
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		resourceproviderregistration.Setup,
		subscriptionazure.Setup,
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
		appservicecertificateorder.Setup,
		accountcognitiveservices.Setup,
		service.Setup,
		availabilityset.Setup,
		capacityreservation.Setup,
		capacityreservationgroup.Setup,
		dedicatedhost.Setup,
		diskaccess.Setup,
		diskencryptionset.Setup,
		galleryapplication.Setup,
		galleryapplicationversion.Setup,
		image.Setup,
		linuxvirtualmachine.Setup,
		linuxvirtualmachinescaleset.Setup,
		manageddisk.Setup,
		manageddisksastoken.Setup,
		orchestratedvirtualmachinescaleset.Setup,
		proximityplacementgroup.Setup,
		sharedimage.Setup,
		sharedimagegallery.Setup,
		snapshot.Setup,
		sshpublickey.Setup,
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
		tokenpassword.Setup,
		webhookcontainerregistry.Setup,
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
		sqldedicatedgateway.Setup,
		sqlfunction.Setup,
		sqlroleassignment.Setup,
		sqlroledefinition.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		table.Setup,
		costanomalyalert.Setup,
		resourcegroupcostmanagementexport.Setup,
		subscriptioncostmanagementexport.Setup,
		customprovider.Setup,
		device.Setup,
		accessconnector.Setup,
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
		databasemigrationproject.Setup,
		databasemigrationservice.Setup,
		backupinstanceblobstorage.Setup,
		backupinstancedisk.Setup,
		backupinstancepostgresql.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydisk.Setup,
		backuppolicypostgresql.Setup,
		backupvault.Setup,
		resourceguard.Setup,
		accountdatashare.Setup,
		datasetblobstorage.Setup,
		datasetdatalakegen2.Setup,
		datasetkustocluster.Setup,
		datasetkustodatabase.Setup,
		datashare.Setup,
		configurationdbformariadb.Setup,
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
		iothubcertificate.Setup,
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
		scheduledevtestlab.Setup,
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
		namespaceschemagroup.Setup,
		serverfluidrelay.Setup,
		policyvirtualmachineconfigurationassignment.Setup,
		hadoopcluster.Setup,
		hbasecluster.Setup,
		interactivequerycluster.Setup,
		kafkacluster.Setup,
		sparkcluster.Setup,
		healthbot.Setup,
		healthcaredicomservice.Setup,
		healthcarefhirservice.Setup,
		healthcaremedtechservice.Setup,
		healthcaremedtechservicefhirdestination.Setup,
		healthcareservice.Setup,
		healthcareworkspace.Setup,
		applicationinsights.Setup,
		applicationinsightsanalyticsitem.Setup,
		applicationinsightsapikey.Setup,
		applicationinsightssmartdetectionrule.Setup,
		applicationinsightsstandardwebtest.Setup,
		applicationinsightswebtest.Setup,
		applicationinsightsworkbook.Setup,
		applicationinsightsworkbooktemplate.Setup,
		monitoractiongroup.Setup,
		monitoractivitylogalert.Setup,
		monitorautoscalesetting.Setup,
		monitordatacollectionendpoint.Setup,
		monitordatacollectionrule.Setup,
		monitordatacollectionruleassociation.Setup,
		monitormetricalert.Setup,
		monitorprivatelinkscope.Setup,
		monitorprivatelinkscopedservice.Setup,
		monitorscheduledqueryrulesalert.Setup,
		monitorscheduledqueryrulesalertv2.Setup,
		monitorscheduledqueryruleslog.Setup,
		application.Setup,
		applicationnetworkruleset.Setup,
		accesspolicy.Setup,
		certificatekeyvault.Setup,
		certificatecontacts.Setup,
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
		labservicelab.Setup,
		labserviceplan.Setup,
		appactioncustom.Setup,
		appactionhttp.Setup,
		appintegrationaccount.Setup,
		appintegrationaccountbatchconfiguration.Setup,
		appintegrationaccountpartner.Setup,
		appintegrationaccountschema.Setup,
		appintegrationaccountsession.Setup,
		apptriggercustom.Setup,
		apptriggerhttprequest.Setup,
		apptriggerrecurrence.Setup,
		appworkflow.Setup,
		integrationserviceenvironment.Setup,
		monitor.Setup,
		subaccount.Setup,
		subaccounttagrule.Setup,
		tagrule.Setup,
		computecluster.Setup,
		computeinstance.Setup,
		synapsespark.Setup,
		workspacemachinelearningservices.Setup,
		maintenanceassignmentdedicatedhost.Setup,
		maintenanceassignmentvirtualmachine.Setup,
		maintenanceconfiguration.Setup,
		federatedidentitycredential.Setup,
		userassignedidentity.Setup,
		managementgroup.Setup,
		managementgroupsubscriptionassociation.Setup,
		accountmaps.Setup,
		creator.Setup,
		marketplaceagreement.Setup,
		asset.Setup,
		assetfilter.Setup,
		contentkeypolicy.Setup,
		job.Setup,
		liveevent.Setup,
		liveeventoutput.Setup,
		servicesaccount.Setup,
		servicesaccountfilter.Setup,
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
		manager.Setup,
		managermanagementgroupconnection.Setup,
		managernetworkgroup.Setup,
		managerstaticmember.Setup,
		managersubscriptionconnection.Setup,
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
		privatednsresolver.Setup,
		privatednssrvrecord.Setup,
		privatednstxtrecord.Setup,
		privatednszone.Setup,
		privatednszonevirtualnetworklink.Setup,
		privateendpoint.Setup,
		privateendpointapplicationsecuritygroupassociation.Setup,
		privatelinkservice.Setup,
		profilenetwork.Setup,
		publicip.Setup,
		publicipprefix.Setup,
		route.Setup,
		routefilter.Setup,
		routemap.Setup,
		routeserver.Setup,
		routeserverbgpconnection.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securityrule.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		trafficmanagerazureendpoint.Setup,
		trafficmanagerexternalendpoint.Setup,
		trafficmanagernestedendpoint.Setup,
		trafficmanagerprofile.Setup,
		virtualhub.Setup,
		virtualhubconnection.Setup,
		virtualhubip.Setup,
		virtualhubroutetable.Setup,
		virtualhubroutetableroute.Setup,
		virtualhubsecuritypartnerprovider.Setup,
		virtualnetworknetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
		vpngateway.Setup,
		vpngatewayconnection.Setup,
		vpnserverconfiguration.Setup,
		vpnserverconfigurationpolicygroup.Setup,
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
		loganalyticsquerypack.Setup,
		loganalyticsquerypackquery.Setup,
		loganalyticssavedsearch.Setup,
		workspaceoperationalinsights.Setup,
		loganalyticssolution.Setup,
		contactprofile.Setup,
		spacecraft.Setup,
		resourcepolicyremediation.Setup,
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
		siterecoverynetworkmapping.Setup,
		siterecoveryprotectioncontainer.Setup,
		siterecoveryprotectioncontainermapping.Setup,
		siterecoveryreplicationpolicy.Setup,
		vaultrecoveryservices.Setup,
		eventrelaynamespace.Setup,
		hybridconnection.Setup,
		hybridconnectionauthorizationrule.Setup,
		namespaceauthorizationrulerelay.Setup,
		resourcedeploymentscriptazurecli.Setup,
		resourcedeploymentscriptazurepowershell.Setup,
		resourcegrouptemplatedeployment.Setup,
		subscriptiontemplatedeployment.Setup,
		servicesearch.Setup,
		sharedprivatelinkservice.Setup,
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
		sentinelalertrulefusion.Setup,
		sentinelalertrulemachinelearningbehavioranalytics.Setup,
		sentinelalertrulemssecurityincident.Setup,
		sentinelautomationrule.Setup,
		sentineldataconnectoriot.Setup,
		sentinelloganalyticsworkspaceonboarding.Setup,
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
		springcloudconnection.Setup,
		networkacl.Setup,
		servicesignalrservice.Setup,
		signalrsharedprivatelinkresource.Setup,
		webpubsub.Setup,
		webpubsubhub.Setup,
		webpubsubnetworkacl.Setup,
		managedapplicationdefinition.Setup,
		cloudapplicationliveview.Setup,
		mssqldatabase.Setup,
		mssqldatabaseextendedauditingpolicy.Setup,
		mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		mssqlelasticpool.Setup,
		mssqlfailovergroup.Setup,
		mssqlfirewallrule.Setup,
		mssqljobagent.Setup,
		mssqljobcredential.Setup,
		mssqlmanageddatabase.Setup,
		mssqlmanagedinstance.Setup,
		mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		mssqlmanagedinstancefailovergroup.Setup,
		mssqlmanagedinstancevulnerabilityassessment.Setup,
		mssqloutboundfirewallrule.Setup,
		mssqlserver.Setup,
		mssqlserverdnsalias.Setup,
		mssqlservermicrosoftsupportauditingpolicy.Setup,
		mssqlserversecurityalertpolicy.Setup,
		mssqlservertransparentdataencryption.Setup,
		mssqlservervulnerabilityassessment.Setup,
		mssqlvirtualnetworkrule.Setup,
		accountstorage.Setup,
		accountlocaluser.Setup,
		accountnetworkrules.Setup,
		blob.Setup,
		blobinventorypolicy.Setup,
		container.Setup,
		datalakegen2filesystem.Setup,
		datalakegen2path.Setup,
		encryptionscope.Setup,
		managementpolicy.Setup,
		objectreplication.Setup,
		queuestorage.Setup,
		share.Setup,
		sharedirectory.Setup,
		tablestorage.Setup,
		tableentity.Setup,
		hpccache.Setup,
		hpccacheaccesspolicy.Setup,
		hpccacheblobnfstarget.Setup,
		hpccacheblobtarget.Setup,
		hpccachenfstarget.Setup,
		diskpool.Setup,
		storagesync.Setup,
		clusterstreamanalytics.Setup,
		functionjavascriptuda.Setup,
		jobstreamanalytics.Setup,
		managedprivateendpointstreamanalytics.Setup,
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
		firewallrulesynapse.Setup,
		integrationruntimeazuresynapse.Setup,
		integrationruntimeselfhostedsynapse.Setup,
		linkedservice.Setup,
		managedprivateendpointsynapse.Setup,
		privatelinkhub.Setup,
		roleassignmentsynapse.Setup,
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
		workspacesqlaadadmin.Setup,
		workspacevulnerabilityassessment.Setup,
		eventsourceeventhub.Setup,
		eventsourceiothub.Setup,
		gen2environment.Setup,
		referencedataset.Setup,
		standardenvironment.Setup,
		appactiveslot.Setup,
		apphybridconnection.Setup,
		appserviceplan.Setup,
		functionapp.Setup,
		functionappactiveslot.Setup,
		functionappfunction.Setup,
		functionapphybridconnection.Setup,
		functionappslot.Setup,
		linuxfunctionapp.Setup,
		linuxfunctionappslot.Setup,
		linuxwebapp.Setup,
		linuxwebappslot.Setup,
		serviceplan.Setup,
		sourcecontroltoken.Setup,
		staticsite.Setup,
		windowsfunctionapp.Setup,
		windowsfunctionappslot.Setup,
		windowswebapp.Setup,
		windowswebappslot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
