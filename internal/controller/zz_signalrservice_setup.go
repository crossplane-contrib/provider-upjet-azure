/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	networkacl "github.com/upbound/provider-azure/internal/controller/signalrservice/networkacl"
	service "github.com/upbound/provider-azure/internal/controller/signalrservice/service"
	signalrsharedprivatelinkresource "github.com/upbound/provider-azure/internal/controller/signalrservice/signalrsharedprivatelinkresource"
	webpubsub "github.com/upbound/provider-azure/internal/controller/signalrservice/webpubsub"
	webpubsubhub "github.com/upbound/provider-azure/internal/controller/signalrservice/webpubsubhub"
	webpubsubnetworkacl "github.com/upbound/provider-azure/internal/controller/signalrservice/webpubsubnetworkacl"
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
