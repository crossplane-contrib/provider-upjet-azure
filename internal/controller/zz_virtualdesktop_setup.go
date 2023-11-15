// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/upbound/provider-azure/internal/controller/virtualdesktop/application"
	applicationgroup "github.com/upbound/provider-azure/internal/controller/virtualdesktop/applicationgroup"
	hostpool "github.com/upbound/provider-azure/internal/controller/virtualdesktop/hostpool"
	hostpoolregistrationinfo "github.com/upbound/provider-azure/internal/controller/virtualdesktop/hostpoolregistrationinfo"
	scalingplan "github.com/upbound/provider-azure/internal/controller/virtualdesktop/scalingplan"
	workspace "github.com/upbound/provider-azure/internal/controller/virtualdesktop/workspace"
	workspaceapplicationgroupassociation "github.com/upbound/provider-azure/internal/controller/virtualdesktop/workspaceapplicationgroupassociation"
)

// Setup_virtualdesktop creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_virtualdesktop(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationgroup.Setup,
		hostpool.Setup,
		hostpoolregistrationinfo.Setup,
		scalingplan.Setup,
		workspace.Setup,
		workspaceapplicationgroupassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
