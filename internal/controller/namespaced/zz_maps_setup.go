// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-azure/v2/internal/controller/namespaced/maps/account"
	creator "github.com/upbound/provider-azure/v2/internal/controller/namespaced/maps/creator"
)

// Setup_maps creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_maps(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		creator.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_maps creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_maps(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		creator.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
