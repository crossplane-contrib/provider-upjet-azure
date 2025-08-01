// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	monitoralertprocessingruleactiongroup "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoralertprocessingruleactiongroup"
	monitoralertprocessingrulesuppression "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitoralertprocessingrulesuppression"
	monitorsmartdetectoralertrule "github.com/upbound/provider-azure/internal/controller/cluster/alertsmanagement/monitorsmartdetectoralertrule"
	server "github.com/upbound/provider-azure/internal/controller/cluster/analysisservices/server"
	api "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/api"
	apidiagnostic "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apidiagnostic"
	apioperation "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apioperation"
	apioperationpolicy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apioperationpolicy"
	apioperationtag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apioperationtag"
	apipolicy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apipolicy"
	apirelease "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apirelease"
	apischema "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apischema"
	apitag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apitag"
	apiversionset "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/apiversionset"
	authorizationserver "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/authorizationserver"
	backend "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/backend"
	certificate "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/certificate"
	customdomain "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/customdomain"
	diagnostic "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/diagnostic"
	emailtemplate "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/emailtemplate"
	gateway "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/gateway"
	gatewayapi "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/gatewayapi"
	globalschema "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/globalschema"
	identityprovideraad "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovideraad"
	identityproviderfacebook "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/identityprovidertwitter"
	logger "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/logger"
	management "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/management"
	namedvalue "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/namedvalue"
	notificationrecipientemail "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/notificationrecipientemail"
	notificationrecipientuser "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/notificationrecipientuser"
	openidconnectprovider "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/openidconnectprovider"
	policy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/policy"
	product "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/product"
	productapi "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/productapi"
	productpolicy "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/productpolicy"
	producttag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/producttag"
	rediscache "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/rediscache"
	subscription "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/subscription"
	tag "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/cluster/apimanagement/user"
	configuration "github.com/upbound/provider-azure/internal/controller/cluster/appconfiguration/configuration"
	springcloudaccelerator "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudaccelerator"
	springcloudactivedeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudactivedeployment"
	springcloudapiportal "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapiportal"
	springcloudapiportalcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapiportalcustomdomain"
	springcloudapp "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappredisassociation"
	springcloudbuilddeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuilddeployment"
	springcloudbuilder "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuilder"
	springcloudbuildpackbinding "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuildpackbinding"
	springcloudcertificate "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcertificate"
	springcloudconfigurationservice "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudconfigurationservice"
	springcloudcontainerdeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcontainerdeployment"
	springcloudcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcustomdomain"
	springcloudcustomizedaccelerator "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcustomizedaccelerator"
	springclouddevtoolportal "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springclouddevtoolportal"
	springcloudgateway "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudgateway"
	springcloudgatewaycustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudgatewaycustomdomain"
	springcloudjavadeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudservice"
	springcloudstorage "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudstorage"
	provider "github.com/upbound/provider-azure/internal/controller/cluster/attestation/provider"
	managementgrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/managementgrouppolicyassignment"
	managementgrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/managementgrouppolicyexemption"
	managementlock "github.com/upbound/provider-azure/internal/controller/cluster/authorization/managementlock"
	pimactiveroleassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/pimactiveroleassignment"
	pimeligibleroleassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/pimeligibleroleassignment"
	policydefinition "github.com/upbound/provider-azure/internal/controller/cluster/authorization/policydefinition"
	policysetdefinition "github.com/upbound/provider-azure/internal/controller/cluster/authorization/policysetdefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcegrouppolicyassignment"
	resourcegrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcegrouppolicyexemption"
	resourcepolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/resourcepolicyexemption"
	roleassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/roleassignment"
	roledefinition "github.com/upbound/provider-azure/internal/controller/cluster/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/internal/controller/cluster/authorization/subscriptionpolicyexemption"
	trustedaccessrolebinding "github.com/upbound/provider-azure/internal/controller/cluster/authorization/trustedaccessrolebinding"
	account "github.com/upbound/provider-azure/internal/controller/cluster/automation/account"
	connection "github.com/upbound/provider-azure/internal/controller/cluster/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/internal/controller/cluster/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/internal/controller/cluster/automation/connectiontype"
	credential "github.com/upbound/provider-azure/internal/controller/cluster/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/internal/controller/cluster/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/internal/controller/cluster/automation/module"
	runbook "github.com/upbound/provider-azure/internal/controller/cluster/automation/runbook"
	schedule "github.com/upbound/provider-azure/internal/controller/cluster/automation/schedule"
	variablebool "github.com/upbound/provider-azure/internal/controller/cluster/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/internal/controller/cluster/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/internal/controller/cluster/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/cluster/automation/variablestring"
	webhook "github.com/upbound/provider-azure/internal/controller/cluster/automation/webhook"
	resourcegroup "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourceproviderregistration"
	subscriptionazure "github.com/upbound/provider-azure/internal/controller/cluster/azure/subscription"
	cluster "github.com/upbound/provider-azure/internal/controller/cluster/azurestackhci/cluster"
	botchannelalexa "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelalexa"
	botchanneldirectline "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchanneldirectline"
	botchannelline "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelline"
	botchannelmsteams "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelmsteams"
	botchannelslack "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelslack"
	botchannelsms "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelsms"
	botchannelsregistration "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelsregistration"
	botchannelwebchat "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botchannelwebchat"
	botconnection "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botconnection"
	botwebapp "github.com/upbound/provider-azure/internal/controller/cluster/botservice/botwebapp"
	rediscachecache "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscache"
	rediscacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscacheaccesspolicy"
	rediscacheaccesspolicyassignment "github.com/upbound/provider-azure/internal/controller/cluster/cache/rediscacheaccesspolicyassignment"
	redisenterprisecluster "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/internal/controller/cluster/cache/redislinkedserver"
	endpoint "github.com/upbound/provider-azure/internal/controller/cluster/cdn/endpoint"
	frontdoorcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorcustomdomain"
	frontdoorcustomdomainassociation "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorcustomdomainassociation"
	frontdoorendpoint "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorendpoint"
	frontdoorfirewallpolicy "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorfirewallpolicy"
	frontdoororigin "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoororigin"
	frontdoororigingroup "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoororigingroup"
	frontdoorprofile "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorprofile"
	frontdoorroute "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorroute"
	frontdoorrule "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorrule"
	frontdoorruleset "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorruleset"
	frontdoorsecuritypolicy "github.com/upbound/provider-azure/internal/controller/cluster/cdn/frontdoorsecuritypolicy"
	profile "github.com/upbound/provider-azure/internal/controller/cluster/cdn/profile"
	appservicecertificateorder "github.com/upbound/provider-azure/internal/controller/cluster/certificateregistration/appservicecertificateorder"
	accountcognitiveservices "github.com/upbound/provider-azure/internal/controller/cluster/cognitiveservices/account"
	deployment "github.com/upbound/provider-azure/internal/controller/cluster/cognitiveservices/deployment"
	service "github.com/upbound/provider-azure/internal/controller/cluster/communication/service"
	availabilityset "github.com/upbound/provider-azure/internal/controller/cluster/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/cluster/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/cluster/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/cluster/compute/snapshot"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/cluster/compute/sshpublickey"
	virtualmachinedatadiskattachment "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineextension"
	virtualmachineruncommand "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineruncommand"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/windowsvirtualmachinescaleset"
	ledger "github.com/upbound/provider-azure/internal/controller/cluster/confidentialledger/ledger"
	budgetmanagementgroup "github.com/upbound/provider-azure/internal/controller/cluster/consumption/budgetmanagementgroup"
	budgetresourcegroup "github.com/upbound/provider-azure/internal/controller/cluster/consumption/budgetresourcegroup"
	budgetsubscription "github.com/upbound/provider-azure/internal/controller/cluster/consumption/budgetsubscription"
	containerapp "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/containerapp"
	customdomaincontainerapp "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/customdomain"
	environment "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environment"
	environmentcertificate "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentcertificate"
	environmentcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentcustomdomain"
	environmentdaprcomponent "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentdaprcomponent"
	environmentstorage "github.com/upbound/provider-azure/internal/controller/cluster/containerapp/environmentstorage"
	agentpool "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/registry"
	scopemap "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/scopemap"
	token "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/token"
	tokenpassword "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/tokenpassword"
	webhookcontainerregistry "github.com/upbound/provider-azure/internal/controller/cluster/containerregistry/webhook"
	kubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetescluster"
	kubernetesclusterextension "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetesclusterextension"
	kubernetesclusternodepool "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetesclusternodepool"
	kubernetesfleetmanager "github.com/upbound/provider-azure/internal/controller/cluster/containerservice/kubernetesfleetmanager"
	accountcosmosdb "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/account"
	cassandracluster "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/provider-azure/internal/controller/cluster/cosmosdb/gremlingraph"
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
	costanomalyalert "github.com/upbound/provider-azure/internal/controller/cluster/costmanagement/costanomalyalert"
	resourcegroupcostmanagementexport "github.com/upbound/provider-azure/internal/controller/cluster/costmanagement/resourcegroupcostmanagementexport"
	subscriptioncostmanagementexport "github.com/upbound/provider-azure/internal/controller/cluster/costmanagement/subscriptioncostmanagementexport"
	customprovider "github.com/upbound/provider-azure/internal/controller/cluster/customproviders/customprovider"
	device "github.com/upbound/provider-azure/internal/controller/cluster/databoxedge/device"
	accessconnector "github.com/upbound/provider-azure/internal/controller/cluster/databricks/accessconnector"
	workspace "github.com/upbound/provider-azure/internal/controller/cluster/databricks/workspace"
	workspacerootdbfscustomermanagedkey "github.com/upbound/provider-azure/internal/controller/cluster/databricks/workspacerootdbfscustomermanagedkey"
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
	databasemigrationproject "github.com/upbound/provider-azure/internal/controller/cluster/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/upbound/provider-azure/internal/controller/cluster/datamigration/databasemigrationservice"
	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancedisk"
	backupinstancekubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancekubernetescluster"
	backupinstancepostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupinstancepostgresql"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicydisk"
	backuppolicykubernetescluster "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicykubernetescluster"
	backuppolicypostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backuppolicypostgresql"
	backupvault "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/internal/controller/cluster/dataprotection/resourceguard"
	accountdatashare "github.com/upbound/provider-azure/internal/controller/cluster/datashare/account"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datashare"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/dbformysql/flexibleserverfirewallrule"
	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/configuration"
	database "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/firewallrule"
	flexibleserverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserver"
	flexibleserveractivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserveractivedirectoryadministrator"
	flexibleserverconfigurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/server"
	serverkey "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/serverkey"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/cluster/dbforpostgresql/virtualnetworkrule"
	iothub "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothub"
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubcertificate"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubsharedaccesspolicy"
	iothubdeviceupdateaccount "github.com/upbound/provider-azure/internal/controller/cluster/deviceupdate/iothubdeviceupdateaccount"
	iothubdeviceupdateinstance "github.com/upbound/provider-azure/internal/controller/cluster/deviceupdate/iothubdeviceupdateinstance"
	globalvmshutdownschedule "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/globalvmshutdownschedule"
	lab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/lab"
	linuxvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/linuxvirtualmachine"
	policydevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/policy"
	scheduledevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/schedule"
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/virtualnetwork"
	windowsvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/cluster/devtestlab/windowsvirtualmachine"
	instance "github.com/upbound/provider-azure/internal/controller/cluster/digitaltwins/instance"
	cloudelasticsearch "github.com/upbound/provider-azure/internal/controller/cluster/elastic/cloudelasticsearch"
	domain "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/domain"
	domaintopic "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/domaintopic"
	eventsubscription "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/eventsubscription"
	systemtopic "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/systemtopic"
	topic "github.com/upbound/provider-azure/internal/controller/cluster/eventgrid/topic"
	authorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/authorizationrule"
	consumergroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespacedisasterrecoveryconfig"
	namespaceschemagroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespaceschemagroup"
	serverfluidrelay "github.com/upbound/provider-azure/internal/controller/cluster/fluidrelay/server"
	policyvirtualmachineconfigurationassignment "github.com/upbound/provider-azure/internal/controller/cluster/guestconfiguration/policyvirtualmachineconfigurationassignment"
	hadoopcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hadoopcluster"
	hbasecluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hbasecluster"
	interactivequerycluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/interactivequerycluster"
	kafkacluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/kafkacluster"
	sparkcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/sparkcluster"
	healthbot "github.com/upbound/provider-azure/internal/controller/cluster/healthbot/healthbot"
	healthcaredicomservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaredicomservice"
	healthcarefhirservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcarefhirservice"
	healthcaremedtechservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaremedtechservice"
	healthcaremedtechservicefhirdestination "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcaremedtechservicefhirdestination"
	healthcareservice "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcareservice"
	healthcareworkspace "github.com/upbound/provider-azure/internal/controller/cluster/healthcareapis/healthcareworkspace"
	applicationinsights "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/internal/controller/cluster/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordatacollectionruleassociation"
	monitordiagnosticsetting "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitordiagnosticsetting"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/cluster/insights/monitorscheduledqueryruleslog"
	application "github.com/upbound/provider-azure/internal/controller/cluster/iotcentral/application"
	applicationnetworkruleset "github.com/upbound/provider-azure/internal/controller/cluster/iotcentral/applicationnetworkruleset"
	accesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/accesspolicy"
	certificatekeyvault "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/certificate"
	certificatecontacts "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/certificatecontacts"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/secret"
	vault "github.com/upbound/provider-azure/internal/controller/cluster/keyvault/vault"
	attacheddatabaseconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/kusto/attacheddatabaseconfiguration"
	clusterkusto "github.com/upbound/provider-azure/internal/controller/cluster/kusto/cluster"
	clustermanagedprivateendpoint "github.com/upbound/provider-azure/internal/controller/cluster/kusto/clustermanagedprivateendpoint"
	clusterprincipalassignment "github.com/upbound/provider-azure/internal/controller/cluster/kusto/clusterprincipalassignment"
	databasekusto "github.com/upbound/provider-azure/internal/controller/cluster/kusto/database"
	databaseprincipalassignment "github.com/upbound/provider-azure/internal/controller/cluster/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/upbound/provider-azure/internal/controller/cluster/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/upbound/provider-azure/internal/controller/cluster/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/upbound/provider-azure/internal/controller/cluster/kusto/iothubdataconnection"
	loadtest "github.com/upbound/provider-azure/internal/controller/cluster/loadtestservice/loadtest"
	appactioncustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactioncustom"
	appactionhttp "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactionhttp"
	appintegrationaccount "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccount"
	appintegrationaccountbatchconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountpartner "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountpartner"
	appintegrationaccountschema "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountsession"
	apptriggercustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggercustom"
	apptriggerhttprequest "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerrecurrence"
	appworkflow "github.com/upbound/provider-azure/internal/controller/cluster/logic/appworkflow"
	computecluster "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/computecluster"
	computeinstance "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/computeinstance"
	synapsespark "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/synapsespark"
	workspacemachinelearningservices "github.com/upbound/provider-azure/internal/controller/cluster/machinelearningservices/workspace"
	maintenanceassignmentdedicatedhost "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/maintenance/maintenanceconfiguration"
	federatedidentitycredential "github.com/upbound/provider-azure/internal/controller/cluster/managedidentity/federatedidentitycredential"
	userassignedidentity "github.com/upbound/provider-azure/internal/controller/cluster/managedidentity/userassignedidentity"
	managementgroup "github.com/upbound/provider-azure/internal/controller/cluster/management/managementgroup"
	managementgroupsubscriptionassociation "github.com/upbound/provider-azure/internal/controller/cluster/management/managementgroupsubscriptionassociation"
	accountmaps "github.com/upbound/provider-azure/internal/controller/cluster/maps/account"
	creator "github.com/upbound/provider-azure/internal/controller/cluster/maps/creator"
	marketplaceagreement "github.com/upbound/provider-azure/internal/controller/cluster/marketplaceordering/marketplaceagreement"
	spatialanchorsaccount "github.com/upbound/provider-azure/internal/controller/cluster/mixedreality/spatialanchorsaccount"
	accountnetapp "github.com/upbound/provider-azure/internal/controller/cluster/netapp/account"
	pool "github.com/upbound/provider-azure/internal/controller/cluster/netapp/pool"
	snapshotnetapp "github.com/upbound/provider-azure/internal/controller/cluster/netapp/snapshot"
	snapshotpolicy "github.com/upbound/provider-azure/internal/controller/cluster/netapp/snapshotpolicy"
	volume "github.com/upbound/provider-azure/internal/controller/cluster/netapp/volume"
	applicationgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/applicationgateway"
	applicationsecuritygroup "github.com/upbound/provider-azure/internal/controller/cluster/network/applicationsecuritygroup"
	bastionhost "github.com/upbound/provider-azure/internal/controller/cluster/network/bastionhost"
	connectionmonitor "github.com/upbound/provider-azure/internal/controller/cluster/network/connectionmonitor"
	ddosprotectionplan "github.com/upbound/provider-azure/internal/controller/cluster/network/ddosprotectionplan"
	dnsaaaarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsaaaarecord"
	dnsarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsarecord"
	dnscaarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnscaarecord"
	dnscnamerecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnscnamerecord"
	dnsmxrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsmxrecord"
	dnsnsrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsnsrecord"
	dnsptrrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnsptrrecord"
	dnssrvrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnssrvrecord"
	dnstxtrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/dnstxtrecord"
	dnszone "github.com/upbound/provider-azure/internal/controller/cluster/network/dnszone"
	expressroutecircuit "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/expressrouteconnection"
	expressroutegateway "github.com/upbound/provider-azure/internal/controller/cluster/network/expressroutegateway"
	expressrouteport "github.com/upbound/provider-azure/internal/controller/cluster/network/expressrouteport"
	firewall "github.com/upbound/provider-azure/internal/controller/cluster/network/firewall"
	firewallapplicationrulecollection "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/upbound/provider-azure/internal/controller/cluster/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicynetwork "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/upbound/provider-azure/internal/controller/cluster/network/frontdoorrulesengine"
	ipgroup "github.com/upbound/provider-azure/internal/controller/cluster/network/ipgroup"
	loadbalancer "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancer"
	loadbalancerbackendaddresspool "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerbackendaddresspool"
	loadbalancerbackendaddresspooladdress "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerbackendaddresspooladdress"
	loadbalancernatpool "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancernatpool"
	loadbalancernatrule "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancernatrule"
	loadbalanceroutboundrule "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalanceroutboundrule"
	loadbalancerprobe "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerprobe"
	loadbalancerrule "github.com/upbound/provider-azure/internal/controller/cluster/network/loadbalancerrule"
	localnetworkgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/localnetworkgateway"
	manager "github.com/upbound/provider-azure/internal/controller/cluster/network/manager"
	managermanagementgroupconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/managermanagementgroupconnection"
	managernetworkgroup "github.com/upbound/provider-azure/internal/controller/cluster/network/managernetworkgroup"
	managerstaticmember "github.com/upbound/provider-azure/internal/controller/cluster/network/managerstaticmember"
	managersubscriptionconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/managersubscriptionconnection"
	natgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/natgateway"
	natgatewaypublicipassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/networkinterfacesecuritygroupassociation"
	packetcapture "github.com/upbound/provider-azure/internal/controller/cluster/network/packetcapture"
	pointtositevpngateway "github.com/upbound/provider-azure/internal/controller/cluster/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsaaaarecord"
	privatednsarecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsarecord"
	privatednscnamerecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednscnamerecord"
	privatednsmxrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsmxrecord"
	privatednsptrrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsptrrecord"
	privatednsresolver "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolver"
	privatednsresolverdnsforwardingruleset "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolverdnsforwardingruleset"
	privatednsresolverforwardingrule "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolverforwardingrule"
	privatednsresolverinboundendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolverinboundendpoint"
	privatednsresolveroutboundendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednsresolveroutboundendpoint"
	privatednssrvrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednstxtrecord"
	privatednszone "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/provider-azure/internal/controller/cluster/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/privateendpoint"
	privateendpointapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/privateendpointapplicationsecuritygroupassociation"
	privatelinkservice "github.com/upbound/provider-azure/internal/controller/cluster/network/privatelinkservice"
	profilenetwork "github.com/upbound/provider-azure/internal/controller/cluster/network/profile"
	publicip "github.com/upbound/provider-azure/internal/controller/cluster/network/publicip"
	publicipprefix "github.com/upbound/provider-azure/internal/controller/cluster/network/publicipprefix"
	route "github.com/upbound/provider-azure/internal/controller/cluster/network/route"
	routefilter "github.com/upbound/provider-azure/internal/controller/cluster/network/routefilter"
	routemap "github.com/upbound/provider-azure/internal/controller/cluster/network/routemap"
	routeserver "github.com/upbound/provider-azure/internal/controller/cluster/network/routeserver"
	routeserverbgpconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/routeserverbgpconnection"
	routetable "github.com/upbound/provider-azure/internal/controller/cluster/network/routetable"
	securitygroup "github.com/upbound/provider-azure/internal/controller/cluster/network/securitygroup"
	securityrule "github.com/upbound/provider-azure/internal/controller/cluster/network/securityrule"
	subnet "github.com/upbound/provider-azure/internal/controller/cluster/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/provider-azure/internal/controller/cluster/network/subnetserviceendpointstoragepolicy"
	trafficmanagerazureendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagerazureendpoint"
	trafficmanagerexternalendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagerexternalendpoint"
	trafficmanagernestedendpoint "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagernestedendpoint"
	trafficmanagerprofile "github.com/upbound/provider-azure/internal/controller/cluster/network/trafficmanagerprofile"
	virtualhub "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhub"
	virtualhubconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubconnection"
	virtualhubip "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubip"
	virtualhubroutetable "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubroutetable"
	virtualhubroutetableroute "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubroutetableroute"
	virtualhubsecuritypartnerprovider "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualhubsecuritypartnerprovider"
	virtualnetworknetwork "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/internal/controller/cluster/network/virtualwan"
	vpngateway "github.com/upbound/provider-azure/internal/controller/cluster/network/vpngateway"
	vpngatewayconnection "github.com/upbound/provider-azure/internal/controller/cluster/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/network/vpnserverconfiguration"
	vpnserverconfigurationpolicygroup "github.com/upbound/provider-azure/internal/controller/cluster/network/vpnserverconfigurationpolicygroup"
	vpnsite "github.com/upbound/provider-azure/internal/controller/cluster/network/vpnsite"
	watcher "github.com/upbound/provider-azure/internal/controller/cluster/network/watcher"
	watcherflowlog "github.com/upbound/provider-azure/internal/controller/cluster/network/watcherflowlog"
	webapplicationfirewallpolicy "github.com/upbound/provider-azure/internal/controller/cluster/network/webapplicationfirewallpolicy"
	authorizationrulenotificationhubs "github.com/upbound/provider-azure/internal/controller/cluster/notificationhubs/authorizationrule"
	notificationhub "github.com/upbound/provider-azure/internal/controller/cluster/notificationhubs/notificationhub"
	notificationhubnamespace "github.com/upbound/provider-azure/internal/controller/cluster/notificationhubs/notificationhubnamespace"
	loganalyticsdataexportrule "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticsquerypack "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsquerypack"
	loganalyticsquerypackquery "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticsquerypackquery"
	loganalyticssavedsearch "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/loganalyticssavedsearch"
	workspaceoperationalinsights "github.com/upbound/provider-azure/internal/controller/cluster/operationalinsights/workspace"
	loganalyticssolution "github.com/upbound/provider-azure/internal/controller/cluster/operationsmanagement/loganalyticssolution"
	contactprofile "github.com/upbound/provider-azure/internal/controller/cluster/orbital/contactprofile"
	spacecraft "github.com/upbound/provider-azure/internal/controller/cluster/orbital/spacecraft"
	resourcepolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/resourcepolicyremediation"
	subscriptionpolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/subscriptionpolicyremediation"
	dashboard "github.com/upbound/provider-azure/internal/controller/cluster/portal/dashboard"
	powerbiembedded "github.com/upbound/provider-azure/internal/controller/cluster/powerbidedicated/powerbiembedded"
	providerconfig "github.com/upbound/provider-azure/internal/controller/cluster/providerconfig"
	accountpurview "github.com/upbound/provider-azure/internal/controller/cluster/purview/account"
	backupcontainerstorageaccount "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyvm"
	backuppolicyvmworkload "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backuppolicyvmworkload"
	backupprotectedfileshare "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicationpolicy "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/siterecoveryreplicationpolicy"
	vaultrecoveryservices "github.com/upbound/provider-azure/internal/controller/cluster/recoveryservices/vault"
	eventrelaynamespace "github.com/upbound/provider-azure/internal/controller/cluster/relay/eventrelaynamespace"
	hybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/relay/hybridconnectionauthorizationrule"
	namespaceauthorizationrulerelay "github.com/upbound/provider-azure/internal/controller/cluster/relay/namespaceauthorizationrule"
	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/cluster/resources/subscriptiontemplatedeployment"
	servicesearch "github.com/upbound/provider-azure/internal/controller/cluster/search/service"
	sharedprivatelinkservice "github.com/upbound/provider-azure/internal/controller/cluster/search/sharedprivatelinkservice"
	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/cluster/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/cluster/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/cluster/security/iotsecuritysolution"
	securitycenterassessment "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterassessment"
	securitycenterassessmentpolicy "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycentercontact"
	securitycenterservervulnerabilityassessmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterservervulnerabilityassessmentvirtualmachine"
	securitycentersetting "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/upbound/provider-azure/internal/controller/cluster/security/securitycenterworkspace"
	storagedefender "github.com/upbound/provider-azure/internal/controller/cluster/security/storagedefender"
	sentinelalertrulefusion "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulemssecurityincident"
	sentinelautomationrule "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelautomationrule"
	sentineldataconnectoriot "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentineldataconnectoriot"
	sentinelloganalyticsworkspaceonboarding "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelloganalyticsworkspaceonboarding"
	sentinelwatchlist "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelwatchlist"
	namespaceauthorizationruleservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfigservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/namespacedisasterrecoveryconfig"
	queue "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/queue"
	queueauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/queueauthorizationrule"
	servicebusnamespace "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/servicebusnamespace"
	subscriptionservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/subscription"
	subscriptionrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/subscriptionrule"
	topicservicebus "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/topic"
	topicauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/servicebus/topicauthorizationrule"
	clusterservicefabric "github.com/upbound/provider-azure/internal/controller/cluster/servicefabric/cluster"
	managedcluster "github.com/upbound/provider-azure/internal/controller/cluster/servicefabric/managedcluster"
	springcloudconnection "github.com/upbound/provider-azure/internal/controller/cluster/servicelinker/springcloudconnection"
	networkacl "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/networkacl"
	servicesignalrservice "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/service"
	signalrsharedprivatelinkresource "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/signalrsharedprivatelinkresource"
	webpubsub "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/webpubsub"
	webpubsubhub "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/webpubsubhub"
	webpubsubnetworkacl "github.com/upbound/provider-azure/internal/controller/cluster/signalrservice/webpubsubnetworkacl"
	managedapplicationdefinition "github.com/upbound/provider-azure/internal/controller/cluster/solutions/managedapplicationdefinition"
	cloudapplicationliveview "github.com/upbound/provider-azure/internal/controller/cluster/spring/cloudapplicationliveview"
	mssqldatabase "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlfirewallrule"
	mssqljobagent "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqljobagent"
	mssqljobcredential "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqljobcredential"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancetransparentdataencryption "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancetransparentdataencryption"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlvirtualnetworkrule"
	accountstorage "github.com/upbound/provider-azure/internal/controller/cluster/storage/account"
	accountlocaluser "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountlocaluser"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/cluster/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/cluster/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/cluster/storage/container"
	containerimmutabilitypolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/containerimmutabilitypolicy"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2filesystem"
	datalakegen2path "github.com/upbound/provider-azure/internal/controller/cluster/storage/datalakegen2path"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/cluster/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/cluster/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/cluster/storage/objectreplication"
	queuestorage "github.com/upbound/provider-azure/internal/controller/cluster/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/cluster/storage/share"
	sharedirectory "github.com/upbound/provider-azure/internal/controller/cluster/storage/sharedirectory"
	tablestorage "github.com/upbound/provider-azure/internal/controller/cluster/storage/table"
	tableentity "github.com/upbound/provider-azure/internal/controller/cluster/storage/tableentity"
	hpccache "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccachenfstarget"
	storagesync "github.com/upbound/provider-azure/internal/controller/cluster/storagesync/storagesync"
	clusterstreamanalytics "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/functionjavascriptuda"
	job "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/job"
	managedprivateendpointstreamanalytics "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/managedprivateendpoint"
	outputblob "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputblob"
	outputeventhub "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputeventhub"
	outputfunction "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputfunction"
	outputmssql "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputmssql"
	outputpowerbi "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputpowerbi"
	outputservicebusqueue "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputservicebustopic"
	outputsynapse "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputsynapse"
	outputtable "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/outputtable"
	referenceinputblob "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/referenceinputblob"
	referenceinputmssql "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/referenceinputmssql"
	streaminputblob "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/upbound/provider-azure/internal/controller/cluster/streamanalytics/streaminputiothub"
	firewallrulesynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/firewallrule"
	integrationruntimeazuresynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/integrationruntimeazure"
	integrationruntimeselfhostedsynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/internal/controller/cluster/synapse/linkedservice"
	managedprivateendpointsynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/internal/controller/cluster/synapse/privatelinkhub"
	roleassignmentsynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/roleassignment"
	sparkpool "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sparkpool"
	sqlpool "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolworkloadclassifier "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolworkloadclassifier"
	sqlpoolworkloadgroup "github.com/upbound/provider-azure/internal/controller/cluster/synapse/sqlpoolworkloadgroup"
	workspacesynapse "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspace"
	workspaceaadadmin "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspaceaadadmin"
	workspaceextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspaceextendedauditingpolicy"
	workspacesecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspacesecurityalertpolicy"
	workspacesqlaadadmin "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspacesqlaadadmin"
	workspacevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/synapse/workspacevulnerabilityassessment"
	appactiveslot "github.com/upbound/provider-azure/internal/controller/cluster/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/web/apphybridconnection"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/cluster/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/functionapp"
	functionappactiveslot "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappactiveslot"
	functionappfunction "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/internal/controller/cluster/web/functionapphybridconnection"
	functionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/functionappslot"
	linuxfunctionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/internal/controller/cluster/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/cluster/web/sourcecontroltoken"
	staticsite "github.com/upbound/provider-azure/internal/controller/cluster/web/staticsite"
	windowsfunctionapp "github.com/upbound/provider-azure/internal/controller/cluster/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/internal/controller/cluster/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/internal/controller/cluster/web/windowswebappslot"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
		customdomain.Setup,
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
		managementgrouppolicyassignment.Setup,
		managementgrouppolicyexemption.Setup,
		managementlock.Setup,
		pimactiveroleassignment.Setup,
		pimeligibleroleassignment.Setup,
		policydefinition.Setup,
		policysetdefinition.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcegrouppolicyexemption.Setup,
		resourcepolicyassignment.Setup,
		resourcepolicyexemption.Setup,
		roleassignment.Setup,
		roledefinition.Setup,
		subscriptionpolicyassignment.Setup,
		subscriptionpolicyexemption.Setup,
		trustedaccessrolebinding.Setup,
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
		rediscacheaccesspolicy.Setup,
		rediscacheaccesspolicyassignment.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
		endpoint.Setup,
		frontdoorcustomdomain.Setup,
		frontdoorcustomdomainassociation.Setup,
		frontdoorendpoint.Setup,
		frontdoorfirewallpolicy.Setup,
		frontdoororigin.Setup,
		frontdoororigingroup.Setup,
		frontdoorprofile.Setup,
		frontdoorroute.Setup,
		frontdoorrule.Setup,
		frontdoorruleset.Setup,
		frontdoorsecuritypolicy.Setup,
		profile.Setup,
		appservicecertificateorder.Setup,
		accountcognitiveservices.Setup,
		deployment.Setup,
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
		virtualmachinedatadiskattachment.Setup,
		virtualmachineextension.Setup,
		virtualmachineruncommand.Setup,
		windowsvirtualmachine.Setup,
		windowsvirtualmachinescaleset.Setup,
		ledger.Setup,
		budgetmanagementgroup.Setup,
		budgetresourcegroup.Setup,
		budgetsubscription.Setup,
		containerapp.Setup,
		customdomaincontainerapp.Setup,
		environment.Setup,
		environmentcertificate.Setup,
		environmentcustomdomain.Setup,
		environmentdaprcomponent.Setup,
		environmentstorage.Setup,
		agentpool.Setup,
		containerconnectedregistry.Setup,
		registry.Setup,
		scopemap.Setup,
		token.Setup,
		tokenpassword.Setup,
		webhookcontainerregistry.Setup,
		kubernetescluster.Setup,
		kubernetesclusterextension.Setup,
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
		costanomalyalert.Setup,
		resourcegroupcostmanagementexport.Setup,
		subscriptioncostmanagementexport.Setup,
		customprovider.Setup,
		device.Setup,
		accessconnector.Setup,
		workspace.Setup,
		workspacerootdbfscustomermanagedkey.Setup,
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
		databasemigrationproject.Setup,
		databasemigrationservice.Setup,
		backupinstanceblobstorage.Setup,
		backupinstancedisk.Setup,
		backupinstancekubernetescluster.Setup,
		backupinstancepostgresql.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydisk.Setup,
		backuppolicykubernetescluster.Setup,
		backuppolicypostgresql.Setup,
		backupvault.Setup,
		resourceguard.Setup,
		accountdatashare.Setup,
		datasetblobstorage.Setup,
		datasetdatalakegen2.Setup,
		datasetkustocluster.Setup,
		datasetkustodatabase.Setup,
		datashare.Setup,
		flexibledatabase.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverfirewallrule.Setup,
		activedirectoryadministrator.Setup,
		configurationdbforpostgresql.Setup,
		database.Setup,
		firewallrule.Setup,
		flexibleserverdbforpostgresql.Setup,
		flexibleserveractivedirectoryadministrator.Setup,
		flexibleserverconfigurationdbforpostgresql.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallruledbforpostgresql.Setup,
		serverdbforpostgresql.Setup,
		serverkey.Setup,
		virtualnetworkrule.Setup,
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
		monitordiagnosticsetting.Setup,
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
		loadtest.Setup,
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
		spatialanchorsaccount.Setup,
		accountnetapp.Setup,
		pool.Setup,
		snapshotnetapp.Setup,
		snapshotpolicy.Setup,
		volume.Setup,
		applicationgateway.Setup,
		applicationsecuritygroup.Setup,
		bastionhost.Setup,
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
		frontdoorfirewallpolicynetwork.Setup,
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
		privatednsresolverdnsforwardingruleset.Setup,
		privatednsresolverforwardingrule.Setup,
		privatednsresolverinboundendpoint.Setup,
		privatednsresolveroutboundendpoint.Setup,
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
		securitycenterservervulnerabilityassessmentvirtualmachine.Setup,
		securitycentersetting.Setup,
		securitycentersubscriptionpricing.Setup,
		securitycenterworkspace.Setup,
		storagedefender.Setup,
		sentinelalertrulefusion.Setup,
		sentinelalertrulemachinelearningbehavioranalytics.Setup,
		sentinelalertrulemssecurityincident.Setup,
		sentinelautomationrule.Setup,
		sentineldataconnectoriot.Setup,
		sentinelloganalyticsworkspaceonboarding.Setup,
		sentinelwatchlist.Setup,
		namespaceauthorizationruleservicebus.Setup,
		namespacedisasterrecoveryconfigservicebus.Setup,
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
		mssqlmanagedinstancetransparentdataencryption.Setup,
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
		containerimmutabilitypolicy.Setup,
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
		storagesync.Setup,
		clusterstreamanalytics.Setup,
		functionjavascriptuda.Setup,
		job.Setup,
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
