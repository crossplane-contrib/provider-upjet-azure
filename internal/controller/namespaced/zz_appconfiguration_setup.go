// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	configuration "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appconfiguration/configuration"
	key "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appconfiguration/key"
)

// Setup_appconfiguration creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appconfiguration(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.Setup,
		key.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_appconfiguration creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appconfiguration(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.SetupGated,
		key.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_appconfiguration registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_appconfiguration(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		configuration.SetupWebhookWithManager,
		key.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
