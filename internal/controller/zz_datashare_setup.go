/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/upbound/provider-azure/internal/controller/datashare/account"
	datasetblobstorage "github.com/upbound/provider-azure/internal/controller/datashare/datasetblobstorage"
	datasetdatalakegen2 "github.com/upbound/provider-azure/internal/controller/datashare/datasetdatalakegen2"
	datasetkustocluster "github.com/upbound/provider-azure/internal/controller/datashare/datasetkustocluster"
	datasetkustodatabase "github.com/upbound/provider-azure/internal/controller/datashare/datasetkustodatabase"
	datashare "github.com/upbound/provider-azure/internal/controller/datashare/datashare"
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
