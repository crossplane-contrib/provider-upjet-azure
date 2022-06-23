package config

import "github.com/upbound/upjet/pkg/config"

var ExternalNameConfigs = map[string]config.ExternalName{

	// base group
	"azurerm_subscription":                   config.TemplatedStringAsIdentifier("alias", "/providers/Microsoft.Subscription/aliases/{{ .externalName }}"),
	"azurerm_resource_provider_registration": config.IdentifierFromProvider,
	"azurerm_resource_group":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .externalName }}"),

	// compute
	"azurerm_disk_encryption_set":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/diskEncryptionSets/{{ .externalName }}"),
	"azurerm_linux_virtual_machine":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachines/{{ .externalName }}"),
	"azurerm_windows_virtual_machine":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachines/{{ .externalName }}"),
	"azurerm_linux_virtual_machine_scale_set":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .externalName }}"),
	"azurerm_windows_virtual_machine_scale_set":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .externalName }}"),
	"azurerm_availability_set":                       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/availabilitySets/{{ .externalName }}"),
	"azurerm_disk_access":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/diskAccesses/{{ .externalName }}"),
	"azurerm_image":                                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/images/{{ .externalName }}"),
	"azurerm_managed_disk":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/disks/{{ .externalName }}"),
	"azurerm_orchestrated_virtual_machine_scale_set": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .externalName }}"),
	"azurerm_proximity_placement_group":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/proximityPlacementGroups/{{ .externalName }}"),
	"azurerm_shared_image_gallery":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .externalName }}"),
	"azurerm_snapshot":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/snapshots/{{ .externalName }}"),
	// Arguments make up the ID.
	"azurerm_marketplace_agreement": config.IdentifierFromProvider,
	"azurerm_dedicated_host":        config.TemplatedStringAsIdentifier("name", "{{ .parameters.dedicated_host_group_id }}/hosts/{ .externalName }}"),

	// cosmosdb
	"azurerm_cosmosdb_sql_database":         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_container":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .externalName }}"),
	"azurerm_cosmosdb_mongo_database":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .externalName }}"),
	"azurerm_cosmosdb_mongo_collection":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .parameters.database_name }}/collections/{{ .externalName }}"),
	"azurerm_cosmosdb_cassandra_keyspace":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/cassandraKeyspaces/{{ .externalName }}"),
	"azurerm_cosmosdb_cassandra_table":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/cassandraKeyspaces/{{ .parameters.cassandra_keyspace_id }}/tables/{{ .externalName }}"),
	"azurerm_cosmosdb_gremlin_database":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .externalName }}"),
	"azurerm_cosmosdb_gremlin_graph":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .parameters.database_name }}/graphs/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_stored_procedure": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .parameters.container_name }}/storedProcedures/{{ .externalName }}"),
	// Contains only container_id as parameter which includes other information like resource group name etc.
	"azurerm_cosmosdb_sql_function":       config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/userDefinedFunctions/{{ .externalName }}"),
	"azurerm_cosmosdb_table":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/tables/{{ .externalName }}"),
	"azurerm_cosmosdb_account":            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .externalName }}"),
	"azurerm_cosmosdb_notebook_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/notebookWorkspaces/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_trigger":        config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/triggers/{{ .externalName }}"),

	// datashare
	"azurerm_data_share_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataShare/accounts/{{ .externalName }}"),
	"azurerm_data_share":         config.TemplatedStringAsIdentifier("name", "{{ .parameters.account_id }}/shares/{{ .externalName }}"),

	// devices
	"azurerm_iothub":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .externalName }}"),
	"azurerm_iothub_consumer_group":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/eventHubEndpoints/{{ .parameters.eventhub_endpoint_name }}/ConsumerGroups/{{ .externalName }}"),
	"azurerm_iothub_dps":                        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .externalName }}"),
	"azurerm_iothub_dps_certificate":            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .parameters.iot_dps_name }}/certificates/{{ .externalName }}"),
	"azurerm_iothub_dps_shared_access_policy":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .parameters.iot_dps_name }}/keys/{{ .externalName }}"),
	"azurerm_iothub_shared_access_policy":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/IotHubKeys/{{ .externalName }}"),
	"azurerm_iothub_endpoint_storage_container": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/Endpoints/{{ .externalName }}"),
	// Configured individually since it doesn't have a naming argument.
	"azurerm_iothub_fallback_route": {},

	// ip
	"azurerm_ip_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/ipGroups/{{ .externalName }}"),

	// eventhub
	"azurerm_eventhub_namespace":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .externalName }}"),
	"azurerm_eventhub":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .externalName }}"),
	"azurerm_eventhub_consumer_group":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/consumergroups/{{ .externalName }}"),
	"azurerm_eventhub_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/authorizationRules/{{ .externalName }}"),

	// apimanagement
	//
	"azurerm_api_management": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .externalName }}"),

	// mssql
	//
	"azurerm_mssql_server": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .providerConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .externalName }}"),
}

func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
			r.Version = "v1beta1"
		}
	}
}
