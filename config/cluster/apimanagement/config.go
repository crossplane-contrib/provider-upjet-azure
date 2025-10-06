// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apimanagement

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-azure/apis/cluster/rconfig"
	"github.com/upbound/provider-azure/config/cluster/common"
)

// Configure configures apimanagement group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_api_management", func(r *config.Resource) {
		r.Kind = "Management"
		// Mutually exclusive with azurerm_api_management_custom_domain
		config.MoveToStatus(r.TerraformResource, "hostname_configuration")
	})
	p.AddResourceConfigurator("azurerm_api_management_api_operation", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			TerraformName: "azurerm_api_management_api",
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_api_policy", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			TerraformName: "azurerm_api_management_api",
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_api_schema", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			TerraformName: "azurerm_api_management_api",
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_product_api", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			TerraformName: "azurerm_api_management_api",
		}
		r.References["product_id"] = config.Reference{
			TerraformName: "azurerm_api_management_product",
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_product_policy", func(r *config.Resource) {
		r.References["product_id"] = config.Reference{
			TerraformName: "azurerm_api_management_product",
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_subscription", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			TerraformName: "azurerm_api_management_user",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["product_id"] = config.Reference{
			TerraformName: "azurerm_api_management_product",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_authorization_server", func(r *config.Resource) {
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_gateway_api", func(r *config.Resource) {
		r.References["gateway_id"] = config.Reference{
			TerraformName: "azurerm_api_management_gateway",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["api_id"] = config.Reference{
			TerraformName: "azurerm_api_management_api",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}

		r.TerraformCustomDiff = apiIdCustomDiff
	})
	p.AddResourceConfigurator("azurerm_api_management_gateway", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "api_management_id")
	})
	p.AddResourceConfigurator("azurerm_api_management_api_tag", func(r *config.Resource) {
		r.TerraformCustomDiff = apiIdCustomDiff
	})
	p.AddResourceConfigurator("azurerm_api_management_api_release", func(r *config.Resource) {
		r.TerraformCustomDiff = apiIdCustomDiff
	})
	p.AddResourceConfigurator("azurerm_api_management_product_group", func(r *config.Resource) {
		r.References["product_id"] = config.Reference{
			TerraformName: "azurerm_api_management_product",
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
		r.References["group_name"] = config.Reference{
			TerraformName: "azurerm_api_management_group",
		}
	})
	p.AddResourceConfigurator("azurerm_api_management_api_operation_policy", func(r *config.Resource) {
		r.References["api_name"] = config.Reference{
			TerraformName: "azurerm_api_management_api",
		}
		r.References["api_management_name"] = config.Reference{
			TerraformName: "azurerm_api_management",
		}
		r.References["operation_id"] = config.Reference{
			TerraformName: "azurerm_api_management_api_operation",
		}
	})
}

func apiIdCustomDiff(diff *terraform.InstanceDiff, state *terraform.InstanceState, resourceConfig *terraform.ResourceConfig) (*terraform.InstanceDiff, error) { //nolint:gocyclo
	if state == nil || state.Empty() || diff == nil || diff.Empty() || diff.Destroy {
		return diff, nil
	}
	if resourceConfig == nil {
		return nil, errors.New("resource config cannot be nil")
	}
	apiId, ok := resourceConfig.Get("api_id")
	if !ok {
		return diff, nil
	}
	apiIdString, ok := apiId.(string)
	if !ok {
		return diff, errors.New("apiId must be string")
	}
	trimmedApiId := strings.Split(apiIdString, ";")
	stateApiId, ok := state.Attributes["api_id"]
	if !ok {
		return diff, nil
	}
	if trimmedApiId[0] == stateApiId {
		delete(diff.Attributes, "api_id")
	}
	return diff, nil
}
