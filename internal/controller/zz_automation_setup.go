// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/automation/account"
	connection "github.com/upbound/provider-azure/internal/controller/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/internal/controller/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/internal/controller/automation/connectiontype"
	credential "github.com/upbound/provider-azure/internal/controller/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/internal/controller/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/internal/controller/automation/module"
	runbook "github.com/upbound/provider-azure/internal/controller/automation/runbook"
	schedule "github.com/upbound/provider-azure/internal/controller/automation/schedule"
	variablebool "github.com/upbound/provider-azure/internal/controller/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/internal/controller/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/internal/controller/automation/variableint"
	variablestring "github.com/upbound/provider-azure/internal/controller/automation/variablestring"
	webhook "github.com/upbound/provider-azure/internal/controller/automation/webhook"
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
