// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	monitoralertprocessingruleactiongroup "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitoralertprocessingruleactiongroup"
	monitoralertprocessingrulesuppression "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitoralertprocessingrulesuppression"
	monitoralertprometheusrulegroup "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitoralertprometheusrulegroup"
	monitorsmartdetectoralertrule "github.com/upbound/provider-azure/internal/controller/namespaced/alertsmanagement/monitorsmartdetectoralertrule"
	server "github.com/upbound/provider-azure/internal/controller/namespaced/analysisservices/server"
	api "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/api"
	apidiagnostic "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apidiagnostic"
	apioperation "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apioperation"
	apioperationpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apioperationpolicy"
	apioperationtag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apioperationtag"
	apipolicy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apipolicy"
	apirelease "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apirelease"
	apischema "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apischema"
	apitag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apitag"
	apiversionset "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apiversionset"
	authorizationserver "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/authorizationserver"
	backend "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/backend"
	certificate "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/certificate"
	customdomain "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/customdomain"
	diagnostic "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/diagnostic"
	emailtemplate "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/emailtemplate"
	gateway "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/gateway"
	gatewayapi "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/gatewayapi"
	globalschema "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/globalschema"
	group "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/group"
	identityprovideraad "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovideraad"
	identityproviderfacebook "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovidertwitter"
	logger "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/logger"
	management "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/management"
	namedvalue "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/namedvalue"
	notificationrecipientemail "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/notificationrecipientemail"
	notificationrecipientuser "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/notificationrecipientuser"
	openidconnectprovider "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/openidconnectprovider"
	policy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/policy"
	policyfragment "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/policyfragment"
	product "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/product"
	productapi "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/productapi"
	productgroup "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/productgroup"
	productpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/productpolicy"
	producttag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/producttag"
	rediscache "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/rediscache"
	subscription "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/subscription"
	tag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/user"
	configuration "github.com/upbound/provider-azure/internal/controller/namespaced/appconfiguration/configuration"
	springcloudaccelerator "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudaccelerator"
	springcloudactivedeployment "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudactivedeployment"
	springcloudapiportal "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudapiportal"
	springcloudapiportalcustomdomain "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudapiportalcustomdomain"
	springcloudapp "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudappredisassociation"
	springcloudbuilddeployment "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudbuilddeployment"
	springcloudbuilder "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudbuilder"
	springcloudbuildpackbinding "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudbuildpackbinding"
	springcloudcertificate "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudcertificate"
	springcloudconfigurationservice "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudconfigurationservice"
	springcloudcontainerdeployment "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudcontainerdeployment"
	springcloudcustomdomain "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudcustomdomain"
	springcloudcustomizedaccelerator "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudcustomizedaccelerator"
	springclouddevtoolportal "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springclouddevtoolportal"
	springcloudgateway "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudgateway"
	springcloudgatewaycustomdomain "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudgatewaycustomdomain"
	springcloudjavadeployment "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudservice"
	springcloudstorage "github.com/upbound/provider-azure/internal/controller/namespaced/appplatform/springcloudstorage"
	provider "github.com/upbound/provider-azure/internal/controller/namespaced/attestation/provider"
	managementgrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/managementgrouppolicyassignment"
	managementgrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/managementgrouppolicyexemption"
	managementlock "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/managementlock"
	pimactiveroleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/pimactiveroleassignment"
	pimeligibleroleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/pimeligibleroleassignment"
	policydefinition "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/policydefinition"
	policysetdefinition "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/policysetdefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcegrouppolicyassignment"
	resourcegrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcegrouppolicyexemption"
	resourcepolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/resourcepolicyexemption"
	roleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/roleassignment"
	roledefinition "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/subscriptionpolicyexemption"
	trustedaccessrolebinding "github.com/upbound/provider-azure/internal/controller/namespaced/authorization/trustedaccessrolebinding"
	account "github.com/upbound/provider-azure/internal/controller/namespaced/automation/account"
	connection "github.com/upbound/provider-azure/internal/controller/namespaced/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/internal/controller/namespaced/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/internal/controller/namespaced/automation/connectiontype"
	credential "github.com/upbound/provider-azure/internal/controller/namespaced/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/internal/controller/namespaced/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/internal/controller/namespaced/automation/module"
	runbook "github.com/upbound/provider-azure/internal/controller/namespaced/automation/runbook"
	schedule "github.com/upbound/provider-azure/internal/controller/namespaced/automation/schedule"
	variablebool "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/namespaced/automation/variablestring"
	webhook "github.com/upbound/provider-azure/internal/controller/namespaced/automation/webhook"
	resourcegroup "github.com/upbound/provider-azure/internal/controller/namespaced/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/namespaced/azure/resourceproviderregistration"
	subscriptionazure "github.com/upbound/provider-azure/internal/controller/namespaced/azure/subscription"
	cluster "github.com/upbound/provider-azure/internal/controller/namespaced/azurestackhci/cluster"
	botchannelalexa "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelalexa"
	botchanneldirectline "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchanneldirectline"
	botchannelline "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelline"
	botchannelmsteams "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelmsteams"
	botchannelslack "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelslack"
	botchannelsms "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelsms"
	botchannelsregistration "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelsregistration"
	botchannelwebchat "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelwebchat"
	botconnection "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botconnection"
	botwebapp "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botwebapp"
	rediscachecache "github.com/upbound/provider-azure/internal/controller/namespaced/cache/rediscache"
	rediscacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/cache/rediscacheaccesspolicy"
	rediscacheaccesspolicyassignment "github.com/upbound/provider-azure/internal/controller/namespaced/cache/rediscacheaccesspolicyassignment"
	redisenterprisecluster "github.com/upbound/provider-azure/internal/controller/namespaced/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/internal/controller/namespaced/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/internal/controller/namespaced/cache/redislinkedserver"
	endpoint "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/endpoint"
	frontdoorcustomdomain "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorcustomdomain"
	frontdoorcustomdomainassociation "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorcustomdomainassociation"
	frontdoorendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorendpoint"
	frontdoorfirewallpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorfirewallpolicy"
	frontdoororigin "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoororigin"
	frontdoororigingroup "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoororigingroup"
	frontdoorprofile "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorprofile"
	frontdoorroute "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorroute"
	frontdoorrule "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorrule"
	frontdoorruleset "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorruleset"
	frontdoorsecret "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorsecret"
	frontdoorsecuritypolicy "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/frontdoorsecuritypolicy"
	profile "github.com/upbound/provider-azure/internal/controller/namespaced/cdn/profile"
	appservicecertificateorder "github.com/upbound/provider-azure/internal/controller/namespaced/certificateregistration/appservicecertificateorder"
	accountcognitiveservices "github.com/upbound/provider-azure/internal/controller/namespaced/cognitiveservices/account"
	accountraiblocklist "github.com/upbound/provider-azure/internal/controller/namespaced/cognitiveservices/accountraiblocklist"
	accountraipolicy "github.com/upbound/provider-azure/internal/controller/namespaced/cognitiveservices/accountraipolicy"
	aiservices "github.com/upbound/provider-azure/internal/controller/namespaced/cognitiveservices/aiservices"
	deployment "github.com/upbound/provider-azure/internal/controller/namespaced/cognitiveservices/deployment"
	service "github.com/upbound/provider-azure/internal/controller/namespaced/communication/service"
	availabilityset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/namespaced/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/namespaced/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/namespaced/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/namespaced/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/namespaced/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/namespaced/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/namespaced/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/namespaced/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/namespaced/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/namespaced/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/namespaced/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/namespaced/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/namespaced/compute/snapshot"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/namespaced/compute/sshpublickey"
	virtualmachinedatadiskattachment "github.com/upbound/provider-azure/internal/controller/namespaced/compute/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/upbound/provider-azure/internal/controller/namespaced/compute/virtualmachineextension"
	virtualmachineruncommand "github.com/upbound/provider-azure/internal/controller/namespaced/compute/virtualmachineruncommand"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/windowsvirtualmachinescaleset"
	ledger "github.com/upbound/provider-azure/internal/controller/namespaced/confidentialledger/ledger"
	budgetmanagementgroup "github.com/upbound/provider-azure/internal/controller/namespaced/consumption/budgetmanagementgroup"
	budgetresourcegroup "github.com/upbound/provider-azure/internal/controller/namespaced/consumption/budgetresourcegroup"
	budgetsubscription "github.com/upbound/provider-azure/internal/controller/namespaced/consumption/budgetsubscription"
	containerapp "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/containerapp"
	containerjob "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/containerjob"
	customdomaincontainerapp "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/customdomain"
	environment "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environment"
	environmentcertificate "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentcertificate"
	environmentcustomdomain "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentcustomdomain"
	environmentdaprcomponent "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentdaprcomponent"
	environmentstorage "github.com/upbound/provider-azure/internal/controller/namespaced/containerapp/environmentstorage"
	agentpool "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/registry"
	scopemap "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/scopemap"
	token "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/token"
	tokenpassword "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/tokenpassword"
	webhookcontainerregistry "github.com/upbound/provider-azure/internal/controller/namespaced/containerregistry/webhook"
	kubernetescluster "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetescluster"
	kubernetesclusterextension "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetesclusterextension"
	kubernetesclusternodepool "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetesclusternodepool"
	kubernetesfleetmanager "github.com/upbound/provider-azure/internal/controller/namespaced/containerservice/kubernetesfleetmanager"
	accountcosmosdb "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/account"
	cassandracluster "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandracluster"
	cassandradatacenter "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandradatacenter"
	cassandrakeyspace "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandrakeyspace"
	cassandratable "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/cassandratable"
	gremlindatabase "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/gremlindatabase"
	gremlingraph "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/gremlingraph"
	mongocluster "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongocluster"
	mongocollection "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongocollection"
	mongodatabase "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongodatabase"
	mongoroledefinition "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongoroledefinition"
	mongouserdefinition "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/mongouserdefinition"
	sqlcontainer "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlcontainer"
	sqldatabase "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqldatabase"
	sqldedicatedgateway "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqldedicatedgateway"
	sqlfunction "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlfunction"
	sqlroleassignment "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlroleassignment"
	sqlroledefinition "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqlstoredprocedure"
	sqltrigger "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/sqltrigger"
	table "github.com/upbound/provider-azure/internal/controller/namespaced/cosmosdb/table"
	costanomalyalert "github.com/upbound/provider-azure/internal/controller/namespaced/costmanagement/costanomalyalert"
	resourcegroupcostmanagementexport "github.com/upbound/provider-azure/internal/controller/namespaced/costmanagement/resourcegroupcostmanagementexport"
	subscriptioncostmanagementexport "github.com/upbound/provider-azure/internal/controller/namespaced/costmanagement/subscriptioncostmanagementexport"
	customprovider "github.com/upbound/provider-azure/internal/controller/namespaced/customproviders/customprovider"
	device "github.com/upbound/provider-azure/internal/controller/namespaced/databoxedge/device"
	accessconnector "github.com/upbound/provider-azure/internal/controller/namespaced/databricks/accessconnector"
	workspace "github.com/upbound/provider-azure/internal/controller/namespaced/databricks/workspace"
	workspacerootdbfscustomermanagedkey "github.com/upbound/provider-azure/internal/controller/namespaced/databricks/workspacerootdbfscustomermanagedkey"
	customdataset "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/customdataset"
	dataflow "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/dataflow"
	datasetazureblob "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetazureblob"
	datasetbinary "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetbinary"
	datasetcosmosdbsqlapi "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetcosmosdbsqlapi"
	datasetdelimitedtext "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetdelimitedtext"
	datasethttp "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasethttp"
	datasetjson "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetjson"
	datasetmysql "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetmysql"
	datasetparquet "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetparquet"
	datasetpostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetpostgresql"
	datasetsnowflake "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetsnowflake"
	datasetsqlservertable "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/datasetsqlservertable"
	factory "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/factory"
	integrationruntimeazure "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/integrationruntimeazure"
	integrationruntimeazuressis "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/integrationruntimeazuressis"
	integrationruntimeselfhosted "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/integrationruntimeselfhosted"
	linkedcustomservice "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedcustomservice"
	linkedserviceazureblobstorage "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceazureblobstorage"
	linkedserviceazuredatabricks "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceazuredatabricks"
	linkedserviceazurefilestorage "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceazurefilestorage"
	linkedserviceazurefunction "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceazurefunction"
	linkedserviceazuresearch "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceazuresearch"
	linkedserviceazuresqldatabase "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceazuresqldatabase"
	linkedserviceazuretablestorage "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceazuretablestorage"
	linkedservicecosmosdb "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicecosmosdb"
	linkedservicecosmosdbmongoapi "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicecosmosdbmongoapi"
	linkedservicedatalakestoragegen2 "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicedatalakestoragegen2"
	linkedservicekeyvault "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicekeyvault"
	linkedservicekusto "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicekusto"
	linkedservicemysql "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicemysql"
	linkedserviceodata "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceodata"
	linkedserviceodbc "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceodbc"
	linkedservicepostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicepostgresql"
	linkedservicesftp "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicesftp"
	linkedservicesnowflake "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicesnowflake"
	linkedservicesqlserver "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicesqlserver"
	linkedservicesynapse "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedservicesynapse"
	linkedserviceweb "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/linkedserviceweb"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/managedprivateendpoint"
	pipeline "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/pipeline"
	triggerblobevent "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/triggerblobevent"
	triggercustomevent "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/triggercustomevent"
	triggerschedule "github.com/upbound/provider-azure/internal/controller/namespaced/datafactory/triggerschedule"
	databasemigrationproject "github.com/upbound/provider-azure/internal/controller/namespaced/datamigration/databasemigrationproject"
	databasemigrationservice "github.com/upbound/provider-azure/internal/controller/namespaced/datamigration/databasemigrationservice"
	backupinstanceblobstorage "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backupinstanceblobstorage"
	backupinstancedisk "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backupinstancedisk"
	backupinstancekubernetescluster "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backupinstancekubernetescluster"
	backupinstancepostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backupinstancepostgresql"
	backupinstancepostgresqlflexibleserver "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backupinstancepostgresqlflexibleserver"
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backuppolicyblobstorage"
	backuppolicydisk "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backuppolicydisk"
	backuppolicykubernetescluster "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backuppolicykubernetescluster"
	backuppolicypostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backuppolicypostgresql"
	backuppolicypostgresqlflexibleserver "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backuppolicypostgresqlflexibleserver"
	backupvault "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/internal/controller/namespaced/dataprotection/resourceguard"
	accountdatashare "github.com/upbound/provider-azure/internal/controller/namespaced/datashare/account"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/namespaced/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/namespaced/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/namespaced/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/namespaced/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/namespaced/datashare/datashare"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/namespaced/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/namespaced/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/dbformysql/flexibleserverfirewallrule"
	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/configuration"
	database "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/firewallrule"
	flexibleserverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/flexibleserver"
	flexibleserveractivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/flexibleserveractivedirectoryadministrator"
	flexibleserverconfigurationdbforpostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallruledbforpostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/server"
	serverkey "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/serverkey"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/namespaced/dbforpostgresql/virtualnetworkrule"
	iothub "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothub"
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubcertificate"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubsharedaccesspolicy"
	iothubdeviceupdateaccount "github.com/upbound/provider-azure/internal/controller/namespaced/deviceupdate/iothubdeviceupdateaccount"
	iothubdeviceupdateinstance "github.com/upbound/provider-azure/internal/controller/namespaced/deviceupdate/iothubdeviceupdateinstance"
	globalvmshutdownschedule "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/globalvmshutdownschedule"
	lab "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/lab"
	linuxvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/linuxvirtualmachine"
	policydevtestlab "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/policy"
	scheduledevtestlab "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/schedule"
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/virtualnetwork"
	windowsvirtualmachinedevtestlab "github.com/upbound/provider-azure/internal/controller/namespaced/devtestlab/windowsvirtualmachine"
	instance "github.com/upbound/provider-azure/internal/controller/namespaced/digitaltwins/instance"
	cloudelasticsearch "github.com/upbound/provider-azure/internal/controller/namespaced/elastic/cloudelasticsearch"
	domain "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/domain"
	domaintopic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/domaintopic"
	eventsubscription "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/eventsubscription"
	systemtopic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/systemtopic"
	systemtopiceventsubscription "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/systemtopiceventsubscription"
	topic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/topic"
	authorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/authorizationrule"
	consumergroup "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/namespacedisasterrecoveryconfig"
	namespaceschemagroup "github.com/upbound/provider-azure/internal/controller/namespaced/eventhub/namespaceschemagroup"
	serverfluidrelay "github.com/upbound/provider-azure/internal/controller/namespaced/fluidrelay/server"
	policyvirtualmachineconfigurationassignment "github.com/upbound/provider-azure/internal/controller/namespaced/guestconfiguration/policyvirtualmachineconfigurationassignment"
	hadoopcluster "github.com/upbound/provider-azure/internal/controller/namespaced/hdinsight/hadoopcluster"
	hbasecluster "github.com/upbound/provider-azure/internal/controller/namespaced/hdinsight/hbasecluster"
	interactivequerycluster "github.com/upbound/provider-azure/internal/controller/namespaced/hdinsight/interactivequerycluster"
	kafkacluster "github.com/upbound/provider-azure/internal/controller/namespaced/hdinsight/kafkacluster"
	sparkcluster "github.com/upbound/provider-azure/internal/controller/namespaced/hdinsight/sparkcluster"
	healthbot "github.com/upbound/provider-azure/internal/controller/namespaced/healthbot/healthbot"
	healthcaredicomservice "github.com/upbound/provider-azure/internal/controller/namespaced/healthcareapis/healthcaredicomservice"
	healthcarefhirservice "github.com/upbound/provider-azure/internal/controller/namespaced/healthcareapis/healthcarefhirservice"
	healthcaremedtechservice "github.com/upbound/provider-azure/internal/controller/namespaced/healthcareapis/healthcaremedtechservice"
	healthcaremedtechservicefhirdestination "github.com/upbound/provider-azure/internal/controller/namespaced/healthcareapis/healthcaremedtechservicefhirdestination"
	healthcareservice "github.com/upbound/provider-azure/internal/controller/namespaced/healthcareapis/healthcareservice"
	healthcareworkspace "github.com/upbound/provider-azure/internal/controller/namespaced/healthcareapis/healthcareworkspace"
	applicationinsights "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsights"
	applicationinsightsanalyticsitem "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsanalyticsitem"
	applicationinsightsapikey "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsapikey"
	applicationinsightssmartdetectionrule "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightssmartdetectionrule"
	applicationinsightsstandardwebtest "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsstandardwebtest"
	applicationinsightswebtest "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightswebtest"
	applicationinsightsworkbook "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsworkbook"
	applicationinsightsworkbooktemplate "github.com/upbound/provider-azure/internal/controller/namespaced/insights/applicationinsightsworkbooktemplate"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitoractiongroup"
	monitoractivitylogalert "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitoractivitylogalert"
	monitorautoscalesetting "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorautoscalesetting"
	monitordatacollectionendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordatacollectionendpoint"
	monitordatacollectionrule "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordatacollectionrule"
	monitordatacollectionruleassociation "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordatacollectionruleassociation"
	monitordiagnosticsetting "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitordiagnosticsetting"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorprivatelinkscopedservice"
	monitorscheduledqueryrulesalert "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorscheduledqueryrulesalert"
	monitorscheduledqueryrulesalertv2 "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorscheduledqueryrulesalertv2"
	monitorscheduledqueryruleslog "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorscheduledqueryruleslog"
	monitorworkspace "github.com/upbound/provider-azure/internal/controller/namespaced/insights/monitorworkspace"
	application "github.com/upbound/provider-azure/internal/controller/namespaced/iotcentral/application"
	applicationnetworkruleset "github.com/upbound/provider-azure/internal/controller/namespaced/iotcentral/applicationnetworkruleset"
	accesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/accesspolicy"
	certificatekeyvault "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificate"
	certificatecontacts "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificatecontacts"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/secret"
	vault "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/vault"
	attacheddatabaseconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/attacheddatabaseconfiguration"
	clusterkusto "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/cluster"
	clustermanagedprivateendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/clustermanagedprivateendpoint"
	clusterprincipalassignment "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/clusterprincipalassignment"
	databasekusto "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/database"
	databaseprincipalassignment "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/databaseprincipalassignment"
	eventgriddataconnection "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/eventgriddataconnection"
	eventhubdataconnection "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/eventhubdataconnection"
	iothubdataconnection "github.com/upbound/provider-azure/internal/controller/namespaced/kusto/iothubdataconnection"
	loadtest "github.com/upbound/provider-azure/internal/controller/namespaced/loadtestservice/loadtest"
	appactioncustom "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appactioncustom"
	appactionhttp "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appactionhttp"
	appintegrationaccount "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appintegrationaccount"
	appintegrationaccountbatchconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountpartner "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appintegrationaccountpartner"
	appintegrationaccountschema "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appintegrationaccountsession"
	apptriggercustom "github.com/upbound/provider-azure/internal/controller/namespaced/logic/apptriggercustom"
	apptriggerhttprequest "github.com/upbound/provider-azure/internal/controller/namespaced/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/upbound/provider-azure/internal/controller/namespaced/logic/apptriggerrecurrence"
	appworkflow "github.com/upbound/provider-azure/internal/controller/namespaced/logic/appworkflow"
	aifoundry "github.com/upbound/provider-azure/internal/controller/namespaced/machinelearningservices/aifoundry"
	computecluster "github.com/upbound/provider-azure/internal/controller/namespaced/machinelearningservices/computecluster"
	computeinstance "github.com/upbound/provider-azure/internal/controller/namespaced/machinelearningservices/computeinstance"
	synapsespark "github.com/upbound/provider-azure/internal/controller/namespaced/machinelearningservices/synapsespark"
	workspacemachinelearningservices "github.com/upbound/provider-azure/internal/controller/namespaced/machinelearningservices/workspace"
	maintenanceassignmentdedicatedhost "github.com/upbound/provider-azure/internal/controller/namespaced/maintenance/maintenanceassignmentdedicatedhost"
	maintenanceassignmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/maintenance/maintenanceassignmentvirtualmachine"
	maintenanceconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/maintenance/maintenanceconfiguration"
	federatedidentitycredential "github.com/upbound/provider-azure/internal/controller/namespaced/managedidentity/federatedidentitycredential"
	userassignedidentity "github.com/upbound/provider-azure/internal/controller/namespaced/managedidentity/userassignedidentity"
	managementgroup "github.com/upbound/provider-azure/internal/controller/namespaced/management/managementgroup"
	managementgroupsubscriptionassociation "github.com/upbound/provider-azure/internal/controller/namespaced/management/managementgroupsubscriptionassociation"
	accountmaps "github.com/upbound/provider-azure/internal/controller/namespaced/maps/account"
	creator "github.com/upbound/provider-azure/internal/controller/namespaced/maps/creator"
	marketplaceagreement "github.com/upbound/provider-azure/internal/controller/namespaced/marketplaceordering/marketplaceagreement"
	spatialanchorsaccount "github.com/upbound/provider-azure/internal/controller/namespaced/mixedreality/spatialanchorsaccount"
	accountnetapp "github.com/upbound/provider-azure/internal/controller/namespaced/netapp/account"
	pool "github.com/upbound/provider-azure/internal/controller/namespaced/netapp/pool"
	snapshotnetapp "github.com/upbound/provider-azure/internal/controller/namespaced/netapp/snapshot"
	snapshotpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/netapp/snapshotpolicy"
	volume "github.com/upbound/provider-azure/internal/controller/namespaced/netapp/volume"
	applicationgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/applicationgateway"
	applicationsecuritygroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/applicationsecuritygroup"
	bastionhost "github.com/upbound/provider-azure/internal/controller/namespaced/network/bastionhost"
	connectionmonitor "github.com/upbound/provider-azure/internal/controller/namespaced/network/connectionmonitor"
	ddosprotectionplan "github.com/upbound/provider-azure/internal/controller/namespaced/network/ddosprotectionplan"
	dnsaaaarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsaaaarecord"
	dnsarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsarecord"
	dnscaarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnscaarecord"
	dnscnamerecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnscnamerecord"
	dnsmxrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsmxrecord"
	dnsnsrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsnsrecord"
	dnsptrrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnsptrrecord"
	dnssrvrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnssrvrecord"
	dnstxtrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnstxtrecord"
	dnszone "github.com/upbound/provider-azure/internal/controller/namespaced/network/dnszone"
	expressroutecircuit "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressrouteconnection"
	expressroutegateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressroutegateway"
	expressrouteport "github.com/upbound/provider-azure/internal/controller/namespaced/network/expressrouteport"
	firewall "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewall"
	firewallapplicationrulecollection "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicynetwork "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/upbound/provider-azure/internal/controller/namespaced/network/frontdoorrulesengine"
	ipgroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/ipgroup"
	loadbalancer "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancer"
	loadbalancerbackendaddresspool "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerbackendaddresspool"
	loadbalancerbackendaddresspooladdress "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerbackendaddresspooladdress"
	loadbalancernatpool "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancernatpool"
	loadbalancernatrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancernatrule"
	loadbalanceroutboundrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalanceroutboundrule"
	loadbalancerprobe "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerprobe"
	loadbalancerrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/loadbalancerrule"
	localnetworkgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/localnetworkgateway"
	manager "github.com/upbound/provider-azure/internal/controller/namespaced/network/manager"
	manageripampool "github.com/upbound/provider-azure/internal/controller/namespaced/network/manageripampool"
	managermanagementgroupconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/managermanagementgroupconnection"
	managernetworkgroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/managernetworkgroup"
	managerroutingconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/network/managerroutingconfiguration"
	managerstaticmember "github.com/upbound/provider-azure/internal/controller/namespaced/network/managerstaticmember"
	managersubscriptionconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/managersubscriptionconnection"
	managerverifierworkspace "github.com/upbound/provider-azure/internal/controller/namespaced/network/managerverifierworkspace"
	natgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/natgateway"
	natgatewaypublicipassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/networkinterfacesecuritygroupassociation"
	packetcapture "github.com/upbound/provider-azure/internal/controller/namespaced/network/packetcapture"
	pointtositevpngateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsaaaarecord"
	privatednsarecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsarecord"
	privatednscnamerecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednscnamerecord"
	privatednsmxrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsmxrecord"
	privatednsptrrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsptrrecord"
	privatednsresolver "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolver"
	privatednsresolverdnsforwardingruleset "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolverdnsforwardingruleset"
	privatednsresolverforwardingrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolverforwardingrule"
	privatednsresolverinboundendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolverinboundendpoint"
	privatednsresolveroutboundendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolveroutboundendpoint"
	privatednsresolvervirtualnetworklink "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednsresolvervirtualnetworklink"
	privatednssrvrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednstxtrecord"
	privatednszone "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/privateendpoint"
	privateendpointapplicationsecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/privateendpointapplicationsecuritygroupassociation"
	privatelinkservice "github.com/upbound/provider-azure/internal/controller/namespaced/network/privatelinkservice"
	profilenetwork "github.com/upbound/provider-azure/internal/controller/namespaced/network/profile"
	publicip "github.com/upbound/provider-azure/internal/controller/namespaced/network/publicip"
	publicipprefix "github.com/upbound/provider-azure/internal/controller/namespaced/network/publicipprefix"
	route "github.com/upbound/provider-azure/internal/controller/namespaced/network/route"
	routefilter "github.com/upbound/provider-azure/internal/controller/namespaced/network/routefilter"
	routemap "github.com/upbound/provider-azure/internal/controller/namespaced/network/routemap"
	routeserver "github.com/upbound/provider-azure/internal/controller/namespaced/network/routeserver"
	routeserverbgpconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/routeserverbgpconnection"
	routetable "github.com/upbound/provider-azure/internal/controller/namespaced/network/routetable"
	securitygroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/securitygroup"
	securityrule "github.com/upbound/provider-azure/internal/controller/namespaced/network/securityrule"
	subnet "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/provider-azure/internal/controller/namespaced/network/subnetserviceendpointstoragepolicy"
	trafficmanagerazureendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagerazureendpoint"
	trafficmanagerexternalendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagerexternalendpoint"
	trafficmanagernestedendpoint "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagernestedendpoint"
	trafficmanagerprofile "github.com/upbound/provider-azure/internal/controller/namespaced/network/trafficmanagerprofile"
	virtualhub "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhub"
	virtualhubconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubconnection"
	virtualhubip "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubip"
	virtualhubroutetable "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubroutetable"
	virtualhubroutetableroute "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubroutetableroute"
	virtualhubsecuritypartnerprovider "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualhubsecuritypartnerprovider"
	virtualnetworknetwork "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetwork"
	virtualnetworkdnsservers "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetworkdnsservers"
	virtualnetworkgateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/internal/controller/namespaced/network/virtualwan"
	vpngateway "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpngateway"
	vpngatewayconnection "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpnserverconfiguration"
	vpnserverconfigurationpolicygroup "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpnserverconfigurationpolicygroup"
	vpnsite "github.com/upbound/provider-azure/internal/controller/namespaced/network/vpnsite"
	watcher "github.com/upbound/provider-azure/internal/controller/namespaced/network/watcher"
	watcherflowlog "github.com/upbound/provider-azure/internal/controller/namespaced/network/watcherflowlog"
	webapplicationfirewallpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/network/webapplicationfirewallpolicy"
	authorizationrulenotificationhubs "github.com/upbound/provider-azure/internal/controller/namespaced/notificationhubs/authorizationrule"
	notificationhub "github.com/upbound/provider-azure/internal/controller/namespaced/notificationhubs/notificationhub"
	notificationhubnamespace "github.com/upbound/provider-azure/internal/controller/namespaced/notificationhubs/notificationhubnamespace"
	loganalyticsdataexportrule "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsdataexportrule"
	loganalyticsdatasourcewindowsevent "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsdatasourcewindowsevent"
	loganalyticsdatasourcewindowsperformancecounter "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsdatasourcewindowsperformancecounter"
	loganalyticslinkedservice "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticslinkedservice"
	loganalyticslinkedstorageaccount "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticslinkedstorageaccount"
	loganalyticsquerypack "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsquerypack"
	loganalyticsquerypackquery "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticsquerypackquery"
	loganalyticssavedsearch "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/loganalyticssavedsearch"
	workspaceoperationalinsights "github.com/upbound/provider-azure/internal/controller/namespaced/operationalinsights/workspace"
	loganalyticssolution "github.com/upbound/provider-azure/internal/controller/namespaced/operationsmanagement/loganalyticssolution"
	contactprofile "github.com/upbound/provider-azure/internal/controller/namespaced/orbital/contactprofile"
	spacecraft "github.com/upbound/provider-azure/internal/controller/namespaced/orbital/spacecraft"
	resourcepolicyremediation "github.com/upbound/provider-azure/internal/controller/namespaced/policyinsights/resourcepolicyremediation"
	subscriptionpolicyremediation "github.com/upbound/provider-azure/internal/controller/namespaced/policyinsights/subscriptionpolicyremediation"
	dashboard "github.com/upbound/provider-azure/internal/controller/namespaced/portal/dashboard"
	powerbiembedded "github.com/upbound/provider-azure/internal/controller/namespaced/powerbidedicated/powerbiembedded"
	providerconfig "github.com/upbound/provider-azure/internal/controller/namespaced/providerconfig"
	accountpurview "github.com/upbound/provider-azure/internal/controller/namespaced/purview/account"
	backupcontainerstorageaccount "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backupcontainerstorageaccount"
	backuppolicyfileshare "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backuppolicyfileshare"
	backuppolicyvm "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backuppolicyvm"
	backuppolicyvmworkload "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backuppolicyvmworkload"
	backupprotectedfileshare "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backupprotectedfileshare"
	backupprotectedvm "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/backupprotectedvm"
	siterecoveryfabric "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryfabric"
	siterecoverynetworkmapping "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoverynetworkmapping"
	siterecoveryprotectioncontainer "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryprotectioncontainer"
	siterecoveryprotectioncontainermapping "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryprotectioncontainermapping"
	siterecoveryreplicationpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/siterecoveryreplicationpolicy"
	vaultrecoveryservices "github.com/upbound/provider-azure/internal/controller/namespaced/recoveryservices/vault"
	eventrelaynamespace "github.com/upbound/provider-azure/internal/controller/namespaced/relay/eventrelaynamespace"
	hybridconnection "github.com/upbound/provider-azure/internal/controller/namespaced/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/relay/hybridconnectionauthorizationrule"
	namespaceauthorizationrulerelay "github.com/upbound/provider-azure/internal/controller/namespaced/relay/namespaceauthorizationrule"
	resourcedeploymentscriptazurecli "github.com/upbound/provider-azure/internal/controller/namespaced/resources/resourcedeploymentscriptazurecli"
	resourcedeploymentscriptazurepowershell "github.com/upbound/provider-azure/internal/controller/namespaced/resources/resourcedeploymentscriptazurepowershell"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/namespaced/resources/resourcegrouptemplatedeployment"
	subscriptiontemplatedeployment "github.com/upbound/provider-azure/internal/controller/namespaced/resources/subscriptiontemplatedeployment"
	servicesearch "github.com/upbound/provider-azure/internal/controller/namespaced/search/service"
	sharedprivatelinkservice "github.com/upbound/provider-azure/internal/controller/namespaced/search/sharedprivatelinkservice"
	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/namespaced/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/namespaced/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/namespaced/security/iotsecuritysolution"
	securitycenterassessment "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterassessment"
	securitycenterassessmentpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentercontact"
	securitycenterservervulnerabilityassessmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterservervulnerabilityassessmentvirtualmachine"
	securitycentersetting "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterworkspace"
	storagedefender "github.com/upbound/provider-azure/internal/controller/namespaced/security/storagedefender"
	sentinelalertrulefusion "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelalertrulemssecurityincident"
	sentinelautomationrule "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelautomationrule"
	sentineldataconnectoriot "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentineldataconnectoriot"
	sentinelloganalyticsworkspaceonboarding "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelloganalyticsworkspaceonboarding"
	sentinelwatchlist "github.com/upbound/provider-azure/internal/controller/namespaced/securityinsights/sentinelwatchlist"
	namespaceauthorizationruleservicebus "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfigservicebus "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/namespacedisasterrecoveryconfig"
	queue "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/queue"
	queueauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/queueauthorizationrule"
	servicebusnamespace "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/servicebusnamespace"
	subscriptionservicebus "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/subscription"
	subscriptionrule "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/subscriptionrule"
	topicservicebus "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/topic"
	topicauthorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/servicebus/topicauthorizationrule"
	clusterservicefabric "github.com/upbound/provider-azure/internal/controller/namespaced/servicefabric/cluster"
	managedcluster "github.com/upbound/provider-azure/internal/controller/namespaced/servicefabric/managedcluster"
	springcloudconnection "github.com/upbound/provider-azure/internal/controller/namespaced/servicelinker/springcloudconnection"
	networkacl "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/networkacl"
	servicesignalrservice "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/service"
	signalrsharedprivatelinkresource "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/signalrsharedprivatelinkresource"
	webpubsub "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/webpubsub"
	webpubsubhub "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/webpubsubhub"
	webpubsubnetworkacl "github.com/upbound/provider-azure/internal/controller/namespaced/signalrservice/webpubsubnetworkacl"
	managedapplicationdefinition "github.com/upbound/provider-azure/internal/controller/namespaced/solutions/managedapplicationdefinition"
	cloudapplicationliveview "github.com/upbound/provider-azure/internal/controller/namespaced/spring/cloudapplicationliveview"
	mssqldatabase "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlfirewallrule"
	mssqljobagent "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqljobagent"
	mssqljobcredential "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqljobcredential"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancetransparentdataencryption "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancetransparentdataencryption"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlvirtualnetworkrule"
	accountstorage "github.com/upbound/provider-azure/internal/controller/namespaced/storage/account"
	accountlocaluser "github.com/upbound/provider-azure/internal/controller/namespaced/storage/accountlocaluser"
	accountnetworkrules "github.com/upbound/provider-azure/internal/controller/namespaced/storage/accountnetworkrules"
	blob "github.com/upbound/provider-azure/internal/controller/namespaced/storage/blob"
	blobinventorypolicy "github.com/upbound/provider-azure/internal/controller/namespaced/storage/blobinventorypolicy"
	container "github.com/upbound/provider-azure/internal/controller/namespaced/storage/container"
	containerimmutabilitypolicy "github.com/upbound/provider-azure/internal/controller/namespaced/storage/containerimmutabilitypolicy"
	datalakegen2filesystem "github.com/upbound/provider-azure/internal/controller/namespaced/storage/datalakegen2filesystem"
	datalakegen2path "github.com/upbound/provider-azure/internal/controller/namespaced/storage/datalakegen2path"
	encryptionscope "github.com/upbound/provider-azure/internal/controller/namespaced/storage/encryptionscope"
	managementpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/storage/managementpolicy"
	objectreplication "github.com/upbound/provider-azure/internal/controller/namespaced/storage/objectreplication"
	queuestorage "github.com/upbound/provider-azure/internal/controller/namespaced/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/namespaced/storage/share"
	sharedirectory "github.com/upbound/provider-azure/internal/controller/namespaced/storage/sharedirectory"
	tablestorage "github.com/upbound/provider-azure/internal/controller/namespaced/storage/table"
	tableentity "github.com/upbound/provider-azure/internal/controller/namespaced/storage/tableentity"
	hpccache "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/namespaced/storagecache/hpccachenfstarget"
	storagesync "github.com/upbound/provider-azure/internal/controller/namespaced/storagesync/storagesync"
	clusterstreamanalytics "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/functionjavascriptuda"
	job "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/job"
	managedprivateendpointstreamanalytics "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/managedprivateendpoint"
	outputblob "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputblob"
	outputeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputeventhub"
	outputfunction "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputfunction"
	outputmssql "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputmssql"
	outputpowerbi "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputpowerbi"
	outputservicebusqueue "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputservicebusqueue"
	outputservicebustopic "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputservicebustopic"
	outputsynapse "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputsynapse"
	outputtable "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/outputtable"
	referenceinputblob "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/referenceinputblob"
	referenceinputmssql "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/referenceinputmssql"
	streaminputblob "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/streaminputblob"
	streaminputeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/streaminputeventhub"
	streaminputiothub "github.com/upbound/provider-azure/internal/controller/namespaced/streamanalytics/streaminputiothub"
	firewallrulesynapse "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/firewallrule"
	integrationruntimeazuresynapse "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/integrationruntimeazure"
	integrationruntimeselfhostedsynapse "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/linkedservice"
	managedprivateendpointsynapse "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/privatelinkhub"
	roleassignmentsynapse "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/roleassignment"
	sparkpool "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sparkpool"
	sqlpool "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolworkloadclassifier "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolworkloadclassifier"
	sqlpoolworkloadgroup "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/sqlpoolworkloadgroup"
	workspacesynapse "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspace"
	workspaceaadadmin "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspaceaadadmin"
	workspaceextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspaceextendedauditingpolicy"
	workspacesecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspacesecurityalertpolicy"
	workspacesqlaadadmin "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspacesqlaadadmin"
	workspacevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/synapse/workspacevulnerabilityassessment"
	appactiveslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/appactiveslot"
	apphybridconnection "github.com/upbound/provider-azure/internal/controller/namespaced/web/apphybridconnection"
	appserviceplan "github.com/upbound/provider-azure/internal/controller/namespaced/web/appserviceplan"
	functionapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionapp"
	functionappactiveslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionappactiveslot"
	functionappfunction "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionappfunction"
	functionapphybridconnection "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionapphybridconnection"
	functionappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/functionappslot"
	linuxfunctionapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxfunctionapp"
	linuxfunctionappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxfunctionappslot"
	linuxwebapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxwebapp"
	linuxwebappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/linuxwebappslot"
	serviceplan "github.com/upbound/provider-azure/internal/controller/namespaced/web/serviceplan"
	sourcecontroltoken "github.com/upbound/provider-azure/internal/controller/namespaced/web/sourcecontroltoken"
	staticsite "github.com/upbound/provider-azure/internal/controller/namespaced/web/staticsite"
	windowsfunctionapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowsfunctionapp"
	windowsfunctionappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowsfunctionappslot"
	windowswebapp "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowswebapp"
	windowswebappslot "github.com/upbound/provider-azure/internal/controller/namespaced/web/windowswebappslot"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoralertprocessingruleactiongroup.Setup,
		monitoralertprocessingrulesuppression.Setup,
		monitoralertprometheusrulegroup.Setup,
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
		group.Setup,
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
		policyfragment.Setup,
		product.Setup,
		productapi.Setup,
		productgroup.Setup,
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
		frontdoorsecret.Setup,
		frontdoorsecuritypolicy.Setup,
		profile.Setup,
		appservicecertificateorder.Setup,
		accountcognitiveservices.Setup,
		accountraiblocklist.Setup,
		accountraipolicy.Setup,
		aiservices.Setup,
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
		containerjob.Setup,
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
		backupinstancepostgresqlflexibleserver.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydisk.Setup,
		backuppolicykubernetescluster.Setup,
		backuppolicypostgresql.Setup,
		backuppolicypostgresqlflexibleserver.Setup,
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
		systemtopiceventsubscription.Setup,
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
		monitorworkspace.Setup,
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
		aifoundry.Setup,
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
		manageripampool.Setup,
		managermanagementgroupconnection.Setup,
		managernetworkgroup.Setup,
		managerroutingconfiguration.Setup,
		managerstaticmember.Setup,
		managersubscriptionconnection.Setup,
		managerverifierworkspace.Setup,
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
		privatednsresolvervirtualnetworklink.Setup,
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
		virtualnetworkdnsservers.Setup,
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

// SetupGated_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoralertprocessingruleactiongroup.SetupGated,
		monitoralertprocessingrulesuppression.SetupGated,
		monitoralertprometheusrulegroup.SetupGated,
		monitorsmartdetectoralertrule.SetupGated,
		server.SetupGated,
		api.SetupGated,
		apidiagnostic.SetupGated,
		apioperation.SetupGated,
		apioperationpolicy.SetupGated,
		apioperationtag.SetupGated,
		apipolicy.SetupGated,
		apirelease.SetupGated,
		apischema.SetupGated,
		apitag.SetupGated,
		apiversionset.SetupGated,
		authorizationserver.SetupGated,
		backend.SetupGated,
		certificate.SetupGated,
		customdomain.SetupGated,
		diagnostic.SetupGated,
		emailtemplate.SetupGated,
		gateway.SetupGated,
		gatewayapi.SetupGated,
		globalschema.SetupGated,
		group.SetupGated,
		identityprovideraad.SetupGated,
		identityproviderfacebook.SetupGated,
		identityprovidergoogle.SetupGated,
		identityprovidermicrosoft.SetupGated,
		identityprovidertwitter.SetupGated,
		logger.SetupGated,
		management.SetupGated,
		namedvalue.SetupGated,
		notificationrecipientemail.SetupGated,
		notificationrecipientuser.SetupGated,
		openidconnectprovider.SetupGated,
		policy.SetupGated,
		policyfragment.SetupGated,
		product.SetupGated,
		productapi.SetupGated,
		productgroup.SetupGated,
		productpolicy.SetupGated,
		producttag.SetupGated,
		rediscache.SetupGated,
		subscription.SetupGated,
		tag.SetupGated,
		user.SetupGated,
		configuration.SetupGated,
		springcloudaccelerator.SetupGated,
		springcloudactivedeployment.SetupGated,
		springcloudapiportal.SetupGated,
		springcloudapiportalcustomdomain.SetupGated,
		springcloudapp.SetupGated,
		springcloudappcosmosdbassociation.SetupGated,
		springcloudappmysqlassociation.SetupGated,
		springcloudappredisassociation.SetupGated,
		springcloudbuilddeployment.SetupGated,
		springcloudbuilder.SetupGated,
		springcloudbuildpackbinding.SetupGated,
		springcloudcertificate.SetupGated,
		springcloudconfigurationservice.SetupGated,
		springcloudcontainerdeployment.SetupGated,
		springcloudcustomdomain.SetupGated,
		springcloudcustomizedaccelerator.SetupGated,
		springclouddevtoolportal.SetupGated,
		springcloudgateway.SetupGated,
		springcloudgatewaycustomdomain.SetupGated,
		springcloudjavadeployment.SetupGated,
		springcloudservice.SetupGated,
		springcloudstorage.SetupGated,
		provider.SetupGated,
		managementgrouppolicyassignment.SetupGated,
		managementgrouppolicyexemption.SetupGated,
		managementlock.SetupGated,
		pimactiveroleassignment.SetupGated,
		pimeligibleroleassignment.SetupGated,
		policydefinition.SetupGated,
		policysetdefinition.SetupGated,
		resourcegrouppolicyassignment.SetupGated,
		resourcegrouppolicyexemption.SetupGated,
		resourcepolicyassignment.SetupGated,
		resourcepolicyexemption.SetupGated,
		roleassignment.SetupGated,
		roledefinition.SetupGated,
		subscriptionpolicyassignment.SetupGated,
		subscriptionpolicyexemption.SetupGated,
		trustedaccessrolebinding.SetupGated,
		account.SetupGated,
		connection.SetupGated,
		connectionclassiccertificate.SetupGated,
		connectiontype.SetupGated,
		credential.SetupGated,
		hybridrunbookworkergroup.SetupGated,
		module.SetupGated,
		runbook.SetupGated,
		schedule.SetupGated,
		variablebool.SetupGated,
		variabledatetime.SetupGated,
		variableint.SetupGated,
		variablestring.SetupGated,
		webhook.SetupGated,
		resourcegroup.SetupGated,
		resourceproviderregistration.SetupGated,
		subscriptionazure.SetupGated,
		cluster.SetupGated,
		botchannelalexa.SetupGated,
		botchanneldirectline.SetupGated,
		botchannelline.SetupGated,
		botchannelmsteams.SetupGated,
		botchannelslack.SetupGated,
		botchannelsms.SetupGated,
		botchannelsregistration.SetupGated,
		botchannelwebchat.SetupGated,
		botconnection.SetupGated,
		botwebapp.SetupGated,
		rediscachecache.SetupGated,
		rediscacheaccesspolicy.SetupGated,
		rediscacheaccesspolicyassignment.SetupGated,
		redisenterprisecluster.SetupGated,
		redisenterprisedatabase.SetupGated,
		redisfirewallrule.SetupGated,
		redislinkedserver.SetupGated,
		endpoint.SetupGated,
		frontdoorcustomdomain.SetupGated,
		frontdoorcustomdomainassociation.SetupGated,
		frontdoorendpoint.SetupGated,
		frontdoorfirewallpolicy.SetupGated,
		frontdoororigin.SetupGated,
		frontdoororigingroup.SetupGated,
		frontdoorprofile.SetupGated,
		frontdoorroute.SetupGated,
		frontdoorrule.SetupGated,
		frontdoorruleset.SetupGated,
		frontdoorsecret.SetupGated,
		frontdoorsecuritypolicy.SetupGated,
		profile.SetupGated,
		appservicecertificateorder.SetupGated,
		accountcognitiveservices.SetupGated,
		accountraiblocklist.SetupGated,
		accountraipolicy.SetupGated,
		aiservices.SetupGated,
		deployment.SetupGated,
		service.SetupGated,
		availabilityset.SetupGated,
		capacityreservation.SetupGated,
		capacityreservationgroup.SetupGated,
		dedicatedhost.SetupGated,
		diskaccess.SetupGated,
		diskencryptionset.SetupGated,
		galleryapplication.SetupGated,
		galleryapplicationversion.SetupGated,
		image.SetupGated,
		linuxvirtualmachine.SetupGated,
		linuxvirtualmachinescaleset.SetupGated,
		manageddisk.SetupGated,
		manageddisksastoken.SetupGated,
		orchestratedvirtualmachinescaleset.SetupGated,
		proximityplacementgroup.SetupGated,
		sharedimage.SetupGated,
		sharedimagegallery.SetupGated,
		snapshot.SetupGated,
		sshpublickey.SetupGated,
		virtualmachinedatadiskattachment.SetupGated,
		virtualmachineextension.SetupGated,
		virtualmachineruncommand.SetupGated,
		windowsvirtualmachine.SetupGated,
		windowsvirtualmachinescaleset.SetupGated,
		ledger.SetupGated,
		budgetmanagementgroup.SetupGated,
		budgetresourcegroup.SetupGated,
		budgetsubscription.SetupGated,
		containerapp.SetupGated,
		containerjob.SetupGated,
		customdomaincontainerapp.SetupGated,
		environment.SetupGated,
		environmentcertificate.SetupGated,
		environmentcustomdomain.SetupGated,
		environmentdaprcomponent.SetupGated,
		environmentstorage.SetupGated,
		agentpool.SetupGated,
		containerconnectedregistry.SetupGated,
		registry.SetupGated,
		scopemap.SetupGated,
		token.SetupGated,
		tokenpassword.SetupGated,
		webhookcontainerregistry.SetupGated,
		kubernetescluster.SetupGated,
		kubernetesclusterextension.SetupGated,
		kubernetesclusternodepool.SetupGated,
		kubernetesfleetmanager.SetupGated,
		accountcosmosdb.SetupGated,
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
		costanomalyalert.SetupGated,
		resourcegroupcostmanagementexport.SetupGated,
		subscriptioncostmanagementexport.SetupGated,
		customprovider.SetupGated,
		device.SetupGated,
		accessconnector.SetupGated,
		workspace.SetupGated,
		workspacerootdbfscustomermanagedkey.SetupGated,
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
		databasemigrationproject.SetupGated,
		databasemigrationservice.SetupGated,
		backupinstanceblobstorage.SetupGated,
		backupinstancedisk.SetupGated,
		backupinstancekubernetescluster.SetupGated,
		backupinstancepostgresql.SetupGated,
		backupinstancepostgresqlflexibleserver.SetupGated,
		backuppolicyblobstorage.SetupGated,
		backuppolicydisk.SetupGated,
		backuppolicykubernetescluster.SetupGated,
		backuppolicypostgresql.SetupGated,
		backuppolicypostgresqlflexibleserver.SetupGated,
		backupvault.SetupGated,
		resourceguard.SetupGated,
		accountdatashare.SetupGated,
		datasetblobstorage.SetupGated,
		datasetdatalakegen2.SetupGated,
		datasetkustocluster.SetupGated,
		datasetkustodatabase.SetupGated,
		datashare.SetupGated,
		flexibledatabase.SetupGated,
		flexibleserver.SetupGated,
		flexibleserverconfiguration.SetupGated,
		flexibleserverfirewallrule.SetupGated,
		activedirectoryadministrator.SetupGated,
		configurationdbforpostgresql.SetupGated,
		database.SetupGated,
		firewallrule.SetupGated,
		flexibleserverdbforpostgresql.SetupGated,
		flexibleserveractivedirectoryadministrator.SetupGated,
		flexibleserverconfigurationdbforpostgresql.SetupGated,
		flexibleserverdatabase.SetupGated,
		flexibleserverfirewallruledbforpostgresql.SetupGated,
		serverdbforpostgresql.SetupGated,
		serverkey.SetupGated,
		virtualnetworkrule.SetupGated,
		iothub.SetupGated,
		iothubcertificate.SetupGated,
		iothubconsumergroup.SetupGated,
		iothubdps.SetupGated,
		iothubdpscertificate.SetupGated,
		iothubdpssharedaccesspolicy.SetupGated,
		iothubendpointeventhub.SetupGated,
		iothubendpointservicebusqueue.SetupGated,
		iothubendpointservicebustopic.SetupGated,
		iothubendpointstoragecontainer.SetupGated,
		iothubenrichment.SetupGated,
		iothubfallbackroute.SetupGated,
		iothubroute.SetupGated,
		iothubsharedaccesspolicy.SetupGated,
		iothubdeviceupdateaccount.SetupGated,
		iothubdeviceupdateinstance.SetupGated,
		globalvmshutdownschedule.SetupGated,
		lab.SetupGated,
		linuxvirtualmachinedevtestlab.SetupGated,
		policydevtestlab.SetupGated,
		scheduledevtestlab.SetupGated,
		virtualnetwork.SetupGated,
		windowsvirtualmachinedevtestlab.SetupGated,
		instance.SetupGated,
		cloudelasticsearch.SetupGated,
		domain.SetupGated,
		domaintopic.SetupGated,
		eventsubscription.SetupGated,
		systemtopic.SetupGated,
		systemtopiceventsubscription.SetupGated,
		topic.SetupGated,
		authorizationrule.SetupGated,
		consumergroup.SetupGated,
		eventhub.SetupGated,
		eventhubnamespace.SetupGated,
		namespaceauthorizationrule.SetupGated,
		namespacedisasterrecoveryconfig.SetupGated,
		namespaceschemagroup.SetupGated,
		serverfluidrelay.SetupGated,
		policyvirtualmachineconfigurationassignment.SetupGated,
		hadoopcluster.SetupGated,
		hbasecluster.SetupGated,
		interactivequerycluster.SetupGated,
		kafkacluster.SetupGated,
		sparkcluster.SetupGated,
		healthbot.SetupGated,
		healthcaredicomservice.SetupGated,
		healthcarefhirservice.SetupGated,
		healthcaremedtechservice.SetupGated,
		healthcaremedtechservicefhirdestination.SetupGated,
		healthcareservice.SetupGated,
		healthcareworkspace.SetupGated,
		applicationinsights.SetupGated,
		applicationinsightsanalyticsitem.SetupGated,
		applicationinsightsapikey.SetupGated,
		applicationinsightssmartdetectionrule.SetupGated,
		applicationinsightsstandardwebtest.SetupGated,
		applicationinsightswebtest.SetupGated,
		applicationinsightsworkbook.SetupGated,
		applicationinsightsworkbooktemplate.SetupGated,
		monitoractiongroup.SetupGated,
		monitoractivitylogalert.SetupGated,
		monitorautoscalesetting.SetupGated,
		monitordatacollectionendpoint.SetupGated,
		monitordatacollectionrule.SetupGated,
		monitordatacollectionruleassociation.SetupGated,
		monitordiagnosticsetting.SetupGated,
		monitormetricalert.SetupGated,
		monitorprivatelinkscope.SetupGated,
		monitorprivatelinkscopedservice.SetupGated,
		monitorscheduledqueryrulesalert.SetupGated,
		monitorscheduledqueryrulesalertv2.SetupGated,
		monitorscheduledqueryruleslog.SetupGated,
		monitorworkspace.SetupGated,
		application.SetupGated,
		applicationnetworkruleset.SetupGated,
		accesspolicy.SetupGated,
		certificatekeyvault.SetupGated,
		certificatecontacts.SetupGated,
		certificateissuer.SetupGated,
		key.SetupGated,
		managedhardwaresecuritymodule.SetupGated,
		managedstorageaccount.SetupGated,
		managedstorageaccountsastokendefinition.SetupGated,
		secret.SetupGated,
		vault.SetupGated,
		attacheddatabaseconfiguration.SetupGated,
		clusterkusto.SetupGated,
		clustermanagedprivateendpoint.SetupGated,
		clusterprincipalassignment.SetupGated,
		databasekusto.SetupGated,
		databaseprincipalassignment.SetupGated,
		eventgriddataconnection.SetupGated,
		eventhubdataconnection.SetupGated,
		iothubdataconnection.SetupGated,
		loadtest.SetupGated,
		appactioncustom.SetupGated,
		appactionhttp.SetupGated,
		appintegrationaccount.SetupGated,
		appintegrationaccountbatchconfiguration.SetupGated,
		appintegrationaccountpartner.SetupGated,
		appintegrationaccountschema.SetupGated,
		appintegrationaccountsession.SetupGated,
		apptriggercustom.SetupGated,
		apptriggerhttprequest.SetupGated,
		apptriggerrecurrence.SetupGated,
		appworkflow.SetupGated,
		aifoundry.SetupGated,
		computecluster.SetupGated,
		computeinstance.SetupGated,
		synapsespark.SetupGated,
		workspacemachinelearningservices.SetupGated,
		maintenanceassignmentdedicatedhost.SetupGated,
		maintenanceassignmentvirtualmachine.SetupGated,
		maintenanceconfiguration.SetupGated,
		federatedidentitycredential.SetupGated,
		userassignedidentity.SetupGated,
		managementgroup.SetupGated,
		managementgroupsubscriptionassociation.SetupGated,
		accountmaps.SetupGated,
		creator.SetupGated,
		marketplaceagreement.SetupGated,
		spatialanchorsaccount.SetupGated,
		accountnetapp.SetupGated,
		pool.SetupGated,
		snapshotnetapp.SetupGated,
		snapshotpolicy.SetupGated,
		volume.SetupGated,
		applicationgateway.SetupGated,
		applicationsecuritygroup.SetupGated,
		bastionhost.SetupGated,
		connectionmonitor.SetupGated,
		ddosprotectionplan.SetupGated,
		dnsaaaarecord.SetupGated,
		dnsarecord.SetupGated,
		dnscaarecord.SetupGated,
		dnscnamerecord.SetupGated,
		dnsmxrecord.SetupGated,
		dnsnsrecord.SetupGated,
		dnsptrrecord.SetupGated,
		dnssrvrecord.SetupGated,
		dnstxtrecord.SetupGated,
		dnszone.SetupGated,
		expressroutecircuit.SetupGated,
		expressroutecircuitauthorization.SetupGated,
		expressroutecircuitconnection.SetupGated,
		expressroutecircuitpeering.SetupGated,
		expressrouteconnection.SetupGated,
		expressroutegateway.SetupGated,
		expressrouteport.SetupGated,
		firewall.SetupGated,
		firewallapplicationrulecollection.SetupGated,
		firewallnatrulecollection.SetupGated,
		firewallnetworkrulecollection.SetupGated,
		firewallpolicy.SetupGated,
		firewallpolicyrulecollectiongroup.SetupGated,
		frontdoor.SetupGated,
		frontdoorcustomhttpsconfiguration.SetupGated,
		frontdoorfirewallpolicynetwork.SetupGated,
		frontdoorrulesengine.SetupGated,
		ipgroup.SetupGated,
		loadbalancer.SetupGated,
		loadbalancerbackendaddresspool.SetupGated,
		loadbalancerbackendaddresspooladdress.SetupGated,
		loadbalancernatpool.SetupGated,
		loadbalancernatrule.SetupGated,
		loadbalanceroutboundrule.SetupGated,
		loadbalancerprobe.SetupGated,
		loadbalancerrule.SetupGated,
		localnetworkgateway.SetupGated,
		manager.SetupGated,
		manageripampool.SetupGated,
		managermanagementgroupconnection.SetupGated,
		managernetworkgroup.SetupGated,
		managerroutingconfiguration.SetupGated,
		managerstaticmember.SetupGated,
		managersubscriptionconnection.SetupGated,
		managerverifierworkspace.SetupGated,
		natgateway.SetupGated,
		natgatewaypublicipassociation.SetupGated,
		natgatewaypublicipprefixassociation.SetupGated,
		networkinterface.SetupGated,
		networkinterfaceapplicationsecuritygroupassociation.SetupGated,
		networkinterfacebackendaddresspoolassociation.SetupGated,
		networkinterfacenatruleassociation.SetupGated,
		networkinterfacesecuritygroupassociation.SetupGated,
		packetcapture.SetupGated,
		pointtositevpngateway.SetupGated,
		privatednsaaaarecord.SetupGated,
		privatednsarecord.SetupGated,
		privatednscnamerecord.SetupGated,
		privatednsmxrecord.SetupGated,
		privatednsptrrecord.SetupGated,
		privatednsresolver.SetupGated,
		privatednsresolverdnsforwardingruleset.SetupGated,
		privatednsresolverforwardingrule.SetupGated,
		privatednsresolverinboundendpoint.SetupGated,
		privatednsresolveroutboundendpoint.SetupGated,
		privatednsresolvervirtualnetworklink.SetupGated,
		privatednssrvrecord.SetupGated,
		privatednstxtrecord.SetupGated,
		privatednszone.SetupGated,
		privatednszonevirtualnetworklink.SetupGated,
		privateendpoint.SetupGated,
		privateendpointapplicationsecuritygroupassociation.SetupGated,
		privatelinkservice.SetupGated,
		profilenetwork.SetupGated,
		publicip.SetupGated,
		publicipprefix.SetupGated,
		route.SetupGated,
		routefilter.SetupGated,
		routemap.SetupGated,
		routeserver.SetupGated,
		routeserverbgpconnection.SetupGated,
		routetable.SetupGated,
		securitygroup.SetupGated,
		securityrule.SetupGated,
		subnet.SetupGated,
		subnetnatgatewayassociation.SetupGated,
		subnetnetworksecuritygroupassociation.SetupGated,
		subnetroutetableassociation.SetupGated,
		subnetserviceendpointstoragepolicy.SetupGated,
		trafficmanagerazureendpoint.SetupGated,
		trafficmanagerexternalendpoint.SetupGated,
		trafficmanagernestedendpoint.SetupGated,
		trafficmanagerprofile.SetupGated,
		virtualhub.SetupGated,
		virtualhubconnection.SetupGated,
		virtualhubip.SetupGated,
		virtualhubroutetable.SetupGated,
		virtualhubroutetableroute.SetupGated,
		virtualhubsecuritypartnerprovider.SetupGated,
		virtualnetworknetwork.SetupGated,
		virtualnetworkdnsservers.SetupGated,
		virtualnetworkgateway.SetupGated,
		virtualnetworkgatewayconnection.SetupGated,
		virtualnetworkpeering.SetupGated,
		virtualwan.SetupGated,
		vpngateway.SetupGated,
		vpngatewayconnection.SetupGated,
		vpnserverconfiguration.SetupGated,
		vpnserverconfigurationpolicygroup.SetupGated,
		vpnsite.SetupGated,
		watcher.SetupGated,
		watcherflowlog.SetupGated,
		webapplicationfirewallpolicy.SetupGated,
		authorizationrulenotificationhubs.SetupGated,
		notificationhub.SetupGated,
		notificationhubnamespace.SetupGated,
		loganalyticsdataexportrule.SetupGated,
		loganalyticsdatasourcewindowsevent.SetupGated,
		loganalyticsdatasourcewindowsperformancecounter.SetupGated,
		loganalyticslinkedservice.SetupGated,
		loganalyticslinkedstorageaccount.SetupGated,
		loganalyticsquerypack.SetupGated,
		loganalyticsquerypackquery.SetupGated,
		loganalyticssavedsearch.SetupGated,
		workspaceoperationalinsights.SetupGated,
		loganalyticssolution.SetupGated,
		contactprofile.SetupGated,
		spacecraft.SetupGated,
		resourcepolicyremediation.SetupGated,
		subscriptionpolicyremediation.SetupGated,
		dashboard.SetupGated,
		powerbiembedded.SetupGated,
		providerconfig.SetupGated,
		accountpurview.SetupGated,
		backupcontainerstorageaccount.SetupGated,
		backuppolicyfileshare.SetupGated,
		backuppolicyvm.SetupGated,
		backuppolicyvmworkload.SetupGated,
		backupprotectedfileshare.SetupGated,
		backupprotectedvm.SetupGated,
		siterecoveryfabric.SetupGated,
		siterecoverynetworkmapping.SetupGated,
		siterecoveryprotectioncontainer.SetupGated,
		siterecoveryprotectioncontainermapping.SetupGated,
		siterecoveryreplicationpolicy.SetupGated,
		vaultrecoveryservices.SetupGated,
		eventrelaynamespace.SetupGated,
		hybridconnection.SetupGated,
		hybridconnectionauthorizationrule.SetupGated,
		namespaceauthorizationrulerelay.SetupGated,
		resourcedeploymentscriptazurecli.SetupGated,
		resourcedeploymentscriptazurepowershell.SetupGated,
		resourcegrouptemplatedeployment.SetupGated,
		subscriptiontemplatedeployment.SetupGated,
		servicesearch.SetupGated,
		sharedprivatelinkservice.SetupGated,
		advancedthreatprotection.SetupGated,
		iotsecuritydevicegroup.SetupGated,
		iotsecuritysolution.SetupGated,
		securitycenterassessment.SetupGated,
		securitycenterassessmentpolicy.SetupGated,
		securitycenterautoprovisioning.SetupGated,
		securitycentercontact.SetupGated,
		securitycenterservervulnerabilityassessmentvirtualmachine.SetupGated,
		securitycentersetting.SetupGated,
		securitycentersubscriptionpricing.SetupGated,
		securitycenterworkspace.SetupGated,
		storagedefender.SetupGated,
		sentinelalertrulefusion.SetupGated,
		sentinelalertrulemachinelearningbehavioranalytics.SetupGated,
		sentinelalertrulemssecurityincident.SetupGated,
		sentinelautomationrule.SetupGated,
		sentineldataconnectoriot.SetupGated,
		sentinelloganalyticsworkspaceonboarding.SetupGated,
		sentinelwatchlist.SetupGated,
		namespaceauthorizationruleservicebus.SetupGated,
		namespacedisasterrecoveryconfigservicebus.SetupGated,
		queue.SetupGated,
		queueauthorizationrule.SetupGated,
		servicebusnamespace.SetupGated,
		subscriptionservicebus.SetupGated,
		subscriptionrule.SetupGated,
		topicservicebus.SetupGated,
		topicauthorizationrule.SetupGated,
		clusterservicefabric.SetupGated,
		managedcluster.SetupGated,
		springcloudconnection.SetupGated,
		networkacl.SetupGated,
		servicesignalrservice.SetupGated,
		signalrsharedprivatelinkresource.SetupGated,
		webpubsub.SetupGated,
		webpubsubhub.SetupGated,
		webpubsubnetworkacl.SetupGated,
		managedapplicationdefinition.SetupGated,
		cloudapplicationliveview.SetupGated,
		mssqldatabase.SetupGated,
		mssqldatabaseextendedauditingpolicy.SetupGated,
		mssqldatabasevulnerabilityassessmentrulebaseline.SetupGated,
		mssqlelasticpool.SetupGated,
		mssqlfailovergroup.SetupGated,
		mssqlfirewallrule.SetupGated,
		mssqljobagent.SetupGated,
		mssqljobcredential.SetupGated,
		mssqlmanageddatabase.SetupGated,
		mssqlmanagedinstance.SetupGated,
		mssqlmanagedinstanceactivedirectoryadministrator.SetupGated,
		mssqlmanagedinstancefailovergroup.SetupGated,
		mssqlmanagedinstancetransparentdataencryption.SetupGated,
		mssqlmanagedinstancevulnerabilityassessment.SetupGated,
		mssqloutboundfirewallrule.SetupGated,
		mssqlserver.SetupGated,
		mssqlserverdnsalias.SetupGated,
		mssqlservermicrosoftsupportauditingpolicy.SetupGated,
		mssqlserversecurityalertpolicy.SetupGated,
		mssqlservertransparentdataencryption.SetupGated,
		mssqlservervulnerabilityassessment.SetupGated,
		mssqlvirtualnetworkrule.SetupGated,
		accountstorage.SetupGated,
		accountlocaluser.SetupGated,
		accountnetworkrules.SetupGated,
		blob.SetupGated,
		blobinventorypolicy.SetupGated,
		container.SetupGated,
		containerimmutabilitypolicy.SetupGated,
		datalakegen2filesystem.SetupGated,
		datalakegen2path.SetupGated,
		encryptionscope.SetupGated,
		managementpolicy.SetupGated,
		objectreplication.SetupGated,
		queuestorage.SetupGated,
		share.SetupGated,
		sharedirectory.SetupGated,
		tablestorage.SetupGated,
		tableentity.SetupGated,
		hpccache.SetupGated,
		hpccacheaccesspolicy.SetupGated,
		hpccacheblobnfstarget.SetupGated,
		hpccacheblobtarget.SetupGated,
		hpccachenfstarget.SetupGated,
		storagesync.SetupGated,
		clusterstreamanalytics.SetupGated,
		functionjavascriptuda.SetupGated,
		job.SetupGated,
		managedprivateendpointstreamanalytics.SetupGated,
		outputblob.SetupGated,
		outputeventhub.SetupGated,
		outputfunction.SetupGated,
		outputmssql.SetupGated,
		outputpowerbi.SetupGated,
		outputservicebusqueue.SetupGated,
		outputservicebustopic.SetupGated,
		outputsynapse.SetupGated,
		outputtable.SetupGated,
		referenceinputblob.SetupGated,
		referenceinputmssql.SetupGated,
		streaminputblob.SetupGated,
		streaminputeventhub.SetupGated,
		streaminputiothub.SetupGated,
		firewallrulesynapse.SetupGated,
		integrationruntimeazuresynapse.SetupGated,
		integrationruntimeselfhostedsynapse.SetupGated,
		linkedservice.SetupGated,
		managedprivateendpointsynapse.SetupGated,
		privatelinkhub.SetupGated,
		roleassignmentsynapse.SetupGated,
		sparkpool.SetupGated,
		sqlpool.SetupGated,
		sqlpoolextendedauditingpolicy.SetupGated,
		sqlpoolsecurityalertpolicy.SetupGated,
		sqlpoolworkloadclassifier.SetupGated,
		sqlpoolworkloadgroup.SetupGated,
		workspacesynapse.SetupGated,
		workspaceaadadmin.SetupGated,
		workspaceextendedauditingpolicy.SetupGated,
		workspacesecurityalertpolicy.SetupGated,
		workspacesqlaadadmin.SetupGated,
		workspacevulnerabilityassessment.SetupGated,
		appactiveslot.SetupGated,
		apphybridconnection.SetupGated,
		appserviceplan.SetupGated,
		functionapp.SetupGated,
		functionappactiveslot.SetupGated,
		functionappfunction.SetupGated,
		functionapphybridconnection.SetupGated,
		functionappslot.SetupGated,
		linuxfunctionapp.SetupGated,
		linuxfunctionappslot.SetupGated,
		linuxwebapp.SetupGated,
		linuxwebappslot.SetupGated,
		serviceplan.SetupGated,
		sourcecontroltoken.SetupGated,
		staticsite.SetupGated,
		windowsfunctionapp.SetupGated,
		windowsfunctionappslot.SetupGated,
		windowswebapp.SetupGated,
		windowswebappslot.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
