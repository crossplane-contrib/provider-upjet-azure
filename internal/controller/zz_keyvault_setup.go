// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesspolicy "github.com/upbound/provider-azure/internal/controller/keyvault/accesspolicy"
	certificate "github.com/upbound/provider-azure/internal/controller/keyvault/certificate"
	certificatecontacts "github.com/upbound/provider-azure/internal/controller/keyvault/certificatecontacts"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/keyvault/secret"
	vault "github.com/upbound/provider-azure/internal/controller/keyvault/vault"
)

// Setup_keyvault creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_keyvault(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspolicy.Setup,
		certificate.Setup,
		certificatecontacts.Setup,
		certificateissuer.Setup,
		key.Setup,
		managedhardwaresecuritymodule.Setup,
		managedstorageaccount.Setup,
		managedstorageaccountsastokendefinition.Setup,
		secret.Setup,
		vault.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
