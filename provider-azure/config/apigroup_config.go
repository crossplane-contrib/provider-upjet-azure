/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
		"azurerm_resource_provider_registration":                         "policyinsights",
		"azurerm_storage_blob":                                           "storage",
		"azurerm_storage_container":                                      "storage",
		"azurerm_storage_table":                                          "storage",
		"azurerm_storage_queue":                                          "storage",
		"azurerm_storage_share":                                          "storage",
		"azurerm_storage_data_lake_gen2_filesystem":                      "storage",
	}

	// this table holds overrides of Microsoft provider API groups
	// for example API group extracted from Microsoft provider "DocumentDB"
	// is overridden as "cosmodb".
	apiGroupOverrides = map[string]string{
		"documentdb": "cosmosdb",
	}
)
