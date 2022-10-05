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

import (
	"regexp"
	"strings"

	tjconfig "github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/name"
)

var (
	// mappings from group name to # of '_'-separated tokens to skip
	// while constructing Kind names under that group
	kindNameRuleMap = map[string]int{
		"dbformariadb":            2,
		"dbformysql":              2,
		"dbforpostgresql":         2,
		"devspaces":               2,
		"devtestlab":              3,
		"logic":                   2,
		"machinelearningservices": 3,
		"security":                1,
		"timeseriesinsights":      5,
		"azurestackhci":           3,
		"maintenance":             1,
		"cognitiveservices":       2,
		"notificationhubs":        3,
	}

	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/microsoft.botservice/botservices/botservice1
	// /providers/microsoft.subscription/aliases/subscription1
	regexMSServiceName = regexp.MustCompile(`.*/microsoft.([^/]+)/.*`)
)

func extractMicrosoftServiceName(r *tjconfig.Resource) string {
	if r == nil || r.MetaResource == nil || len(r.MetaResource.ImportStatements) < 1 {
		return ""
	}
	groups := regexMSServiceName.FindStringSubmatch(strings.ToLower(r.MetaResource.ImportStatements[0]))
	if len(groups) < 2 {
		return ""
	}
	return groups[1]
}

// default api-group & kind configuration for all resources
func groupKindOverride(r *tjconfig.Resource) {
	g := extractMicrosoftServiceName(r)
	if apiGroup, ok := apiGroupOverrides[g]; ok {
		g = apiGroup
	}
	apiGroup, ok := resourceAPIGroupMap[r.Name]
	if !ok {
		if g == "" {
			return
		}
		apiGroup = g
	}

	r.ShortGroup = apiGroup
	parts := strings.Split(r.Name, "_")
	i, ok := kindNameRuleMap[apiGroup]
	if !ok {
		i = 1 // by default we drop only the first token (azurerm)
		// check if group name is a prefix for the resource name
		for j := 2; j <= len(parts); j++ {
			// do not include azurerm in comparison
			if strings.Join(parts[1:j], "") == apiGroup {
				// if group name is a prefix for resource name,
				// we do not include it in Kind name
				i = j
			}
		}
	}
	if i >= len(parts) {
		i = len(parts) - 1
	}
	r.Kind = name.NewFromSnake(strings.Join(parts[i:], "_")).Camel
}

func init() {
	name.AddAcronym("iothub", "IOTHub")
	name.AddAcronym("iot", "IOT")
	name.AddAcronym("servicebus", "ServiceBus")
	name.AddAcronym("hbase", "HBase")
	name.AddAcronym("rserver", "RServer")
	name.AddAcronym("dsc", "DSC")
	name.AddAcronym("datetime", "DateTime")
	name.AddAcronym("postgresql", "PostgreSQL")
	name.AddAcronym("aaaa", "AAAA")
	name.AddAcronym("caa", "CAA")
	name.AddAcronym("cname", "CNAME")
	name.AddAcronym("mx", "MX")
	name.AddAcronym("ns", "NS")
	name.AddAcronym("ptr", "PTR")
	name.AddAcronym("srv", "SRV")
	name.AddAcronym("txt", "TXT")
	name.AddAcronym("ddos", "DDOS")
	name.AddAcronym("elasticpool", "ElasticPool")
	name.AddAcronym("eventgrid", "EventGrid")
	name.AddAcronym("eventhub", "EventHub")
	name.AddAcronym("keyspace", "KeySpace")
	name.AddAcronym("aad", "AAD")
	name.AddAcronym("hci", "HCI")
	name.AddAcronym("udf", "UDF")
	name.AddAcronym("mssql", "MSSQL")
	name.AddAcronym("aadb2c", "AADB2C")
	name.AddAcronym("cosmosdb", "CosmosDB")
	name.AddAcronym("nodeconfiguration", "NodeConfiguration")
	name.AddAcronym("runbook", "RunBook")
	name.AddAcronym("vmware", "VMware")
	name.AddAcronym("directline", "DirectLine")
	name.AddAcronym("ms", "MS")
	name.AddAcronym("dataset", "DataSet")
	name.AddAcronym("sqlapi", "SQLAPI")
	name.AddAcronym("ssis", "SSIS")
	name.AddAcronym("odata", "OData")
	name.AddAcronym("sftp", "SFTP")
	name.AddAcronym("mariadb", "MariaDB")
	name.AddAcronym("dps", "DPS")
	name.AddAcronym("sas", "SAS")
	name.AddAcronym("powerbi", "PowerBI")
	name.AddAcronym("aws", "AWS")
	name.AddAcronym("filesystem", "FileSystem")
	name.AddAcronym("hpc", "HPC")
	name.AddAcronym("hostname", "HostName")
	name.AddAcronym("datasource", "DataSource")
	name.AddAcronym("healthbot", "HealthBot")
}
