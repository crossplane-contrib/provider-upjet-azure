// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

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
)

// Setup_cdn creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cdn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
