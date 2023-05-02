/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	configuration "github.com/upbound/provider-azure/internal/controller/dbformariadb/configuration"
	database "github.com/upbound/provider-azure/internal/controller/dbformariadb/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/dbformariadb/firewallrule"
	server "github.com/upbound/provider-azure/internal/controller/dbformariadb/server"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/dbformariadb/virtualnetworkrule"
)

// Setup_dbformariadb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbformariadb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
