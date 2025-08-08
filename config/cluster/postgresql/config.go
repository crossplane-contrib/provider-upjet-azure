// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package postgresql

import (
	"fmt"
	"strconv"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/types/comments"

	"github.com/upbound/provider-azure/apis/cluster/rconfig"
	"github.com/upbound/provider-azure/config/cluster/common"
)

const (
	postgresqlServerPort = 5432
)

// Configure configures postgresql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_postgresql_server", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return map[string][]byte{
				xpv1.ResourceCredentialsSecretUserKey:     []byte(fmt.Sprintf("%s@%s", attr["administrator_login"], attr["name"])),
				xpv1.ResourceCredentialsSecretPasswordKey: []byte(attr["administrator_login_password"].(string)),
				xpv1.ResourceCredentialsSecretEndpointKey: []byte(attr["fqdn"].(string)),
				xpv1.ResourceCredentialsSecretPortKey:     []byte(strconv.Itoa(postgresqlServerPort)),
			}, nil
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_configuration", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_postgresql_flexible_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_database", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_postgresql_server",
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_active_directory_administrator", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_postgresql_server",
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_active_directory_administrator", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_postgresql_flexible_server",
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_database", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_postgresql_flexible_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_firewall_rule", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_postgresql_server",
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_firewall_rule", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_postgresql_flexible_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{
				xpv1.ResourceCredentialsSecretEndpointKey: []byte(attr["fqdn"].(string)),
				xpv1.ResourceCredentialsSecretPortKey:     []byte(strconv.Itoa(postgresqlServerPort)),
			}
			// administrator_login is optional and may not be available depending on the auth type.
			if _, ok := attr["administrator_login"].(string); ok {
				conn[xpv1.ResourceCredentialsSecretUserKey] = []byte(fmt.Sprintf("%s@%s", attr["administrator_login"].(string), attr["name"].(string)))
			}
			// administrator_password is optional and may not be available depending on the auth type.
			if _, ok := attr["administrator_password"].(string); ok {
				conn[xpv1.ResourceCredentialsSecretPasswordKey] = []byte(attr["administrator_password"].(string))
			}
			return conn, nil
		}
		desc, _ := comments.New("If true, the password will be auto-generated and"+
			" stored in the Secret referenced by the administratorPasswordSecretRef field.",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["auto_generate_password"] = &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Description: desc.String(),
		}
		r.InitializerFns = append(r.InitializerFns,
			common.PasswordGenerator(
				"spec.forProvider.administratorPasswordSecretRef",
				"spec.forProvider.autoGeneratePassword",
			))
		r.TerraformResource.Schema["administrator_password"].Description = "Password for the " +
			"master DB user. If you set autoGeneratePassword to true, the Secret" +
			" referenced here will be created or updated with generated password" +
			" if it does not already contain one."
	})

	p.AddResourceConfigurator("azurerm_postgresql_virtual_network_rule", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_postgresql_server",
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_server_key", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_postgresql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["key_vault_key_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_key",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_configuration", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_postgresql_server",
		}
	})
}
