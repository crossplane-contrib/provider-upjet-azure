// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	kubernetescluster "github.com/upbound/provider-azure/v2/internal/controller/namespaced/containerservice/kubernetescluster"
	kubernetesclusterextension "github.com/upbound/provider-azure/v2/internal/controller/namespaced/containerservice/kubernetesclusterextension"
	kubernetesclusternodepool "github.com/upbound/provider-azure/v2/internal/controller/namespaced/containerservice/kubernetesclusternodepool"
	kubernetesfleetmanager "github.com/upbound/provider-azure/v2/internal/controller/namespaced/containerservice/kubernetesfleetmanager"
)

// Setup_containerservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kubernetescluster.Setup,
		kubernetesclusterextension.Setup,
		kubernetesclusternodepool.Setup,
		kubernetesfleetmanager.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_containerservice creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_containerservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kubernetescluster.SetupGated,
		kubernetesclusterextension.SetupGated,
		kubernetesclusternodepool.SetupGated,
		kubernetesfleetmanager.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
