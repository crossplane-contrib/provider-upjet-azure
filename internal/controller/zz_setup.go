/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	management "github.com/upbound/provider-azure/internal/controller/apimanagement/management"
	policydefinition "github.com/upbound/provider-azure/internal/controller/authorization/policydefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	roleassignment "github.com/upbound/provider-azure/internal/controller/authorization/roleassignment"
	resourcegroup "github.com/upbound/provider-azure/internal/controller/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/azure/resourceproviderregistration"
	subscription "github.com/upbound/provider-azure/internal/controller/azure/subscription"
	rediscache "github.com/upbound/provider-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "github.com/upbound/provider-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/internal/controller/cache/redislinkedserver"
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
	agentpool "github.com/upbound/provider-azure/internal/controller/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/provider-azure/internal/controller/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/provider-azure/internal/controller/containerregistry/registry"
	scopemap "github.com/upbound/provider-azure/internal/controller/containerregistry/scopemap"
	token "github.com/upbound/provider-azure/internal/controller/containerregistry/token"
	webhook "github.com/upbound/provider-azure/internal/controller/containerregistry/webhook"
	kubernetescluster "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetescluster"
	kubernetesclusternodepool "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetesclusternodepool"
	account "github.com/upbound/provider-azure/internal/controller/cosmosdb/account"
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
	backuppolicyblobstorage "github.com/upbound/provider-azure/internal/controller/dataprotection/backuppolicyblobstorage"
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
	server "github.com/upbound/provider-azure/internal/controller/dbformariadb/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/dbformariadb/virtualnetworkrule"
	configurationdbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/configuration"
	firewallruledbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/firewallrule"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverfirewallrule"
	serverdbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/server"
	virtualnetworkruledbformysql "github.com/upbound/provider-azure/internal/controller/dbformysql/virtualnetworkrule"
	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
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
	authorizationrule "github.com/upbound/provider-azure/internal/controller/eventhub/authorizationrule"
	consumergroup "github.com/upbound/provider-azure/internal/controller/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/eventhub/eventhubnamespace"
	applicationinsights "github.com/upbound/provider-azure/internal/controller/insights/applicationinsights"
	monitoractiongroup "github.com/upbound/provider-azure/internal/controller/insights/monitoractiongroup"
	monitormetricalert "github.com/upbound/provider-azure/internal/controller/insights/monitormetricalert"
	monitorprivatelinkscope "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscope"
	monitorprivatelinkscopedservice "github.com/upbound/provider-azure/internal/controller/insights/monitorprivatelinkscopedservice"
	accesspolicy "github.com/upbound/provider-azure/internal/controller/keyvault/accesspolicy"
	certificate "github.com/upbound/provider-azure/internal/controller/keyvault/certificate"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/keyvault/secret"
	vault "github.com/upbound/provider-azure/internal/controller/keyvault/vault"
	cluster "github.com/upbound/provider-azure/internal/controller/kusto/cluster"
	databasekusto "github.com/upbound/provider-azure/internal/controller/kusto/database"
	integrationserviceenvironment "github.com/upbound/provider-azure/internal/controller/logic/integrationserviceenvironment"
	managementgroup "github.com/upbound/provider-azure/internal/controller/management/managementgroup"
	marketplaceagreement "github.com/upbound/provider-azure/internal/controller/marketplaceordering/marketplaceagreement"
	asset "github.com/upbound/provider-azure/internal/controller/media/asset"
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
	profile "github.com/upbound/provider-azure/internal/controller/network/profile"
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
	virtualhub "github.com/upbound/provider-azure/internal/controller/network/virtualhub"
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkgateway "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/internal/controller/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/internal/controller/network/virtualwan"
	vpnserverconfiguration "github.com/upbound/provider-azure/internal/controller/network/vpnserverconfiguration"
	watcher "github.com/upbound/provider-azure/internal/controller/network/watcher"
	watcherflowlog "github.com/upbound/provider-azure/internal/controller/network/watcherflowlog"
	notificationhub "github.com/upbound/provider-azure/internal/controller/notificationhubs/notificationhub"
	workspace "github.com/upbound/provider-azure/internal/controller/operationalinsights/workspace"
	providerconfig "github.com/upbound/provider-azure/internal/controller/providerconfig"
	resourcegrouptemplatedeployment "github.com/upbound/provider-azure/internal/controller/resources/resourcegrouptemplatedeployment"
	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/security/iotsecuritysolution"
	servicebusnamespace "github.com/upbound/provider-azure/internal/controller/servicebus/servicebusnamespace"
	mssqldatabase "github.com/upbound/provider-azure/internal/controller/sql/mssqldatabase"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/sql/mssqlfailovergroup"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/sql/mssqlserverdnsalias"
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
	queue "github.com/upbound/provider-azure/internal/controller/storage/queue"
	share "github.com/upbound/provider-azure/internal/controller/storage/share"
	tablestorage "github.com/upbound/provider-azure/internal/controller/storage/table"
	hpccache "github.com/upbound/provider-azure/internal/controller/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/storagecache/hpccachenfstarget"
	storagesync "github.com/upbound/provider-azure/internal/controller/storagesync/storagesync"
	clusterstreamanalytics "github.com/upbound/provider-azure/internal/controller/streamanalytics/cluster"
	functionjavascriptuda "github.com/upbound/provider-azure/internal/controller/streamanalytics/functionjavascriptuda"
	job "github.com/upbound/provider-azure/internal/controller/streamanalytics/job"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/streamanalytics/managedprivateendpoint"
	outputblob "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputblob"
	outputfunction "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputfunction"
	outputsynapse "github.com/upbound/provider-azure/internal/controller/streamanalytics/outputsynapse"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		management.Setup,
		policydefinition.Setup,
		resourcegrouppolicyassignment.Setup,
		roleassignment.Setup,
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
		sqlcontainer.Setup,
		sqldatabase.Setup,
		sqlfunction.Setup,
		sqlroleassignment.Setup,
		sqlroledefinition.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		table.Setup,
		backuppolicyblobstorage.Setup,
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
		server.Setup,
		virtualnetworkrule.Setup,
		configurationdbformysql.Setup,
		firewallruledbformysql.Setup,
		flexibledatabase.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverfirewallrule.Setup,
		serverdbformysql.Setup,
		virtualnetworkruledbformysql.Setup,
		activedirectoryadministrator.Setup,
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
		authorizationrule.Setup,
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		applicationinsights.Setup,
		monitoractiongroup.Setup,
		monitormetricalert.Setup,
		monitorprivatelinkscope.Setup,
		monitorprivatelinkscopedservice.Setup,
		accesspolicy.Setup,
		certificate.Setup,
		certificateissuer.Setup,
		key.Setup,
		managedhardwaresecuritymodule.Setup,
		managedstorageaccount.Setup,
		managedstorageaccountsastokendefinition.Setup,
		secret.Setup,
		vault.Setup,
		cluster.Setup,
		databasekusto.Setup,
		integrationserviceenvironment.Setup,
		managementgroup.Setup,
		marketplaceagreement.Setup,
		asset.Setup,
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
		profile.Setup,
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
		servicebusnamespace.Setup,
		mssqldatabase.Setup,
		mssqlfailovergroup.Setup,
		mssqlmanageddatabase.Setup,
		mssqlmanagedinstance.Setup,
		mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		mssqlmanagedinstancefailovergroup.Setup,
		mssqlmanagedinstancevulnerabilityassessment.Setup,
		mssqloutboundfirewallrule.Setup,
		mssqlserver.Setup,
		mssqlserverdnsalias.Setup,
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
		clusterstreamanalytics.Setup,
		functionjavascriptuda.Setup,
		job.Setup,
		managedprivateendpoint.Setup,
		outputblob.Setup,
		outputfunction.Setup,
		outputsynapse.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
