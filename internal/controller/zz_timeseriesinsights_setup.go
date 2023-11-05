// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	eventsourceeventhub "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/eventsourceeventhub"
	eventsourceiothub "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/eventsourceiothub"
	gen2environment "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/gen2environment"
	referencedataset "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/referencedataset"
	standardenvironment "github.com/upbound/provider-azure/internal/controller/timeseriesinsights/standardenvironment"
)

// Setup_timeseriesinsights creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_timeseriesinsights(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		eventsourceeventhub.Setup,
		eventsourceiothub.Setup,
		gen2environment.Setup,
		referencedataset.Setup,
		standardenvironment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
