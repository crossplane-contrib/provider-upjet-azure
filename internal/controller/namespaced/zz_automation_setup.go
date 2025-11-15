// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/account"
	connection "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/connection"
	connectionclassiccertificate "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/connectionclassiccertificate"
	connectiontype "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/connectiontype"
	credential "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/credential"
	hybridrunbookworkergroup "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/hybridrunbookworkergroup"
	module "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/module"
	runbook "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/runbook"
	schedule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/schedule"
	variablebool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/variablebool"
	variabledatetime "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/variabledatetime"
	variableint "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/variableint"
	variablestring "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/variablestring"
	webhook "github.com/upbound/provider-azure/v2/internal/controller/namespaced/automation/webhook"
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

// SetupGated_automation creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_automation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		connection.SetupGated,
		connectionclassiccertificate.SetupGated,
		connectiontype.SetupGated,
		credential.SetupGated,
		hybridrunbookworkergroup.SetupGated,
		module.SetupGated,
		runbook.SetupGated,
		schedule.SetupGated,
		variablebool.SetupGated,
		variabledatetime.SetupGated,
		variableint.SetupGated,
		variablestring.SetupGated,
		webhook.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
