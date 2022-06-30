package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs is a map of external name configurations for the whole
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// apimanagement
	"azurerm_api_management": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .externalName }}"),

	// authorization
	"azurerm_resource_group_policy_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Authorization/policyAssignments/{{ .externalName }}"),

	// base group
	"azurerm_subscription":                   config.TemplatedStringAsIdentifier("alias", "/providers/Microsoft.Subscription/aliases/{{ .externalName }}"),
	"azurerm_resource_provider_registration": config.IdentifierFromProvider,
	"azurerm_resource_group":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .externalName }}"),

	// compute
	"azurerm_disk_encryption_set":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/diskEncryptionSets/{{ .externalName }}"),
	"azurerm_linux_virtual_machine":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachines/{{ .externalName }}"),
	"azurerm_windows_virtual_machine":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachines/{{ .externalName }}"),
	"azurerm_linux_virtual_machine_scale_set":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .externalName }}"),
	"azurerm_windows_virtual_machine_scale_set":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .externalName }}"),
	"azurerm_availability_set":                       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/availabilitySets/{{ .externalName }}"),
	"azurerm_disk_access":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/diskAccesses/{{ .externalName }}"),
	"azurerm_image":                                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/images/{{ .externalName }}"),
	"azurerm_managed_disk":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/disks/{{ .externalName }}"),
	"azurerm_orchestrated_virtual_machine_scale_set": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .externalName }}"),
	"azurerm_proximity_placement_group":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/proximityPlacementGroups/{{ .externalName }}"),
	"azurerm_shared_image_gallery":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .externalName }}"),
	"azurerm_snapshot":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/snapshots/{{ .externalName }}"),
	// Arguments make up the ID.
	"azurerm_marketplace_agreement": config.IdentifierFromProvider,
	"azurerm_dedicated_host":        config.TemplatedStringAsIdentifier("name", "{{ .parameters.dedicated_host_group_id }}/hosts/{ .externalName }}"),

	// cosmosdb
	"azurerm_cosmosdb_sql_database":         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_container":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_role_assignment":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlRoleAssignments/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_role_definition":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlRoleDefinitions/{{ .externalName }}"),
	"azurerm_cosmosdb_mongo_database":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .externalName }}"),
	"azurerm_cosmosdb_mongo_collection":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .parameters.database_name }}/collections/{{ .externalName }}"),
	"azurerm_cosmosdb_cassandra_keyspace":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/cassandraKeyspaces/{{ .externalName }}"),
	"azurerm_cosmosdb_cassandra_table":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/cassandraKeyspaces/{{ .parameters.cassandra_keyspace_id }}/tables/{{ .externalName }}"),
	"azurerm_cosmosdb_gremlin_database":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .externalName }}"),
	"azurerm_cosmosdb_gremlin_graph":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .parameters.database_name }}/graphs/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_stored_procedure": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .parameters.container_name }}/storedProcedures/{{ .externalName }}"),
	// Contains only container_id as parameter which includes other information like resource group name etc.
	"azurerm_cosmosdb_sql_function":         config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/userDefinedFunctions/{{ .externalName }}"),
	"azurerm_cosmosdb_table":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/tables/{{ .externalName }}"),
	"azurerm_cosmosdb_account":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .externalName }}"),
	"azurerm_cosmosdb_notebook_workspace":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/notebookWorkspaces/{{ .externalName }}"),
	"azurerm_cosmosdb_sql_trigger":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/triggers/{{ .externalName }}"),
	"azurerm_cosmosdb_cassandra_cluster":    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/cassandraClusters/{{ .externalName }}"),
	"azurerm_cosmosdb_cassandra_datacenter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cassandra_cluster_id }}/dataCenters/{{ .externalName }}"),

	// datashare
	"azurerm_data_share_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataShare/accounts/{{ .externalName }}"),
	"azurerm_data_share":         config.TemplatedStringAsIdentifier("name", "{{ .parameters.account_id }}/shares/{{ .externalName }}"),

	// devices
	"azurerm_iothub":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .externalName }}"),
	"azurerm_iothub_consumer_group":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/eventHubEndpoints/{{ .parameters.eventhub_endpoint_name }}/ConsumerGroups/{{ .externalName }}"),
	"azurerm_iothub_dps":                        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .externalName }}"),
	"azurerm_iothub_dps_certificate":            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .parameters.iot_dps_name }}/certificates/{{ .externalName }}"),
	"azurerm_iothub_dps_shared_access_policy":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .parameters.iot_dps_name }}/keys/{{ .externalName }}"),
	"azurerm_iothub_shared_access_policy":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/IotHubKeys/{{ .externalName }}"),
	"azurerm_iothub_endpoint_storage_container": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/Endpoints/{{ .externalName }}"),
	"azurerm_iothub_fallback_route":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/FallbackRoute/default"),
	"azurerm_iothub_endpoint_eventhub":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/Endpoints/{{ .externalName }}"),
	"azurerm_iothub_endpoint_servicebus_queue":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/Endpoints/{{ .externalName }}"),
	"azurerm_iothub_endpoint_servicebus_topic":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/Endpoints/{{ .externalName }}"),
	"azurerm_iothub_route":                      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/IotHubs/{{ .parameters.iothub_name }}/Routes/{{ .externalName }}"),
	// No name field.
	"azurerm_iothub_enrichment": config.IdentifierFromProvider,

	// eventhub
	"azurerm_eventhub_namespace":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .externalName }}"),
	"azurerm_eventhub":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .externalName }}"),
	"azurerm_eventhub_consumer_group":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/consumergroups/{{ .externalName }}"),
	"azurerm_eventhub_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Eventhub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/authorizationRules/{{ .externalName }}"),

	// keyvault
	"azurerm_key_vault":                                              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.KeyVault/vaults/{{ .externalName }}"),
	"azurerm_key_vault_secret":                                       config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.key_vault_id }}.vault.azure.net/secrets/{{ .externalName }}/{{ .parameters.version }}"),
	"azurerm_key_vault_key":                                          config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.key_vault_id }}.vault.azure.net/keys/{{ .externalName }}/{{ .parameters.version }}"),
	"azurerm_key_vault_certificate":                                  config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.key_vault_id }}.vault.azure.net/certificates/{{ .externalName }}/{{ .parameters.version }}"),
	"azurerm_key_vault_certificate_issuer":                           config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.key_vault_id }}.vault.azure.net/certificates/issuers/{{ .externalName }}"),
	"azurerm_key_vault_managed_storage_account":                      config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.key_vault_id }}.vault.azure.net/storage/{{ .externalName }}"),
	"azurerm_key_vault_managed_storage_account_sas_token_definition": config.TemplatedStringAsIdentifier("name", "{{ .parameters.managed_storage_account_id }}/sas/{{ .externalName }}"),
	"azurerm_key_vault_managed_hardware_security_module":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.KeyVault/managedHSMs//{{ .externalName }}"),
	"azurerm_key_vault_access_policy":                                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.KeyVault/managedHSMs//{{ .externalName }}"),

	// containerservice
	"azurerm_kubernetes_cluster":           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerService/managedClusters/{{ .externalName }}"),
	"azurerm_kubernetes_cluster_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .parameters.kubernetes_cluster_id }}/agentPools/{{ .externalName }}"),

	// operationalinsights
	"azurerm_log_analytics_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/workspaces/{{ .externalName }}"),

	// logic
	"azurerm_integration_service_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationServiceEnvironments/{{ .externalName }}"),

	// management
	"azurerm_management_group": config.TemplatedStringAsIdentifier("name", "/providers/Microsoft.Management/managementGroups/{{ .externalName }}"),

	// mariadb
	"azurerm_mariadb_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .externalName }}"),
	"azurerm_mariadb_database":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/databases/{{ .externalName }}"),
	"azurerm_mariadb_firewall_rule":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/firewallRules/{{ .externalName }}"),
	"azurerm_mariadb_virtual_network_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .externalName }}"),
	"azurerm_mariadb_configuration":        config.IdentifierFromProvider,

	// monitor
	"azurerm_monitor_metric_alert": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/metricAlerts/{{ .externalName }}"),

	// network
	"azurerm_virtual_network":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .externalName }}"),
	"azurerm_ip_group":                                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/ipGroups/{{ .externalName }}"),
	"azurerm_subnet":                                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .parameters.virtual_network_name }}/subnets/{{ .externalName }}"),
	"azurerm_subnet_nat_gateway_association":            config.IdentifierFromProvider,
	"azurerm_subnet_network_security_group_association": config.IdentifierFromProvider,
	"azurerm_subnet_service_endpoint_storage_policy":    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/serviceEndpointPolicies/{{ .externalName }}"),
	"azurerm_subnet_route_table_association":            config.IdentifierFromProvider,
	"azurerm_network_interface":                         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkInterfaces/{{ .externalName }}"),
	"azurerm_lb":                                        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/loadBalancers/{{ .externalName }}"),
	"azurerm_lb_backend_address_pool":                   config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/backendAddressPools/{{ .externalName }}"),
	"azurerm_lb_nat_pool":                               config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/inboundNatPools/{{ .externalName }}"),
	"azurerm_lb_nat_rule":                               config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/inboundNatRules/{{ .externalName }}"),
	"azurerm_lb_backend_address_pool_address":           config.TemplatedStringAsIdentifier("name", "{{ .parameters.backend_address_pool_id }}/addresses/{{ .externalName }}"),
	"azurerm_lb_outbound_rule":                          config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/outboundRules/{{ .externalName }}"),
	"azurerm_lb_probe":                                  config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/probes/{{ .externalName }}"),
	"azurerm_lb_rule":                                   config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/loadBalancingRules/{{ .externalName }}"),
	"azurerm_local_network_gateway":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/localNetworkGateways/{{ .externalName }}"),
	"azurerm_nat_gateway":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/natGateways/{{ .externalName }}"),
	"azurerm_network_watcher":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkWatchers/{{ .externalName }}"),
	"azurerm_network_watcher_flow_log":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkWatchers/{{ .parameters.network_watcher_name }}/flowLogs/{{ .externalName }}"),
	"azurerm_network_connection_monitor":                config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_watcher_id }}/backendAddressPools/{{ .connectionMonitors }}"),
	"azurerm_network_ddos_protection_plan":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/ddosProtectionPlans/{{ .externalName }}"),
	"azurerm_application_security_group":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/applicationSecurityGroups/{{ .externalName }}"),
	"azurerm_network_security_group":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkSecurityGroups/{{ .externalName }}"),
	"azurerm_network_security_rule":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkSecurityGroups/{{ .parameters.network_security_group_name }}/securityRules/{{ .externalName }}"),
	"azurerm_virtual_network_gateway":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworkGateways/{{ .externalName }}"),
	"azurerm_virtual_network_peering":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .parameters.virtual_network_name }}/virtualNetworkPeerings/{{ .externalName }}"),
	"azurerm_virtual_network_gateway_connection":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/connections/{{ .externalName }}"),
	"azurerm_virtual_wan":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualWans/{{ .externalName }}"),
	"azurerm_virtual_hub":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualHubs/{{ .externalName }}"),
	"azurerm_public_ip":                                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/publicIPAddresses/{{ .externalName }}"),
	"azurerm_public_ip_prefix":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/publicIPPrefixes/{{ .externalName }}"),
	"azurerm_network_profile":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkProfiles/{{ .externalName }}"),
	"azurerm_private_dns_zone":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .externalName }}"),
	"azurerm_private_dns_a_record":                      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/A/{{ .externalName }}"),
	"azurerm_private_dns_aaaa_record":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/AAAA/{{ .externalName }}"),
	"azurerm_private_dns_cname_record":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/CNAME/{{ .externalName }}"),
	"azurerm_private_dns_mx_record":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/MX/{{ .externalName }}"),
	"azurerm_private_dns_ptr_record":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/PTR/{{ .externalName }}"),
	"azurerm_private_dns_srv_record":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/SRV/{{ .externalName }}"),
	"azurerm_private_dns_txt_record":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/TXT/{{ .externalName }}"),
	"azurerm_private_dns_zone_virtual_network_link":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.private_dns_zone_name }}/virtualNetworkLinks/{ .externalName }}"),
	"azurerm_private_link_service":                      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateLinkServices/{{ .externalName }}"),
	"azurerm_private_endpoint":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateEndpoints/{{ .externalName }}"),
	"azurerm_network_packet_capture":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkWatchers/{{ .parameters.network_watcher_name }}/packetCaptures/{{ .externalName }}"),
	"azurerm_vpn_server_configuration":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/vpnServerConfigurations/{{ .externalName }}"),
	"azurerm_point_to_site_vpn_gateway":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/p2sVpnGateways/{{ .externalName }}"),
	// Made up of arguments.
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
	"azurerm_network_interface_security_group_association": config.IdentifierFromProvider,
	// Made up of arguments.
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1
	"azurerm_network_interface_nat_rule_association": config.IdentifierFromProvider,
	// Made up of arguments.
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1
	"azurerm_network_interface_backend_address_pool_association": config.IdentifierFromProvider,
	// Made up of arguments.
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationSecurityGroups/securityGroup1
	"azurerm_network_interface_application_security_group_association": config.IdentifierFromProvider,
	// Made up of arguments.
	"azurerm_nat_gateway_public_ip_association": config.IdentifierFromProvider,
	// Made up of arguments.
	"azurerm_nat_gateway_public_ip_prefix_association": config.IdentifierFromProvider,

	// notification
	"azurerm_notification_hub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NotificationHubs/namespaces/{{ .parameters.namespace_name }}/notificationHubs/{{ .externalName }}"),

	// postgresql
	"azurerm_postgresql_server":                         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .externalName }}"),
	"azurerm_postgresql_flexible_server_configuration":  config.IdentifierFromProvider,
	"azurerm_postgresql_database":                       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/databases/{{ .externalName }}"),
	"azurerm_postgresql_active_directory_administrator": config.TemplatedStringAsIdentifier("login", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/administrators/{{ .externalName }}"),
	"azurerm_postgresql_flexible_server_database":       config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/databases/{{ .externalName }}"),
	"azurerm_postgresql_flexible_server_firewall_rule":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/firewallRules/{{ .externalName }}"),
	"azurerm_postgresql_firewall_rule":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/firewallRules/{{ .externalName }}"),
	"azurerm_postgresql_flexible_server":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{ .externalName }}"),
	"azurerm_postgresql_virtual_network_rule":           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .externalName }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
	"azurerm_postgresql_server_key": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/configurations/backslash_quote
	"azurerm_postgresql_configuration": config.IdentifierFromProvider,

	// redis
	"azurerm_redis_cache":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/Redis/{{ .externalName }}"),
	"azurerm_redis_firewall_rule":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/Redis/{{ .parameters.redis_cache_name }}/firewallRules/{{ .externalName }}"),
	"azurerm_redis_enterprise_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redisEnterprise/{{ .externalName }}"),
	// Done individually.
	"azurerm_redis_enterprise_database": {},
	// Unknown format.
	"azurerm_redis_linked_server": config.IdentifierFromProvider,

	// resource
	"azurerm_resource_group_template_deployment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Resources/deployments/{{ .externalName }}"),

	// security
	"azurerm_advanced_threat_protection": config.TemplatedStringAsIdentifier("", "{{ .parameters.target_resource_id }}/providers/Microsoft.Security/advancedThreatProtectionSettings/default"),
	"azurerm_iot_security_device_group":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.target_resource_id }}/providers/Microsoft.Security/deviceSecurityGroups/{{ .externalName }}"),
	"azurerm_iot_security_solution":      config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Security/IoTSecuritySolutions/{{ .externalName }}"),

	// sql
	"azurerm_mssql_server":                             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .externalName }}"),
	"azurerm_mssql_server_transparent_data_encryption": config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/encryptionProtector/current"),

	// storage
	"azurerm_storage_account":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Storage/storageAccounts/{{ .externalName }}"),
	"azurerm_storage_blob":      config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.storage_account_name }}.blob.core.windows.net/{{ .parameters.storage_container_name }}/{{ .externalName }}"),
	"azurerm_storage_container": config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.storage_account_name }}.blob.core.windows.net/{{ .externalName }}"),

	// storagesync
	"azurerm_storage_sync": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .terraformProviderConfig.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StorageSync/storageSyncServices/{{ .externalName }}"),
}

// ExternalNameConfigurations adds all external name configurations from
// the main table ExternalNameConfigs.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
			r.Version = "v1beta1"
		}
	}
}

// ResourcesWithExternalNameConfig returns a list of resources that have external
// name config defined in the external name table.
func ResourcesWithExternalNameConfig() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for r := range ExternalNameConfigs {
		l[i] = r + "$"
		i++
	}
	return l
}
