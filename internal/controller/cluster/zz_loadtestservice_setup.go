// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	loadtest "github.com/upbound/provider-azure/v2/internal/controller/cluster/loadtestservice/loadtest"
)

// Setup_loadtestservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_loadtestservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loadtest.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_loadtestservice creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_loadtestservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loadtest.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
