// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	availabilityset "github.com/upbound/provider-azure/internal/controller/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/compute/snapshot"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/compute/sshpublickey"
	virtualmachinedatadiskattachment "github.com/upbound/provider-azure/internal/controller/compute/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/upbound/provider-azure/internal/controller/compute/virtualmachineextension"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/compute/windowsvirtualmachinescaleset"
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		availabilityset.Setup,
		capacityreservation.Setup,
		capacityreservationgroup.Setup,
		dedicatedhost.Setup,
		diskaccess.Setup,
		diskencryptionset.Setup,
		galleryapplication.Setup,
		galleryapplicationversion.Setup,
		image.Setup,
		linuxvirtualmachine.Setup,
		linuxvirtualmachinescaleset.Setup,
		manageddisk.Setup,
		manageddisksastoken.Setup,
		orchestratedvirtualmachinescaleset.Setup,
		proximityplacementgroup.Setup,
		sharedimage.Setup,
		sharedimagegallery.Setup,
		snapshot.Setup,
		sshpublickey.Setup,
		virtualmachinedatadiskattachment.Setup,
		virtualmachineextension.Setup,
		windowsvirtualmachine.Setup,
		windowsvirtualmachinescaleset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
