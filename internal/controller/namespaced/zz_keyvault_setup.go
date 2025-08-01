// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/accesspolicy"
	certificate "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificate"
	certificatecontacts "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificatecontacts"
	certificateissuer "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/certificateissuer"
	key "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/key"
	managedhardwaresecuritymodule "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedhardwaresecuritymodule"
	managedstorageaccount "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedstorageaccount"
	managedstorageaccountsastokendefinition "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/managedstorageaccountsastokendefinition"
	secret "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/secret"
	vault "github.com/upbound/provider-azure/internal/controller/namespaced/keyvault/vault"
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

// SetupGated_keyvault creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_keyvault(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesspolicy.SetupGated,
		certificate.SetupGated,
		certificatecontacts.SetupGated,
		certificateissuer.SetupGated,
		key.SetupGated,
		managedhardwaresecuritymodule.SetupGated,
		managedstorageaccount.SetupGated,
		managedstorageaccountsastokendefinition.SetupGated,
		secret.SetupGated,
		vault.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
