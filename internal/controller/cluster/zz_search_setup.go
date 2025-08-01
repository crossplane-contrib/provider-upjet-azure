// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	service "github.com/upbound/provider-azure/internal/controller/cluster/search/service"
	sharedprivatelinkservice "github.com/upbound/provider-azure/internal/controller/cluster/search/sharedprivatelinkservice"
)

// Setup_search creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_search(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		service.Setup,
		sharedprivatelinkservice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
