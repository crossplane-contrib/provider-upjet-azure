// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/cluster/cognitiveservices/account"
	aiservices "github.com/upbound/provider-azure/internal/controller/cluster/cognitiveservices/aiservices"
	deployment "github.com/upbound/provider-azure/internal/controller/cluster/cognitiveservices/deployment"
)

// Setup_cognitiveservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cognitiveservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		aiservices.Setup,
		deployment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cognitiveservices creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cognitiveservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		aiservices.SetupGated,
		deployment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
