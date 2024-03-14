// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

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
