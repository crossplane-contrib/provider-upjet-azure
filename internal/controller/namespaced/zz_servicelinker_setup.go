// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	springcloudconnection "github.com/upbound/provider-azure/internal/controller/namespaced/servicelinker/springcloudconnection"
)

// Setup_servicelinker creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicelinker(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		springcloudconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_servicelinker creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_servicelinker(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		springcloudconnection.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
