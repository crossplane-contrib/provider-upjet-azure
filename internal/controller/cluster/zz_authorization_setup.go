// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	managementgrouppolicyassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/managementgrouppolicyassignment"
	managementgrouppolicyexemption "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/managementgrouppolicyexemption"
	managementlock "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/managementlock"
	pimactiveroleassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/pimactiveroleassignment"
	pimeligibleroleassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/pimeligibleroleassignment"
	policydefinition "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/policydefinition"
	policysetdefinition "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/policysetdefinition"
	resourcegrouppolicyassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/resourcegrouppolicyassignment"
	resourcegrouppolicyexemption "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/resourcegrouppolicyexemption"
	resourcepolicyassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/resourcepolicyassignment"
	resourcepolicyexemption "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/resourcepolicyexemption"
	roleassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/roleassignment"
	roledefinition "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/roledefinition"
	subscriptionpolicyassignment "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/subscriptionpolicyassignment"
	subscriptionpolicyexemption "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/subscriptionpolicyexemption"
	trustedaccessrolebinding "github.com/upbound/provider-azure/v2/internal/controller/cluster/authorization/trustedaccessrolebinding"
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

// SetupGated_authorization creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_authorization(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementgrouppolicyassignment.SetupGated,
		managementgrouppolicyexemption.SetupGated,
		managementlock.SetupGated,
		pimactiveroleassignment.SetupGated,
		pimeligibleroleassignment.SetupGated,
		policydefinition.SetupGated,
		policysetdefinition.SetupGated,
		resourcegrouppolicyassignment.SetupGated,
		resourcegrouppolicyexemption.SetupGated,
		resourcepolicyassignment.SetupGated,
		resourcepolicyexemption.SetupGated,
		roleassignment.SetupGated,
		roledefinition.SetupGated,
		subscriptionpolicyassignment.SetupGated,
		subscriptionpolicyexemption.SetupGated,
		trustedaccessrolebinding.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
