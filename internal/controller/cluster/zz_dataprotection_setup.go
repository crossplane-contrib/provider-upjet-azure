// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backupinstanceblobstorage "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backupinstanceblobstorage"
	backupinstancedatalakestorage "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backupinstancedatalakestorage"
	backupinstancedisk "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backupinstancedisk"
	backupinstancekubernetescluster "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backupinstancekubernetescluster"
	backupinstancepostgresql "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backupinstancepostgresql"
	backupinstancepostgresqlflexibleserver "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backupinstancepostgresqlflexibleserver"
	backuppolicyblobstorage "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backuppolicyblobstorage"
	backuppolicydatalakestorage "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backuppolicydatalakestorage"
	backuppolicydisk "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backuppolicydisk"
	backuppolicykubernetescluster "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backuppolicykubernetescluster"
	backuppolicypostgresql "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backuppolicypostgresql"
	backuppolicypostgresqlflexibleserver "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backuppolicypostgresqlflexibleserver"
	backupvault "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/backupvault"
	resourceguard "github.com/upbound/provider-azure/v2/internal/controller/cluster/dataprotection/resourceguard"
)

// Setup_dataprotection creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataprotection(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupinstanceblobstorage.Setup,
		backupinstancedatalakestorage.Setup,
		backupinstancedisk.Setup,
		backupinstancekubernetescluster.Setup,
		backupinstancepostgresql.Setup,
		backupinstancepostgresqlflexibleserver.Setup,
		backuppolicyblobstorage.Setup,
		backuppolicydatalakestorage.Setup,
		backuppolicydisk.Setup,
		backuppolicykubernetescluster.Setup,
		backuppolicypostgresql.Setup,
		backuppolicypostgresqlflexibleserver.Setup,
		backupvault.Setup,
		resourceguard.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dataprotection creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dataprotection(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupinstanceblobstorage.SetupGated,
		backupinstancedatalakestorage.SetupGated,
		backupinstancedisk.SetupGated,
		backupinstancekubernetescluster.SetupGated,
		backupinstancepostgresql.SetupGated,
		backupinstancepostgresqlflexibleserver.SetupGated,
		backuppolicyblobstorage.SetupGated,
		backuppolicydatalakestorage.SetupGated,
		backuppolicydisk.SetupGated,
		backuppolicykubernetescluster.SetupGated,
		backuppolicypostgresql.SetupGated,
		backuppolicypostgresqlflexibleserver.SetupGated,
		backupvault.SetupGated,
		resourceguard.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_dataprotection registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_dataprotection(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		backupinstanceblobstorage.SetupWebhookWithManager,
		backupinstancedatalakestorage.SetupWebhookWithManager,
		backupinstancedisk.SetupWebhookWithManager,
		backupinstancekubernetescluster.SetupWebhookWithManager,
		backupinstancepostgresql.SetupWebhookWithManager,
		backupinstancepostgresqlflexibleserver.SetupWebhookWithManager,
		backuppolicyblobstorage.SetupWebhookWithManager,
		backuppolicydatalakestorage.SetupWebhookWithManager,
		backuppolicydisk.SetupWebhookWithManager,
		backuppolicykubernetescluster.SetupWebhookWithManager,
		backuppolicypostgresql.SetupWebhookWithManager,
		backuppolicypostgresqlflexibleserver.SetupWebhookWithManager,
		backupvault.SetupWebhookWithManager,
		resourceguard.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
