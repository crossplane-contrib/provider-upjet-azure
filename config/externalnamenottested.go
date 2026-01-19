// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{
	// aad
	//
	// Domain Services can be imported using the resource ID, together with the
	// Replica Set ID that you wish to designate as the initial replica set
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/initialReplicaSetId/00000000-0000-0000-0000-000000000000
	"azurerm_active_directory_domain_service": config.IdentifierFromProvider,
	// Domain Service Replica Sets can be imported using the resource ID of the
	// parent Domain Service and the Replica Set ID
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/replicaSets/00000000-0000-0000-0000-000000000000
	"azurerm_active_directory_domain_service_replica_set": config.IdentifierFromProvider,

	// aadb2c
	//
	// AAD B2C Directories can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureActiveDirectory/b2cDirectories/directory-name
	"azurerm_aadb2c_directory": config.IdentifierFromProvider,

	// aadiam
	//
	// Monitor Azure Active Directory Diagnostic Settings can be imported using the resource id
	// /providers/Microsoft.AADIAM/diagnosticSettings/setting1
	"azurerm_monitor_aad_diagnostic_setting": config.TemplatedStringAsIdentifier("name", "/providers/Microsoft.AADIAM/diagnosticSettings/{{ .external_name }}"),

	// api
	//
	// API Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.Web/connections/example-connection
	"azurerm_api_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/connections/{{ .external_name }}"),

	// apimanagement
	//
	// API Management Azure AD B2C Identity Providers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/service1/identityProviders/AadB2C
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_aadb2c": config.IdentifierFromProvider,

	// appconfiguration
	// There are two different syntaxes:
	// App Configuration Features can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/appConfFeature1/Label/label1
	// If you wish to import a key with an empty label then sustitute the label's name with %00, like this
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/appConfFeature1/Label/%00
	"azurerm_app_configuration_feature": config.IdentifierFromProvider,
	// There are two different syntaxes:
	// App Configuration Keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/label1
	// If you wish to import a key with an empty label then sustitute the label's name with %00, like this
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/%00
	"azurerm_app_configuration_key": config.IdentifierFromProvider,

	// kusto
	//
	// Customer Managed Keys for a Kusto Cluster can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
	"azurerm_kusto_cluster_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.cluster_id }}"),
	// Kusto Scripts can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/Scripts/script1
	"azurerm_kusto_script": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/Databases/{{ .parameters.database_id }}/Scripts/{{ .external_name }}"),

	// appplatform
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/routeConfigs/routeConfig1
	"azurerm_spring_cloud_gateway_route_config": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_gateway_id }}/routeConfigs/{{ .external_name }}"),

	// avs
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/clusters/cluster1
	"azurerm_vmware_cluster": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vmware_cloud_id }}/clusters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/authorizations/authorization1
	"azurerm_vmware_express_route_authorization": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vmware_cloud_id }}/clusters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/PrivateClouds/privateCloud1
	"azurerm_vmware_private_cloud": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AVS/PrivateClouds/{{ .external_name }}"),

	// blueprint
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint
	"azurerm_blueprint_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Blueprint/blueprintAssignments/{{ .external_name }}"),

	// cdn
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customDomains/domain1
	"azurerm_cdn_endpoint_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_endpoint_id }}/customDomains/{{ .external_name }}"),

	// cognitiveservices
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
	"azurerm_cognitive_account_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .paramteres.cognitive_account_id }}"),

	// authorization
	//
	// Policy Remediations can be imported using the resource id
	//  /providers/Microsoft.Management/managementGroups/my-mgmt-group-id/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_management_group_policy_remediation": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/Microsoft.PolicyInsights/remediations/{{ .external_name }}"),
	// /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName
	"azurerm_management_group_policy_set_definition": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Authorization/policySetDefinitions/{{ .external_name }}"),

	// automation
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/certificates/certificate1
	"azurerm_automation_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/certificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_service_principal": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),

	// botservice
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/DirectLineSpeechChannel
	"azurerm_bot_channel_direct_line_speech": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/DirectLineSpeechChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/EmailChannel
	"azurerm_bot_channel_email": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/EmailChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/FacebookChannel
	"azurerm_bot_channel_facebook": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/FacebookChannel"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.BotService/botServices/botService1
	"azurerm_bot_service_azure_bot": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .external_name }}"),

	// compute
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/1.2.3
	"azurerm_shared_image_version": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .parameters.gallery_name }}/images/{{ .parameters.image_name }}/versions/{{ .external_name }}"),

	// batch
	//
	// Batch Account can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Batch/batchAccounts/account1
	"azurerm_batch_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .externalName }}"),
	// Batch Application can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Batch/batchAccounts/exampleba/applications/example-batch-application
	"azurerm_batch_application": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/applications/{{ .externalName }}"),
	// Batch Certificate can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Batch/batchAccounts/batch1/certificates/certificate1
	"azurerm_batch_certificate": config.TemplatedStringAsIdentifier("name", "/subscBriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/certificates/{{ .externalName }}"),
	// Batch Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Batch/batchAccounts/myBatchAccount1/pools/myBatchPool1
	"azurerm_batch_pool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/pools/{{ .externalName }}"),
	// Batch job can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Batch/batchAccounts/account1/pools/pool1/jobs/job1
	"azurerm_batch_job": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/pools/{{ .parameters.batch_pool_name }}/jobs/{{ .externalName }}"),

	// iot
	//
	// Azure IoT Time Series Insights EventHub Event Source can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/eventSources/example
	"azurerm_time_series_insights_event_source_eventhub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.TimeSeriesInsights/environments/{{ .parameters.environment_name }}/eventSources/{{ .externalName }}"),

	// logic
	//
	// Logic App Integration Account Agreements can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/agreements/agreement1
	"azurerm_logic_app_integration_account_agreement": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/agreements/{{ .external_name }}"),
	// Logic App Integration Account Assemblies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/assemblies/assembly1
	"azurerm_logic_app_integration_account_assembly": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/assemblies/{{ .external_name }}"),
	// Logic App Integration Account Maps can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/maps/map1
	"azurerm_logic_app_integration_account_map": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/maps/{{ .external_name }}"),
	// Logic Apps can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/logicapp1
	"azurerm_logic_app_standard": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),

	// operationalinsights
	//
	// Log Analytics Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
	"azurerm_log_analytics_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/clusters/{{ .external_name }}"),
	// Log Analytics Cluster Customer Managed Keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
	"azurerm_log_analytics_cluster_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.log_analytics_cluster_id }}"),

	// portal
	//
	// Portal Tenant Configurations can be imported using the resource id
	// /providers/Microsoft.Portal/tenantConfigurations/default
	"azurerm_portal_tenant_configuration": config.IdentifierFromProvider,

	// synapse
	//
	// Synapse SQL Pool Vulnerability Assessment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default
	"azurerm_synapse_sql_pool_vulnerability_assessment": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_id }}/vulnerabilityAssessments/default"),
	// Synapse SQL Pool Vulnerability Assessment Rule Baselines can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default/rules/rule1/baselines/baseline1
	"azurerm_synapse_sql_pool_vulnerability_assessment_baseline": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_vulnerability_assessment_id }}/rules/{{ .parameters.rule_name }}/baselines/{{ .external_name }}"),
	// Synapse Workspace Keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/keys/key1
	"azurerm_synapse_workspace_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/keys/{{ .external_name }}"),

	// digitaltwins
	//
	// Digital Twins Eventgrid Endpoints can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
	"azurerm_digital_twins_endpoint_eventgrid": config.TemplatedStringAsIdentifier("name", "{{ .parameters.digital_twins_id }}/endpoints/{{ .external_name }}"),
	// Digital Twins Eventhub Endpoints can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
	"azurerm_digital_twins_endpoint_eventhub": config.IdentifierFromProvider,
	// config.TemplatedStringAsIdentifier("name", "{{ .parameters.digital_twins_id }}/endpoints/{{ .external_name }}"),
	// Digital Twins Service Bus Endpoints can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
	"azurerm_digital_twins_endpoint_servicebus": config.TemplatedStringAsIdentifier("name", "{{ .parameters.digital_twins_id }}/endpoints/{{ .external_name }}"),

	// sqlvirtualmachine
	//
	// Microsoft SQL Virtual Machines can be imported using the resource id
	/// subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/example1
	"azurerm_mssql_virtual_machine": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{{ .external_name }}"),

	// static
	// Static Site Custom Domains can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1/customDomains/name.contoso.com
	"azurerm_static_site_custom_domain": config.IdentifierFromProvider,

	// storage
	//
	// Customer Managed Keys for a Storage Account can be imported using the resource id of the Storage Account
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
	"azurerm_storage_account_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.storage_account_id }}"),

	// securityinsights
	//
	// AWS CloudTrail Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_aws_cloud_trail": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Azure Active Directory Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_azure_active_directory": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Azure Advanced Threat Protection Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_azure_advanced_threat_protection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Azure Security Center Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_azure_security_center": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Microsoft Cloud App Security Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_microsoft_cloud_app_security": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Office 365 Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_office_365": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Threat Intelligence Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_threat_intelligence": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),

	// sentinel
	//
	// AWS S3 Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_aws_s3": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Sentinel Watchlist Items can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1/watchlistItems/item1
	"azurerm_sentinel_watchlist_item": config.TemplatedStringAsIdentifier("", "{{ .parameters.watchlist_id }}/watchlistItems/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_nrt": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),
	// Sentinel Scheduled Alert Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_scheduled": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),

	// solutions
	//
	// Managed Application can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applications/app1
	"azurerm_managed_application": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Solutions/applications/{{ .external_name }}"),

	// resources
	//
	// Management Group Template Deployments can be imported using the resource id
	// /providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/deploy1
	"azurerm_management_group_template_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Resources/deployments/{{ .external_name }}"),
	// Tenant Template Deployments can be imported using the resource id
	// /providers/Microsoft.Resources/deployments/deploy1
	"azurerm_tenant_template_deployment": config.TemplatedStringAsIdentifier("name", "/providers/Microsoft.Resources/deployments/{{ .external_name }}"),

	// hardwaresecuritymodules
	//
	// Dedicated Hardware Security Module can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1
	"azurerm_dedicated_hardware_security_module": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{{ .external_name }}"),

	// app
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/site1/publicCertificates/publicCertificate1
	"azurerm_app_service_public_certificate": config.TemplatedStringAsIdentifier("certificate_name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .parameters.app_service_name }}/publicCertificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/staging/hostNameBindings/mywebsite.com
	"azurerm_app_service_slot_custom_hostname_binding": config.TemplatedStringAsIdentifier("hostname", "{{ .parameters.app_service_slot_id }}/hostNameBindings/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_app_service_source_control_slot": config.IdentifierFromProvider,

	// vpn
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/vpnGateways/vpnGateway1/natRules/natRule1
	"azurerm_vpn_gateway_nat_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vpn_gateway_id }}/natRules/{{ .external_name }}"),

	// web
	//
	// DEPRECATED
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1
	"azurerm_app_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),
	// DEPRECATED
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_app_service_active_slot": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com
	"azurerm_app_service_custom_hostname_binding": config.TemplatedStringAsIdentifier("hostname", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .parameters.app_service_name }}/hostNameBindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Web/hostingEnvironments/myAppServiceEnv
	"azurerm_app_service_environment_v3": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/hostingEnvironments/{{ .external_name }}"),
	// DEPRECATED
	// /subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/exampleResourceGroup1/providers/Microsoft.Web/sites/exampleAppService1/hybridConnectionNamespaces/exampleRN1/relays/exampleRHC1
	"azurerm_app_service_hybrid_connection": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Web/certificates/customhost.contoso.com
	"azurerm_app_service_managed_certificate": config.IdentifierFromProvider,
	// DEPRECATED
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/website1/slots/instance1
	"azurerm_app_service_slot": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .parameters.app_service_name }}/slots/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/staging/config/virtualNetwork
	"azurerm_app_service_slot_virtual_network_swift_connection": config.TemplatedStringAsIdentifier("", "{{ .parameters.app_service_id }}/slots/{{ .parameters.slot_name }}/config/virtualNetwork"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_app_service_source_control": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/config/virtualNetwork
	"azurerm_app_service_virtual_network_swift_connection": config.TemplatedStringAsIdentifier("", "{{ .parameters.app_service_id }}/config/virtualNetwork"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/certificate1
	"azurerm_app_service_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/certificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/hostNameBindings/mywebsite.com|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/certificates/mywebsite.com
	"azurerm_app_service_certificate_binding": config.TemplatedStringAsIdentifier("", "{{ .parameters.hostname_binding_id }}|{{ .parameters.certificate_id }}"),

	// eventhub
	//
	// EventHub Cluster's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/clusters/cluster1
	"azurerm_eventhub_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/clusters/{{ .external_name }}"),
	// Customer Managed Keys for a EventHub Namespace can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1
	"azurerm_eventhub_namespace_customer_managed_key": config.ParameterAsIdentifier("eventhub_namespace_id"),

	// active_directory
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AAD/domainServices/DomainService1/trusts/trust1
	"azurerm_active_directory_domain_service_trust": config.TemplatedStringAsIdentifier("name", "{{ .parameters.domain_service_id }}/trusts/{{ .external_name }}"),

	// api_management
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/tagDescriptions/tagDescriptionId1
	"azurerm_api_management_api_tag_description": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/certificateAuthorities/cert1
	"azurerm_api_management_gateway_certificate_authority": config.TemplatedStringAsIdentifier("certificate_name", "{{ .parameters.api_management_id }}/gateways/{{ .parameters.gateway_name }}/certificateAuthorities/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/hostnameConfigurations/hc1
	"azurerm_api_management_gateway_host_name_configuration": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_management_id }}/gateways/{{ .parameters.gateway_name }}/hostnameConfigurations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/groups/groupId/users/user123
	"azurerm_api_management_group_user": config.TemplatedStringAsIdentifier("user_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/groups/{{ .parameters.group_name }}/users/{{ .external_name }}"),

	// app
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/webapp/providers/Microsoft.ServiceLinker/linkers/serviceconnector1
	"azurerm_app_service_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.app_service_id }}/providers/Microsoft.ServiceLinker/linkers/{{ .external_name }}"),

	// automation
	//
	"azurerm_automation_hybrid_runbook_worker": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/hybridRunbookWorkerGroups/grp1
	"azurerm_automation_software_update_configuration": config.TemplatedStringAsIdentifier("name", "{{ .parameters.automation_account_id }}/softwareUpdateConfigurations/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/sourceControls/sc1
	"azurerm_automation_source_control": config.TemplatedStringAsIdentifier("name", "{{ .parameters.automation_account_id }}/sourceControls/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/watchers/watch1
	"azurerm_automation_watcher": config.TemplatedStringAsIdentifier("name", "{{ .parameters.automation_account_id }}/watchers/{{ .external_name }}"),

	// cost_management
	//
	// /providers/Microsoft.Billing/billingAccounts/12345678/providers/Microsoft.CostManagement/exports/export1
	"azurerm_billing_account_cost_management_export": config.TemplatedStringAsIdentifier("name", "{{ .parameters.billing_account_id }}/providers/Microsoft.CostManagement/{{ .external_name }}"),

	// data_factory
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/dataflows/example
	"azurerm_data_factory_flowlet_data_flow": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/dataflows/{{ .external_name }}"),

	// datadog_monitor
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Datadog/monitors/monitor1
	"azurerm_datadog_monitor": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Datadog/monitors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Datadog/monitors/monitor1/singleSignOnConfigurations/default
	"azurerm_datadog_monitor_sso_configuration": config.TemplatedStringAsIdentifier("", "{{ .parameters.datadog_monitor_id }}/singleSignOnConfigurations/default"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Datadog/monitors/monitor1/tagRules/default
	"azurerm_datadog_monitor_tag_rule": config.TemplatedStringAsIdentifier("", "{{ .parameters.datadog_monitor_id }}/tagRules/default"),

	// digital_twins
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/timeSeriesDatabaseConnections/connection1
	"azurerm_digital_twins_time_series_database_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.digital_twins_id }}/timeSeriesDatabaseConnections/{{ .external_name }}"),

	// maintenance
	//
	// Maintenance Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Maintenance/configurationAssignments/assign1
	"azurerm_maintenance_assignment_virtual_machine_scale_set": config.TemplatedStringAsIdentifier("", "{{ .parameters.maintenance_configuration_id }}/configurationAssignments/default"),

	// managedservices
	//
	// Lighthouse Assignments can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ManagedServices/registrationAssignments/00000000-0000-0000-0000-000000000000
	"azurerm_lighthouse_assignment": config.IdentifierFromProvider,
	// Lighthouse Definitions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ManagedServices/registrationDefinitions/00000000-0000-0000-0000-000000000000
	"azurerm_lighthouse_definition": config.IdentifierFromProvider,

	// resource_group
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_resource_group_policy_remediation": config.TemplatedStringAsIdentifier("name", "{{ .parameters.resource_group_id }}/providers/Microsoft.PolicyInsights/remediations/{{ .external_name }}"),

	// sentinel
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_dynamics_365": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_microsoft_threat_protection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_office_365_project": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_office_atp": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_office_irm": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_office_power_bi": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),

	// stream_analytics
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/schedule/default
	"azurerm_stream_analytics_job_schedule": config.TemplatedStringAsIdentifier("", "{{ .parameters.stream_analytics_job_id }}/schedule/default"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_eventhub_v2": config.TemplatedStringAsIdentifier("name", "{{ .parameters.stream_analytics_job_id }}/inputs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1
	"azurerm_stream_analytics_job_storage_account": config.TemplatedStringAsIdentifier("", "{{ .parameters.stream_analytics_job_id }}"),

	// virtual_machine
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/packetCaptures/capture1
	"azurerm_virtual_machine_packet_capture": config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_watcher_id }}/packetCaptures/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkWatchers/watcher1/packetCaptures/capture1
	"azurerm_virtual_machine_scale_set_packet_capture": config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_watcher_id }}/packetCaptures/{{ .external_name }}"),

	// virtual_network
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/virtualNetworkGateways/gw1/natRules/rule1
	"azurerm_virtual_network_gateway_nat_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.virtual_network_gateway_id }}/natRules/{{ .external_name }}"),

	// vmware_netapp
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/privateCloud1/clusters/Cluster1/dataStores/datastore1
	"azurerm_vmware_netapp_volume_attachment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vmware_cluster_id }}/dataStores/{{ .external_name }}"),

	// web_pubsub
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webPubSub1/sharedPrivateLinkResources/resource1
	"azurerm_web_pubsub_shared_private_link_resource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.web_pubsub_id }}/sharedPrivateLinkResources/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
	"azurerm_web_pubsub_socketio": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.SignalRService/webPubSub/{{ .external_name }}"),

	// mssql
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/managedInstances/instance1/securityAlertPolicies/Default
	"azurerm_mssql_managed_instance_security_alert_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/managedInstances/{{ .parameters.managed_instance_name }}/securityAlertPolicies/{{ .external_name }}"),

	// nginx
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/deploy1/certificates/cer1
	"azurerm_nginx_certificate": config.TemplatedStringAsIdentifier("name", "{{ .parameters.nginx_deployment_id }}/certificates/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/dep1/configurations/default
	"azurerm_nginx_configuration": config.TemplatedStringAsIdentifier("", "{{ .parameters.nginx_deployment_id }}/configurations/default"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/dep1
	"azurerm_nginx_deployment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Nginx.NginxPlus/nginxDeployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/deploy1/apiKeys/key1
	"azurerm_nginx_api_key": config.TemplatedStringAsIdentifier("name", "{{ .parameters.nginx_deployment_id }}/apiKeys/{{ .external_name }}"),

	// private_dns
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1/inboundEndpoints/inboundEndpoint1
	"azurerm_dns_resolver_inbound_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.private_dns_resolver_id }}/inboundEndpoints/{{ .external_name }}"),

	// machinelearning
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/mlw1/datastores/datastore1
	"azurerm_machine_learning_datastore_blobstorage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workspace_id }}/datastores/{{ .external_name }}"),
	// AIFoundryProject parses ai_services_hub_id to get required resourcegroup
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/project1
	"azurerm_ai_foundry_project": config.IdentifierFromProvider,

	// network
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/scopeConnections/scopeConnection1
	"azurerm_network_manager_scope_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_manager_id }}/scopeConnections/{{ .external_name }}"),

	// orbital
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/spacecrafts/spacecraft1/contacts/contact1
	"azurerm_orbital_contact": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spacecraft_id }}/contacts/{{ .external_name }}"),

	// sentinel
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_threat_intelligence_taxii": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/onboardingStates/defaults
	"azurerm_security_insights_sentinel_onboarding": config.IdentifierFromProvider,

	// recoveryservices
	//
	// /subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/groupName/providers/Microsoft.RecoveryServices/vaults/vaultName/replicationRecoveryPlans/planName
	"azurerm_site_recovery_replication_recovery_plan": config.TemplatedStringAsIdentifier("name", "{{ .parameters.recovery_vault_id }}/replicationRecoveryPlans/{{ .external_name }}"),

	// advisor
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Advisor/recommendations/00000000-0000-0000-0000-000000000000/suppressions/name
	"azurerm_advisor_suppression": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Advisor/recommendations/{{ .parameters.recommendation_id }}/supressions/{{ .external_name }}"),

	// arckubernetes
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/cluster1
	"azurerm_arc_kubernetes_provisioned_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kubernetes/connectedClusters/{{ .external_name }}"),

	// hybridcompute
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.HybridCompute/machines/machine1
	"azurerm_arc_machine": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HybridCompute/machines/{{ .external_name }}"),

	// automanage
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HybridCompute/machines/machine1/providers/Microsoft.AutoManage/configurationProfileAssignments/default
	"azurerm_arc_machine_automanage_configuration_assignment": config.TemplatedStringAsIdentifier("", "{{ .parameters.arc_machine_id }}/providers/Microsoft.AutoManage/configurationProfileAssignments/default"),

	// dataprotection
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
	"azurerm_data_protection_backup_instance_mysql_flexible_server": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupInstances/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupPolicies/backupPolicy1
	"azurerm_data_protection_backup_policy_mysql_flexible_server": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupPolicies/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1
	"azurerm_data_protection_backup_vault_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.data_protection_backup_vault_id }}"),

	// devcenter
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/devCenters/dc1/attachedNetworks/et1
	"azurerm_dev_center_attached_network": config.TemplatedStringAsIdentifier("name", "{{ .parameters.dev_center_id }}/attachedNetworks/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevCenter/projects/project1/pools/pool1
	"azurerm_dev_center_project_pool": config.TemplatedStringAsIdentifier("name", "{{ .parameters.dev_center_project_id }}/pools/{{ .external_name }}"),

	// dynatrace
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Dynatrace.Observability/monitors/monitor1
	"azurerm_dynatrace_monitor": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Dynatrace.Observability/monitors/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Dynatrace.Observability/monitors/monitor1/tagRules/tagRules1
	"azurerm_dynatrace_tag_rules": config.TemplatedStringAsIdentifier("name", "{{ .parameters.monitor_id }}/tagRules/{{ .external_name }}"),

	// communication
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/emailServices/service1/domains/domain1/senderUsernames/username1
	"azurerm_email_communication_service_domain_sender_username": config.TemplatedStringAsIdentifier("name", "{{ .parameters.email_service_domain_id }}/senderUsernames/{{ .external_name }}"),

	// fabric
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Fabric/capacities/capacity1
	"azurerm_fabric_capacity": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Fabric/capacities/{{ .external_name }}"),

	// keyvault
	//
	// https://example-hsm.managedhsm.azure.net/keys/example
	"azurerm_key_vault_managed_hardware_security_module_key_rotation_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.managed_hsm_key_id }}"),

	// sql
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/jobs/myjob1
	"azurerm_mssql_job": config.TemplatedStringAsIdentifier("name", "{{ .parameters.job_agent_id }}/jobs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/jobs/myjob1
	"azurerm_mssql_job_schedule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.job_agent_id }}/jobs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/jobs/myjob1/steps/myjobstep1
	"azurerm_mssql_job_step": config.TemplatedStringAsIdentifier("name", "{{ .parameters.job_id }}/steps/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/targetGroups/mytargetgroup1
	"azurerm_mssql_job_target_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.job_agent_id }}/targetGroups/{{ .external_name }}"),

	// netapp
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolicies/backuppolicy1
	"azurerm_netapp_backup_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .parameters.account_name }}/backupPolicies/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1
	"azurerm_netapp_backup_vault": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .parameters.account_name }}/backupVaults/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mytest-rg/providers/Microsoft.NetApp/netAppAccounts/netapp-account-test/volumeGroups/netapp-volumegroup-test
	"azurerm_netapp_volume_group_oracle": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .parameters.account_name }}/volumeGroups/{{ .external_name }}"),

	// oracle
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/autonomousDatabases/autonomousDatabases1
	"azurerm_oracle_autonomous_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Oracle.Database/autonomousDatabases/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/cloudVmClusters/cloudVmClusters1
	"azurerm_oracle_cloud_vm_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Oracle.Database/cloudVmClusters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Oracle.Database/cloudExadataInfrastructures/cloudExadataInfrastructures1
	"azurerm_oracle_exadata_infrastructure": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Oracle.Database/cloudExadataInfrastructures/{{ .external_name }}"),

	// qumulo
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Qumulo.Storage/fileSystems/example
	"azurerm_qumulo_file_system": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Qumulo.Storage/fileSystems/{{ .external_name }}"),

	// servicebus
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1
	"azurerm_servicebus_namespace_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.namespace_id }}"),

	// azurestackhci
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/clus1/deploymentSettings/default
	"azurerm_stack_hci_deployment_setting": config.TemplatedStringAsIdentifier("", "{{ .parameters.stack_hci_cluster_id }}/deploymentSettings/default"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/cluster1/arcSettings/default/extensions/extension1
	"azurerm_stack_hci_extension": config.TemplatedStringAsIdentifier("name", "{{ .parameters.arc_setting_id }}/extensions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/image1
	"azurerm_stack_hci_marketplace_gallery_image": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/networkInterfaces/ni1
	"azurerm_stack_hci_network_interface": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AzureStackHCI/networkInterfaces/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/storageContainers/storage1
	"azurerm_stack_hci_storage_path": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AzureStackHCI/storageContainers/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.AzureStackHCI/virtualHardDisks/disk1
	"azurerm_stack_hci_virtual_hard_disk": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AzureStackHCI/virtualHardDisks/{{ .external_name }}"),

	// scvmm
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.HybridCompute/machines/machine1/providers/Microsoft.ScVmm/virtualMachineInstances/default
	"azurerm_system_center_virtual_machine_manager_virtual_machine_instance": config.TemplatedStringAsIdentifier("", "{{ .parameters.scoped_resource_id }}/providers/Microsoft.ScVmm/virtualMachineInstances/default"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.HybridCompute/machines/machine1/providers/Microsoft.ScVmm/virtualMachineInstances/default/guestAgents/default
	"azurerm_system_center_virtual_machine_manager_virtual_machine_instance_guest_agent": config.TemplatedStringAsIdentifier("", "{{ .parameters.scoped_resource_id }}/providers/Microsoft.ScVmm/virtualMachineInstances/default/guestAgents/default"),

	// trustedsigning
	//
	// /subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.CodeSigning/codeSigningAccounts/example-account
	"azurerm_trusted_signing_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CodeSigning/codeSigningAccounts/{{ .external_name }}"),

	// videoindexer
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.VideoIndexer/accounts/example-account-name
	"azurerm_video_indexer_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.VideoIndexer/accounts/{{ .external_name }}"),
}
