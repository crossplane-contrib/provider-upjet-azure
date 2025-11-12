// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	domain "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/domain"
	domaintopic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/domaintopic"
	eventgridnamespace "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/eventgridnamespace"
	eventsubscription "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/eventsubscription"
	partnerconfiguration "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/partnerconfiguration"
	systemtopic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/systemtopic"
	systemtopiceventsubscription "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/systemtopiceventsubscription"
	topic "github.com/upbound/provider-azure/internal/controller/namespaced/eventgrid/topic"
)

// Setup_eventgrid creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventgrid(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domaintopic.Setup,
		eventgridnamespace.Setup,
		eventsubscription.Setup,
		partnerconfiguration.Setup,
		systemtopic.Setup,
		systemtopiceventsubscription.Setup,
		topic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_eventgrid creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_eventgrid(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.SetupGated,
		domaintopic.SetupGated,
		eventgridnamespace.SetupGated,
		eventsubscription.SetupGated,
		partnerconfiguration.SetupGated,
		systemtopic.SetupGated,
		systemtopiceventsubscription.SetupGated,
		topic.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
