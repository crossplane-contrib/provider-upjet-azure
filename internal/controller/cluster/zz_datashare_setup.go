// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/cluster/datashare/account"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/cluster/datashare/datashare"
)

// Setup_datashare creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datashare(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		datasetblobstorage.Setup,
		datasetdatalakegen2.Setup,
		datasetkustocluster.Setup,
		datasetkustodatabase.Setup,
		datashare.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datashare creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datashare(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.SetupGated,
		datasetblobstorage.SetupGated,
		datasetdatalakegen2.SetupGated,
		datasetkustocluster.SetupGated,
		datasetkustodatabase.SetupGated,
		datashare.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
