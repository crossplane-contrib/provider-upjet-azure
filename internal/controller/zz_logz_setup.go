// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	monitor "github.com/upbound/provider-azure/internal/controller/logz/monitor"
	subaccount "github.com/upbound/provider-azure/internal/controller/logz/subaccount"
	subaccounttagrule "github.com/upbound/provider-azure/internal/controller/logz/subaccounttagrule"
	tagrule "github.com/upbound/provider-azure/internal/controller/logz/tagrule"
)

// Setup_logz creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logz(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitor.Setup,
		subaccount.Setup,
		subaccounttagrule.Setup,
		tagrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
