/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	management "github.com/upbound/official-providers/provider-azure/internal/controller/apimanagement/management"
	resourcegrouppolicyassignment "github.com/upbound/official-providers/provider-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	resourcegroup "github.com/upbound/official-providers/provider-azure/internal/controller/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/official-providers/provider-azure/internal/controller/azure/resourceproviderregistration"
	subscription "github.com/upbound/official-providers/provider-azure/internal/controller/azure/subscription"
	rediscache "github.com/upbound/official-providers/provider-azure/internal/controller/cache/rediscache"
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
	agentpool "github.com/upbound/official-providers/provider-azure/internal/controller/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/official-providers/provider-azure/internal/controller/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/official-providers/provider-azure/internal/controller/containerregistry/registry"
	scopemap "github.com/upbound/official-providers/provider-azure/internal/controller/containerregistry/scopemap"
	token "github.com/upbound/official-providers/provider-azure/internal/controller/containerregistry/token"
	webhook "github.com/upbound/official-providers/provider-azure/internal/controller/containerregistry/webhook"
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
	consumergroup "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/consumergroup"
	eventhub "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/official-providers/provider-azure/internal/controller/eventhub/eventhubnamespace"
	monitormetricalert "github.com/upbound/official-providers/provider-azure/internal/controller/insights/monitormetricalert"
	accesspolicy "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/accesspolicy"
	certificate "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/certificate"
	certificateissuer "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/certificateissuer"
	key "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/secret"
	vault "github.com/upbound/official-providers/provider-azure/internal/controller/keyvault/vault"
	integrationserviceenvironment "github.com/upbound/official-providers/provider-azure/internal/controller/logic/integrationserviceenvironment"
	managementgroup "github.com/upbound/official-providers/provider-azure/internal/controller/management/managementgroup"
	marketplaceagreement "github.com/upbound/official-providers/provider-azure/internal/controller/marketplaceordering/marketplaceagreement"
	applicationsecuritygroup "github.com/upbound/official-providers/provider-azure/internal/controller/network/applicationsecuritygroup"
	connectionmonitor "github.com/upbound/official-providers/provider-azure/internal/controller/network/connectionmonitor"
	ddosprotectionplan "github.com/upbound/official-providers/provider-azure/internal/controller/network/ddosprotectionplan"
	dnsaaaarecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnsaaaarecord"
	dnsarecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnsarecord"
	dnscaarecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnscaarecord"
	dnscnamerecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnscnamerecord"
	dnsmxrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnsmxrecord"
	dnsnsrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnsnsrecord"
	dnsptrrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnsptrrecord"
	dnssrvrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnssrvrecord"
	dnstxtrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnstxtrecord"
	dnszone "github.com/upbound/official-providers/provider-azure/internal/controller/network/dnszone"
	expressroutecircuit "github.com/upbound/official-providers/provider-azure/internal/controller/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/upbound/official-providers/provider-azure/internal/controller/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/upbound/official-providers/provider-azure/internal/controller/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/upbound/official-providers/provider-azure/internal/controller/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/upbound/official-providers/provider-azure/internal/controller/network/expressrouteconnection"
	expressroutegateway "github.com/upbound/official-providers/provider-azure/internal/controller/network/expressroutegateway"
	expressrouteport "github.com/upbound/official-providers/provider-azure/internal/controller/network/expressrouteport"
	firewall "github.com/upbound/official-providers/provider-azure/internal/controller/network/firewall"
	firewallapplicationrulecollection "github.com/upbound/official-providers/provider-azure/internal/controller/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/upbound/official-providers/provider-azure/internal/controller/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/upbound/official-providers/provider-azure/internal/controller/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/upbound/official-providers/provider-azure/internal/controller/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/upbound/official-providers/provider-azure/internal/controller/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/upbound/official-providers/provider-azure/internal/controller/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/upbound/official-providers/provider-azure/internal/controller/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicy "github.com/upbound/official-providers/provider-azure/internal/controller/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/upbound/official-providers/provider-azure/internal/controller/network/frontdoorrulesengine"
	ipgroup "github.com/upbound/official-providers/provider-azure/internal/controller/network/ipgroup"
	loadbalancer "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancer"
	loadbalancerbackendaddresspool "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancerbackendaddresspool"
	loadbalancerbackendaddresspooladdress "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancerbackendaddresspooladdress"
	loadbalancernatpool "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancernatpool"
	loadbalancernatrule "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancernatrule"
	loadbalanceroutboundrule "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalanceroutboundrule"
	loadbalancerprobe "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancerprobe"
	loadbalancerrule "github.com/upbound/official-providers/provider-azure/internal/controller/network/loadbalancerrule"
	localnetworkgateway "github.com/upbound/official-providers/provider-azure/internal/controller/network/localnetworkgateway"
	natgateway "github.com/upbound/official-providers/provider-azure/internal/controller/network/natgateway"
	natgatewaypublicipassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/upbound/official-providers/provider-azure/internal/controller/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/networkinterfacesecuritygroupassociation"
	packetcapture "github.com/upbound/official-providers/provider-azure/internal/controller/network/packetcapture"
	pointtositevpngateway "github.com/upbound/official-providers/provider-azure/internal/controller/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednsaaaarecord"
	privatednsarecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednsarecord"
	privatednscnamerecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednscnamerecord"
	privatednsmxrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednsmxrecord"
	privatednsptrrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednsptrrecord"
	privatednssrvrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednstxtrecord"
	privatednszone "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/official-providers/provider-azure/internal/controller/network/privateendpoint"
	privatelinkservice "github.com/upbound/official-providers/provider-azure/internal/controller/network/privatelinkservice"
	profile "github.com/upbound/official-providers/provider-azure/internal/controller/network/profile"
	publicip "github.com/upbound/official-providers/provider-azure/internal/controller/network/publicip"
	publicipprefix "github.com/upbound/official-providers/provider-azure/internal/controller/network/publicipprefix"
	securitygroup "github.com/upbound/official-providers/provider-azure/internal/controller/network/securitygroup"
	securityrule "github.com/upbound/official-providers/provider-azure/internal/controller/network/securityrule"
	subnet "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/official-providers/provider-azure/internal/controller/network/subnetserviceendpointstoragepolicy"
	virtualhub "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualhub"
	virtualnetwork "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/official-providers/provider-azure/internal/controller/network/virtualwan"
	vpnserverconfiguration "github.com/upbound/official-providers/provider-azure/internal/controller/network/vpnserverconfiguration"
	watcher "github.com/upbound/official-providers/provider-azure/internal/controller/network/watcher"
	watcherflowlog "github.com/upbound/official-providers/provider-azure/internal/controller/network/watcherflowlog"
	notificationhub "github.com/upbound/official-providers/provider-azure/internal/controller/notificationhubs/notificationhub"
	workspace "github.com/upbound/official-providers/provider-azure/internal/controller/operationalinsights/workspace"
	providerconfig "github.com/upbound/official-providers/provider-azure/internal/controller/providerconfig"
	resourcegrouptemplatedeployment "github.com/upbound/official-providers/provider-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	advancedthreatprotection "github.com/upbound/official-providers/provider-azure/internal/controller/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/official-providers/provider-azure/internal/controller/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/official-providers/provider-azure/internal/controller/security/iotsecuritysolution"
	mssqlserver "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlserver"
	mssqlservertransparentdataencryption "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlservertransparentdataencryption"
	mssqlvirtualnetworkrule "github.com/upbound/official-providers/provider-azure/internal/controller/sql/mssqlvirtualnetworkrule"
	accountstorage "github.com/upbound/official-providers/provider-azure/internal/controller/storage/account"
	accountnetworkrules "github.com/upbound/official-providers/provider-azure/internal/controller/storage/accountnetworkrules"
	blob "github.com/upbound/official-providers/provider-azure/internal/controller/storage/blob"
	blobinventorypolicy "github.com/upbound/official-providers/provider-azure/internal/controller/storage/blobinventorypolicy"
	container "github.com/upbound/official-providers/provider-azure/internal/controller/storage/container"
	datalakegen2filesystem "github.com/upbound/official-providers/provider-azure/internal/controller/storage/datalakegen2filesystem"
	encryptionscope "github.com/upbound/official-providers/provider-azure/internal/controller/storage/encryptionscope"
	managementpolicy "github.com/upbound/official-providers/provider-azure/internal/controller/storage/managementpolicy"
	objectreplication "github.com/upbound/official-providers/provider-azure/internal/controller/storage/objectreplication"
	queue "github.com/upbound/official-providers/provider-azure/internal/controller/storage/queue"
	share "github.com/upbound/official-providers/provider-azure/internal/controller/storage/share"
	tablestorage "github.com/upbound/official-providers/provider-azure/internal/controller/storage/table"
	hpccache "github.com/upbound/official-providers/provider-azure/internal/controller/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/official-providers/provider-azure/internal/controller/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/official-providers/provider-azure/internal/controller/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/official-providers/provider-azure/internal/controller/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/official-providers/provider-azure/internal/controller/storagecache/hpccachenfstarget"
	storagesync "github.com/upbound/official-providers/provider-azure/internal/controller/storagesync/storagesync"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		management.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		subscription.Setup,
		rediscache.Setup,
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
		agentpool.Setup,
		containerconnectedregistry.Setup,
		registry.Setup,
		scopemap.Setup,
		token.Setup,
		webhook.Setup,
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
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		monitormetricalert.Setup,
		accesspolicy.Setup,
		certificate.Setup,
		certificateissuer.Setup,
		key.Setup,
		managedhardwaresecuritymodule.Setup,
		managedstorageaccount.Setup,
		managedstorageaccountsastokendefinition.Setup,
		secret.Setup,
		vault.Setup,
		integrationserviceenvironment.Setup,
		managementgroup.Setup,
		marketplaceagreement.Setup,
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
		profile.Setup,
		publicip.Setup,
		publicipprefix.Setup,
		securitygroup.Setup,
		securityrule.Setup,
		subnet.Setup,
		subnetnatgatewayassociation.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		subnetserviceendpointstoragepolicy.Setup,
		virtualhub.Setup,
		virtualnetwork.Setup,
		virtualnetworkgateway.Setup,
		virtualnetworkgatewayconnection.Setup,
		virtualnetworkpeering.Setup,
		virtualwan.Setup,
		vpnserverconfiguration.Setup,
		watcher.Setup,
		watcherflowlog.Setup,
		notificationhub.Setup,
		workspace.Setup,
		providerconfig.Setup,
		resourcegrouptemplatedeployment.Setup,
		advancedthreatprotection.Setup,
		iotsecuritydevicegroup.Setup,
		iotsecuritysolution.Setup,
		mssqlserver.Setup,
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
		queue.Setup,
		share.Setup,
		tablestorage.Setup,
		hpccache.Setup,
		hpccacheaccesspolicy.Setup,
		hpccacheblobnfstarget.Setup,
		hpccacheblobtarget.Setup,
		hpccachenfstarget.Setup,
		storagesync.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
