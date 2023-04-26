/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accessconnector "github.com/upbound/provider-azure/internal/controller/databricks/accessconnector"
	workspace "github.com/upbound/provider-azure/internal/controller/databricks/workspace"
	workspacecustomermanagedkey "github.com/upbound/provider-azure/internal/controller/databricks/workspacecustomermanagedkey"
)

// Setup_databricks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_databricks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessconnector.Setup,
		workspace.Setup,
		workspacecustomermanagedkey.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
