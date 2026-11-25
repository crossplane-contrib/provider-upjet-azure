// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package communication

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_communication_service", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
	})
}
