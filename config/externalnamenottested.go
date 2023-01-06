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

	// analysisservices
	//
	// Analysis Services Server can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AnalysisServices/servers/server1
	"azurerm_analysis_services_server": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AnalysisServices/servers/{{ .external_name }}"),

	// api
	//
	// API Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.Web/connections/example-connection
	"azurerm_api_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/connections/{{ .external_name }}"),
	// API Management API Tags can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/tags/tag1
	"azurerm_api_management_api_tag": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}/tags/{{ .external_name }}"),
	// API Management Notification Recipient Users can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientUsers/userid1
	"azurerm_api_management_notification_recipient_user": config.TemplatedStringAsIdentifier("user_id", "{{ .parameters.api_management_id }}/recipientUsers/{{ .external_name }}"),

	// apimanagement
	// API Management Services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1
	"azurerm_api_management": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .external_name }}"),
	// API Management API's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1
	"azurerm_api_management_api": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/apis/{{ .external_name }}"),
	// API Management Service API Diagnostics Logs can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/diagnostics/diagnostic1
	"azurerm_api_management_api_diagnostic": config.TemplatedStringAsIdentifier("identifier", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/apis/{{ .parameters.api_name }}/diagnostics/{{ .external_name }}"),
	// API Management API Operation's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/operations/operation1
	"azurerm_api_management_api_operation": config.TemplatedStringAsIdentifier("operation_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/apis/{{ .parameters.api_name }}/operations/{{ .external_name }}"),
	// API Management API Operation Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/operations/operation1/policies/policy
	"azurerm_api_management_api_operation_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/apis/{{ .parameters.api_name }}/operations/{{ .parameters.operation_id }}/policies/policy"),
	// API Management API Operation Tags can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/operations/operation1/tags/tag1
	"azurerm_api_management_api_operation_tag": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_operation_id }}/tags/{{ .external_name }}"),
	// API Management API Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/exampleId/policies/policy
	"azurerm_api_management_api_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/apis/{{ .parameters.api_name }}/policies/policy"),
	// API Management API Releases can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/releases/release1
	"azurerm_api_management_api_release": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}/releases/{{ .external_name }}"),
	// API Management API Schema's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/schemas/schema1
	"azurerm_api_management_api_schema": config.TemplatedStringAsIdentifier("schema_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/apis/{{ .parameters.api_name }}/schemas/{{ .external_name }}"),
	// API Version Set can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apiVersionSets/set1
	"azurerm_api_management_api_version_set": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/apiVersionSets/{{ .external_name }}"),
	// API Management Authorization Servers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/authorizationServers/server1
	"azurerm_api_management_authorization_server": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/authorizationServers/{{ .external_name }}"),
	// API Management backends can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/backends/backend1
	"azurerm_api_management_backend": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/backends/{{ .external_name }}"),
	// API Management Certificates can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/certificates/certificate1
	"azurerm_api_management_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/certificates/{{ .external_name }}"),
	// API Management Custom Domains can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/customDomains/default
	"azurerm_api_management_custom_domain": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_management_id }}/customDomains/default"),
	// API Management Diagnostics can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/diagnostics/applicationinsights
	"azurerm_api_management_diagnostic": config.TemplatedStringAsIdentifier("identifier", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/diagnostics/{{ .external_name }}"),
	// API Management Email Templates can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/templates/template1
	"azurerm_api_management_email_template": config.TemplatedStringAsIdentifier("template_name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/templates/{{ .external_name }}"),
	// API Management Gateways can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1
	"azurerm_api_management_gateway": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/gateways/{{ .external_name }}"),
	// API Management AAD Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/aad
	"azurerm_api_management_identity_provider_aad": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/identityProviders/aad"),
	// API Management Azure AD B2C Identity Providers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/service1/identityProviders/AadB2C
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_aadb2c": config.IdentifierFromProvider,
	// API Management Facebook Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/facebook
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_facebook": config.IdentifierFromProvider,
	// API Management Google Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/google
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_google": config.IdentifierFromProvider,
	// API Management Microsoft Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/microsoft
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_microsoft": config.IdentifierFromProvider,
	// API Management Twitter Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/twitter
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_twitter": config.IdentifierFromProvider,
	// API Management Loggers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/example-rg/providers/Microsoft.ApiManagement/service/example-apim/loggers/example-logger
	"azurerm_api_management_logger": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/loggers/{{ .external_name }}"),
	// API Management Properties can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/namedValues/example-apimp
	"azurerm_api_management_named_value": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/namedValues/{{ .external_name }}"),
	// API Management Notification Recipient Emails can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientEmails/email1
	"azurerm_api_management_notification_recipient_email": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_management_id }}/notifications/{{ .parameters.notification_type }}/recipientEmails/{{ .parameters.email }}"),
	// API Management OpenID Connect Providers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/openidConnectProviders/provider1
	"azurerm_api_management_openid_connect_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/openidConnectProviders/{{ .external_name }}"),
	// API Management service Policys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/policies/policy
	"azurerm_api_management_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_management_id }}/policies/policy"),
	// API Management Products can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/products/myproduct
	"azurerm_api_management_product": config.TemplatedStringAsIdentifier("product_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/products/{{ .external_name }}"),
	// API Management Product API's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/apis/apiId
	"azurerm_api_management_product_api": config.TemplatedStringAsIdentifier("api_name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/products/{{ .parameters.product_id }}/apis/{{ .external_name }}"),
	// API Management Product Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/policies/policy
	"azurerm_api_management_product_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/products/{{ .parameters.product_id }}/policies/policy"),
	// API Management Redis Caches can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/caches/cache1
	"azurerm_api_management_redis_cache": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_management_id }}/caches/{{ .external_name }}"),
	// API Management Subscriptions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/example-apim/subscriptions/subscription-name
	"azurerm_api_management_subscription": config.TemplatedStringAsIdentifier("display_name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/subscriptions/{{ .external_name }}"),
	// API Management Tags can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/tags/tag1
	"azurerm_api_management_tag": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_management_id }}/tags/{{ .external_name }}"),
	// API Management Users can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/users/abc123
	"azurerm_api_management_user": config.TemplatedStringAsIdentifier("user_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/users/{{ .external_name }}"),

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

	// appplatform
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1
	// TODO: Please see the ID while testing and if possible normalize the API for this resource.
	"azurerm_spring_cloud_active_deployment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1
	"azurerm_spring_cloud_api_portal": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/apiPortals/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/spring/myservice/apps/myapp
	"azurerm_spring_cloud_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/spring/{{ .parameters.service_name }}/apps/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
	"azurerm_spring_cloud_app_cosmosdb_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
	"azurerm_spring_cloud_app_mysql_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/spring/myservice/apps/myapp/bindings/bind1
	"azurerm_spring_cloud_app_redis_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_build_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1/buildPackBindings/binding1
	"azurerm_spring_cloud_build_pack_binding": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_builder_id }}/buildPackBindings/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1
	"azurerm_spring_cloud_builder": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/spring1/certificates/cert1
	"azurerm_spring_cloud_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/spring/{{ .parameters.service_name }}/certificates/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/configurationServices/configurationService1
	"azurerm_spring_cloud_configuration_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/configurationServices/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_container_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/domains/domain.com
	"azurerm_spring_cloud_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/domains/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1
	"azurerm_spring_cloud_gateway": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/gateways/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/domains/domain1
	"azurerm_spring_cloud_gateway_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_gateway_id }}/domains/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/routeConfigs/routeConfig1
	"azurerm_spring_cloud_gateway_route_config": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_gateway_id }}/routeConfigs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_java_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/spring/spring1
	"azurerm_spring_cloud_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/spring/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/storages/storage1
	"azurerm_spring_cloud_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/storages/{{ .external_name }}"),

	// avs
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/clusters/cluster1
	"azurerm_vmware_cluster": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vmware_cloud_id }}/clusters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/authorizations/authorization1
	"azurerm_vmware_express_route_authorization": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vmware_cloud_id }}/clusters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/PrivateClouds/privateCloud1
	"azurerm_vmware_private_cloud": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AVS/PrivateClouds/{{ .external_name }}"),

	// azurestackhci
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/cluster1
	"azurerm_stack_hci_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AzureStackHCI/clusters/{{ .external_name }}"),

	// blueprint
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint
	"azurerm_blueprint_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Blueprint/blueprintAssignments/{{ .external_name }}"),

	// cdn
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1/endpoints/myendpoint1
	"azurerm_cdn_endpoint": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cdn/profiles/{{ .parameters.profile_name }}/endpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customDomains/domain1
	"azurerm_cdn_endpoint_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_endpoint_id }}/customDomains/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1
	"azurerm_cdn_profile": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cdn/profiles/{{ .external_name }}"),

	// certificateregistration
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.CertificateRegistration/certificateOrders/certificateorder1
	"azurerm_app_service_certificate_order": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CertificateRegistration/certificateOrders/{{ .external_name }}"),

	// cognitiveservices
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
	"azurerm_cognitive_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CognitiveServices/accounts/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
	"azurerm_cognitive_account_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .paramteres.cognitive_account_id }}"),

	// authorization
	//
	// /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyAssignments/assignment1
	"azurerm_management_group_policy_assignment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Authorization/policyAssignments/{{ .external_name }}"),
	// /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_management_group_policy_exemption": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Authorization/locks/lock1
	"azurerm_management_lock": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policySetDefinitions/testPolicySet
	"azurerm_policy_set_definition": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policySetDefinitions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_resource_group_policy_exemption": config.TemplatedStringAsIdentifier("name", "{{ .parameters.resource_group_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),
	// {resource}/providers/Microsoft.Authorization/policyAssignments/assignment1
	"azurerm_resource_policy_assignment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_resource_policy_exemption": config.IdentifierFromProvider,
	// for scope Subscription, the id format is /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleAssignments/00000000-0000-0000-0000-000000000000
	// for scope Resource Group, the id format is /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Authorization/roleAssignments/00000000-0000-0000-0000-000000000000
	"azurerm_role_assignment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}|{scope}
	"azurerm_role_definition": config.TemplatedStringAsIdentifier("role_definition_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/roleDefinitions/{{ .external_name }}|{{ .parameters.scope }}"),
	// /subscriptions/00000000-0000-0000-000000000000/providers/Microsoft.Authorization/policyAssignments/assignment1
	"azurerm_subscription_policy_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policyAssignments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-000000000000/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_subscription_policy_exemption": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),

	// automation
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1
	"azurerm_automation_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .external_name }}"),
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
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/credentials/credential1
	"azurerm_automation_credential": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/credentials/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/nodeConfigurations/configuration1
	"azurerm_automation_dsc_nodeconfiguration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/nodeConfigurations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/jobSchedules/10000000-1001-1001-1001-000000000001
	"azurerm_automation_job_schedule": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/modules/module1
	"azurerm_automation_module": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/modules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/Get-AzureVMTutorial
	"azurerm_automation_runbook": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/runbooks/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/schedules/schedule1
	"azurerm_automation_schedule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_bool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_int": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_string": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webHooks/TestRunbook_webhook
	"azurerm_automation_webhook": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),

	// botservice
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/AlexaChannel
	"azurerm_bot_channel_alexa": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/AlexaChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/DirectLineSpeechChannel
	"azurerm_bot_channel_direct_line_speech": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/DirectLineSpeechChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/DirectlineChannel
	"azurerm_bot_channel_directline": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/DirectlineChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/EmailChannel
	"azurerm_bot_channel_email": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/EmailChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/FacebookChannel
	"azurerm_bot_channel_facebook": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/FacebookChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/LineChannel
	"azurerm_bot_channel_line": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/LineChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/MsTeamsChannel
	"azurerm_bot_channel_ms_teams": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/MsTeamsChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/SlackChannel
	"azurerm_bot_channel_slack": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/SlackChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/SmsChannel
	"azurerm_bot_channel_sms": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/SmsChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/WebChatChannel
	"azurerm_bot_channel_web_chat": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/WebChatChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
	"azurerm_bot_channels_registration": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/connections/example
	"azurerm_bot_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/connections/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.BotService/botServices/botService1
	"azurerm_bot_service_azure_bot": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
	"azurerm_bot_web_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .external_name }}"),

	// communication
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/CommunicationServices/communicationService1
	"azurerm_communication_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Communication/CommunicationServices/{{ .external_name }}"),

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

	// confidentialledger
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-group/providers/Microsoft.ConfidentialLedger/ledgers/example-ledger
	"azurerm_confidential_ledger": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ConfidentialLedger/ledgers/{{ .external_name }}"),

	// consumption
	//
	// /providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/budgets/budget1
	"azurerm_consumption_budget_management_group": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Consumption/budgets/resourceGroup1
	"azurerm_consumption_budget_resource_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Consumption/budgets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/budgets/subscription1
	"azurerm_consumption_budget_subscription": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Consumption/budgets/{{ .external_name }}"),

	// costmanagement
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.CostManagement/exports/export1
	"azurerm_resource_group_cost_management_export": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CostManagement/exports/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/exports/export1
	"azurerm_subscription_cost_management_export": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.CostManagement/exports/{{ .external_name }}"),

	// customproviders
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.CustomProviders/resourceProviders/example
	"azurerm_custom_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CustomProviders/resourceProviders/{{ .external_name }}"),
}
