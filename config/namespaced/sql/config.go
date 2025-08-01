// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sql

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/upbound/provider-azure/apis/namespaced/rconfig"
)

func msSQLConnectionDetails(attr map[string]interface{}) (map[string][]byte, error) {
	dbName, ok := attr["name"].(string)
	if !ok {
		return nil, errors.New("cannot get name attribute")
	}
	username, ok := attr["administrator_login"].(string)
	if !ok {
		return nil, errors.New("cannot get administrator_login attribute")
	}

	endpoint, ok := attr["fully_qualified_domain_name"].(string)
	if !ok {
		return nil, errors.New("cannot get fully_qualified_domain_name attribute")
	}
	conn := map[string][]byte{
		xpv1.ResourceCredentialsSecretUserKey:     []byte(fmt.Sprintf("%s@%s", username, dbName)),
		xpv1.ResourceCredentialsSecretEndpointKey: []byte(endpoint),
	}

	// Note(turkenh): include password only if available, might not be available
	//  depending on the auth type.
	if password, ok := attr["administrator_login_password"].(string); ok {
		conn[xpv1.ResourceCredentialsSecretPasswordKey] = []byte(password)
	}

	return conn, nil
}

// Configure configures sql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_mssql_server", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = msSQLConnectionDetails
	})

	p.AddResourceConfigurator("azurerm_mssql_server_transparent_data_encryption", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["key_vault_key_id"] = config.Reference{
			TerraformName: "azurerm_key_vault_key",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_virtual_network_rule", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance", func(r *config.Resource) {
		r.References["dns_zone_partner_id"] = config.Reference{
			TerraformName: "azurerm_mssql_managed_instance",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_database", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			TerraformName: "azurerm_mssql_managed_instance",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance_active_directory_administrator", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			TerraformName: "azurerm_mssql_managed_instance",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_failover_group", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["partner_server.id"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["databases"] = config.Reference{
			TerraformName: "azurerm_mssql_database",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance_failover_group", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			TerraformName: "azurerm_mssql_managed_instance",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["partner_managed_instance_id"] = config.Reference{
			TerraformName: "azurerm_mssql_managed_instance",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance_vulnerability_assessment", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			TerraformName: "azurerm_mssql_managed_instance",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_database", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"maintenance_configuration_name", "elastic_pool_id"},
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_outbound_firewall_rule", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_server_dns_alias", func(r *config.Resource) {
		r.References["mssql_server_id"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_server_security_alert_policy", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_elasticpool", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"max_size_bytes"},
		}
		r.References["server_name"] = config.Reference{
			TerraformName: "azurerm_mssql_server",
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_database_vulnerability_assessment_rule_baseline", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			TerraformName: "azurerm_mssql_database",
		}
	})
}
