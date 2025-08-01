// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	iothub "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothub"
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubcertificate"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/namespaced/devices/iothubsharedaccesspolicy"
)

// Setup_devices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		iothub.Setup,
		iothubcertificate.Setup,
		iothubconsumergroup.Setup,
		iothubdps.Setup,
		iothubdpscertificate.Setup,
		iothubdpssharedaccesspolicy.Setup,
		iothubendpointeventhub.Setup,
		iothubendpointservicebusqueue.Setup,
		iothubendpointservicebustopic.Setup,
		iothubendpointstoragecontainer.Setup,
		iothubenrichment.Setup,
		iothubfallbackroute.Setup,
		iothubroute.Setup,
		iothubsharedaccesspolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_devices creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_devices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		iothub.SetupGated,
		iothubcertificate.SetupGated,
		iothubconsumergroup.SetupGated,
		iothubdps.SetupGated,
		iothubdpscertificate.SetupGated,
		iothubdpssharedaccesspolicy.SetupGated,
		iothubendpointeventhub.SetupGated,
		iothubendpointservicebusqueue.SetupGated,
		iothubendpointservicebustopic.SetupGated,
		iothubendpointstoragecontainer.SetupGated,
		iothubenrichment.SetupGated,
		iothubfallbackroute.SetupGated,
		iothubroute.SetupGated,
		iothubsharedaccesspolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
