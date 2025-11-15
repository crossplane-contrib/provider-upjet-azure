// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	globalvmshutdownschedule "github.com/upbound/provider-azure/v2/internal/controller/cluster/devtestlab/globalvmshutdownschedule"
	lab "github.com/upbound/provider-azure/v2/internal/controller/cluster/devtestlab/lab"
	linuxvirtualmachine "github.com/upbound/provider-azure/v2/internal/controller/cluster/devtestlab/linuxvirtualmachine"
	policy "github.com/upbound/provider-azure/v2/internal/controller/cluster/devtestlab/policy"
	schedule "github.com/upbound/provider-azure/v2/internal/controller/cluster/devtestlab/schedule"
	virtualnetwork "github.com/upbound/provider-azure/v2/internal/controller/cluster/devtestlab/virtualnetwork"
	windowsvirtualmachine "github.com/upbound/provider-azure/v2/internal/controller/cluster/devtestlab/windowsvirtualmachine"
)

// Setup_devtestlab creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devtestlab(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		globalvmshutdownschedule.Setup,
		lab.Setup,
		linuxvirtualmachine.Setup,
		policy.Setup,
		schedule.Setup,
		virtualnetwork.Setup,
		windowsvirtualmachine.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_devtestlab creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_devtestlab(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		globalvmshutdownschedule.SetupGated,
		lab.SetupGated,
		linuxvirtualmachine.SetupGated,
		policy.SetupGated,
		schedule.SetupGated,
		virtualnetwork.SetupGated,
		windowsvirtualmachine.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
