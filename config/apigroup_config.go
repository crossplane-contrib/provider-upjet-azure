// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

// nolint:misspell
var (
	// provider-azure uses the Microsoft service names extracted
	// from Terraform import statements, e.g.:
	// /providers/Microsoft.Subscription/aliases/subscription1 -> subscription
	// this table holds only overrides for the resources.
	resourceAPIGroupMap = map[string]string{
		"azurerm_key_vault_certificate":                                  "keyvault",
		"azurerm_key_vault_certificate_issuer":                           "keyvault",
		"azurerm_key_vault_key":                                          "keyvault",
		"azurerm_key_vault_secret":                                       "keyvault",
		"azurerm_key_vault_managed_storage_account":                      "keyvault",
		"azurerm_key_vault_managed_storage_account_sas_token_definition": "keyvault",
		"azurerm_key_vault_certificate_contacts":                         "keyvault",
		"azurerm_resource_provider_registration":                         "policyinsights",
		"azurerm_storage_blob":                                           "storage",
		"azurerm_storage_container":                                      "storage",
		"azurerm_storage_table":                                          "storage",
		"azurerm_storage_queue":                                          "storage",
		"azurerm_storage_share":                                          "storage",
		"azurerm_storage_data_lake_gen2_filesystem":                      "storage",
		"azurerm_monitor_diagnostic_setting":                             "insights",
		"azurerm_monitor_workspace":                                      "insights",
	}

	// this table holds overrides of Microsoft provider API groups
	// for example API group extracted from Microsoft provider "DocumentDB"
	// is overridden as "cosmodb".
	apiGroupOverrides = map[string]string{
		"documentdb": "cosmosdb",
	}
)
