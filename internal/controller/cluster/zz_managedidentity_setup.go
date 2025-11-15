// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	federatedidentitycredential "github.com/upbound/provider-azure/v2/internal/controller/cluster/managedidentity/federatedidentitycredential"
	userassignedidentity "github.com/upbound/provider-azure/v2/internal/controller/cluster/managedidentity/userassignedidentity"
)

// Setup_managedidentity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managedidentity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		federatedidentitycredential.Setup,
		userassignedidentity.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_managedidentity creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_managedidentity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		federatedidentitycredential.SetupGated,
		userassignedidentity.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
