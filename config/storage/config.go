// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package storage

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-azure/apis/rconfig"
	"github.com/upbound/provider-azure/config/common"
)

// Configure configures storage group
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("azurerm_storage_account", func(r *config.Resource) {
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// skip diff customization on create
			if state == nil || state.Empty() {
				return diff, nil
			}
			if config == nil {
				return nil, errors.New("resource config cannot be nil")
			}
			// skip no diff or destroy diffs
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			lengthDiffKeys := []string{
				"blob_properties.#",
				"share_properties.#",
				"queue_properties.#",
			}
			for _, key := range lengthDiffKeys {
				// Note(cem): We should consider filtering effectively-empty
				// (Old = "" and New = "") length keys (`*.#` and `*.%`) in
				// Upjet. Doing so _should be_ safe, but in case we are afraid
				// to deploy at scale, we might tie the functionality to a
				// resource configuration flag, as a precaution.
				if diff.Attributes[key] != nil && diff.Attributes[key].Old == "" && diff.Attributes[key].New == "" {
					delete(diff.Attributes, key)
				}
			}

			return diff, nil
		}
	})

	p.AddResourceConfigurator("azurerm_storage_blob", func(r *config.Resource) {
		r.References["storage_account_name"] = config.Reference{
			TerraformName: "azurerm_storage_account",
		}
		r.References["storage_container_name"] = config.Reference{
			TerraformName: "azurerm_storage_container",
		}
	})

	p.AddResourceConfigurator("azurerm_storage_container", func(r *config.Resource) {
		r.References["storage_account_name"] = config.Reference{
			TerraformName: "azurerm_storage_account",
		}
	})

	p.AddResourceConfigurator("azurerm_storage_data_lake_gen2_filesystem", func(r *config.Resource) {
		r.References["storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_storage_encryption_scope", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "storage_account_id")
	})

	p.AddResourceConfigurator("azurerm_storage_management_policy", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "storage_account_id")
	})
}
