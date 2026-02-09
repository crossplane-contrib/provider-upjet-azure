// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	activedirectoryadministrator "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/activedirectoryadministrator"
	configuration "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/configuration"
	database "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/database"
	firewallrule "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/firewallrule"
	flexibleserver "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/flexibleserver"
	flexibleserveractivedirectoryadministrator "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/flexibleserveractivedirectoryadministrator"
	flexibleserverbackup "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/flexibleserverbackup"
	flexibleserverconfiguration "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/flexibleserverfirewallrule"
	flexibleservervirtualendpoint "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/flexibleservervirtualendpoint"
	server "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/server"
	serverkey "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/serverkey"
	virtualnetworkrule "github.com/upbound/provider-azure/v2/internal/controller/cluster/dbforpostgresql/virtualnetworkrule"
)

// Setup_dbforpostgresql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbforpostgresql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activedirectoryadministrator.Setup,
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		flexibleserver.Setup,
		flexibleserveractivedirectoryadministrator.Setup,
		flexibleserverbackup.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		flexibleservervirtualendpoint.Setup,
		server.Setup,
		serverkey.Setup,
		virtualnetworkrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dbforpostgresql creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dbforpostgresql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activedirectoryadministrator.SetupGated,
		configuration.SetupGated,
		database.SetupGated,
		firewallrule.SetupGated,
		flexibleserver.SetupGated,
		flexibleserveractivedirectoryadministrator.SetupGated,
		flexibleserverbackup.SetupGated,
		flexibleserverconfiguration.SetupGated,
		flexibleserverdatabase.SetupGated,
		flexibleserverfirewallrule.SetupGated,
		flexibleservervirtualendpoint.SetupGated,
		server.SetupGated,
		serverkey.SetupGated,
		virtualnetworkrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
