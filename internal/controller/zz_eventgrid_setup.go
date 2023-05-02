/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	domain "github.com/upbound/provider-azure/internal/controller/eventgrid/domain"
	domaintopic "github.com/upbound/provider-azure/internal/controller/eventgrid/domaintopic"
	eventsubscription "github.com/upbound/provider-azure/internal/controller/eventgrid/eventsubscription"
	systemtopic "github.com/upbound/provider-azure/internal/controller/eventgrid/systemtopic"
	topic "github.com/upbound/provider-azure/internal/controller/eventgrid/topic"
)

// Setup_eventgrid creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventgrid(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domaintopic.Setup,
		eventsubscription.Setup,
		systemtopic.Setup,
		topic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
