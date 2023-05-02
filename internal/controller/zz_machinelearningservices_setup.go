/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	computecluster "github.com/upbound/provider-azure/internal/controller/machinelearningservices/computecluster"
	computeinstance "github.com/upbound/provider-azure/internal/controller/machinelearningservices/computeinstance"
	synapsespark "github.com/upbound/provider-azure/internal/controller/machinelearningservices/synapsespark"
	workspace "github.com/upbound/provider-azure/internal/controller/machinelearningservices/workspace"
)

// Setup_machinelearningservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_machinelearningservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		computecluster.Setup,
		computeinstance.Setup,
		synapsespark.Setup,
		workspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
