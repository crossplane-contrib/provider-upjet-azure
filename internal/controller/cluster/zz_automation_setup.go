// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/cluster/automation/account"
	connection "github.com/upbound/provider-azure/internal/controller/cluster/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/internal/controller/cluster/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/internal/controller/cluster/automation/connectiontype"
	credential "github.com/upbound/provider-azure/internal/controller/cluster/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/internal/controller/cluster/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/internal/controller/cluster/automation/module"
	runbook "github.com/upbound/provider-azure/internal/controller/cluster/automation/runbook"
	schedule "github.com/upbound/provider-azure/internal/controller/cluster/automation/schedule"
	variablebool "github.com/upbound/provider-azure/internal/controller/cluster/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/internal/controller/cluster/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/internal/controller/cluster/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/cluster/automation/variablestring"
	webhook "github.com/upbound/provider-azure/internal/controller/cluster/automation/webhook"
)

// Setup_automation creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_automation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		connection.Setup,
		connectionclassiccertificate.Setup,
		connectiontype.Setup,
		credential.Setup,
		hybridrunbookworkergroup.Setup,
		module.Setup,
		runbook.Setup,
		schedule.Setup,
		variablebool.Setup,
		variabledatetime.Setup,
		variableint.Setup,
		variablestring.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
