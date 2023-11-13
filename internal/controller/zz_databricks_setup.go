// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accessconnector "github.com/upbound/provider-azure/internal/controller/databricks/accessconnector"
	workspace "github.com/upbound/provider-azure/internal/controller/databricks/workspace"
	workspacecustomermanagedkey "github.com/upbound/provider-azure/internal/controller/databricks/workspacecustomermanagedkey"
	workspacerootdbfscustomermanagedkey "github.com/upbound/provider-azure/internal/controller/databricks/workspacerootdbfscustomermanagedkey"
)

// Setup_databricks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_databricks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessconnector.Setup,
		workspace.Setup,
		workspacecustomermanagedkey.Setup,
		workspacerootdbfscustomermanagedkey.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
