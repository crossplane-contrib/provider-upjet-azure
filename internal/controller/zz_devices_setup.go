// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	iothub "github.com/upbound/provider-azure/internal/controller/devices/iothub"
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/devices/iothubcertificate"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/devices/iothubsharedaccesspolicy"
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
