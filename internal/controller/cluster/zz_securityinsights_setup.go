// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	sentinelalertrulefusion "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulefusion"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulemachinelearningbehavioranalytics"
	sentinelalertrulemssecurityincident "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelalertrulemssecurityincident"
	sentinelautomationrule "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelautomationrule"
	sentineldataconnectoriot "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentineldataconnectoriot"
	sentinelloganalyticsworkspaceonboarding "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelloganalyticsworkspaceonboarding"
	sentinelwatchlist "github.com/upbound/provider-azure/internal/controller/cluster/securityinsights/sentinelwatchlist"
)

// Setup_securityinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_securityinsights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		sentinelalertrulefusion.Setup,
		sentinelalertrulemachinelearningbehavioranalytics.Setup,
		sentinelalertrulemssecurityincident.Setup,
		sentinelautomationrule.Setup,
		sentineldataconnectoriot.Setup,
		sentinelloganalyticsworkspaceonboarding.Setup,
		sentinelwatchlist.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
