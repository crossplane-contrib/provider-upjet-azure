// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	rediscache "github.com/upbound/provider-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "github.com/upbound/provider-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "github.com/upbound/provider-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "github.com/upbound/provider-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "github.com/upbound/provider-azure/internal/controller/cache/redislinkedserver"
)

// Setup_cache creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cache(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rediscache.Setup,
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
