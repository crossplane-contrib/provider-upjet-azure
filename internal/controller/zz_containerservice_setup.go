// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	kubernetescluster "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetescluster"
	kubernetesclusternodepool "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetesclusternodepool"
	kubernetesfleetmanager "github.com/upbound/provider-azure/internal/controller/containerservice/kubernetesfleetmanager"
)

// Setup_containerservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kubernetescluster.Setup,
		kubernetesclusternodepool.Setup,
		kubernetesfleetmanager.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
