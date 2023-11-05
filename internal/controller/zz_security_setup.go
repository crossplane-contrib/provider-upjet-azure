// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	advancedthreatprotection "github.com/upbound/provider-azure/internal/controller/security/advancedthreatprotection"
	iotsecuritydevicegroup "github.com/upbound/provider-azure/internal/controller/security/iotsecuritydevicegroup"
	iotsecuritysolution "github.com/upbound/provider-azure/internal/controller/security/iotsecuritysolution"
	securitycenterassessment "github.com/upbound/provider-azure/internal/controller/security/securitycenterassessment"
	securitycenterassessmentpolicy "github.com/upbound/provider-azure/internal/controller/security/securitycenterassessmentpolicy"
	securitycenterautoprovisioning "github.com/upbound/provider-azure/internal/controller/security/securitycenterautoprovisioning"
	securitycentercontact "github.com/upbound/provider-azure/internal/controller/security/securitycentercontact"
	securitycenterservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/security/securitycenterservervulnerabilityassessment"
	securitycenterservervulnerabilityassessmentvirtualmachine "github.com/upbound/provider-azure/internal/controller/security/securitycenterservervulnerabilityassessmentvirtualmachine"
	securitycentersetting "github.com/upbound/provider-azure/internal/controller/security/securitycentersetting"
	securitycentersubscriptionpricing "github.com/upbound/provider-azure/internal/controller/security/securitycentersubscriptionpricing"
	securitycenterworkspace "github.com/upbound/provider-azure/internal/controller/security/securitycenterworkspace"
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
		securitycenterservervulnerabilityassessment.Setup,
		securitycenterservervulnerabilityassessmentvirtualmachine.Setup,
		securitycentersetting.Setup,
		securitycentersubscriptionpricing.Setup,
		securitycenterworkspace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
