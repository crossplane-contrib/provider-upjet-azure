// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	availabilityset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/availabilityset"
	capacityreservation "github.com/upbound/provider-azure/internal/controller/namespaced/compute/capacityreservation"
	capacityreservationgroup "github.com/upbound/provider-azure/internal/controller/namespaced/compute/capacityreservationgroup"
	dedicatedhost "github.com/upbound/provider-azure/internal/controller/namespaced/compute/dedicatedhost"
	diskaccess "github.com/upbound/provider-azure/internal/controller/namespaced/compute/diskaccess"
	diskencryptionset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/diskencryptionset"
	galleryapplication "github.com/upbound/provider-azure/internal/controller/namespaced/compute/galleryapplication"
	galleryapplicationversion "github.com/upbound/provider-azure/internal/controller/namespaced/compute/galleryapplicationversion"
	image "github.com/upbound/provider-azure/internal/controller/namespaced/compute/image"
	linuxvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/compute/linuxvirtualmachine"
	linuxvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/linuxvirtualmachinescaleset"
	manageddisk "github.com/upbound/provider-azure/internal/controller/namespaced/compute/manageddisk"
	manageddisksastoken "github.com/upbound/provider-azure/internal/controller/namespaced/compute/manageddisksastoken"
	orchestratedvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/orchestratedvirtualmachinescaleset"
	proximityplacementgroup "github.com/upbound/provider-azure/internal/controller/namespaced/compute/proximityplacementgroup"
	sharedimage "github.com/upbound/provider-azure/internal/controller/namespaced/compute/sharedimage"
	sharedimagegallery "github.com/upbound/provider-azure/internal/controller/namespaced/compute/sharedimagegallery"
	snapshot "github.com/upbound/provider-azure/internal/controller/namespaced/compute/snapshot"
	sshpublickey "github.com/upbound/provider-azure/internal/controller/namespaced/compute/sshpublickey"
	virtualmachinedatadiskattachment "github.com/upbound/provider-azure/internal/controller/namespaced/compute/virtualmachinedatadiskattachment"
	virtualmachineextension "github.com/upbound/provider-azure/internal/controller/namespaced/compute/virtualmachineextension"
	virtualmachineruncommand "github.com/upbound/provider-azure/internal/controller/namespaced/compute/virtualmachineruncommand"
	windowsvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/compute/windowsvirtualmachine"
	windowsvirtualmachinescaleset "github.com/upbound/provider-azure/internal/controller/namespaced/compute/windowsvirtualmachinescaleset"
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
		windowsvirtualmachine.SetupGated,
		windowsvirtualmachinescaleset.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
