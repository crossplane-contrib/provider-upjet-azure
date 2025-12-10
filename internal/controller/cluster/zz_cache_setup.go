// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	managedredis "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/managedredis"
	rediscache "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/rediscache"
	rediscacheaccesspolicy "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/rediscacheaccesspolicy"
	rediscacheaccesspolicyassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/rediscacheaccesspolicyassignment"
	redisenterprisecluster "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/v2/internal/controller/cluster/cache/redislinkedserver"
)

// Setup_cache creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cache(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managedredis.Setup,
		rediscache.Setup,
		rediscacheaccesspolicy.Setup,
		rediscacheaccesspolicyassignment.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cache creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cache(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managedredis.SetupGated,
		rediscache.SetupGated,
		rediscacheaccesspolicy.SetupGated,
		rediscacheaccesspolicyassignment.SetupGated,
		redisenterprisecluster.SetupGated,
		redisenterprisedatabase.SetupGated,
		redisfirewallrule.SetupGated,
		redislinkedserver.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
