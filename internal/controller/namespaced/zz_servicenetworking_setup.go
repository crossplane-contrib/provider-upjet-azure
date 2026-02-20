// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	applicationloadbalancer "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicenetworking/applicationloadbalancer"
	applicationloadbalancerfrontend "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicenetworking/applicationloadbalancerfrontend"
	applicationloadbalancersecuritypolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicenetworking/applicationloadbalancersecuritypolicy"
	applicationloadbalancersubnetassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicenetworking/applicationloadbalancersubnetassociation"
)

// Setup_servicenetworking creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicenetworking(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationloadbalancer.Setup,
		applicationloadbalancerfrontend.Setup,
		applicationloadbalancersecuritypolicy.Setup,
		applicationloadbalancersubnetassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_servicenetworking creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_servicenetworking(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationloadbalancer.SetupGated,
		applicationloadbalancerfrontend.SetupGated,
		applicationloadbalancersecuritypolicy.SetupGated,
		applicationloadbalancersubnetassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
