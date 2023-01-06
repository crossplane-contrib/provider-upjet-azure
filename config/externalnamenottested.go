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
	// TODO: TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_aadb2c": config.IdentifierFromProvider,
	// API Management Facebook Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/facebook
	// TODO: TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_facebook": config.IdentifierFromProvider,
	// API Management Google Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/google
	// TODO: TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_google": config.IdentifierFromProvider,
	// API Management Microsoft Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/microsoft
	// TODO: TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_microsoft": config.IdentifierFromProvider,
	// API Management Twitter Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/twitter
	// TODO: TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
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
}
