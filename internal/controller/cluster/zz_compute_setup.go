// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	availabilityset "github.com/upbound/provider-azure/internal/controller/cluster/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/cluster/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/cluster/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/cluster/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/cluster/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/cluster/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/cluster/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/cluster/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/cluster/compute/snapshot"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/cluster/compute/sshpublickey"
	virtualmachinedatadiskattachment "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineextension"
	virtualmachineruncommand "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachineruncommand"
	virtualmachinescalesetstandbypool "github.com/upbound/provider-azure/internal/controller/cluster/compute/virtualmachinescalesetstandbypool"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/cluster/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/cluster/compute/windowsvirtualmachinescaleset"
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
		virtualmachineruncommand.Setup,
		virtualmachinescalesetstandbypool.Setup,
		windowsvirtualmachine.Setup,
		windowsvirtualmachinescaleset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_compute creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		availabilityset.SetupGated,
		capacityreservation.SetupGated,
		capacityreservationgroup.SetupGated,
		dedicatedhost.SetupGated,
		diskaccess.SetupGated,
		diskencryptionset.SetupGated,
		galleryapplication.SetupGated,
		galleryapplicationversion.SetupGated,
		image.SetupGated,
		linuxvirtualmachine.SetupGated,
		linuxvirtualmachinescaleset.SetupGated,
		manageddisk.SetupGated,
		manageddisksastoken.SetupGated,
		orchestratedvirtualmachinescaleset.SetupGated,
		proximityplacementgroup.SetupGated,
		sharedimage.SetupGated,
		sharedimagegallery.SetupGated,
		snapshot.SetupGated,
		sshpublickey.SetupGated,
		virtualmachinedatadiskattachment.SetupGated,
		virtualmachineextension.SetupGated,
		virtualmachineruncommand.SetupGated,
		virtualmachinescalesetstandbypool.SetupGated,
		windowsvirtualmachine.SetupGated,
		windowsvirtualmachinescaleset.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
