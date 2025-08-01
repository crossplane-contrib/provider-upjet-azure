// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	marketplaceagreement "github.com/upbound/provider-azure/internal/controller/cluster/marketplaceordering/marketplaceagreement"
)

// Setup_marketplaceordering creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_marketplaceordering(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		marketplaceagreement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_marketplaceordering creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_marketplaceordering(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		marketplaceagreement.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
