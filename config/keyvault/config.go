// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package keyvault

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures keyvault group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_key_vault", func(r *config.Resource) {
		// Mutually exclusive with azurerm_key_vault_access_policy
		config.MoveToStatus(r.TerraformResource, "access_policy")
	})

	p.AddResourceConfigurator("azurerm_key_vault_secret", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_key", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"rotation_policy"},
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_certificate", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_certificate_issuer", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_managed_storage_account", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_access_policy", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.OverrideFieldNames["AccessPolicyParameters"] = "AccessPolicyParameters_2"
		r.OverrideFieldNames["AccessPolicyInitParameters"] = "AccessPolicyInitParameters_2"
		r.OverrideFieldNames["AccessPolicyObservation"] = "AccessPolicyObservation_2"
	})

	p.AddResourceConfigurator("azurerm_key_vault_managed_storage_account_sas_token_definition", func(r *config.Resource) {
		r.References["managed_storage_account_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_managed_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_key_vault_certificate_contacts", func(r *config.Resource) {
		r.OverrideFieldNames["ContactParameters"] = "CertificateContactsContactParameters"
		r.OverrideFieldNames["ContactInitParameters"] = "CertificateContactsContactInitParameters"
		r.OverrideFieldNames["ContactObservation"] = "CertificateContactsContactObservation"
	})
}
