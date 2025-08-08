// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	resourcegroup "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/cluster/azure/resourceproviderregistration"
	subscription "github.com/upbound/provider-azure/internal/controller/cluster/azure/subscription"
	providerconfig "github.com/upbound/provider-azure/internal/controller/cluster/providerconfig"
)

// Setup_config creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_config(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		subscription.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_config creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_config(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcegroup.SetupGated,
		resourceproviderregistration.SetupGated,
		subscription.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
