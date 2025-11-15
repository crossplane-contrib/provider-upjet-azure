// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	networkacl "github.com/upbound/provider-azure/v2/internal/controller/namespaced/signalrservice/networkacl"
	service "github.com/upbound/provider-azure/v2/internal/controller/namespaced/signalrservice/service"
	signalrsharedprivatelinkresource "github.com/upbound/provider-azure/v2/internal/controller/namespaced/signalrservice/signalrsharedprivatelinkresource"
	webpubsub "github.com/upbound/provider-azure/v2/internal/controller/namespaced/signalrservice/webpubsub"
	webpubsubhub "github.com/upbound/provider-azure/v2/internal/controller/namespaced/signalrservice/webpubsubhub"
	webpubsubnetworkacl "github.com/upbound/provider-azure/v2/internal/controller/namespaced/signalrservice/webpubsubnetworkacl"
)

// Setup_signalrservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_signalrservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		networkacl.Setup,
		service.Setup,
		signalrsharedprivatelinkresource.Setup,
		webpubsub.Setup,
		webpubsubhub.Setup,
		webpubsubnetworkacl.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_signalrservice creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_signalrservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		networkacl.SetupGated,
		service.SetupGated,
		signalrsharedprivatelinkresource.SetupGated,
		webpubsub.SetupGated,
		webpubsubhub.SetupGated,
		webpubsubnetworkacl.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
