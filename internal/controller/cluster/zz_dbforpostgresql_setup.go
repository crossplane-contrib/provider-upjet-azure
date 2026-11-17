// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	flexibleserver "github.com/upbound/provider-azure/v3/internal/controller/cluster/dbforpostgresql/flexibleserver"
	flexibleserveractivedirectoryadministrator "github.com/upbound/provider-azure/v3/internal/controller/cluster/dbforpostgresql/flexibleserveractivedirectoryadministrator"
	flexibleserverbackup "github.com/upbound/provider-azure/v3/internal/controller/cluster/dbforpostgresql/flexibleserverbackup"
	flexibleserverconfiguration "github.com/upbound/provider-azure/v3/internal/controller/cluster/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/v3/internal/controller/cluster/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/v3/internal/controller/cluster/dbforpostgresql/flexibleserverfirewallrule"
	flexibleservervirtualendpoint "github.com/upbound/provider-azure/v3/internal/controller/cluster/dbforpostgresql/flexibleservervirtualendpoint"
)

// Setup_dbforpostgresql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbforpostgresql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		flexibleserver.Setup,
		flexibleserveractivedirectoryadministrator.Setup,
		flexibleserverbackup.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		flexibleservervirtualendpoint.Setup,
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
		flexibleserver.SetupGated,
		flexibleserveractivedirectoryadministrator.SetupGated,
		flexibleserverbackup.SetupGated,
		flexibleserverconfiguration.SetupGated,
		flexibleserverdatabase.SetupGated,
		flexibleserverfirewallrule.SetupGated,
		flexibleservervirtualendpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_dbforpostgresql registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_dbforpostgresql(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		flexibleserver.SetupWebhookWithManager,
		flexibleserveractivedirectoryadministrator.SetupWebhookWithManager,
		flexibleserverbackup.SetupWebhookWithManager,
		flexibleserverconfiguration.SetupWebhookWithManager,
		flexibleserverdatabase.SetupWebhookWithManager,
		flexibleserverfirewallrule.SetupWebhookWithManager,
		flexibleservervirtualendpoint.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
