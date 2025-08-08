// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/namespaced/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/namespaced/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/namespaced/security/iotsecuritysolution"
	securitycenterassessment "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterassessment"
	securitycenterassessmentpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentercontact"
	securitycenterservervulnerabilityassessmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterservervulnerabilityassessmentvirtualmachine"
	securitycentersetting "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/upbound/provider-azure/internal/controller/namespaced/security/securitycenterworkspace"
	storagedefender "github.com/upbound/provider-azure/internal/controller/namespaced/security/storagedefender"
)

// Setup_security creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_security(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		advancedthreatprotection.Setup,
		iotsecuritydevicegroup.Setup,
		iotsecuritysolution.Setup,
		securitycenterassessment.Setup,
		securitycenterassessmentpolicy.Setup,
		securitycenterautoprovisioning.Setup,
		securitycentercontact.Setup,
		securitycenterservervulnerabilityassessmentvirtualmachine.Setup,
		securitycentersetting.Setup,
		securitycentersubscriptionpricing.Setup,
		securitycenterworkspace.Setup,
		storagedefender.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_security creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_security(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		advancedthreatprotection.SetupGated,
		iotsecuritydevicegroup.SetupGated,
		iotsecuritysolution.SetupGated,
		securitycenterassessment.SetupGated,
		securitycenterassessmentpolicy.SetupGated,
		securitycenterautoprovisioning.SetupGated,
		securitycentercontact.SetupGated,
		securitycenterservervulnerabilityassessmentvirtualmachine.SetupGated,
		securitycentersetting.SetupGated,
		securitycentersubscriptionpricing.SetupGated,
		securitycenterworkspace.SetupGated,
		storagedefender.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
