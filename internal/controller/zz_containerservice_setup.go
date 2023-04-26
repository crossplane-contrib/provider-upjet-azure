/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

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
