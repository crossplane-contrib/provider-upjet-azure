/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	eventrelaynamespace "github.com/upbound/provider-azure/internal/controller/relay/eventrelaynamespace"
	hybridconnection "github.com/upbound/provider-azure/internal/controller/relay/hybridconnection"
	hybridconnectionauthorizationrule "github.com/upbound/provider-azure/internal/controller/relay/hybridconnectionauthorizationrule"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/relay/namespaceauthorizationrule"
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
