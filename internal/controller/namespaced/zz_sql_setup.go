// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	mssqldatabase "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlfirewallrule"
	mssqljobagent "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqljobagent"
	mssqljobcredential "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqljobcredential"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancetransparentdataencryption "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancetransparentdataencryption"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/namespaced/sql/mssqlvirtualnetworkrule"
)

// Setup_sql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		mssqldatabase.Setup,
		mssqldatabaseextendedauditingpolicy.Setup,
		mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		mssqlelasticpool.Setup,
		mssqlfailovergroup.Setup,
		mssqlfirewallrule.Setup,
		mssqljobagent.Setup,
		mssqljobcredential.Setup,
		mssqlmanageddatabase.Setup,
		mssqlmanagedinstance.Setup,
		mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		mssqlmanagedinstancefailovergroup.Setup,
		mssqlmanagedinstancetransparentdataencryption.Setup,
		mssqlmanagedinstancevulnerabilityassessment.Setup,
		mssqloutboundfirewallrule.Setup,
		mssqlserver.Setup,
		mssqlserverdnsalias.Setup,
		mssqlservermicrosoftsupportauditingpolicy.Setup,
		mssqlserversecurityalertpolicy.Setup,
		mssqlservertransparentdataencryption.Setup,
		mssqlservervulnerabilityassessment.Setup,
		mssqlvirtualnetworkrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_sql creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_sql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		mssqldatabase.SetupGated,
		mssqldatabaseextendedauditingpolicy.SetupGated,
		mssqldatabasevulnerabilityassessmentrulebaseline.SetupGated,
		mssqlelasticpool.SetupGated,
		mssqlfailovergroup.SetupGated,
		mssqlfirewallrule.SetupGated,
		mssqljobagent.SetupGated,
		mssqljobcredential.SetupGated,
		mssqlmanageddatabase.SetupGated,
		mssqlmanagedinstance.SetupGated,
		mssqlmanagedinstanceactivedirectoryadministrator.SetupGated,
		mssqlmanagedinstancefailovergroup.SetupGated,
		mssqlmanagedinstancetransparentdataencryption.SetupGated,
		mssqlmanagedinstancevulnerabilityassessment.SetupGated,
		mssqloutboundfirewallrule.SetupGated,
		mssqlserver.SetupGated,
		mssqlserverdnsalias.SetupGated,
		mssqlservermicrosoftsupportauditingpolicy.SetupGated,
		mssqlserversecurityalertpolicy.SetupGated,
		mssqlservertransparentdataencryption.SetupGated,
		mssqlservervulnerabilityassessment.SetupGated,
		mssqlvirtualnetworkrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
