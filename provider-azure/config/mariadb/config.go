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

package mariadb

import (
	"fmt"
	"strconv"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/config/common"
)

const (
	mariadbServerPort = 3306
)

// Configure configures mariadb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_mariadb_server", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"ssl_enforcement", "storage_profile"},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMariaDB/servers/server1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforMariaDB", "servers", "name")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return map[string][]byte{
				xpv1.ResourceCredentialsSecretUserKey:     []byte(fmt.Sprintf("%s@%s", attr["administrator_login"], attr["name"])),
				xpv1.ResourceCredentialsSecretPasswordKey: []byte(attr["administrator_login_password"].(string)),
				xpv1.ResourceCredentialsSecretEndpointKey: []byte(attr["fqdn"].(string)),
				xpv1.ResourceCredentialsSecretPortKey:     []byte(strconv.Itoa(mariadbServerPort)),
			}, nil
		}
	})

	p.AddResourceConfigurator("azurerm_mariadb_database", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.References = config.References{
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMariaDB/servers/server1/databases/database1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforMariaDB",
			"servers", "server_name",
			"databases", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_mariadb_firewall_rule", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.References = config.References{
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMariaDB/servers/server1/firewallRules/rule1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforMariaDB",
			"servers", "server_name",
			"firewallRules", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_mariadb_virtual_network_rule", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.References = config.References{
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforMariaDB/servers/myserver/virtualNetworkRules/vnetrulename
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DBforMariaDB",
			"servers", "server_name",
			"virtualNetworkRules", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_mariadb_configuration", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.References = config.References{
			"server_name": config.Reference{
				Type: "Server",
			},
		}
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMariaDB/servers/server1/configurations/interactive_timeout
		// which needs to be a valid MariaDB configuration name
		// See https://mariadb.com/kb/en/server-system-variables/
		r.ExternalName = config.IdentifierFromProvider
	})
}
