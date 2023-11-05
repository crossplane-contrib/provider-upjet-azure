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

package dbformysql

import (
	"fmt"
	"strconv"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry"
)

const (
	mysqlServerPort = 3306
)

// Configure configures dbformysql group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_mysql_flexible_server", func(r *config.Resource) {
		r.References["private_dns_zone_id"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/network/v1beta1.PrivateDNSZone",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_mysql_flexible_database", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			Type: "FlexibleServer",
		}
	})

	p.AddResourceConfigurator("azurerm_mysql_flexible_server_configuration", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			Type: "FlexibleServer",
		}
	})

	p.AddResourceConfigurator("azurerm_mysql_flexible_server_firewall_rule", func(r *config.Resource) {
		r.References["server_name"] = config.Reference{
			Type: "FlexibleServer",
		}
	})
	p.AddResourceConfigurator("azurerm_mysql_server", func(r *config.Resource) {
		r.MetaResource.ExternalName = registry.RandRFC1123Subdomain
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return map[string][]byte{
				xpv1.ResourceCredentialsSecretUserKey:     []byte(fmt.Sprintf("%s@%s", attr["administrator_login"], attr["name"])),
				xpv1.ResourceCredentialsSecretPasswordKey: []byte(attr["administrator_login_password"].(string)),
				xpv1.ResourceCredentialsSecretEndpointKey: []byte(attr["fqdn"].(string)),
				xpv1.ResourceCredentialsSecretPortKey:     []byte(strconv.Itoa(mysqlServerPort)),
			}, nil
		}
	})
}
