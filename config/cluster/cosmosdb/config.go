// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cosmosdb

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-azure/apis/cluster/rconfig"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_cosmosdb_account", func(r *config.Resource) {
		delete(r.References, "geo_location.location")
	})
	p.AddResourceConfigurator("azurerm_cosmosdb_sql_container", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			TerraformName: "azurerm_cosmosdb_sql_database",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_collection", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			TerraformName: "azurerm_cosmosdb_mongo_database",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_table", func(r *config.Resource) {
		delete(r.References, "cassandra_keyspace_id")
		r.References["cassandra_keyspace_id"] = config.Reference{
			TerraformName: "azurerm_cosmosdb_cassandra_keyspace",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"analytical_storage_ttl"},
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_graph", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			TerraformName: "azurerm_cosmosdb_gremlin_database",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_function", func(r *config.Resource) {
		r.References["container_id"] = config.Reference{
			TerraformName: "azurerm_cosmosdb_sql_container",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_stored_procedure", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			TerraformName: "azurerm_cosmosdb_sql_database",
		}
		r.References["container_name"] = config.Reference{
			TerraformName: "azurerm_cosmosdb_sql_container",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_trigger", func(r *config.Resource) {
		r.References = config.References{
			"container_id": config.Reference{
				TerraformName: "azurerm_cosmosdb_sql_container",
				Extractor:     rconfig.ExtractResourceIDFuncPath,
			},
		}
	})

	p.AddResourceConfigurator("azurerm_mongo_cluster", func(r *config.Resource) {
		r.References = config.References{
			"source_server_id": config.Reference{
				TerraformName: "azurerm_mongo_cluster",
				Extractor:     rconfig.ExtractResourceIDFuncPath,
			},
			"source_location": config.Reference{
				TerraformName: "azurerm_mongo_cluster",
				Extractor:     rconfig.ExtractResourceLocationFuncPath,
			},
		}
		r.TerraformCustomDiff = mongoClusterCustomDiff
		r.UseAsync = true
	})
}

// MongoCluster can be initialized as a GeoReplica, which has multiple fields that are kept in sync with source MongoCluster
func mongoClusterCustomDiff(diff *terraform.InstanceDiff, state *terraform.InstanceState, resourceConfig *terraform.ResourceConfig) (*terraform.InstanceDiff, error) { //nolint:gocyclo
	if state == nil || state.Empty() || diff == nil || diff.Empty() || diff.Destroy {
		return diff, nil
	}
	if resourceConfig == nil {
		return nil, errors.New("resource config cannot be nil")
	}
	createMode, ok := resourceConfig.Get("create_mode")
	if !ok {
		return diff, nil
	}
	if createMode == "GeoReplica" {
		delete(diff.Attributes, "administrator_username")
		delete(diff.Attributes, "compute_tier")
		delete(diff.Attributes, "high_availability_mode")
		delete(diff.Attributes, "shard_count")
		delete(diff.Attributes, "storage_size_in_gb")
		delete(diff.Attributes, "version")
		delete(diff.Attributes, "preview_features")
	}
	return diff, nil
}
