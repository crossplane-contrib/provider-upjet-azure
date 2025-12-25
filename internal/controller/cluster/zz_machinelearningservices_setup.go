// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aifoundry "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/aifoundry"
	aifoundryproject "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/aifoundryproject"
	computecluster "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/computecluster"
	computeinstance "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/computeinstance"
	synapsespark "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/synapsespark"
	workspace "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/workspace"
	workspaceoutboundrulefqdn "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/workspaceoutboundrulefqdn"
	workspaceoutboundruleprivateendpoint "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/workspaceoutboundruleprivateendpoint"
	workspaceoutboundruleservicetag "github.com/upbound/provider-azure/v2/internal/controller/cluster/machinelearningservices/workspaceoutboundruleservicetag"
)

// Setup_machinelearningservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_machinelearningservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aifoundry.Setup,
		aifoundryproject.Setup,
		computecluster.Setup,
		computeinstance.Setup,
		synapsespark.Setup,
		workspace.Setup,
		workspaceoutboundrulefqdn.Setup,
		workspaceoutboundruleprivateendpoint.Setup,
		workspaceoutboundruleservicetag.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_machinelearningservices creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_machinelearningservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aifoundry.SetupGated,
		aifoundryproject.SetupGated,
		computecluster.SetupGated,
		computeinstance.SetupGated,
		synapsespark.SetupGated,
		workspace.SetupGated,
		workspaceoutboundrulefqdn.SetupGated,
		workspaceoutboundruleprivateendpoint.SetupGated,
		workspaceoutboundruleservicetag.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
