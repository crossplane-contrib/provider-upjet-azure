// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	managementgrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/managementgrouppolicyassignment"
	managementgrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/managementgrouppolicyexemption"
	managementlock "github.com/upbound/provider-azure/internal/controller/authorization/managementlock"
	pimactiveroleassignment "github.com/upbound/provider-azure/internal/controller/authorization/pimactiveroleassignment"
	pimeligibleroleassignment "github.com/upbound/provider-azure/internal/controller/authorization/pimeligibleroleassignment"
	policydefinition "github.com/upbound/provider-azure/internal/controller/authorization/policydefinition"
	policysetdefinition "github.com/upbound/provider-azure/internal/controller/authorization/policysetdefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/resourcegrouppolicyassignment"
	resourcegrouppolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/resourcegrouppolicyexemption"
	resourcepolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/resourcepolicyexemption"
	roleassignment "github.com/upbound/provider-azure/internal/controller/authorization/roleassignment"
	roledefinition "github.com/upbound/provider-azure/internal/controller/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/internal/controller/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/internal/controller/authorization/subscriptionpolicyexemption"
	trustedaccessrolebinding "github.com/upbound/provider-azure/internal/controller/authorization/trustedaccessrolebinding"
)

// Setup_authorization creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_authorization(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementgrouppolicyassignment.Setup,
		managementgrouppolicyexemption.Setup,
		managementlock.Setup,
		pimactiveroleassignment.Setup,
		pimeligibleroleassignment.Setup,
		policydefinition.Setup,
		policysetdefinition.Setup,
		resourcegrouppolicyassignment.Setup,
		resourcegrouppolicyexemption.Setup,
		resourcepolicyassignment.Setup,
		resourcepolicyexemption.Setup,
		roleassignment.Setup,
		roledefinition.Setup,
		subscriptionpolicyassignment.Setup,
		subscriptionpolicyexemption.Setup,
		trustedaccessrolebinding.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
