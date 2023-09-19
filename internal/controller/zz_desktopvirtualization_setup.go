/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	virtualdesktopapplication "github.com/upbound/provider-azure/internal/controller/desktopvirtualization/virtualdesktopapplication"
	virtualdesktopapplicationgroup "github.com/upbound/provider-azure/internal/controller/desktopvirtualization/virtualdesktopapplicationgroup"
	virtualdesktophostpoolregistrationinfo "github.com/upbound/provider-azure/internal/controller/desktopvirtualization/virtualdesktophostpoolregistrationinfo"
	virtualdesktopscalingplan "github.com/upbound/provider-azure/internal/controller/desktopvirtualization/virtualdesktopscalingplan"
	virtualdesktopworkspace "github.com/upbound/provider-azure/internal/controller/desktopvirtualization/virtualdesktopworkspace"
	virtualdesktopworkspaceapplicationgroupassociation "github.com/upbound/provider-azure/internal/controller/desktopvirtualization/virtualdesktopworkspaceapplicationgroupassociation"
)

// Setup_desktopvirtualization creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_desktopvirtualization(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		virtualdesktopapplication.Setup,
		virtualdesktopapplicationgroup.Setup,
		virtualdesktophostpoolregistrationinfo.Setup,
		virtualdesktopscalingplan.Setup,
		virtualdesktopworkspace.Setup,
		virtualdesktopworkspaceapplicationgroupassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
