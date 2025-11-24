// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	flexibledatabase "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbformysql/flexibleserverfirewallrule"
)

// Setup_dbformysql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbformysql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		flexibledatabase.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverfirewallrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dbformysql creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dbformysql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		flexibledatabase.SetupGated,
		flexibleserver.SetupGated,
		flexibleserverconfiguration.SetupGated,
		flexibleserverfirewallrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
