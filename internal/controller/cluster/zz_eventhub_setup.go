// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	authorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/authorizationrule"
	consumergroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/consumergroup"
	eventhub "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhub"
	eventhubnamespace "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/eventhubnamespace"
	namespaceauthorizationrule "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespacedisasterrecoveryconfig"
	namespaceschemagroup "github.com/upbound/provider-azure/internal/controller/cluster/eventhub/namespaceschemagroup"
)

// Setup_eventhub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizationrule.Setup,
		consumergroup.Setup,
		eventhub.Setup,
		eventhubnamespace.Setup,
		namespaceauthorizationrule.Setup,
		namespacedisasterrecoveryconfig.Setup,
		namespaceschemagroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
