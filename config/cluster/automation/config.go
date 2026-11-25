// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package automation

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/config/cluster/common"
)

// Configure configures automation group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_automation_account", func(r *config.Resource) {
		r.Version = "v1beta3"
		r.PreviousVersions = []string{"v1beta2"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta2")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"encryption"},
		}
	})
}
