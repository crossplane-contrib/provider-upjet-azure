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

package sql

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/upbound/provider-azure/apis/rconfig"
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
		r.UseAsync = false
		r.References["server_id"] = config.Reference{
			Type:      "MSSQLServer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["key_vault_key_id"] = config.Reference{
			Type:      rconfig.VaultKeyReferencePath,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_virtual_network_rule", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			Type:      "MSSQLServer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance", func(r *config.Resource) {
		r.References["dns_zone_partner_id"] = config.Reference{
			Type:      "MSSQLManagedInstance",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_database", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			Type:      "MSSQLManagedInstance",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance_active_directory_administrator", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			Type:      "MSSQLManagedInstance",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_failover_group", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			Type:      "MSSQLServer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["partner_server.id"] = config.Reference{
			Type:      "MSSQLServer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["databases"] = config.Reference{
			Type:      "MSSQLDatabase",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance_failover_group", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			Type:      "MSSQLManagedInstance",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["partner_managed_instance_id"] = config.Reference{
			Type:      "MSSQLManagedInstance",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_managed_instance_vulnerability_assessment", func(r *config.Resource) {
		r.References["managed_instance_id"] = config.Reference{
			Type:      "MSSQLManagedInstance",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_database", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			Type:      "MSSQLServer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"maintenance_configuration_name", "elastic_pool_id"},
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_outbound_firewall_rule", func(r *config.Resource) {
		r.References["server_id"] = config.Reference{
			Type:      "MSSQLServer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_server_dns_alias", func(r *config.Resource) {
		r.References["mssql_server_id"] = config.Reference{
			Type:      "MSSQLServer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_server_security_alert_policy", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			Type: "MSSQLServer",
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_elasticpool", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"max_size_bytes"},
		}
		r.References["server_name"] = config.Reference{
			Type: "MSSQLServer",
		}
	})

	p.AddResourceConfigurator("azurerm_mssql_database_vulnerability_assessment_rule_baseline", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			Type: "MSSQLDatabase",
		}
	})
}
