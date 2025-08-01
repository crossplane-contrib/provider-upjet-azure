// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	resourcepolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/resourcepolicyremediation"
	subscriptionpolicyremediation "github.com/upbound/provider-azure/internal/controller/cluster/policyinsights/subscriptionpolicyremediation"
)

// Setup_policyinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_policyinsights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcepolicyremediation.Setup,
		subscriptionpolicyremediation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_policyinsights creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_policyinsights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcepolicyremediation.SetupGated,
		subscriptionpolicyremediation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
