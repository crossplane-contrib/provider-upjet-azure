// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	iothub "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothub"
	iothubcertificate "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubcertificate"
	iothubconsumergroup "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubconsumergroup"
	iothubdps "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdps"
	iothubdpscertificate "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdpscertificate"
	iothubdpssharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubdpssharedaccesspolicy"
	iothubendpointeventhub "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointeventhub"
	iothubendpointservicebusqueue "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointservicebusqueue"
	iothubendpointservicebustopic "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointservicebustopic"
	iothubendpointstoragecontainer "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubendpointstoragecontainer"
	iothubenrichment "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubenrichment"
	iothubfallbackroute "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubfallbackroute"
	iothubroute "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubroute"
	iothubsharedaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/devices/iothubsharedaccesspolicy"
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
