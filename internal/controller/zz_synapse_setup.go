// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	firewallrule "github.com/upbound/provider-azure/internal/controller/synapse/firewallrule"
	integrationruntimeazure "github.com/upbound/provider-azure/internal/controller/synapse/integrationruntimeazure"
	integrationruntimeselfhosted "github.com/upbound/provider-azure/internal/controller/synapse/integrationruntimeselfhosted"
	linkedservice "github.com/upbound/provider-azure/internal/controller/synapse/linkedservice"
	managedprivateendpoint "github.com/upbound/provider-azure/internal/controller/synapse/managedprivateendpoint"
	privatelinkhub "github.com/upbound/provider-azure/internal/controller/synapse/privatelinkhub"
	roleassignment "github.com/upbound/provider-azure/internal/controller/synapse/roleassignment"
	sparkpool "github.com/upbound/provider-azure/internal/controller/synapse/sparkpool"
	sqlpool "github.com/upbound/provider-azure/internal/controller/synapse/sqlpool"
	sqlpoolextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolextendedauditingpolicy"
	sqlpoolsecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolsecurityalertpolicy"
	sqlpoolworkloadclassifier "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolworkloadclassifier"
	sqlpoolworkloadgroup "github.com/upbound/provider-azure/internal/controller/synapse/sqlpoolworkloadgroup"
	workspace "github.com/upbound/provider-azure/internal/controller/synapse/workspace"
	workspaceaadadmin "github.com/upbound/provider-azure/internal/controller/synapse/workspaceaadadmin"
	workspaceextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/synapse/workspaceextendedauditingpolicy"
	workspacesecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/synapse/workspacesecurityalertpolicy"
	workspacesqlaadadmin "github.com/upbound/provider-azure/internal/controller/synapse/workspacesqlaadadmin"
	workspacevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/synapse/workspacevulnerabilityassessment"
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
