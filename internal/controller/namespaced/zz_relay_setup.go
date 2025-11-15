// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	eventrelaynamespace "github.com/upbound/provider-azure/v2/internal/controller/namespaced/relay/eventrelaynamespace"
	hybridconnection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/relay/hybridconnectionauthorizationrule"
	namespaceauthorizationrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/relay/namespaceauthorizationrule"
)

// Setup_relay creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_relay(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		eventrelaynamespace.Setup,
		hybridconnection.Setup,
		hybridconnectionauthorizationrule.Setup,
		namespaceauthorizationrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_relay creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_relay(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		eventrelaynamespace.SetupGated,
		hybridconnection.SetupGated,
		hybridconnectionauthorizationrule.SetupGated,
		namespaceauthorizationrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
