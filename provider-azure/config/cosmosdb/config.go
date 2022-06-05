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

package cosmosdb

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
	"github.com/upbound/official-providers/provider-azure/config/common"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_cosmosdb_sql_container", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
			"database_name": config.Reference{
				Type: "SQLDatabase",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/000-000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/sqlDatabases/db1/containers/container1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"sqlDatabases", "database_name",
			"containers", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_collection", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
			"database_name": config.Reference{
				Type: "MongoDatabase",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1/collections/collection1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"mongodbDatabases", "database_name",
			"collections", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_keyspace", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"cassandraKeyspaces", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_table", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"cassandra_keyspace_id": config.Reference{
				Type:      "CassandraKeySpace",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1/tables/table1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"cassandraKeyspaces", "cassandra_keyspace_id",
			"tables", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_graph", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
			"database_name": config.Reference{
				Type: "GremlinDatabase",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/gremlinDatabases/db1/graphs/graphs1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"gremlinDatabases", "database_name",
			"graphs", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_function", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"container_id": config.Reference{
				Type:      "SQLContainer",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/userDefinedFunctions/userDefinedFunction1
		r.ExternalName.GetIDFn = func(_ context.Context, name string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			containerID, ok := parameters["container_id"]
			if !ok {
				return "", errors.Errorf(common.ErrFmtNoAttribute, "container_id")
			}
			containerIDStr, ok := containerID.(string)
			if !ok {
				return "", errors.Errorf(common.ErrFmtUnexpectedType, "container_id")
			}
			return fmt.Sprintf("%s/userDefinedFunctions/%s", containerIDStr, name), nil
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_stored_procedure", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
			"database_name": config.Reference{
				Type: "SQLDatabase",
			},
			"container_name": config.Reference{
				Type: "SQLContainer",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1/containers/c1/storedProcedures/sp1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"sqlDatabases", "database_name",
			"containers", "container_name",
			"storedProcedures", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_database", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/gremlinDatabases/db1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"gremlinDatabases", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_database", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/mongodbDatabases/db1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"mongodbDatabases", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_database", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/db1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"sqlDatabases", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_table", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/tables/table1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"tables", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_account", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB", "databaseAccounts", "name")
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_notebook_workspace", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"account_name": config.Reference{
				Type: "Account",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/notebookWorkspaces/notebookWorkspace1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.DocumentDB",
			"databaseAccounts", "account_name",
			"notebookWorkspaces", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_trigger", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = config.References{
			"container_id": config.Reference{
				Type:      "SQLContainer",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID

		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/triggers/trigger1
		r.ExternalName.GetIDFn = func(_ context.Context, name string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			containerID, ok := parameters["container_id"]
			if !ok {
				return "", errors.Errorf(common.ErrFmtNoAttribute, "container_id")
			}
			containerIDStr, ok := containerID.(string)
			if !ok {
				return "", errors.Errorf(common.ErrFmtUnexpectedType, "container_id")
			}
			return fmt.Sprintf("%s/triggers/%s", containerIDStr, name), nil
		}
	})
}
