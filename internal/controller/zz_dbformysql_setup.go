// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/dbformysql/activedirectoryadministrator"
	configuration "github.com/upbound/provider-azure/internal/controller/dbformysql/configuration"
	database "github.com/upbound/provider-azure/internal/controller/dbformysql/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/dbformysql/firewallrule"
	flexibledatabase "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibledatabase"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/dbformysql/flexibleserverfirewallrule"
	server "github.com/upbound/provider-azure/internal/controller/dbformysql/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/dbformysql/virtualnetworkrule"
)

// Setup_dbformysql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbformysql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activedirectoryadministrator.Setup,
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		flexibledatabase.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverfirewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
