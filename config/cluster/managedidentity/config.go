// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package managedidentity

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/v2/apis/cluster/rconfig"
	"github.com/upbound/provider-azure/v2/config/cluster/common"
)

// Configure configures managedidentity group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_federated_identity_credential", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.ControllerReconcileVersion = r.Version //nolint:staticcheck // upjet does not default it at runtime
		common.DeprecateVersionForAzureRMV5(r, "v1beta1")
		r.References["user_assigned_identity_id"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
