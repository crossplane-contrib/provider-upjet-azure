/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	authorizationrule "github.com/upbound/provider-azure/internal/controller/notificationhubs/authorizationrule"
	notificationhub "github.com/upbound/provider-azure/internal/controller/notificationhubs/notificationhub"
	notificationhubnamespace "github.com/upbound/provider-azure/internal/controller/notificationhubs/notificationhubnamespace"
)

// Setup_notificationhubs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_notificationhubs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authorizationrule.Setup,
		notificationhub.Setup,
		notificationhubnamespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
