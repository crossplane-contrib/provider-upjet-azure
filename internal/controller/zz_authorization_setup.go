// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	managementlock "github.com/upbound/provider-azure/internal/controller/authorization/managementlock"
	policydefinition "github.com/upbound/provider-azure/internal/controller/authorization/policydefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	resourcepolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/resourcepolicyexemption"
	roleassignment "github.com/upbound/provider-azure/internal/controller/authorization/roleassignment"
	roledefinition "github.com/upbound/provider-azure/internal/controller/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/subscriptionpolicyexemption"
)

// Setup_authorization creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_authorization(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementlock.Setup,
		policydefinition.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcepolicyassignment.Setup,
		resourcepolicyexemption.Setup,
		roleassignment.Setup,
		roledefinition.Setup,
		subscriptionpolicyassignment.Setup,
		subscriptionpolicyexemption.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
