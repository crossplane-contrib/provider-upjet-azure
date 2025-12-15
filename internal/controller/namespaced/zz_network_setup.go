// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	applicationgateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/applicationgateway"
	applicationsecuritygroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/applicationsecuritygroup"
	bastionhost "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/bastionhost"
	connectionmonitor "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/connectionmonitor"
	ddosprotectionplan "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/ddosprotectionplan"
	dnsaaaarecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnsaaaarecord"
	dnsarecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnsarecord"
	dnscaarecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnscaarecord"
	dnscnamerecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnscnamerecord"
	dnsmxrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnsmxrecord"
	dnsnsrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnsnsrecord"
	dnsptrrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnsptrrecord"
	dnssrvrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnssrvrecord"
	dnstxtrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnstxtrecord"
	dnszone "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/dnszone"
	expressroutecircuit "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/expressroutecircuit"
	expressroutecircuitauthorization "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/expressroutecircuitauthorization"
	expressroutecircuitconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/expressroutecircuitconnection"
	expressroutecircuitpeering "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/expressroutecircuitpeering"
	expressrouteconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/expressrouteconnection"
	expressroutegateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/expressroutegateway"
	expressrouteport "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/expressrouteport"
	firewall "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/firewall"
	firewallapplicationrulecollection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/firewallapplicationrulecollection"
	firewallnatrulecollection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/firewallnatrulecollection"
	firewallnetworkrulecollection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/firewallnetworkrulecollection"
	firewallpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/firewallpolicy"
	firewallpolicyrulecollectiongroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/firewallpolicyrulecollectiongroup"
	frontdoor "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/frontdoor"
	frontdoorcustomhttpsconfiguration "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/frontdoorcustomhttpsconfiguration"
	frontdoorfirewallpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/frontdoorfirewallpolicy"
	frontdoorrulesengine "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/frontdoorrulesengine"
	ipgroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/ipgroup"
	ipgroupcidr "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/ipgroupcidr"
	loadbalancer "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalancer"
	loadbalancerbackendaddresspool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalancerbackendaddresspool"
	loadbalancerbackendaddresspooladdress "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalancerbackendaddresspooladdress"
	loadbalancernatpool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalancernatpool"
	loadbalancernatrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalancernatrule"
	loadbalanceroutboundrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalanceroutboundrule"
	loadbalancerprobe "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalancerprobe"
	loadbalancerrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/loadbalancerrule"
	localnetworkgateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/localnetworkgateway"
	manager "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/manager"
	manageripampool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/manageripampool"
	managermanagementgroupconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/managermanagementgroupconnection"
	managernetworkgroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/managernetworkgroup"
	managerroutingconfiguration "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/managerroutingconfiguration"
	managerstaticmember "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/managerstaticmember"
	managersubscriptionconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/managersubscriptionconnection"
	managerverifierworkspace "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/managerverifierworkspace"
	natgateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/natgateway"
	natgatewaypublicipassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/natgatewaypublicipassociation"
	natgatewaypublicipprefixassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/natgatewaypublicipprefixassociation"
	networkinterface "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/networkinterface"
	networkinterfaceapplicationsecuritygroupassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/networkinterfaceapplicationsecuritygroupassociation"
	networkinterfacebackendaddresspoolassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/networkinterfacebackendaddresspoolassociation"
	networkinterfacenatruleassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/networkinterfacenatruleassociation"
	networkinterfacesecuritygroupassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/networkinterfacesecuritygroupassociation"
	packetcapture "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/packetcapture"
	pointtositevpngateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/pointtositevpngateway"
	privatednsaaaarecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsaaaarecord"
	privatednsarecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsarecord"
	privatednscnamerecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednscnamerecord"
	privatednsmxrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsmxrecord"
	privatednsptrrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsptrrecord"
	privatednsresolver "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsresolver"
	privatednsresolverdnsforwardingruleset "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsresolverdnsforwardingruleset"
	privatednsresolverforwardingrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsresolverforwardingrule"
	privatednsresolverinboundendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsresolverinboundendpoint"
	privatednsresolveroutboundendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsresolveroutboundendpoint"
	privatednsresolvervirtualnetworklink "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednsresolvervirtualnetworklink"
	privatednssrvrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednssrvrecord"
	privatednstxtrecord "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednstxtrecord"
	privatednszone "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednszone"
	privatednszonevirtualnetworklink "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatednszonevirtualnetworklink"
	privateendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privateendpoint"
	privateendpointapplicationsecuritygroupassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privateendpointapplicationsecuritygroupassociation"
	privatelinkservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/privatelinkservice"
	profile "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/profile"
	publicip "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/publicip"
	publicipprefix "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/publicipprefix"
	route "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/route"
	routefilter "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/routefilter"
	routemap "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/routemap"
	routeserver "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/routeserver"
	routeserverbgpconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/routeserverbgpconnection"
	routetable "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/routetable"
	securitygroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/securitygroup"
	securityrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/securityrule"
	subnet "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/subnet"
	subnetnatgatewayassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/subnetnatgatewayassociation"
	subnetnetworksecuritygroupassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/subnetroutetableassociation"
	subnetserviceendpointstoragepolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/subnetserviceendpointstoragepolicy"
	trafficmanagerazureendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/trafficmanagerazureendpoint"
	trafficmanagerexternalendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/trafficmanagerexternalendpoint"
	trafficmanagernestedendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/trafficmanagernestedendpoint"
	trafficmanagerprofile "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/trafficmanagerprofile"
	virtualhub "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualhub"
	virtualhubconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualhubconnection"
	virtualhubip "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualhubip"
	virtualhubroutetable "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualhubroutetable"
	virtualhubroutetableroute "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualhubroutetableroute"
	virtualhubroutingintent "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualhubroutingintent"
	virtualhubsecuritypartnerprovider "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualhubsecuritypartnerprovider"
	virtualnetwork "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualnetwork"
	virtualnetworkdnsservers "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualnetworkdnsservers"
	virtualnetworkgateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualnetworkgateway"
	virtualnetworkgatewayconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualnetworkgatewayconnection"
	virtualnetworkpeering "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualnetworkpeering"
	virtualwan "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/virtualwan"
	vpngateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/vpngateway"
	vpngatewayconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/vpngatewayconnection"
	vpnserverconfiguration "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/vpnserverconfiguration"
	vpnserverconfigurationpolicygroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/vpnserverconfigurationpolicygroup"
	vpnsite "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/vpnsite"
	watcher "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/watcher"
	watcherflowlog "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/watcherflowlog"
	webapplicationfirewallpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/network/webapplicationfirewallpolicy"
)

// Setup_network creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_network(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
		frontdoorfirewallpolicy.Setup,
		frontdoorrulesengine.Setup,
		ipgroup.Setup,
		ipgroupcidr.Setup,
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
		virtualhubroutingintent.Setup,
		virtualhubsecuritypartnerprovider.Setup,
		virtualnetwork.Setup,
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_network creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_network(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
		frontdoorfirewallpolicy.SetupGated,
		frontdoorrulesengine.SetupGated,
		ipgroup.SetupGated,
		ipgroupcidr.SetupGated,
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
		profile.SetupGated,
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
		virtualhubroutingintent.SetupGated,
		virtualhubsecuritypartnerprovider.SetupGated,
		virtualnetwork.SetupGated,
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
