// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	budgetmanagementgroup "github.com/upbound/provider-azure/v2/internal/controller/cluster/consumption/budgetmanagementgroup"
	budgetresourcegroup "github.com/upbound/provider-azure/v2/internal/controller/cluster/consumption/budgetresourcegroup"
	budgetsubscription "github.com/upbound/provider-azure/v2/internal/controller/cluster/consumption/budgetsubscription"
)

// Setup_consumption creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_consumption(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budgetmanagementgroup.Setup,
		budgetresourcegroup.Setup,
		budgetsubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_consumption creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_consumption(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budgetmanagementgroup.SetupGated,
		budgetresourcegroup.SetupGated,
		budgetsubscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
