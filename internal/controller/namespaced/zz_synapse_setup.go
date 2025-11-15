// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	firewallrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/firewallrule"
	integrationruntimeazure "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/integrationruntimeazure"
	integrationruntimeselfhosted "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/linkedservice"
	managedprivateendpoint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/privatelinkhub"
	roleassignment "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/roleassignment"
	sparkpool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/sparkpool"
	sqlpool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolworkloadclassifier "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/sqlpoolworkloadclassifier"
	sqlpoolworkloadgroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/sqlpoolworkloadgroup"
	workspace "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/workspace"
	workspaceaadadmin "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/workspaceaadadmin"
	workspaceextendedauditingpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/workspaceextendedauditingpolicy"
	workspacesecurityalertpolicy "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/workspacesecurityalertpolicy"
	workspacesqlaadadmin "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/workspacesqlaadadmin"
	workspacevulnerabilityassessment "github.com/upbound/provider-azure/v2/internal/controller/namespaced/synapse/workspacevulnerabilityassessment"
)

// Setup_synapse creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_synapse(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewallrule.Setup,
		integrationruntimeazure.Setup,
		integrationruntimeselfhosted.Setup,
		linkedservice.Setup,
		managedprivateendpoint.Setup,
		privatelinkhub.Setup,
		roleassignment.Setup,
		sparkpool.Setup,
		sqlpool.Setup,
		sqlpoolextendedauditingpolicy.Setup,
		sqlpoolsecurityalertpolicy.Setup,
		sqlpoolworkloadclassifier.Setup,
		sqlpoolworkloadgroup.Setup,
		workspace.Setup,
		workspaceaadadmin.Setup,
		workspaceextendedauditingpolicy.Setup,
		workspacesecurityalertpolicy.Setup,
		workspacesqlaadadmin.Setup,
		workspacevulnerabilityassessment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_synapse creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_synapse(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewallrule.SetupGated,
		integrationruntimeazure.SetupGated,
		integrationruntimeselfhosted.SetupGated,
		linkedservice.SetupGated,
		managedprivateendpoint.SetupGated,
		privatelinkhub.SetupGated,
		roleassignment.SetupGated,
		sparkpool.SetupGated,
		sqlpool.SetupGated,
		sqlpoolextendedauditingpolicy.SetupGated,
		sqlpoolsecurityalertpolicy.SetupGated,
		sqlpoolworkloadclassifier.SetupGated,
		sqlpoolworkloadgroup.SetupGated,
		workspace.SetupGated,
		workspaceaadadmin.SetupGated,
		workspaceextendedauditingpolicy.SetupGated,
		workspacesecurityalertpolicy.SetupGated,
		workspacesqlaadadmin.SetupGated,
		workspacevulnerabilityassessment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
