package config

import "github.com/upbound/upjet/pkg/config"

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
	// API Management Custom Domains can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/customDomains/default
	"azurerm_api_management_custom_domain": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_management_id }}/customDomains/default"),
	// API Management Azure AD B2C Identity Providers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/service1/identityProviders/AadB2C
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_aadb2c": config.IdentifierFromProvider,

	// appconfiguration
	//
	// App Configurations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1
	"azurerm_app_configuration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppConfiguration/configurationStores/{{ .external_name }}"),
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
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1
	"azurerm_spring_cloud_api_portal": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/apiPortals/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_build_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1/buildPackBindings/binding1
	"azurerm_spring_cloud_build_pack_binding": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_builder_id }}/buildPackBindings/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1
	"azurerm_spring_cloud_builder": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/configurationServices/configurationService1
	"azurerm_spring_cloud_configuration_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/configurationServices/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_container_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1
	"azurerm_spring_cloud_gateway": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/gateways/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/domains/domain1
	"azurerm_spring_cloud_gateway_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_gateway_id }}/domains/{{ .external_name }}"),
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

	// certificateregistration
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.CertificateRegistration/certificateOrders/certificateorder1
	"azurerm_app_service_certificate_order": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CertificateRegistration/certificateOrders/{{ .external_name }}"),

	// cognitiveservices
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
	"azurerm_cognitive_account_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .paramteres.cognitive_account_id }}"),

	// authorization
	//
	// /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyAssignments/assignment1
	"azurerm_management_group_policy_assignment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Authorization/policyAssignments/{{ .external_name }}"),
	// /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_management_group_policy_exemption": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policySetDefinitions/testPolicySet
	"azurerm_policy_set_definition": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policySetDefinitions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_resource_group_policy_exemption": config.TemplatedStringAsIdentifier("name", "{{ .parameters.resource_group_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),
	// Policy Remediations can be imported using the resource id
	//  /providers/Microsoft.Management/managementGroups/my-mgmt-group-id/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_management_group_policy_remediation": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/Microsoft.PolicyInsights/remediations/{{ .external_name }}"),

	// automation
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/certificates/certificate1
	"azurerm_automation_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/certificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_classic_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_service_principal": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/nodeConfigurations/configuration1
	"azurerm_automation_dsc_nodeconfiguration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/nodeConfigurations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/jobSchedules/10000000-1001-1001-1001-000000000001
	"azurerm_automation_job_schedule": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/Get-AzureVMTutorial
	"azurerm_automation_runbook": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/runbooks/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webHooks/TestRunbook_webhook
	"azurerm_automation_webhook": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/schedules/schedule1
	"azurerm_automation_schedule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_datetime": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/variables/{{ .external_name }}"),

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
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StoragePool/diskPools/diskPoolValue/iscsiTargets/iscsiTargetValue/lun|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/disks/disk1
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_disk_pool_iscsi_target_lun": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StoragePool/diskPools/storagePool1/managedDisks|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/disks/disk1
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_disk_pool_managed_disk_attachment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1
	"azurerm_shared_image": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .parameters.gallery_name }}/images/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/1.2.3
	"azurerm_shared_image_version": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .parameters.gallery_name }}/images/{{ .parameters.image_name }}/versions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/SshPublicKeys/mySshPublicKeyName1
	"azurerm_ssh_public_key": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/SshPublicKeys/{{ .external_name }}"),
	// Disk SAS Token can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/disks/manageddisk1
	"azurerm_managed_disk_sas_token": config.IdentifierFromProvider,

	// customproviders
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.CustomProviders/resourceProviders/example
	"azurerm_custom_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CustomProviders/resourceProviders/{{ .external_name }}"),

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

	// insights
	//
	// Diagnostic Settings can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.KeyVault/vaults/vault1|logMonitoring1
	"azurerm_monitor_diagnostic_setting": config.IdentifierFromProvider,

	// iot
	//
	// Azure IoT Time Series Insights EventHub Event Source can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/eventSources/example
	"azurerm_time_series_insights_event_source_eventhub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.TimeSeriesInsights/environments/{{ .parameters.environment_name }}/eventSources/{{ .externalName }}"),

	// load
	//
	// An existing Load Test can be imported into Terraform using the resource id
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}
	"azurerm_load_test": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.LoadTestService/loadTests/{{ .externalName }}"),

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
	// Synapse Role Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1|000000000000
	"azurerm_synapse_role_assignment": config.IdentifierFromProvider,
	// Synapse SQL Pool Vulnerability Assessment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default
	"azurerm_synapse_sql_pool_vulnerability_assessment": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_id }}/vulnerabilityAssessments/default"),
	// Synapse SQL Pool Vulnerability Assessment Rule Baselines can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default/rules/rule1/baselines/baseline1
	"azurerm_synapse_sql_pool_vulnerability_assessment_baseline": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_vulnerability_assessment_id }}/rules/{{ .parameters.rule_name }}/baselines/{{ .external_name }}"),
	// Synapse Workspace Keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/keys/key1
	"azurerm_synapse_workspace_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/keys/{{ .external_name }}"),

	// dbformysql
	// A MySQL Server Key can be imported using the resource id of the MySQL Server Key
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforMySQL/servers/server1/keys/keyvaultname_key-name_keyversion
	"azurerm_mysql_server_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/keys/{{ .external_name }}"),

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

	// sql
	//
	// MS SQL Database Extended Auditing Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/databases/db1/extendedAuditingSettings/default
	"azurerm_mssql_database_extended_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.database_id }}/extendedAuditingSettings/default"),
	// Database Vulnerability Assessment Rule Baseline can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/databases/mysqldatabase/vulnerabilityAssessments/Default/rules/VA2065/baselines/master
	"azurerm_mssql_database_vulnerability_assessment_rule_baseline": config.TemplatedStringAsIdentifier("baseline_name", "{{ .parameters.server_vulnerability_assessment_id }}/rules/{{ .parameters.rule_id }}/baselines/{{ .external_name }}"),
	// SQL Elastic Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/myelasticpoolname
	"azurerm_mssql_elasticpool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/elasticPools/{{ .external_name }}"),
	// SQL Firewall Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/firewallRules/rule1
	"azurerm_mssql_firewall_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/firewallRules/{{ .external_name }}"),
	// Elastic Job Agents can be imported using the id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1
	"azurerm_mssql_job_agent": config.IdentifierFromProvider,
	// Elastic Job Credentials can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/credentials/credential1
	"azurerm_mssql_job_credential": config.TemplatedStringAsIdentifier("name", "{{ .parameters.job_agent_id }}/credentials/{{ .external_name }}"),
	// MS SQL Server Security Alert Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/securityAlertPolicies/Default
	"azurerm_mssql_server_security_alert_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/securityAlertPolicies/default"),
	// MS SQL Server Vulnerability Assessment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/vulnerabilityAssessments/Default
	"azurerm_mssql_server_vulnerability_assessment": config.IdentifierFromProvider,
	// A SQL Active Directory Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/administrators/activeDirectory
	"azurerm_sql_active_directory_administrator": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/administrators/activeDirectory/default"),
	// SQL Databases can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/databases/database1
	"azurerm_sql_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	// SQL Elastic Pool's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/pool1
	"azurerm_sql_elasticpool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/elasticPools/{{ .external_name }}"),
	// SQL Firewall Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/firewallRules/rule1
	"azurerm_sql_firewall_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	// SQL Managed Databases can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver/databases/mydatabase
	"azurerm_sql_managed_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.sql_managed_instance_id }}/databases/{{ .external_name }}"),
	// SQL Servers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver
	"azurerm_sql_managed_instance": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/managedInstances/{{ .external_name }}"),
	// A SQL Active Directory Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/mymanagedinstance/administrators/activeDirectory
	"azurerm_sql_managed_instance_active_directory_administrator": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_managed_instance_id }}/administrators/activeDirectory"),
	// SQL Instance Failover Groups can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Sql/locations/Location/instanceFailoverGroups/failoverGroup1
	"azurerm_sql_managed_instance_failover_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/locations/{{ .parameters.location }}/instanceFailoverGroups/{{ .external_name }}"),

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
	// Storage Account Network Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
	"azurerm_storage_account_network_rules": config.TemplatedStringAsIdentifier("", "{{ .parameters.storage_account_id }}"),

	// databoxedge
	//
	// DEPRECATED
	// Databox Edge Orders can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1/orders/default
	"azurerm_databox_edge_order": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{ .parameters.device_name }}/orders/default"),
	// Databox Edge Devices can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1
	"azurerm_databox_edge_device": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{ .external_name }}"),

	// databricks
	//
	// Databrick Workspaces can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
	"azurerm_databricks_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Databricks/workspaces/{{ .external_name }}"),
	// Databricks Workspace Customer Managed Key can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
	"azurerm_databricks_workspace_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.workspace_id }}"),

	// securityinsights
	//
	// Sentinel Scheduled Alert Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_scheduled": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),
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
	// Sentinel Fusion Alert Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_fusion": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),
	// Sentinel Machine Learning Behavior Analytics Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_machine_learning_behavior_analytics": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}//providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),

	// sentinel
	//
	// AWS S3 Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_aws_s3": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Sentinel Watchlist Items can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1/watchlistItems/item1
	"azurerm_sentinel_watchlist_item": config.TemplatedStringAsIdentifier("", "{{ .parameters.watchlist_id }}/watchlistItems/{{ .external_name }}"),

	// solutions
	//
	// Managed Application can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applications/app1
	"azurerm_managed_application": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Solutions/applications/{{ .external_name }}"),

	// disk
	//
	// iSCSI Targets can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.StoragePool/diskPools/pool1/iscsiTargets/iscsiTarget1
	// SKIP LIST: Azure are officially halting the preview of Azure Disk Pools, and it will not be made generally available.
	"azurerm_disk_pool_iscsi_target": config.TemplatedStringAsIdentifier("name", "{{ .parameters.disks_pool_id }}/iscsiTargets/{{ .external_name }}"),

	// recoveryservices
	//
	// Site Recovery Protection Containers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name
	"azurerm_site_recovery_protection_container": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationFabrics/{{ .parameters.recovery_fabric_name }}/replicationProtectionContainers/{{ .external_name }}"),
	// Site Recovery Protection Container Mappings can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric1/replicationProtectionContainers/container1/replicationProtectionContainerMappings/mapping1
	"azurerm_site_recovery_protection_container_mapping": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationFabrics/{{ .parameters.recovery_fabric_name }}/replicationProtectionContainers/{{ .parameters.recovery_source_protection_container_name }}/replicationProtectionContainerMappings/{{ .external_name }}"),
	// Site Recovery Replicated VM's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name/replicationProtectedItems/vm-replication-name
	"azurerm_site_recovery_replicated_vm": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationFabrics/{{ .parameters.source_recovery_fabric_name }}/replicationProtectionContainers/{{ .parameters.source_recovery_protection_container_name }}/replicationProtectedItems/{{ .external_name }}"),
	// Site Recovery Replication Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationPolicies/policy-name
	"azurerm_site_recovery_replication_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationPolicies/{{ .external_name }}"),

	// relay
	//
	// Relay Hybrid Connection's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1/hybridConnections/hconn1
	"azurerm_relay_hybrid_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Relay/namespaces/{{ .parameters.relay_namespace_name }}/hybridConnections/{{ .extenal_name }}"),
	// Azure Relay Hybrid Connection Authorization Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Relay/namespaces/namespace1/hybridConnections/connection1/authorizationRules/rule1
	"azurerm_relay_hybrid_connection_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Relay/namespaces/{{ .parameters.relay_namespace_name }}/hybridConnections/{{ .parameters.hybrid_connection_name }}/authorizationRules/{{ .external_name }}"),
	// Azure Relay Namespace's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1
	"azurerm_relay_namespace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Relay/namespaces/{{ .external_name }}"),
	// Azure Relay Namespace Authorization Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Relay/namespaces/namespace1/authorizationRules/rule1
	"azurerm_relay_namespace_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Relay/namespaces/{{ .parameters.namespace_name }}/authorizationRules/{{ .external_name }}"),

	// resource
	//
	// Policy Remediations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_resource_policy_remediation": config.IdentifierFromProvider,

	// resources
	//
	// Managements can be imported using the resource id
	// /managementGroup/MyManagementGroup/subscription/12345678-1234-1234-1234-123456789012
	"azurerm_management_group_subscription_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.management_group_id }}/subscription/{{ .parameters.subscription_id }}"),
	// Management Group Template Deployments can be imported using the resource id
	// /providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/deploy1
	"azurerm_management_group_template_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Resources/deployments/{{ .external_name }}"),
	// No Import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_template_deployment": config.IdentifierFromProvider,
	// Tenant Template Deployments can be imported using the resource id
	// /providers/Microsoft.Resources/deployments/deploy1
	"azurerm_tenant_template_deployment": config.TemplatedStringAsIdentifier("name", "/providers/Microsoft.Resources/deployments/{{ .external_name }}"),

	// search
	//
	// Search Services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Search/searchServices/service1
	"azurerm_search_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Search/searchServices/{{ .external_name }}"),

	// function
	//
	// a Function App Active Slot can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_function_app_active_slot": config.IdentifierFromProvider,
	// a Function App Function can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/functions/function1
	"azurerm_function_app_function": config.IdentifierFromProvider,
	// a Function App Hybrid Connection can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/hybridConnectionNamespaces/hybridConnectionNamespace1/relays/relay1
	"azurerm_function_app_hybrid_connection": config.IdentifierFromProvider,

	// hardwaresecuritymodules
	//
	// Dedicated Hardware Security Module can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/hsm1
	"azurerm_dedicated_hardware_security_module": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{{ .external_name }}"),

	//nolint
	// hdinsight
	//
	// HDInsight Hadoop Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
	//nolint
	"azurerm_hdinsight_hadoop_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HDInsight/clusters/{{ .external_name }}"),
	// HDInsight HBase Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
	//nolint
	"azurerm_hdinsight_hbase_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HDInsight/clusters/{{ .external_name }}"),
	// HDInsight Interactive Query Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
	//nolint
	"azurerm_hdinsight_interactive_query_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HDInsight/clusters/{{ .external_name }}"),
	// HDInsight Kafka Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
	//nolint
	"azurerm_hdinsight_kafka_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HDInsight/clusters/{{ .external_name }}"),
	// HDInsight Spark Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
	//nolint
	"azurerm_hdinsight_spark_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HDInsight/clusters/{{ .external_name }}"),

	// healthbot
	//
	// Healthbot Service can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.HealthBot/healthBots/bot1
	"azurerm_healthbot": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HealthBot/healthBots/{{ .external_name }}"),

	// healthcare
	//
	// Healthcare DICOM Service can be imported using the resourceid
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomServices/service1
	"azurerm_healthcare_dicom_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workspace_id }}/dicomServices/{{ .external_name }}"),
	// Healthcare FHIR Service can be imported using the resourceid
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/fhirServices/service1
	"azurerm_healthcare_fhir_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workspace_id }}/fhirServices/{{ .external_name }}"),
	// Healthcare Workspaces can be imported using the resourceid
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1
	"azurerm_healthcare_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HealthcareApis/workspaces/{{ .external_name }}"),

	// app
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/site1/publicCertificates/publicCertificate1
	"azurerm_app_service_public_certificate": config.TemplatedStringAsIdentifier("certificate_name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .parameters.app_service_name }}/publicCertificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/staging/hostNameBindings/mywebsite.com
	"azurerm_app_service_slot_custom_hostname_binding": config.TemplatedStringAsIdentifier("hostname", "{{ .parameters.app_service_slot_id }}/hostNameBindings/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_app_service_source_control_slot": config.IdentifierFromProvider,

	// synapse
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/vulnerabilityAssessments/default
	"azurerm_synapse_workspace_vulnerability_assessment": config.IdentifierFromProvider,

	// timeseriesinsights
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/accessPolicies/example
	"azurerm_iot_time_series_insights_access_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.time_series_insights_environment_id }}/accessPolicies/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/eventSources/example
	"azurerm_iot_time_series_insights_event_source_iothub": config.TemplatedStringAsIdentifier("name", "{{ .parameters.environment_id }}/eventSources/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
	"azurerm_iot_time_series_insights_gen2_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.TimeSeriesInsights/environments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example/referenceDataSets/example
	"azurerm_iot_time_series_insights_reference_data_set": config.TemplatedStringAsIdentifier("name", "{{ .parameters.time_series_insights_environment_id }}/referenceDataSets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
	"azurerm_iot_time_series_insights_standard_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.TimeSeriesInsights/environments/{{ .external_name }}"),

	// traffic
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Network/trafficManagerProfiles/example-profile/AzureEndpoints/example-endpoint
	"azurerm_traffic_manager_azure_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.profile_id }}/AzureEndpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-group/providers/Microsoft.Network/trafficManagerProfiles/example-profile/ExternalEndpoints/example-endpoint
	"azurerm_traffic_manager_external_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.profile_id }}/ExternalEndpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Network/trafficManagerProfiles/example-profile/NestedEndpoints/example-endpoint
	"azurerm_traffic_manager_nested_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.profile_id }}/NestedEndpoints/{{ .external_name }}"),

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
	"azurerm_app_service_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/hostingEnvironments/{{ .external_name }}"),
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

	// eventgrid
	//
	// EventGrid System Topic Event Subscriptions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/systemTopics/topic1/eventSubscriptions/subscription1
	"azurerm_eventgrid_system_topic_event_subscription": config.IdentifierFromProvider,

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
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/apis/api1
	"azurerm_api_management_gateway_api": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/certificateAuthorities/cert1
	"azurerm_api_management_gateway_certificate_authority": config.TemplatedStringAsIdentifier("certificate_name", "{{ .parameters.api_management_id }}/gateways/{{ .parameters.gateway_name }}/certificateAuthorities/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/hostnameConfigurations/hc1
	"azurerm_api_management_gateway_host_name_configuration": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_management_id }}/gateways/{{ .parameters.gateway_name }}/hostnameConfigurations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/schemas/schema1
	"azurerm_api_management_global_schema": config.TemplatedStringAsIdentifier("schema_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/schemas/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/groups/groupId/users/user123
	"azurerm_api_management_group_user": config.TemplatedStringAsIdentifier("user_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/groups/{{ .parameters.group_name }}/users/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/products/myproduct/tags/mytag
	"azurerm_api_management_product_tag": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_management_product_id }}/tags/{{ .external_name }}"),

	// app
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/sites/webapp/providers/Microsoft.ServiceLinker/linkers/serviceconnector1
	"azurerm_app_service_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.app_service_id }}/providers/Microsoft.ServiceLinker/linkers/{{ .external_name }}"),

	// application_insights
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/webTests/my_test
	"azurerm_application_insights_web_test": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/webTests/{{ .external.name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/workbooks/resource1
	"azurerm_application_insights_workbook": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/workbooks/{{ .external.name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/workbookTemplates/resource1
	"azurerm_application_insights_workbook_template": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/workbookTemplates/{{ .external.name }}"),

	// automation
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connectionTypes/type1
	"azurerm_automation_connection_type": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connectionTypes/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/hybridRunbookWorkerGroups/group1/hybridRunbookWorkers/00000000-0000-0000-0000-000000000000
	"azurerm_automation_hybrid_runbook_worker": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/hybridRunbookWorkerGroups/grp1
	"azurerm_automation_hybrid_runbook_worker_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/hybridRunbookWorkerGroups/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/softwareUpdateConfigurations/suc1
	"azurerm_automation_software_update_configuration": config.TemplatedStringAsIdentifier("name", "{{ .parameters.automation_account_id }}/softwareUpdateConfigurations/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/sourceControls/sc1
	"azurerm_automation_source_control": config.TemplatedStringAsIdentifier("name", "{{ .parameters.automation_account_id }}/sourceControls/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/watchers/watch1
	"azurerm_automation_watcher": config.TemplatedStringAsIdentifier("name", "{{ .parameters.automation_account_id }}/watchers/{{ .external_name }}"),

	// network
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/bastionHosts/instance1
	"azurerm_bastion_host": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/bastionHosts/{{ .exteranl_name }}"),

	// cost_management
	//
	// /providers/Microsoft.Billing/billingAccounts/12345678/providers/Microsoft.CostManagement/exports/export1
	"azurerm_billing_account_cost_management_export": config.TemplatedStringAsIdentifier("name", "{{ .parameters.billing_account_id }}/providers/Microsoft.CostManagement/{{ .external_name }}"),

	// compute
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1/capacityReservations/capacityReservation1
	"azurerm_capacity_reservation": config.TemplatedStringAsIdentifier("name", "{{ .parameters.capacity_reservation_group_id }}/capacityReservations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1
	"azurerm_capacity_reservation_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/capacityReservationGroups/{{ .external_name }}"),

	// cdn
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/firewallPolicy1
	"azurerm_cdn_frontdoor_firewall_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/secrets/secrets1
	"azurerm_cdn_frontdoor_secret": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_profile_id }}/secrets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/securityPolicies/policy1
	"azurerm_cdn_frontdoor_security_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_profile_id }}/securityPolicies/{{ .external_name }}"),

	// cognitive_deployment
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.CognitiveServices/accounts/account1/deployments/deployment1
	"azurerm_cognitive_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cognitive_account_id }}/deployments/{{ .external_name }}"),

	// container_registry
	//
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_container_registry_task_schedule_run_now": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/tokens/token1/passwords/password
	"azurerm_container_registry_token_password": config.IdentifierFromProvider,

	// cosmo_db
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/services/SqlDedicatedGateway
	"azurerm_cosmosdb_sql_dedicated_gateway": config.TemplatedStringAsIdentifier("", "{{ .parameters.cosmosdb_account_id }}/services/SqlDedicatedGateway"),

	// dashboard_grafana
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Dashboard/grafana/workspace1
	"azurerm_dashboard_grafana": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Dashboard/grafana/{{ .external_name }}"),

	// data_factory
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_sql_server_table": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/dataflows/example
	"azurerm_data_factory_flowlet_data_flow": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/dataflows/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
	"azurerm_data_factory_integration_runtime_managed": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/integrationruntimes/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_cosmosdb_mongoapi": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),

	// data_protection
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
	"azurerm_data_protection_backup_instance_postgresql": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupInstances/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/resourceGuards/resourceGuard1
	"azurerm_data_protection_resource_guard": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataProtection/resourceGuards/{{ .external_name }}"),

	// databricks
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/accessConnectors/connector1
	"azurerm_databricks_access_connector": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Databricks/accessConnectors/{{ .external_name }}"),

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

	// eventhub
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/schemaGroups/group1
	"azurerm_eventhub_namespace_schema_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/schemaGroups/{{ .external_name }}"),

	// federated_identity
	//
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{parentIdentityName}/federatedIdentityCredentials/{resourceName}
	"azurerm_federated_identity_credential": config.TemplatedStringAsIdentifier("name", "{{ .parameters.parent_id }}/federatedIdentityCredentials/{{ .external_name }}"),

	// fluid_relay
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.FluidRelay/fluidRelayServers/server1
	"azurerm_fluid_relay_server": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.FluidRelay/fluidRelayServers/{{ .external_name }}"),

	// gallery
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1
	"azurerm_gallery_application": config.TemplatedStringAsIdentifier("name", "{{ .parameters.gallery_id }}/applications/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1/versions/galleryApplicationVersion1
	"azurerm_gallery_application_version": config.TemplatedStringAsIdentifier("name", "{{ .parameters.gallery_application_id }}/versions/{{ .external_name }}"),

	// healthcare
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotConnectors/iotconnector1
	"azurerm_healthcare_medtech_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workspace_id }}/iotConnectors/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotConnectors/iotconnector1/fhirDestinations/destination1
	"azurerm_healthcare_medtech_service_fhir_destination": config.TemplatedStringAsIdentifier("name", "{{ .parameters.medtech_service_id }}/fhirDestinations/{{ .external_name }}"),

	// iot
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/eventSources/example
	"azurerm_iot_time_series_insights_event_source_eventhub": config.TemplatedStringAsIdentifier("name", "{{ .parameters.environment_id }}/eventSources/{{ .external_name }}"),

	// logz
	//
	// logz Monitors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1
	"azurerm_logz_monitor": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logz/monitors/{{ .external_name }}"),
	// logz Tag Rules can be imported using the resource id
	//  /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1/tagRules/ruleSet1
	"azurerm_logz_tag_rule": config.TemplatedStringAsIdentifier("", "{{ .parameters.logz_monitor_id }}/tagRules/default"),

	// machinelearningservices
	//
	// Machine Learning Workspace can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1
	"azurerm_machine_learning_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.MachineLearningServices/workspaces/{{ .external_name }}"),
	// Machine Learning Compute Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
	"azurerm_machine_learning_compute_cluster": config.TemplatedStringAsIdentifier("name", "{{ .parameters.machine_learning_workspace_id }}/computes/{{ .external_name }}"),
	// Machine Learning Synapse Sparks can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
	"azurerm_machine_learning_synapse_spark": config.TemplatedStringAsIdentifier("name", "{{ .parameters.machine_learning_workspace_id }}/computes/{{ .external_name }}"),
	// Machine Learning Compute Instances can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/compute1
	"azurerm_machine_learning_compute_instance": config.TemplatedStringAsIdentifier("name", "{{ .parameters.machine_learning_workspace_id }}/computes/{{ .external_name }}"),

	// maintenance
	//
	// Maintenance Configuration can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Maintenance/maintenanceConfigurations/example-mc
	"azurerm_maintenance_configuration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Maintenance/maintenanceConfigurations/{{ .external_name }}"),
	// Maintenance Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/hostGroups/group1/hosts/host1/providers/Microsoft.Maintenance/configurationAssignments/assign1
	"azurerm_maintenance_assignment_dedicated_host": config.TemplatedStringAsIdentifier("", "{{ .parameters.maintenance_configuration_id }}/configurationAssignments/default"),
	// Maintenance Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.Maintenance/configurationAssignments/assign1
	"azurerm_maintenance_assignment_virtual_machine": config.TemplatedStringAsIdentifier("", "{{ .parameters.maintenance_configuration_id }}/configurationAssignments/default"),
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

	// resource_deployment
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Resources/deploymentScripts/script1
	"azurerm_resource_deployment_script_azure_cli": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Resources/deploymentScripts/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Resources/deploymentScripts/script1
	"azurerm_resource_deployment_script_azure_power_shell": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Resources/deploymentScripts/{{ .external_name }}"),

	// resource_group
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_resource_group_policy_remediation": config.TemplatedStringAsIdentifier("name", "{{ .parameters.resource_group_id }}/providers/Microsoft.PolicyInsights/remediations/{{ .external_name }}"),

	// route
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1/routes/myroute1
	"azurerm_route": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/routeTables/{{ .parameters.route_table_name }}/routes/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/routeFilters/routeFilter1
	"azurerm_route_filter": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/routeFilters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1
	"azurerm_route_map": config.TemplatedStringAsIdentifier("name", "{{ .parameters.virtual_hub_id }}/routeMaps/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/routeServer1
	"azurerm_route_server": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualHubs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/routeServer1/bgpConnections/connection1
	"azurerm_route_server_bgp_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.route_server_id }}/bgpConnections/{{ .external_name }}"),

	// search
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Search/searchServices/service1/sharedPrivateLinkResources/resource1
	"azurerm_search_shared_private_link_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.search_service_id }}/sharedPrivateLinkResources/{{ .external_name }}"),

	// sentinel
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_ms_security_incident": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_nrt": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_dynamics_365": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_iot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
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

	// signalr
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1/sharedPrivateLinkResources/resource1
	"azurerm_signalr_shared_private_link_resource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.signalr_service_id }}/sharedPrivateLinkResources/{{ .external_name }}"),

	// site_recovery
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/primary-fabric-name/replicationNetworks/azureNetwork/replicationNetworkMappings/mapping-name
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_site_recovery_network_mapping": config.IdentifierFromProvider,

	// spring_cloud
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/applicationAccelerators/default
	"azurerm_spring_cloud_accelerator": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/applicationAccelerators/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1/domains/domain1
	"azurerm_spring_cloud_api_portal_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_api_portal_id }}/domains/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/applicationLiveViews/default
	"azurerm_spring_cloud_application_live_view": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/applicationLiveViews/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/Spring/springcloud/apps/springcloudapp/deployments/deployment/providers/Microsoft.ServiceLinker/linkers/serviceconnector1
	"azurerm_spring_cloud_connection": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/Spring/spring1/applicationAccelerators/default/customizedAccelerators/customizedAccelerator1
	"azurerm_spring_cloud_customized_accelerator": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_accelerator_id }}/customizedAccelerators/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/DevToolPortals/default
	"azurerm_spring_cloud_dev_tool_portal": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/DevToolPortals/{{ .external_name }}"),

	// static_site
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1
	"azurerm_static_site": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/staticSites/{{ .external_name }}"),

	// storage
	//
	// https://account1.dfs.core.windows.net/fileSystem1/path
	"azurerm_storage_data_lake_gen2_path": config.IdentifierFromProvider,
	// https://tomdevsa20.file.core.windows.net/share1/directory1
	"azurerm_storage_share_directory": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StorageSync/storageSyncServices/sync1/syncGroups/syncgroup1/cloudEndpoints/cloudEndpoint1
	"azurerm_storage_sync_cloud_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.storage_sync_group_id }}/cloudEndpoints/{{ .external_name }}"),
	// https://example.table.core.windows.net/table1(PartitionKey='samplepartition',RowKey='samplerow')
	"azurerm_storage_table_entity": config.IdentifierFromProvider,
	// stream_analytics
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/schedule/default
	"azurerm_stream_analytics_job_schedule": config.TemplatedStringAsIdentifier("", "{{ .parameters.stream_analytics_job_id }}/schedule/default"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
	"azurerm_stream_analytics_output_eventhub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_reference_input_mssql": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_name }}/inputs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_eventhub_v2": config.TemplatedStringAsIdentifier("name", "{{ .parameters.stream_analytics_job_id }}/inputs/{{ .external_name }}"),

	// synapse_workspace
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlAdministrators/activeDirectory
	"azurerm_synapse_workspace_sql_aad_admin": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/sqlAdministrators/activeDirectory"),

	// virtual_desktop
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DesktopVirtualization/hostPools/pool1/registrationInfo/default
	"azurerm_virtual_desktop_host_pool_registration_info": config.TemplatedStringAsIdentifier("", "{{ .parameters.hostpool_id }}/registrationInfo/default"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/plan1
	"azurerm_virtual_desktop_scaling_plan": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DesktopVirtualization/scalingPlans/{{ .external_name }}"),

	// key_vault
	//
	// https://example-keyvault.vault.azure.net/certificates/contacts
	"azurerm_key_vault_certificate_contacts": config.IdentifierFromProvider,

	// iothub
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/example/certificates/example
	"azurerm_iothub_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/iotHubs/{{ .parameters.iothub_name }}/certificates/{{ .external_name }}"),

	// virtual_hub
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/connection1
	"azurerm_virtual_hub_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.virtual_hub_id }}/hubVirtualNetworkConnections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/ipConfigurations/ipConfig1
	"azurerm_virtual_hub_ip": config.TemplatedStringAsIdentifier("name", "{{ .parameters.virtual_hub_id }}/ipConfigurations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/routeTable1
	"azurerm_virtual_hub_route_table": config.TemplatedStringAsIdentifier("name", "{{ .parameters.virtual_hub_id }}/hubRouteTables/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/routeTable1/routes/routeName
	"azurerm_virtual_hub_route_table_route": config.TemplatedStringAsIdentifier("name", "{{ .parameters.route_table_id }}/routes/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/securityPartnerProviders/securityPartnerProvider1
	"azurerm_virtual_hub_security_partner_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/securityPartnerProviders/{{ .external_name }}"),

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

	// vpn_server
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/vpnServerConfigurations/serverConfiguration1/configurationPolicyGroups/configurationPolicyGroup1
	"azurerm_vpn_server_configuration_policy_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vpn_server_configuration_id }}/configurationPolicyGroups/{{ .external_name }}"),

	// web_app
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_web_app_active_slot": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/hybridConnectionNamespaces/hybridConnectionNamespace1/relays/relay1
	"azurerm_web_app_hybrid_connection": config.IdentifierFromProvider,

	// web_pubsub
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
	"azurerm_web_pubsub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.SignalRService/webPubSub/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webPubSub1/hubs/webPubSubhub1
	"azurerm_web_pubsub_hub": config.TemplatedStringAsIdentifier("name", "{{ .parameters.web_pubsub_id }}/hubs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webpubsub1
	// TODO: Bug in documentation. Normalize external_name while testing
	"azurerm_web_pubsub_network_acl": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webPubSub1/sharedPrivateLinkResources/resource1
	"azurerm_web_pubsub_shared_private_link_resource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.web_pubsub_id }}/sharedPrivateLinkResources/{{ .external_name }}"),

	// windows
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_windows_function_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_windows_function_app_slot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.function_app_id }}/slots/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_windows_web_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_windows_web_app_slot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.app_service_id }}/slots/{{ .external_name }}"),

	// logz
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1/accounts/subAccount1
	"azurerm_logz_sub_account": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logz_monitor_id }}/accounts/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1/accounts/subAccount1/tagRules/ruleSet1
	"azurerm_logz_sub_account_tag_rule": config.IdentifierFromProvider,

	// media_job
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Media/mediaServices/account1/transforms/transform1/jobs/job1
	"azurerm_media_job": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/transforms/{{ .parameters.transform_name }}/jobs/{{ .external_name }}"),

	// mssql
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/managedInstances/instance1/securityAlertPolicies/Default
	"azurerm_mssql_managed_instance_security_alert_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/managedInstances/{{ .parameters.managed_instance_name }}/securityAlertPolicies/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Sql/managedInstances/instance1/encryptionProtector/current
	"azurerm_mssql_managed_instance_transparent_data_encryption": config.TemplatedStringAsIdentifier("", "{{ .parameters.managed_instance_id }}/encryptionProtector/current"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/devOpsAuditingSettings/default
	"azurerm_mssql_server_microsoft_support_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/devOpsAuditingSettings/default"),

	// network_manager
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1
	"azurerm_network_manager": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkManagers/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/networkGroups/networkGroup1
	"azurerm_network_manager_network_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_manager_id }}/networkGroups/{{ .external_name }}"),

	// nginx
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/deploy1/certificates/cer1
	"azurerm_nginx_certificate": config.TemplatedStringAsIdentifier("name", "{{ .parameters.nginx_deployment_id }}/certificates/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/dep1/configurations/default
	"azurerm_nginx_configuration": config.TemplatedStringAsIdentifier("", "{{ .parameters.nginx_deployment_id }}/configurations/default"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Nginx.NginxPlus/nginxDeployments/dep1
	"azurerm_nginx_deployment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Nginx.NginxPlus/nginxDeployments/{{ .external_name }}"),

	// orbital_contact
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/contactProfiles/contactProfile1
	"azurerm_orbital_contact_profile": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Orbital/contactProfiles/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/spacecrafts/spacecraft1
	"azurerm_orbital_spacecraft": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Orbital/spacecrafts/{{ .external_name }}"),

	// policy_virtual_machine
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/assignment1
	"azurerm_policy_virtual_machine_configuration_assignment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.virtual_machine_id }}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{ .external_name }}"),

	// postgres
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/myserver/administrators/objectId
	"azurerm_postgresql_flexible_server_active_directory_administrator": config.TemplatedStringAsIdentifier("object_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{ .parameters.server_name }}/administrators/{{ .external_name }}"),

	// private_dns
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1
	"azurerm_private_dns_resolver": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsResolvers/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRuleset1
	"azurerm_private_dns_resolver_dns_forwarding_ruleset": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsForwardingRulesets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRuleset1/forwardingRules/forwardingRule1
	"azurerm_private_dns_resolver_forwarding_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.dns_forwarding_ruleset_id }}/forwardingRules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1/inboundEndpoints/inboundEndpoint1
	"azurerm_dns_resolver_inbound_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.private_dns_resolver_id }}/inboundEndpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1/outboundEndpoints/outboundEndpoint1
	"azurerm_private_dns_resolver_outbound_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.private_dns_resolver_id }}/outboundEndpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsForwardingRulesets/dnsForwardingRuleset1/virtualNetworkLinks/virtualNetworkLink1
	"azurerm_private_dns_resolver_virtual_network_link": config.TemplatedStringAsIdentifier("name", "{{ .parameters.dns_forwarding_ruleset_id }}/virtualNetworkLinks/{{ .external_name }}"),
}
