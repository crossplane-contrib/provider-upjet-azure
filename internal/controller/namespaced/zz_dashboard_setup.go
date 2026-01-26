// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	grafana "github.com/upbound/provider-azure/v2/internal/controller/namespaced/dashboard/grafana"
	grafanamanagedprivateendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/dashboard/grafanamanagedprivateendpoint"
)

// Setup_dashboard creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dashboard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		grafana.Setup,
		grafanamanagedprivateendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dashboard creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dashboard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		grafana.SetupGated,
		grafanamanagedprivateendpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
