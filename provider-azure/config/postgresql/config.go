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

package postgresql

import (
	"context"
	"fmt"
	"strconv"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
	"github.com/upbound/official-providers/provider-azure/config/common"
)

const (
	errFmtNoAttribute    = `"attribute not found: %s`
	errFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`

	postgresqlServerPort = 5432
)

// Configure configures postgresql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_postgresql_server", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL", "servers", "name")
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
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"server_id": config.Reference{
				Type:      "FlexibleServer",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("azurerm_postgresql_database", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/databases/database1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"servers", "server_name",
			"databases", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_active_directory_administrator", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			// TODO(aru): this may have to be a reference to the server's resource group
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.ExternalName{
			SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
				base["login"] = name
			},
			OmittedFields:     []string{"login"},
			GetExternalNameFn: common.GetNameFromFullyQualifiedID,
			// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/myserver/administrators/activeDirectory
			GetIDFn: common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
				"servers", "server_name",
				"administrators", "login",
			),
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_database", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"server_id": config.Reference{
				Type:      "FlexibleServer",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/databases/database1
		r.ExternalName.GetIDFn = func(ctx context.Context, name string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			serverID, ok := parameters["server_id"]
			if !ok {
				return "", errors.Errorf(errFmtNoAttribute, "server_id")
			}
			serverIDStr, ok := serverID.(string)
			if !ok {
				return "", errors.Errorf(errFmtUnexpectedType, "server_id")
			}
			return fmt.Sprintf("%s/databases/%s", serverIDStr, name), nil
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_firewall_rule", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/firewallRules/rule1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"servers", "server_name",
			"firewallRules", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server_firewall_rule", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"server_id": config.Reference{
				Type:      "FlexibleServer",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/flexibleServers/flexibleServer1/firewallRules/firewallRule1
		r.ExternalName.GetIDFn = func(ctx context.Context, name string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			serverID, ok := parameters["server_id"]
			if !ok {
				return "", errors.Errorf(errFmtNoAttribute, "server_id")
			}
			serverIDStr, ok := serverID.(string)
			if !ok {
				return "", errors.Errorf(errFmtUnexpectedType, "server_id")
			}
			return fmt.Sprintf("%s/firewallRules/%s", serverIDStr, name), nil
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_flexible_server", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"delegated_subnet_id": config.Reference{
				Type:      rconfig.SubnetReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/flexibleServers/server1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"flexibleServers", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_virtual_network_rule", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"server_name": config.Reference{
				Type: "Server",
			},
			"subnet_id": config.Reference{
				Type:      rconfig.SubnetReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/myserver/virtualNetworkRules/vnetrulename
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforPostgreSQL",
			"servers", "server_name",
			"virtualNetworkRules", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_postgresql_server_key", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"server_id": config.Reference{
				Type:      "Server",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"key_vault_key_id": config.Reference{
				Type:      rconfig.VaultKeyReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = config.NopSetIdentifierArgument
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			return fmt.Sprintf("%s/keys/%s", parameters["server_id"], externalName), nil
		}
	})

	p.AddResourceConfigurator("azurerm_postgresql_configuration", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/configurations/backslash_quote
		r.ExternalName = config.IdentifierFromProvider
	})
}
