// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
	profile "github.com/upbound/provider-azure/internal/controller/network/profile"
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
	virtualnetwork "github.com/upbound/provider-azure/internal/controller/network/virtualnetwork"
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
)

// Setup_network creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_network(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
		profile.Setup,
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
		virtualnetwork.Setup,
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
