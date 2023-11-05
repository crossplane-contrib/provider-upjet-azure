// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	agentpool "github.com/upbound/provider-azure/internal/controller/containerregistry/agentpool"
	containerconnectedregistry "github.com/upbound/provider-azure/internal/controller/containerregistry/containerconnectedregistry"
	registry "github.com/upbound/provider-azure/internal/controller/containerregistry/registry"
	scopemap "github.com/upbound/provider-azure/internal/controller/containerregistry/scopemap"
	token "github.com/upbound/provider-azure/internal/controller/containerregistry/token"
	tokenpassword "github.com/upbound/provider-azure/internal/controller/containerregistry/tokenpassword"
	webhook "github.com/upbound/provider-azure/internal/controller/containerregistry/webhook"
)

// Setup_containerregistry creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerregistry(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentpool.Setup,
		containerconnectedregistry.Setup,
		registry.Setup,
		scopemap.Setup,
		token.Setup,
		tokenpassword.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
