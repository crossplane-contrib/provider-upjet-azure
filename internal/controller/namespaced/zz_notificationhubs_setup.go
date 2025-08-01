// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	authorizationrule "github.com/upbound/provider-azure/internal/controller/namespaced/notificationhubs/authorizationrule"
	notificationhub "github.com/upbound/provider-azure/internal/controller/namespaced/notificationhubs/notificationhub"
	notificationhubnamespace "github.com/upbound/provider-azure/internal/controller/namespaced/notificationhubs/notificationhubnamespace"
)

// Setup_notificationhubs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_notificationhubs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizationrule.Setup,
		notificationhub.Setup,
		notificationhubnamespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_notificationhubs creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_notificationhubs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizationrule.SetupGated,
		notificationhub.SetupGated,
		notificationhubnamespace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
