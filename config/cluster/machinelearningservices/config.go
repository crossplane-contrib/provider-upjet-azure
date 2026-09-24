// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package machinelearningservices

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/provider-azure/v2/apis/cluster/rconfig"
)

const group = "machinelearningservices"

// Configure configures machinelearningservices group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_ai_foundry", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AIFoundry"
		r.References["key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage.storage_account_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["application_insights_id"] = config.Reference{
			TerraformName: "azurerm_application_insights",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["container_registry_id"] = config.Reference{
			TerraformName: "azurerm_container_registry",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["primary_user_assigned_identity"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["encryption.key_vault_id"] = config.Reference{
			TerraformName: "azurerm_key_vault",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["encryption.user_assigned_identity_id"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_machine_learning_workspace_network_outbound_rule_fqdn", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "WorkspaceOutboundRuleFqdn"
		r.References["workspace_id"] = config.Reference{
			TerraformName: "azurerm_machine_learning_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_machine_learning_workspace_network_outbound_rule_service_tag", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "WorkspaceOutboundRuleServiceTag"
		r.References["workspace_id"] = config.Reference{
			TerraformName: "azurerm_machine_learning_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_machine_learning_workspace_network_outbound_rule_private_endpoint", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "WorkspaceOutboundRulePrivateEndpoint"
		r.References["service_resource_id"] = config.Reference{
			TerraformName: "azurerm_storage_account",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_ai_foundry_project", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AIFoundryProject"
		r.References["ai_services_hub_id"] = config.Reference{
			TerraformName: "azurerm_ai_foundry",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["primary_user_assigned_identity"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	for _, name := range []string{
		"azurerm_machine_learning_datastore_blobstorage",
		"azurerm_machine_learning_datastore_datalake_gen2",
	} {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.ShortGroup = group
			r.References["workspace_id"] = config.Reference{
				TerraformName: "azurerm_machine_learning_workspace",
				Extractor:     rconfig.ExtractResourceIDFuncPath,
			}
			r.References["storage_container_id"] = config.Reference{
				TerraformName: "azurerm_storage_container",
				Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("resource_manager_id",true)`,
			}
		})
	}

	p.AddResourceConfigurator("azurerm_machine_learning_datastore_fileshare", func(r *config.Resource) {
		r.ShortGroup = group
		r.References["workspace_id"] = config.Reference{
			TerraformName: "azurerm_machine_learning_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["storage_fileshare_id"] = config.Reference{
			TerraformName: "azurerm_storage_share",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("rbac_scope_id",true)`,
		}
	})

	p.AddResourceConfigurator("azurerm_machine_learning_inference_cluster", func(r *config.Resource) {
		r.ShortGroup = group
		r.TerraformResource.Schema["ssl"].Elem.(*schema.Resource).
			Schema["key"].Sensitive = true
		r.References["machine_learning_workspace_id"] = config.Reference{
			TerraformName: "azurerm_machine_learning_workspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["kubernetes_cluster_id"] = config.Reference{
			TerraformName: "azurerm_kubernetes_cluster",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["identity.identity_ids"] = config.Reference{
			TerraformName: "azurerm_user_assigned_identity",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
