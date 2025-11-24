// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	healthcaredicomservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/healthcareapis/healthcaredicomservice"
	healthcarefhirservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/healthcareapis/healthcarefhirservice"
	healthcaremedtechservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/healthcareapis/healthcaremedtechservice"
	healthcaremedtechservicefhirdestination "github.com/upbound/provider-azure/v2/internal/controller/namespaced/healthcareapis/healthcaremedtechservicefhirdestination"
	healthcareservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/healthcareapis/healthcareservice"
	healthcareworkspace "github.com/upbound/provider-azure/v2/internal/controller/namespaced/healthcareapis/healthcareworkspace"
)

// Setup_healthcareapis creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_healthcareapis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		healthcaredicomservice.Setup,
		healthcarefhirservice.Setup,
		healthcaremedtechservice.Setup,
		healthcaremedtechservicefhirdestination.Setup,
		healthcareservice.Setup,
		healthcareworkspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_healthcareapis creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_healthcareapis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		healthcaredicomservice.SetupGated,
		healthcarefhirservice.SetupGated,
		healthcaremedtechservice.SetupGated,
		healthcaremedtechservicefhirdestination.SetupGated,
		healthcareservice.SetupGated,
		healthcareworkspace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
