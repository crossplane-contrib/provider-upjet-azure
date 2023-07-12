package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameConfigs is a map of external name configurations for the whole
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
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
	// API Management Diagnostics can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/diagnostics/applicationinsights
	"azurerm_api_management_diagnostic": config.TemplatedStringAsIdentifier("identifier", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/diagnostics/{{ .external_name }}"),
	// API Management Email Templates can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/templates/template1
	"azurerm_api_management_email_template": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/templates/{{ .parameters.template_name }}"),
	// API Management Gateways can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1
	"azurerm_api_management_gateway": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/gateways/{{ .external_name }}"),
	// API Management AAD Identity Provider can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/identityProviders/aad
	"azurerm_api_management_identity_provider_aad": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/identityProviders/aad"),
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
	"azurerm_api_management_product_api": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/products/{{ .parameters.product_id }}/apis/{{ .parameters.api_name }}"),
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
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/schemas/schema1
	"azurerm_api_management_global_schema": config.TemplatedStringAsIdentifier("schema_id", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ApiManagement/service/{{ .parameters.api_management_name }}/schemas/{{ .external_name }}"),

	// authorization
	"azurerm_resource_group_policy_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Authorization/policyAssignments/{{ .external_name }}"),
	"azurerm_role_assignment":                  config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Authorization/locks/lock1
	"azurerm_management_lock": config.IdentifierFromProvider,
	// {resource}/providers/Microsoft.Authorization/policyAssignments/assignment1
	"azurerm_resource_policy_assignment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_resource_policy_exemption": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}|{scope}
	"azurerm_role_definition": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-000000000000/providers/Microsoft.Authorization/policyAssignments/assignment1
	"azurerm_subscription_policy_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policyAssignments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-000000000000/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_subscription_policy_exemption": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),

	// automation
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_datetime": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/variables/{{ .external_name }}"),

	// azurestackhci
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AzureStackHCI/clusters/cluster1
	"azurerm_stack_hci_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AzureStackHCI/clusters/{{ .external_name }}"),

	// base group
	"azurerm_subscription":                   config.TemplatedStringAsIdentifier("alias", "/providers/Microsoft.Subscription/aliases/{{ .external_name }}"),
	"azurerm_resource_provider_registration": config.IdentifierFromProvider,
	"azurerm_resource_group":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .external_name }}"),

	// botservice
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/AlexaChannel
	"azurerm_bot_channel_alexa": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/AlexaChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/DirectlineChannel
	"azurerm_bot_channel_directline": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/DirectlineChannel"),
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
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example
	"azurerm_bot_web_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .external_name }}"),

	// cognitiveservices
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
	"azurerm_cognitive_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CognitiveServices/accounts/{{ .external_name }}"),

	// communication
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Communication/CommunicationServices/communicationService1
	"azurerm_communication_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Communication/communicationServices/{{ .external_name }}"),

	// compute
	"azurerm_disk_encryption_set":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/diskEncryptionSets/{{ .external_name }}"),
	"azurerm_linux_virtual_machine":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachines/{{ .external_name }}"),
	"azurerm_windows_virtual_machine":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachines/{{ .external_name }}"),
	"azurerm_linux_virtual_machine_scale_set":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .external_name }}"),
	"azurerm_windows_virtual_machine_scale_set":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .external_name }}"),
	"azurerm_availability_set":                       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/availabilitySets/{{ .external_name }}"),
	"azurerm_disk_access":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/diskAccesses/{{ .external_name }}"),
	"azurerm_image":                                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/images/{{ .external_name }}"),
	"azurerm_managed_disk":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/disks/{{ .external_name }}"),
	"azurerm_orchestrated_virtual_machine_scale_set": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/virtualMachineScaleSets/{{ .external_name }}"),
	"azurerm_proximity_placement_group":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/proximityPlacementGroups/{{ .external_name }}"),
	"azurerm_shared_image_gallery":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .external_name }}"),
	"azurerm_snapshot":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/snapshots/{{ .external_name }}"),
	// Arguments make up the ID.
	"azurerm_marketplace_agreement": config.IdentifierFromProvider,
	"azurerm_dedicated_host":        config.TemplatedStringAsIdentifier("name", "{{ .parameters.dedicated_host_group_id }}/hosts/{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1
	"azurerm_gallery_application": config.TemplatedStringAsIdentifier("name", "{{ .parameters.gallery_id }}/applications/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/galleries/gallery1/applications/galleryApplication1/versions/galleryApplicationVersion1
	"azurerm_gallery_application_version": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1/capacityReservations/capacityReservation1
	"azurerm_capacity_reservation": config.TemplatedStringAsIdentifier("name", "{{ .parameters.capacity_reservation_group_id }}/capacityReservations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/capacityReservationGroups/capacityReservationGroup1
	"azurerm_capacity_reservation_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/capacityReservationGroups/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1
	"azurerm_shared_image": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .parameters.gallery_name }}/images/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/SshPublicKeys/mySshPublicKeyName1
	"azurerm_ssh_public_key": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/sshPublicKeys/{{ .external_name }}"),
	// Disk SAS Token can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/disks/manageddisk1
	"azurerm_managed_disk_sas_token": config.IdentifierFromProvider,

	// cdn
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1/endpoints/myendpoint1
	"azurerm_cdn_endpoint": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cdn/profiles/{{ .parameters.profile_name }}/endpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1
	"azurerm_cdn_profile": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cdn/profiles/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1
	"azurerm_cdn_frontdoor_profile": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cdn/profiles/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1
	"azurerm_cdn_frontdoor_origin_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_profile_id }}/originGroups/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1/origins/origin1
	"azurerm_cdn_frontdoor_origin": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_origin_group_id }}/origins/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/customDomains/customDomain1
	"azurerm_cdn_frontdoor_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_profile_id }}/customDomains/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/associations/assoc1
	"azurerm_cdn_frontdoor_custom_domain_association": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/afdEndpoints/endpoint1
	"azurerm_cdn_frontdoor_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_profile_id }}/afdEndpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/afdEndpoints/endpoint1/routes/route1
	"azurerm_cdn_frontdoor_route": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_endpoint_id }}/routes/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1
	"azurerm_cdn_frontdoor_rule_set": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_profile_id }}/ruleSets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1/rules/rule1
	"azurerm_cdn_frontdoor_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_frontdoor_rule_set_id }}/rules/{{ .external_name }}"),

	// cosmosdb
	"azurerm_cosmosdb_sql_database":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .external_name }}"),
	"azurerm_cosmosdb_sql_container": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .external_name }}"),
	// We switched to IdentifierFromProvider configuration because of the problem in this issue: https://github.com/upbound/upjet/issues/32
	"azurerm_cosmosdb_sql_role_assignment": config.IdentifierFromProvider,
	// We switched to IdentifierFromProvider configuration because of the problem in this issue: https://github.com/upbound/upjet/issues/32
	"azurerm_cosmosdb_sql_role_definition":  config.IdentifierFromProvider,
	"azurerm_cosmosdb_mongo_database":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .external_name }}"),
	"azurerm_cosmosdb_mongo_collection":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .parameters.database_name }}/collections/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_keyspace":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/cassandraKeyspaces/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_table":      config.TemplatedStringAsIdentifier("name", "{{ .parameters.cassandra_keyspace_id }}/tables/{{ .external_name }}"),
	"azurerm_cosmosdb_gremlin_database":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .external_name }}"),
	"azurerm_cosmosdb_gremlin_graph":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .parameters.database_name }}/graphs/{{ .external_name }}"),
	"azurerm_cosmosdb_sql_stored_procedure": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .parameters.container_name }}/storedProcedures/{{ .external_name }}"),
	// Contains only container_id as parameter which includes other information like resource group name etc.
	"azurerm_cosmosdb_sql_function": config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/userDefinedFunctions/{{ .external_name }}"),
	"azurerm_cosmosdb_table":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/tables/{{ .external_name }}"),
	"azurerm_cosmosdb_account":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .external_name }}"),
	// NOTE: "azurerm_cosmosdb_notebook_workspace" resource removed because this is not available in Azure portal.
	// Please see the note https://learn.microsoft.com/en-us/azure/cosmos-db/sql/enable-notebooks
	// "azurerm_cosmosdb_notebook_workspace":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/notebookWorkspaces/{{ .external_name }}"),
	"azurerm_cosmosdb_sql_trigger":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/triggers/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_cluster":    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/cassandraClusters/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_datacenter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cassandra_cluster_id }}/dataCenters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/services/SqlDedicatedGateway
	"azurerm_cosmosdb_sql_dedicated_gateway": config.TemplatedStringAsIdentifier("", "{{ .parameters.cosmosdb_account_id }}/services/SqlDedicatedGateway"),

	// datashare
	"azurerm_data_share_account":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataShare/accounts/{{ .external_name }}"),
	"azurerm_data_share":                        config.TemplatedStringAsIdentifier("name", "{{ .parameters.account_id }}/shares/{{ .external_name }}"),
	"azurerm_data_share_dataset_blob_storage":   config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_share_id }}/dataSets/{{ .external_name }}"),
	"azurerm_data_share_dataset_data_lake_gen2": config.TemplatedStringAsIdentifier("name", "{{ .parameters.share_id }}/dataSets/{{ .external_name }}"),
	"azurerm_data_share_dataset_kusto_cluster":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.share_id }}/dataSets/{{ .external_name }}"),
	"azurerm_data_share_dataset_kusto_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.share_id }}/dataSets/{{ .external_name }}"),

	// dataprotection
	"azurerm_data_protection_backup_vault":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataProtection/backupVaults/{{ .external_name }}"),
	"azurerm_data_protection_backup_policy_blob_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupPolicies/{{ .external_name }}"),
	"azurerm_data_protection_backup_policy_disk":         config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupPolicies/{{ .external_name }}"),
	"azurerm_data_protection_backup_policy_postgresql":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataProtection/backupVaults/{{ .parameters.vault_name }}/backupPolicies/{{ .external_name }}"),

	// devices
	"azurerm_iothub":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/iotHubs/{{ .external_name }}"),
	"azurerm_iothub_consumer_group":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/iotHubs/{{ .parameters.iothub_name }}/eventHubEndpoints/{{ .parameters.eventhub_endpoint_name }}/consumerGroups/{{ .external_name }}"),
	"azurerm_iothub_dps":                        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .external_name }}"),
	"azurerm_iothub_dps_certificate":            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .parameters.iothub_dps_name }}/certificates/{{ .external_name }}"),
	"azurerm_iothub_dps_shared_access_policy":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/provisioningServices/{{ .parameters.iothub_dps_name }}/keys/{{ .external_name }}"),
	"azurerm_iothub_shared_access_policy":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/iotHubs/{{ .parameters.iothub_name }}/iotHubKeys/{{ .external_name }}"),
	"azurerm_iothub_endpoint_storage_container": config.TemplatedStringAsIdentifier("name", "{{ .parameters.iothub_id }}/endpoints/{{ .external_name }}"),
	"azurerm_iothub_fallback_route":             config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/iotHubs/{{ .parameters.iothub_name }}/fallbackRoute/default"),
	"azurerm_iothub_endpoint_eventhub":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.iothub_id }}/endpoints/{{ .external_name }}"),
	"azurerm_iothub_endpoint_servicebus_queue":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.iothub_id }}/endpoints/{{ .external_name }}"),
	"azurerm_iothub_endpoint_servicebus_topic":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.iothub_id }}/endpoints/{{ .external_name }}"),
	"azurerm_iothub_route":                      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/iotHubs/{{ .parameters.iothub_name }}/routes/{{ .external_name }}"),
	"azurerm_iothub_enrichment":                 config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Devices/iotHubs/example/certificates/example
	"azurerm_iothub_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Devices/iotHubs/{{ .parameters.iothub_name }}/certificates/{{ .external_name }}"),

	// eventhub
	"azurerm_eventhub_namespace":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .external_name }}"),
	"azurerm_eventhub":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .external_name }}"),
	"azurerm_eventhub_consumer_group":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/consumerGroups/{{ .external_name }}"),
	"azurerm_eventhub_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/authorizationRules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/schemaGroups/group1
	"azurerm_eventhub_namespace_schema_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/schemaGroups/{{ .external_name }}"),

	// iotcentral
	//
	// The IoT Central Application can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.IoTCentral/iotApps/app1
	"azurerm_iotcentral_application": config.IdentifierFromProvider,

	// keyvault
	"azurerm_key_vault":                                              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.KeyVault/vaults/{{ .external_name }}"),
	"azurerm_key_vault_secret":                                       keyVaultURLIDConf("secrets"),
	"azurerm_key_vault_key":                                          keyVaultURLIDConf("keys"),
	"azurerm_key_vault_certificate":                                  keyVaultURLIDConf("certificates"),
	"azurerm_key_vault_certificate_issuer":                           keyVaultURLIDWithoutVersionConfFn("certificates/issuers"),
	"azurerm_key_vault_managed_storage_account":                      keyVaultURLIDWithoutVersionConfFn("storage"),
	"azurerm_key_vault_managed_storage_account_sas_token_definition": config.TemplatedStringAsIdentifier("name", "{{ .parameters.managed_storage_account_id }}/sas/{{ .external_name }}"),
	"azurerm_key_vault_managed_hardware_security_module":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.KeyVault/managedHSMs/{{ .external_name }}"),
	"azurerm_key_vault_access_policy":                                keyVaultAccessPolicy(),
	// Key Vault Certificate Contacts can be imported using the resource id
	// https://example-keyvault.vault.azure.net/certificates/contacts
	"azurerm_key_vault_certificate_contacts": config.IdentifierFromProvider,

	// kusto
	"azurerm_kusto_cluster":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/clusters/{{ .external_name }}"),
	"azurerm_kusto_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/clusters/{{ .parameters.cluster_name }}/databases/{{ .external_name }}"),
	// Kusto Attached Database Configurations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/AttachedDatabaseConfigurations/configuration1
	"azurerm_kusto_attached_database_configuration": config.IdentifierFromProvider,
	// Data Explorer Cluster Principal Assignments can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/PrincipalAssignments/assignment1
	"azurerm_kusto_cluster_principal_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/clusters/{{ .parameters.cluster_name }}/principalAssignments/{{ .external_name }}"),
	// Kusto Database Principal Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/PrincipalAssignments/assignment1
	"azurerm_kusto_database_principal_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/clusters/{{ .parameters.cluster_name }}/databases/{{ .parameters.database_name }}/principalAssignments/{{ .external_name }}"),
	// Kusto Event Grid Data Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/dataConnection1
	"azurerm_kusto_eventgrid_data_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/clusters/{{ .parameters.cluster_name }}/databases/{{ .parameters.database_name }}/dataConnections/{{ .external_name }}"),
	// Kusto EventHub Data Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/eventHubConnection1
	"azurerm_kusto_eventhub_data_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/clusters/{{ .parameters.cluster_name }}/databases/{{ .parameters.database_name }}/dataConnections/{{ .external_name }}"),
	// Kusto IotHub Data Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/dataConnection1
	"azurerm_kusto_iothub_data_connection": config.IdentifierFromProvider,

	// linux
	//
	// AppServices can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/serverfarms/farm1
	"azurerm_service_plan": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/serverfarms/{{ .external_name }}"),
	// Linux Function Apps can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_linux_function_app": config.IdentifierFromProvider,
	// A Linux Function App Slot can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_linux_function_app_slot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.function_app_id }}/slots/{{ .external_name }}"),
	// Linux Web Apps can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_linux_web_app_slot": config.IdentifierFromProvider,

	// containerservice
	"azurerm_kubernetes_cluster":           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerService/managedClusters/{{ .external_name }}"),
	"azurerm_kubernetes_cluster_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .parameters.kubernetes_cluster_id }}/agentPools/{{ .external_name }}"),

	// containerregistry
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/mygroup1/providers/Microsoft.ContainerRegistry/registries/myregistry1
	"azurerm_container_registry":            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerRegistry/registries/{{ .external_name }}"),
	"azurerm_container_registry_agent_pool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerRegistry/registries/{{ .parameters.container_registry_name }}/agentPools/{{ .external_name }}"),
	"azurerm_container_registry_scope_map":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerRegistry/registries/{{ .parameters.container_registry_name }}/scopeMaps/{{ .external_name }}"),
	"azurerm_container_registry_token":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerRegistry/registries/{{ .parameters.container_registry_name }}/tokens/{{ .external_name }}"),
	"azurerm_container_registry_webhook":    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerRegistry/registries/{{ .parameters.container_registry_name }}/webHooks/{{ .external_name }}"),
	"azurerm_container_connected_registry":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerRegistry/registries/{{ .parameters.container_registry_name }}/connectedRegistries/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.ContainerRegistry/registries/registry1/tokens/token1/passwords/password
	"azurerm_container_registry_token_password": config.IdentifierFromProvider,

	// operationalinsights
	"azurerm_log_analytics_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/workspaces/{{ .external_name }}"),
	// Log Analytics Data Export Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataExports/dataExport1
	"azurerm_log_analytics_data_export_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workspace_resource_id }}/dataExports/{{ .external_name }}"),
	// Log Analytics Windows Event DataSources can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataSources/datasource1
	"azurerm_log_analytics_datasource_windows_event": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/workspaces/{{ .parameters.workspace_name }}/dataSources/{{ .external_name }}"),
	// Log Analytics Windows Performance Counter DataSources can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataSources/datasource1
	"azurerm_log_analytics_datasource_windows_performance_counter": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/workspaces/{{ .parameters.workspace_name }}/dataSources/{{ .external_name }}"),
	// Log Analytics Workspaces can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/linkedServices/Automation
	"azurerm_log_analytics_linked_service": config.TemplatedStringAsIdentifier("", "{{ .parameters.workspace_id }}/linkedServices/{{ .external_name }}"),
	// Log Analytics Linked Storage Accounts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/linkedStorageAccounts/{dataSourceType}
	"azurerm_log_analytics_linked_storage_account": config.IdentifierFromProvider,
	// Log Analytics Saved Searches can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/savedSearches/search1
	"azurerm_log_analytics_saved_search": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/savedSearches/{{ .external_name }}"),
	// Log Analytics Solutions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationsManagement/solutions/solution1
	"azurerm_log_analytics_solution": config.IdentifierFromProvider,

	// insights
	"azurerm_application_insights": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/components/{{ .external_name }}"),
	// Application Insights API keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/instance1/apiKeys/00000000-0000-0000-0000-000000000000
	"azurerm_application_insights_api_key":        config.IdentifierFromProvider,
	"azurerm_monitor_action_group":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/actionGroups/{{ .external_name }}"),
	"azurerm_monitor_metric_alert":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/metricAlerts/{{ .external_name }}"),
	"azurerm_monitor_private_link_scope":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/privateLinkScopes/{{ .external_name }}"),
	"azurerm_monitor_private_link_scoped_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/privateLinkScopes/{{ .parameters.scope_name }}/scopedResources/{{ .external_name }}"),
	// Activity log alerts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/activityLogAlerts/myalertname
	"azurerm_monitor_activity_log_alert": config.IdentifierFromProvider,
	// AutoScale Setting can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/autoScaleSettings/setting1
	"azurerm_monitor_autoscale_setting": config.IdentifierFromProvider,
	// Scheduled Query Rule Alerts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
	"azurerm_monitor_scheduled_query_rules_alert": config.IdentifierFromProvider,
	// Scheduled Query Rule Log can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
	"azurerm_monitor_scheduled_query_rules_log": config.IdentifierFromProvider,
	// Application Insights Analytics Items can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/mycomponent1/analyticsItems/11111111-1111-1111-1111-111111111111
	"azurerm_application_insights_analytics_item": config.IdentifierFromProvider,
	// Application Insights Smart Detection Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/mycomponent1/smartDetectionRule/myrule1
	"azurerm_application_insights_smart_detection_rule": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/webTests/my_test
	"azurerm_application_insights_web_test": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/workbooks/resource1
	"azurerm_application_insights_workbook": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/workbookTemplates/resource1
	"azurerm_application_insights_workbook_template": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Insights/webTests/appinsightswebtest
	"azurerm_application_insights_standard_web_test": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/webTests/{{ .external_name }}"),
	// logic
	"azurerm_integration_service_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationServiceEnvironments/{{ .external_name }}"),

	// management
	// /providers/Microsoft.Management/managementGroups/group1
	"azurerm_management_group": config.TemplatedStringAsIdentifier("name", "/providers/Microsoft.Management/managementGroups/{{ .external_name }}"),
	// /managementGroup/MyManagementGroup/subscription/12345678-1234-1234-1234-123456789012
	"azurerm_management_group_subscription_association": managementGroupSubscriptionAssociation(),

	// mariadb
	"azurerm_mariadb_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .external_name }}"),
	"azurerm_mariadb_database":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	"azurerm_mariadb_firewall_rule":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_mariadb_virtual_network_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .external_name }}"),
	"azurerm_mariadb_configuration":        config.IdentifierFromProvider,

	// media
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Media/mediaServices/account1
	"azurerm_media_services_account":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .external_name }}"),
	"azurerm_media_asset":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/assets/{{ .external_name }}"),
	"azurerm_media_live_event":         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/liveEvents/{{ .external_name }}"),
	"azurerm_media_live_event_output":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.live_event_id }}/liveOutputs/{{ .external_name }}"),
	"azurerm_media_streaming_endpoint": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/streamingEndpoints/{{ .external_name }}"),
	"azurerm_media_streaming_locator":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/streamingLocators/{{ .external_name }}"),
	"azurerm_media_streaming_policy":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/streamingPolicies/{{ .external_name }}"),
	"azurerm_media_transform":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/transforms/{{ .external_name }}"),
	// Asset Filters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaServices/account1/assets/asset1/assetFilters/filter1
	"azurerm_media_asset_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.asset_id }}/assetFilters/{{ .external_name }}"),
	// Content Key Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaServices/account1/contentKeyPolicies/policy1
	"azurerm_media_content_key_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/contentKeyPolicies/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaServices/account1/accountFilters/filter1
	"azurerm_media_services_account_filter": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/accountFilters/{{ .external_name }}"),

	// mixedreality
	"azurerm_spatial_anchors_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{{ .external_name }}"),

	// dbformysql
	"azurerm_mysql_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .external_name }}"),
	"azurerm_mysql_configuration":        config.IdentifierFromProvider,
	"azurerm_mysql_firewall_rule":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_mysql_virtual_network_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .external_name }}"),
	// A MySQL Active Directory Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforMySQL/servers/myserver/administrators/activeDirectory
	"azurerm_mysql_active_directory_administrator": config.IdentifierFromProvider,
	// MySQL Database's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1/databases/database1
	"azurerm_mysql_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	// dbformysql flexible variants
	"azurerm_mysql_flexible_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .external_name }}"),
	"azurerm_mysql_flexible_database":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	"azurerm_mysql_flexible_server_configuration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/configurations/{{ .external_name }}"),
	"azurerm_mysql_flexible_server_firewall_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),

	// digitaltwins
	//
	// Digital Twins instances can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1
	"azurerm_digital_twins_instance": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{ .external_name }}"),

	// disk
	//
	// Disk Pools can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StoragePool/diskPools/diskPool1
	"azurerm_disk_pool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StoragePool/diskPools/{{ .external_name }}"),

	// netapp
	"azurerm_netapp_account":         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .external_name }}"),
	"azurerm_netapp_pool":            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .parameters.account_name }}/capacityPools/{{ .external_name }}"),
	"azurerm_netapp_snapshot_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .parameters.account_name }}/snapshotPolicies/{{ .external_name }}"),
	"azurerm_netapp_snapshot":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .parameters.account_name }}/capacityPools/{{ .parameters.pool_name }}/volumes/{{ .parameters.volume_name }}/snapshots/{{ .external_name }}"),
	"azurerm_netapp_volume":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NetApp/netAppAccounts/{{ .parameters.account_name }}/capacityPools/{{ .parameters.pool_name }}/volumes/{{ .external_name }}"),

	// network
	"azurerm_virtual_network":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .external_name }}"),
	"azurerm_ip_group":                                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/ipGroups/{{ .external_name }}"),
	"azurerm_subnet":                                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .parameters.virtual_network_name }}/subnets/{{ .external_name }}"),
	"azurerm_subnet_nat_gateway_association":            config.IdentifierFromProvider,
	"azurerm_subnet_network_security_group_association": config.IdentifierFromProvider,
	"azurerm_subnet_service_endpoint_storage_policy":    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/serviceEndpointPolicies/{{ .external_name }}"),
	"azurerm_subnet_route_table_association":            config.IdentifierFromProvider,
	"azurerm_network_interface":                         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkInterfaces/{{ .external_name }}"),
	"azurerm_lb":                                        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/loadBalancers/{{ .external_name }}"),
	"azurerm_lb_backend_address_pool":                   config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/backendAddressPools/{{ .external_name }}"),
	"azurerm_lb_nat_pool":                               config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/inboundNatPools/{{ .external_name }}"),
	"azurerm_lb_nat_rule":                               config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/inboundNatRules/{{ .external_name }}"),
	"azurerm_lb_backend_address_pool_address":           config.TemplatedStringAsIdentifier("name", "{{ .parameters.backend_address_pool_id }}/addresses/{{ .external_name }}"),
	"azurerm_lb_outbound_rule":                          config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/outboundRules/{{ .external_name }}"),
	"azurerm_lb_probe":                                  config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/probes/{{ .external_name }}"),
	"azurerm_lb_rule":                                   config.TemplatedStringAsIdentifier("name", "{{ .parameters.loadbalancer_id }}/loadBalancingRules/{{ .external_name }}"),
	"azurerm_local_network_gateway":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/localNetworkGateways/{{ .external_name }}"),
	"azurerm_nat_gateway":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/natGateways/{{ .external_name }}"),
	"azurerm_network_watcher":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkWatchers/{{ .external_name }}"),
	"azurerm_network_watcher_flow_log":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkWatchers/{{ .parameters.network_watcher_name }}/flowLogs/{{ .external_name }}"),
	"azurerm_network_connection_monitor":                config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_watcher_id }}/backendAddressPools/{{ .connectionMonitors }}"),
	"azurerm_network_ddos_protection_plan":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/ddosProtectionPlans/{{ .external_name }}"),
	"azurerm_application_security_group":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/applicationSecurityGroups/{{ .external_name }}"),
	"azurerm_network_security_group":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkSecurityGroups/{{ .external_name }}"),
	"azurerm_network_security_rule":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkSecurityGroups/{{ .parameters.network_security_group_name }}/securityRules/{{ .external_name }}"),
	"azurerm_virtual_network_gateway":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworkGateways/{{ .external_name }}"),
	"azurerm_virtual_network_peering":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .parameters.virtual_network_name }}/virtualNetworkPeerings/{{ .external_name }}"),
	"azurerm_virtual_network_gateway_connection":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/connections/{{ .external_name }}"),
	"azurerm_virtual_wan":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualWans/{{ .external_name }}"),
	"azurerm_virtual_hub":                               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualHubs/{{ .external_name }}"),
	"azurerm_public_ip":                                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/publicIPAddresses/{{ .external_name }}"),
	"azurerm_public_ip_prefix":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/publicIPPrefixes/{{ .external_name }}"),
	"azurerm_network_profile":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkProfiles/{{ .external_name }}"),
	"azurerm_dns_zone":                                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .external_name }}"),
	"azurerm_dns_a_record":                              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/A/{{ .external_name }}"),
	"azurerm_dns_aaaa_record":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/AAAA/{{ .external_name }}"),
	"azurerm_dns_caa_record":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/CAA/{{ .external_name }}"),
	"azurerm_dns_cname_record":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/CNAME/{{ .external_name }}"),
	"azurerm_dns_mx_record":                             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/MX/{{ .external_name }}"),
	"azurerm_dns_ns_record":                             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/NS/{{ .external_name }}"),
	"azurerm_dns_ptr_record":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/PTR/{{ .external_name }}"),
	"azurerm_dns_srv_record":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/SRV/{{ .external_name }}"),
	"azurerm_dns_txt_record":                            config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsZones/{{ .parameters.zone_name }}/TXT/{{ .external_name }}"),
	"azurerm_private_dns_zone":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .external_name }}"),
	"azurerm_private_dns_a_record":                      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/A/{{ .external_name }}"),
	"azurerm_private_dns_aaaa_record":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/AAAA/{{ .external_name }}"),
	"azurerm_private_dns_cname_record":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/CNAME/{{ .external_name }}"),
	"azurerm_private_dns_mx_record":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/MX/{{ .external_name }}"),
	"azurerm_private_dns_ptr_record":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/PTR/{{ .external_name }}"),
	"azurerm_private_dns_srv_record":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/SRV/{{ .external_name }}"),
	"azurerm_private_dns_txt_record":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.zone_name }}/TXT/{{ .external_name }}"),
	"azurerm_private_dns_zone_virtual_network_link":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.private_dns_zone_name }}/virtualNetworkLinks/{{ .external_name }}"),
	"azurerm_private_link_service":                      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateLinkServices/{{ .external_name }}"),
	"azurerm_private_endpoint":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateEndpoints/{{ .external_name }}"),
	"azurerm_network_packet_capture":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkWatchers/{{ .parameters.network_watcher_name }}/packetCaptures/{{ .external_name }}"),
	"azurerm_vpn_server_configuration":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/vpnServerConfigurations/{{ .external_name }}"),
	"azurerm_point_to_site_vpn_gateway":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/p2sVpnGateways/{{ .external_name }}"),
	"azurerm_express_route_circuit":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/expressRouteCircuits/{{ .external_name }}"),
	"azurerm_express_route_circuit_authorization":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/expressRouteCircuits/{{ .parameters.express_route_circuit_name }}/authorizations/{{ .external_name }}"),
	"azurerm_express_route_circuit_peering":             config.TemplatedStringAsIdentifier("peering_type", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/expressRouteCircuits/{{ .parameters.express_route_circuit_name }}/peerings/{{ .external_name }}"),
	"azurerm_express_route_port":                        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/expressRoutePorts/{{ .external_name }}"),
	"azurerm_express_route_gateway":                     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/expressRouteGateways/{{ .external_name }}"),
	"azurerm_express_route_circuit_connection":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.peering_id }}/connections/{{ .external_name }}"),
	"azurerm_express_route_connection":                  config.TemplatedStringAsIdentifier("name", "{{ .parameters.express_route_gateway_id }}/expressRouteConnections/{{ .external_name }}"),
	"azurerm_firewall":                                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/azureFirewalls/{{ .external_name }}"),
	"azurerm_firewall_application_rule_collection":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/azureFirewalls/{{ .parameters.azure_firewall_name }}/applicationRuleCollections/{{ .external_name }}"),
	"azurerm_firewall_nat_rule_collection":              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/azureFirewalls/{{ .parameters.azure_firewall_name }}/natRuleCollections/{{ .external_name }}"),
	"azurerm_firewall_network_rule_collection":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/azureFirewalls/{{ .parameters.azure_firewall_name }}/networkRuleCollections/{{ .external_name }}"),
	"azurerm_firewall_policy":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/firewallPolicies/{{ .external_name }}"),
	"azurerm_firewall_policy_rule_collection_group":     config.TemplatedStringAsIdentifier("name", "{{ .parameters.firewall_policy_id }}/ruleCollectionGroups/{{ .external_name }}"),
	"azurerm_frontdoor":                                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/frontDoors/{{ .external_name }}"),
	"azurerm_frontdoor_firewall_policy":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/{{ .external_name }}"),
	"azurerm_frontdoor_rules_engine":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/frontdoors/{{ .parameters.frontdoor_name }}/rulesengines/{{ .external_name }}"),
	"azurerm_frontdoor_custom_https_configuration":      config.IdentifierFromProvider,
	"azurerm_application_gateway":                       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/applicationGateways/{{ .external_name }}"),
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
	"azurerm_route_table":                              config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/routeTables/{{ .external_name }}"),
	// Traffic Manager Profiles can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/trafficManagerProfiles/mytrafficmanagerprofile1
	"azurerm_traffic_manager_profile": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/trafficManagerProfiles/{{ .external_name }}"),
	// VPN Gateways can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/vpnGateways/gateway1
	"azurerm_vpn_gateway": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/vpnGateways/{{ .external_name }}"),
	// VPN Gateway Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/conn1
	"azurerm_vpn_gateway_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vpn_gateway_id }}/vpnConnections/{{ .external_name }}"),
	// VPN Sites can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/vpnSites/site1
	"azurerm_vpn_site": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/vpnSites/{{ .external_name }}"),
	// Web Application Firewall Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/example-wafpolicy
	"azurerm_web_application_firewall_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Network/vpnServerConfigurations/serverConfiguration1/configurationPolicyGroups/configurationPolicyGroup1
	"azurerm_vpn_server_configuration_policy_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vpn_server_configuration_id }}/configurationPolicyGroups/{{ .external_name }}"),

	// notification
	"azurerm_notification_hub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NotificationHubs/namespaces/{{ .parameters.namespace_name }}/notificationHubs/{{ .external_name }}"),
	// Notification Hub Authorization Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1/notificationHubs/hub1/authorizationRules/rule1
	"azurerm_notification_hub_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NotificationHubs/namespaces/{{ .parameters.namespace_name }}/notificationHubs/{{ .parameters.notification_hub_name }}/authorizationRules/{{ .external_name }}"),
	// Notification Hub Namespaces can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.NotificationHubs/namespaces/namespace1
	"azurerm_notification_hub_namespace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NotificationHubs/namespaces/{{ .external_name }}"),

	// postgresql
	"azurerm_postgresql_server":                         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .external_name }}"),
	"azurerm_postgresql_flexible_server_configuration":  config.IdentifierFromProvider,
	"azurerm_postgresql_database":                       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	"azurerm_postgresql_active_directory_administrator": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}"),
	"azurerm_postgresql_flexible_server_database":       config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/databases/{{ .external_name }}"),
	"azurerm_postgresql_flexible_server_firewall_rule":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/firewallRules/{{ .external_name }}"),
	"azurerm_postgresql_firewall_rule":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_postgresql_flexible_server":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{ .external_name }}"),
	"azurerm_postgresql_virtual_network_rule":           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
	"azurerm_postgresql_server_key": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/configurations/backslash_quote
	"azurerm_postgresql_configuration": config.IdentifierFromProvider,

	// redis
	"azurerm_redis_cache":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redis/{{ .external_name }}"),
	"azurerm_redis_firewall_rule":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redis/{{ .parameters.redis_cache_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_redis_enterprise_cluster":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redisEnterprise/{{ .external_name }}"),
	"azurerm_redis_enterprise_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cluster_id }}/databases/{{ .external_name }}"),
	"azurerm_redis_linked_server":       config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redis/{{ .parameters.target_redis_cache_name }}/linkedServers/{{ .parameters.linked_redis_cache_name }}"),

	// resource
	"azurerm_resource_group_template_deployment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Resources/deployments/{{ .external_name }}"),

	// synapse
	//
	// Synapse Firewall Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.Synapse/workspaces/workspace1/firewallRules/rule1
	"azurerm_synapse_firewall_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/firewallRules/{{ .external_name }}"),
	// Synapse Azure Integration Runtimes can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/integrationRuntimes/IntegrationRuntime1
	"azurerm_synapse_integration_runtime_azure": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/integrationRuntimes/{{ .external_name }}"),
	// Synapse Self-hosted Integration Runtimes can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/integrationRuntimes/IntegrationRuntime1
	"azurerm_synapse_integration_runtime_self_hosted": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/integrationRuntimes/{{ .external_name }}"),
	// Synapse Linked Services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/linkedServices/linkedservice1
	"azurerm_synapse_linked_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/linkedServices/{{ .external_name }}"),
	// Synapse Managed Private Endpoint can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
	"azurerm_synapse_managed_private_endpoint": config.IdentifierFromProvider,
	// Synapse Private Link Hub can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/privateLinkHubs/privateLinkHub1
	"azurerm_synapse_private_link_hub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Synapse/privateLinkHubs/{{ .external_name }}"),
	// Synapse Spark Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/sparkPool1
	"azurerm_synapse_spark_pool": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/bigDataPools/{{ .external_name }}"),
	// Synapse SQL Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1
	"azurerm_synapse_sql_pool": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/sqlPools/{{ .external_name }}"),
	// Synapse SQL Pool Extended Auditing Policys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/extendedAuditingSettings/default
	"azurerm_synapse_sql_pool_extended_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_id }}/extendedAuditingSettings/default"),
	// Synapse SQL Pool Security Alert Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/securityAlertPolicies/default
	"azurerm_synapse_sql_pool_security_alert_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_id }}/securityAlertPolicies/default"),
	// Synapse SQL Pool Workload Classifiers can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/workloadGroups/workloadGroup1/workloadClassifiers/workloadClassifier1
	"azurerm_synapse_sql_pool_workload_classifier": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workload_group_id }}/workloadClassifiers/{{ .external_name }}"),
	// Synapse SQL Pool Workload Groups can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/workloadGroups/workloadGroup1
	"azurerm_synapse_sql_pool_workload_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.sql_pool_id }}/workloadGroups/{{ .external_name }}"),
	// Synapse Workspace can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
	"azurerm_synapse_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Synapse/workspaces/{{ .external_name }}"),
	// Synapse Workspace Azure AD Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/administrators/activeDirectory
	"azurerm_synapse_workspace_aad_admin": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/administrators/{{ .external_name }}"),
	// Synapse Workspace Extended Auditing Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/extendedAuditingSettings/default
	"azurerm_synapse_workspace_extended_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/extendedAuditingSettings/default"),
	// Synapse Workspace Security Alert Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/securityAlertPolicies/Default
	"azurerm_synapse_workspace_security_alert_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/securityAlertPolicies/Default"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/vulnerabilityAssessments/default
	"azurerm_synapse_workspace_vulnerability_assessment": config.IdentifierFromProvider,
	// Synapse Role Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1|000000000000
	"azurerm_synapse_role_assignment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlAdministrators/activeDirectory
	"azurerm_synapse_workspace_sql_aad_admin": config.IdentifierFromProvider,

	// security
	// We switched to IdentifierFromProvider configuration because of the problem in this issue: https://github.com/upbound/upjet/issues/32
	"azurerm_advanced_threat_protection": config.IdentifierFromProvider,
	// We switched to IdentifierFromProvider configuration because of the problem in this issue: https://github.com/upbound/upjet/issues/32
	"azurerm_iot_security_device_group": config.IdentifierFromProvider,
	"azurerm_iot_security_solution":     config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Security/iotSecuritySolutions/{{ .external_name }}"),

	// servicebus
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1
	"azurerm_servicebus_namespace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ServiceBus/namespaces/{{ .external_name }}"),

	// sql
	"azurerm_mssql_server":                                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .external_name }}"),
	"azurerm_mssql_database":                                        config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/databases/{{ .external_name }}"),
	"azurerm_mssql_failover_group":                                  config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/failoverGroups/{{ .external_name }}"),
	"azurerm_mssql_server_transparent_data_encryption":              config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/encryptionProtector/current"),
	"azurerm_mssql_virtual_network_rule":                            config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/virtualNetworkRules/{{ .external_name }}"),
	"azurerm_mssql_managed_instance":                                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/managedInstances/{{ .external_name }}"),
	"azurerm_mssql_managed_database":                                config.TemplatedStringAsIdentifier("name", "{{ .parameters.managed_instance_id }}/databases/{{ .external_name }}"),
	"azurerm_mssql_managed_instance_active_directory_administrator": config.TemplatedStringAsIdentifier("", "{{ .parameters.managed_instance_id }}/administrators/activeDirectory"),
	"azurerm_mssql_managed_instance_failover_group":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/locations/{{ .parameters.location }}/instanceFailoverGroups/{{ .external_name }}"),
	"azurerm_mssql_managed_instance_vulnerability_assessment":       config.TemplatedStringAsIdentifier("", "{{ .parameters.managed_instance_id }}/vulnerabilityAssessments/Default"),
	"azurerm_mssql_outbound_firewall_rule":                          config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/outboundFirewallRules/{{ .external_name }}"),
	"azurerm_mssql_server_dns_alias":                                config.TemplatedStringAsIdentifier("name", "{{ .parameters.mssql_server_id }}/dnsAliases/{{ .external_name }}"),
	// MS SQL Database Extended Auditing Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/databases/db1/extendedAuditingSettings/default
	"azurerm_mssql_database_extended_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.database_id }}/extendedAuditingSettings/default"),
	// MS SQL Server Security Alert Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/securityAlertPolicies/Default
	"azurerm_mssql_server_security_alert_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/securityAlertPolicies/default"),
	// SQL Firewall Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/firewallRules/rule1
	"azurerm_mssql_firewall_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/firewallRules/{{ .external_name }}"),
	// Database Vulnerability Assessment Rule Baseline can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/databases/mysqldatabase/vulnerabilityAssessments/Default/rules/VA2065/baselines/master
	"azurerm_mssql_database_vulnerability_assessment_rule_baseline": config.IdentifierFromProvider,
	// SQL Elastic Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/myelasticpoolname
	"azurerm_mssql_elasticpool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/elasticPools/{{ .external_name }}"),
	// Elastic Job Agents can be imported using the id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1
	"azurerm_mssql_job_agent": config.IdentifierFromProvider,
	// Elastic Job Credentials can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/credentials/credential1
	"azurerm_mssql_job_credential": config.TemplatedStringAsIdentifier("name", "{{ .parameters.job_agent_id }}/credentials/{{ .external_name }}"),
	// MS SQL Server Vulnerability Assessment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/vulnerabilityAssessments/Default
	"azurerm_mssql_server_vulnerability_assessment": config.IdentifierFromProvider,

	// storage
	"azurerm_hpc_cache_access_policy":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StorageCache/caches/{{ .hpc_cache_id }}/cacheAccessPolicies/{{ .external_name }}"),
	"azurerm_hpc_cache_blob_nfs_target": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StorageCache/caches/{{ .cacheName }}/storageTargets/{{ .external_name }}"),
	"azurerm_hpc_cache_blob_target":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StorageCache/caches/{{ .cacheName }}/storageTargets/{{ .external_name }}"),
	"azurerm_hpc_cache":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StorageCache/caches/{{ .external_name }}"),
	"azurerm_hpc_cache_nfs_target":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StorageCache/caches/{{ .cacheName }}/storageTargets/{{ .external_name }}"),
	"azurerm_storage_account":           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Storage/storageAccounts/{{ .external_name }}"),
	// This resource does not have its own identifier, but rather uses the name of storage account.
	// Following _Case 6_ from the **Adding a New Resource** guide
	"azurerm_storage_account_network_rules":     config.IdentifierFromProvider,
	"azurerm_storage_blob":                      config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.storage_account_name }}.blob.core.windows.net/{{ .parameters.storage_container_name }}/{{ .external_name }}"),
	"azurerm_storage_blob_inventory_policy":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Storage/storageAccounts/{{ .storage_account_id }}/inventoryPolicies/Default"),
	"azurerm_storage_container":                 config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.storage_account_name }}.blob.core.windows.net/{{ .external_name }}"),
	"azurerm_storage_data_lake_gen2_filesystem": storageDataLakeGen2Filesystem(),
	"azurerm_storage_encryption_scope":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Storage/storageAccounts/{{ .storage_account_id }}/encryptionScopes/{{ .external_name }}"),
	"azurerm_storage_management_policy":         config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Storage/storageAccounts/{{ .storage_account_id }}/managementPolicies/default"),
	// The id of this resource is a concatenation of 2 resource names, but in the terraform documentation
	// this reasource does not have a name so instead it concatenates destination and target storage account IDs
	"azurerm_storage_object_replication": config.IdentifierFromProvider,
	"azurerm_storage_queue":              config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.storage_account_name }}.queue.core.windows.net/{{ .external_name }}"),
	"azurerm_storage_share":              config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.storage_account_name }}.file.core.windows.net/{{ .external_name }}"),
	// Table ID comes with an unusual https format there the name attribute is not separated by "/",
	// but fits this remplate Tables('replace-with-table-name')
	"azurerm_storage_table": config.IdentifierFromProvider,

	// storagesync
	"azurerm_storage_sync": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StorageSync/storageSyncServices/{{ .external_name }}"),

	// streamanalytics
	"azurerm_stream_analytics_cluster":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/clusters/{{ .external_name }}"),
	"azurerm_stream_analytics_function_javascript_uda":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.job_name }}/functions/{{ .external_name }}"),
	"azurerm_stream_analytics_job":                      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .external_name }}"),
	"azurerm_stream_analytics_managed_private_endpoint": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/clusters/{{ .parameters.stream_analytics_cluster_name }}/privateEndpoints/{{ .external_name }}"),
	"azurerm_stream_analytics_output_function":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_id }}/outputs/{{ .external_name }}"),
	"azurerm_stream_analytics_output_synapse":           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	"azurerm_stream_analytics_output_table":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	// "azurerm_stream_analytics_function_javascript_udf": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.job_name }}/outputs/{{ .external_name }}"),
	// "azurerm_stream_analytics_output_cosmosdb":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_id }}/outputs/{{ .external_name }}"),
	// Does not work, tracked in https://github.com/upbound/official-providers/issues/368
	"azurerm_stream_analytics_output_powerbi": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	"azurerm_stream_analytics_output_blob":    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_id }}/outputs/{{ .external_name }}"),

	// timeseriesinsights
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/eventSources/example
	"azurerm_iot_time_series_insights_event_source_iothub": config.TemplatedStringAsIdentifier("name", "{{ .parameters.environment_id }}/eventSources/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
	"azurerm_iot_time_series_insights_gen2_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.TimeSeriesInsights/environments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example/referenceDataSets/example
	"azurerm_iot_time_series_insights_reference_data_set": config.TemplatedStringAsIdentifier("name", "{{ .parameters.time_series_insights_environment_id }}/referenceDataSets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/example
	"azurerm_iot_time_series_insights_standard_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.TimeSeriesInsights/environments/{{ .external_name }}"),

	// azurerm_policy_definition can be imported
	// azurerm_policy_definition.examplePolicy /subscriptions/<SUBSCRIPTION_ID>/providers/Microsoft.Authorization/policyDefinitions/<POLICY_NAME>
	"azurerm_policy_definition": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policyDefinitions/{{ .external_name }}"),

	// alertsmanagement
	//
	// Monitor Action Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
	"azurerm_monitor_action_rule_action_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AlertsManagement/actionRules/{{ .external_name }}"),
	// Monitor Action Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
	"azurerm_monitor_action_rule_suppression": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AlertsManagement/actionRules/{{ .external_name }}"),
	// Monitor Smart Detector Alert Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AlertsManagement/smartDetectorAlertRules/rule1
	"azurerm_monitor_smart_detector_alert_rule": config.IdentifierFromProvider,

	// appplatform
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1
	"azurerm_spring_cloud_active_deployment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/spring/myservice/apps/myapp
	"azurerm_spring_cloud_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/spring/{{ .parameters.service_name }}/apps/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
	"azurerm_spring_cloud_app_cosmosdb_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
	"azurerm_spring_cloud_app_mysql_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/spring/myservice/apps/myapp/bindings/bind1
	"azurerm_spring_cloud_app_redis_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/spring1/certificates/cert1
	"azurerm_spring_cloud_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/spring/{{ .parameters.service_name }}/certificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/domains/domain.com
	"azurerm_spring_cloud_custom_domain": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_java_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/spring/spring1
	"azurerm_spring_cloud_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/spring/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/storages/storage1
	"azurerm_spring_cloud_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/storages/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1/domains/domain1
	"azurerm_spring_cloud_api_portal_custom_domain": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/Spring/springcloud/apps/springcloudapp/deployments/deployment/providers/Microsoft.ServiceLinker/linkers/serviceconnector1
	"azurerm_spring_cloud_connection": config.IdentifierFromProvider,

	// analysisservices
	//
	// Analysis Services Server can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AnalysisServices/servers/server1
	"azurerm_analysis_services_server": config.IdentifierFromProvider,

	// automation
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1
	"azurerm_automation_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/credentials/credential1
	"azurerm_automation_credential": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/credentials/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/modules/module1
	"azurerm_automation_module": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/modules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_bool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/variables/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_int": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/variables/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tfex-example-rg/providers/Microsoft.Automation/automationAccounts/tfex-example-account/variables/tfex-example-var
	"azurerm_automation_variable_string": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/variables/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/Get-AzureVMTutorial
	"azurerm_automation_runbook": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_classic_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webHooks/TestRunbook_webhook
	"azurerm_automation_webhook": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/schedules/schedule1
	"azurerm_automation_schedule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),

	// recovery
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1/backupPolicies/policy1
	"azurerm_backup_policy_vm_workload": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/backupPolicies/{{ .external_name }}"),

	// confidentialledger
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-group/providers/Microsoft.ConfidentialLedger/ledgers/example-ledger
	"azurerm_confidential_ledger": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ConfidentialLedger/ledgers/{{ .external_name }}"),

	// consumption
	//
	// /providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/budgets/budget1
	"azurerm_consumption_budget_management_group": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Consumption/budgets/resourceGroup1
	"azurerm_consumption_budget_resource_group": config.IdentifierFromProvider,
	// "azurerm_consumption_budget_resource_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Consumption/budgets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/budgets/subscription1
	"azurerm_consumption_budget_subscription": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Consumption/budgets/{{ .external_name }}"),

	// costmanagement
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.CostManagement/exports/export1
	"azurerm_resource_group_cost_management_export": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CostManagement/exports/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CostManagement/exports/export1
	"azurerm_subscription_cost_management_export": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/scheduledActions/dailyanomalybyresourcegroup
	"azurerm_cost_anomaly_alert": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.CostManagement/scheduledActions/{{ .external_name }}"),

	// api
	//
	// API Management API Tags can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/tags/tag1
	"azurerm_api_management_api_tag": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}/tags/{{ .external_name }}"),
	// API Management Notification Recipient Users can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientUsers/userid1
	"azurerm_api_management_notification_recipient_user": config.IdentifierFromProvider,

	// logic
	//
	// Logic App Integration Account Batch Configurations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/batchConfigurations/batchConfiguration1
	"azurerm_logic_app_integration_account_batch_configuration": config.IdentifierFromProvider,
	// Logic App Integration Accounts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1
	"azurerm_logic_app_integration_account": config.IdentifierFromProvider,

	// streamanalytics
	//
	// Stream Analytics Outputs to Microsoft SQL Server Database can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
	"azurerm_stream_analytics_output_mssql": config.IdentifierFromProvider,
	// Stream Analytics Reference Input Blob's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
	"azurerm_stream_analytics_reference_input_blob": config.IdentifierFromProvider,
	// Stream Analytics Stream Input Blob's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_blob": config.IdentifierFromProvider,
	// Stream Analytics Stream Input EventHub's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_eventhub": config.IdentifierFromProvider,
	// Stream Analytics Stream Input IoTHub's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_iothub": config.IdentifierFromProvider,
	// Stream Analytics Output ServiceBus Queue's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
	"azurerm_stream_analytics_output_servicebus_queue": config.IdentifierFromProvider,
	// Stream Analytics Output ServiceBus Topic's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
	"azurerm_stream_analytics_output_servicebus_topic": config.IdentifierFromProvider,

	// subscription
	//
	// Policy Remediations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_subscription_policy_remediation": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.PolicyInsights/remediations/{{ .external_name }}"),

	// databoxedge
	//
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
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/accessConnectors/connector1
	"azurerm_databricks_access_connector": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Databricks/accessConnectors/{{ .external_name }}"),

	// purview
	//
	// Purview Accounts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Purview/accounts/account1
	"azurerm_purview_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Purview/accounts/{{ .external_name }}"),

	// powerbidedicated
	// PowerBI Embedded can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.PowerBIDedicated/capacities/capacity1
	"azurerm_powerbi_embedded": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.PowerBIDedicated/capacities/{{ .external_name }}"),

	// recoveryservices
	//
	// Recovery Services Vaults can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/vault1
	"azurerm_recovery_services_vault": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .external_name }}"),
	// Site Recovery Fabric can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name
	"azurerm_site_recovery_fabric": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationFabrics/{{ .external_name }}"),
	// VM Backup Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
	"azurerm_backup_policy_vm": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/backupPolicies/{{ .external_name }}"),
	// Azure File Share Backup Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupPolicies/policy1
	"azurerm_backup_policy_file_share": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/backupPolicies/{{ .external_name }}"),
	// Backup Storage Account Containers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/backupFabrics/Azure/protectionContainers/StorageContainer;storage;storage-rg-name;storage-account
	"azurerm_backup_container_storage_account": config.IdentifierFromProvider,
	// Azure Backup Protected File Shares can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupFabrics/Azure/protectionContainers/StorageContainer;storage;group2;example-storage-account/protectedItems/AzureFileShare;3f6e3108a45793581bcbd1c61c87a3b2ceeb4ff4bc02a95ce9d1022b23722935
	"azurerm_backup_protected_file_share": config.IdentifierFromProvider,
	// Recovery Services Protected VMs can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.RecoveryServices/vaults/example-recovery-vault/backupFabrics/Azure/protectionContainers/iaasvmcontainer;iaasvmcontainerv2;group1;vm1/protectedItems/vm;iaasvmcontainerv2;group1;vm1
	"azurerm_backup_protected_vm": config.IdentifierFromProvider,
	// Site Recovery Protection Containers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric-name/replicationProtectionContainers/protection-container-name
	"azurerm_site_recovery_protection_container": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationFabrics/{{ .parameters.recovery_fabric_name }}/replicationProtectionContainers/{{ .external_name }}"),
	// Site Recovery Protection Container Mappings can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/fabric1/replicationProtectionContainers/container1/replicationProtectionContainerMappings/mapping1
	"azurerm_site_recovery_protection_container_mapping": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationFabrics/{{ .parameters.recovery_fabric_name }}/replicationProtectionContainers/{{ .parameters.recovery_source_protection_container_name }}/replicationProtectionContainerMappings/{{ .external_name }}"),
	// Site Recovery Replication Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationPolicies/policy-name
	"azurerm_site_recovery_replication_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.RecoveryServices/vaults/{{ .parameters.recovery_vault_name }}/replicationPolicies/{{ .external_name }}"),

	// relay
	//
	// Azure Relay Namespace's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1
	"azurerm_relay_namespace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Relay/namespaces/{{ .external_name }}"),
	// Azure Relay Namespace Authorization Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Relay/namespaces/namespace1/authorizationRules/rule1
	"azurerm_relay_namespace_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Relay/namespaces/{{ .parameters.namespace_name }}/authorizationRules/{{ .external_name }}"),
	// Relay Hybrid Connection's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Relay/namespaces/relay1/hybridConnections/hconn1
	"azurerm_relay_hybrid_connection": config.IdentifierFromProvider,
	// Azure Relay Hybrid Connection Authorization Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Relay/namespaces/namespace1/hybridConnections/connection1/authorizationRules/rule1
	"azurerm_relay_hybrid_connection_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Relay/namespaces/{{ .parameters.relay_namespace_name }}/hybridConnections/{{ .parameters.hybrid_connection_name }}/authorizationRules/{{ .external_name }}"),

	// resource
	//
	// Policy Remediations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_resource_policy_remediation": config.IdentifierFromProvider,

	// datafactory
	//
	// Data Factory ODBC Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_odbc": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example
	"azurerm_data_factory": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataFactory/factories/{{ .external_name }}"),
	// Data Factory Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_custom_dataset": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Data Flow can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/dataflows/example
	"azurerm_data_factory_data_flow": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/dataflows/{{ .external_name }}"),
	// Data Factory Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_azure_blob": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Binary Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_binary": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_cosmosdb_sqlapi": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_delimited_text": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_http": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_json": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory MySQL Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_mysql": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_parquet": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory PostgreSQL Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_postgresql": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Snowflake Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_snowflake": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Azure Integration Runtimes can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
	"azurerm_data_factory_integration_runtime_azure": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/integrationruntimes/{{ .external_name }}"),
	// Data Factory Azure-SSIS Integration Runtimes can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
	"azurerm_data_factory_integration_runtime_azure_ssis": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/integrationruntimes/{{ .external_name }}"),
	// Data Factory Pipeline's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/pipelines/example
	"azurerm_data_factory_pipeline": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/pipelines/{{ .external_name }}"),
	// Data Factory Blob Event Trigger can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
	"azurerm_data_factory_trigger_blob_event": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/triggers/{{ .external_name }}"),
	// Data Factory Schedule Trigger can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
	"azurerm_data_factory_trigger_schedule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/triggers/{{ .external_name }}"),
	// Data Factories can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
	"azurerm_data_factory_integration_runtime_self_hosted": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/integrationruntimes/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_custom_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_azure_blob_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_azure_databricks": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_azure_file_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_azure_function": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_azure_search": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Azure SQL Database Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_azure_sql_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_azure_table_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_cosmosdb": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Data Lake Storage Gen2 Linked Services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_data_lake_storage_gen2": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Key Vault Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_key_vault": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_kusto": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory MySQL Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_mysql": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory OData Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_odata": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory PostgreSQL Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_postgresql": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_sftp": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Snowflake Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_snowflake": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory SQL Server Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_sql_server": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Synapse Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_synapse": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_web": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory Managed Private Endpoint can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
	"azurerm_data_factory_managed_private_endpoint": config.IdentifierFromProvider,
	// Data Factory Custom Event Trigger can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/triggers/example
	"azurerm_data_factory_trigger_custom_event": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/triggers/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_sql_server_table": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
	"azurerm_data_factory_integration_runtime_managed": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/integrationruntimes/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_cosmosdb_mongoapi": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),

	// security
	//
	// Security Assessment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/00000000-0000-0000-0000-000000000000
	"azurerm_security_center_assessment": config.IdentifierFromProvider,
	// Security Assessments Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/assessmentMetadata/metadata1
	"azurerm_security_center_assessment_policy": config.IdentifierFromProvider,
	// Security Center Auto Provisioning can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/autoProvisioningSettings/default
	"azurerm_security_center_auto_provisioning": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Security/autoProvisioningSettings/default"),
	// The contact can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/securityContacts/default1
	"azurerm_security_center_contact": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Security/securityContacts/default1"),
	// DEPRECATED
	// Server Vulnerability Assessments can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.Compute/virtualMachines/vm-name/providers/Microsoft.Security/serverVulnerabilityAssessments/Default
	"azurerm_security_center_server_vulnerability_assessment": config.IdentifierFromProvider,
	// Server Vulnerability Assessments can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.Compute/virtualMachines/vm-name/providers/Microsoft.Security/serverVulnerabilityAssessments/Default
	"azurerm_security_center_server_vulnerability_assessment_virtual_machine": config.TemplatedStringAsIdentifier("", "{{ .parameters.virtual_machine_id }}/providers/Microsoft.Security/serverVulnerabilityAssessments/Default"),
	// The setting can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/settings/<setting_name>
	"azurerm_security_center_setting": config.IdentifierFromProvider,
	// The pricing tier can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/pricings/<resource_type>
	"azurerm_security_center_subscription_pricing": config.IdentifierFromProvider,
	// The contact can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/workspaceSettings/default
	"azurerm_security_center_workspace": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Security/workspaceSettings/default"),

	// servicebus
	//
	// ServiceBus Namespace authorization rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/authorizationRules/rule1
	"azurerm_servicebus_namespace_authorization_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/authorizationRules/{{ .external_name }}"),
	// Service Bus DR configs can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/disasterRecoveryConfigs/config1
	"azurerm_servicebus_namespace_disaster_recovery_config": config.TemplatedStringAsIdentifier("name", "{{ .parameters.primary_namespace_id }}/disasterRecoveryConfigs/{{ .external_name }}"),
	// Service Bus Namespace can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/sbns1
	// TODO: Check documentation, it seems there is a bug
	"azurerm_servicebus_namespace_network_rule_set": config.IdentifierFromProvider,
	// Service Bus Queue can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/queues/snqueue1
	"azurerm_servicebus_queue": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/queues/{{ .external_name }}"),
	// ServiceBus Queue Authorization Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/queues/queue1/authorizationRules/rule1
	"azurerm_servicebus_queue_authorization_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.queue_id }}/authorizationRules/{{ .external_name }}"),
	// Service Bus Subscriptions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1
	"azurerm_servicebus_subscription": config.TemplatedStringAsIdentifier("name", "{{ .parameters.topic_id }}/subscriptions/{{ .external_name }}"),
	// Service Bus Subscription Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
	"azurerm_servicebus_subscription_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.subscription_id }}/rules/{{ .external_name }}"),
	// Service Bus Topics can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1
	"azurerm_servicebus_topic": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/topics/{{ .external_name }}"),
	// ServiceBus Topic authorization rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/topics/topic1/authorizationRules/rule1
	"azurerm_servicebus_topic_authorization_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.topic_id }}/authorizationRules/{{ .external_name }}"),

	// traffic
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Network/trafficManagerProfiles/example-profile/AzureEndpoints/example-endpoint
	"azurerm_traffic_manager_azure_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.profile_id }}/AzureEndpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-group/providers/Microsoft.Network/trafficManagerProfiles/example-profile/ExternalEndpoints/example-endpoint
	"azurerm_traffic_manager_external_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.profile_id }}/ExternalEndpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Network/trafficManagerProfiles/example-profile/NestedEndpoints/example-endpoint
	"azurerm_traffic_manager_nested_endpoint": config.TemplatedStringAsIdentifier("name", "{{ .parameters.profile_id }}/NestedEndpoints/{{ .external_name }}"),

	// portal
	//
	// Dashboards can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Portal/dashboards/00000000-0000-0000-0000-000000000000
	"azurerm_portal_dashboard": config.IdentifierFromProvider,

	// resources
	//
	// Subscription Template Deployments can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources/deployments/template1
	"azurerm_subscription_template_deployment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Resources/deployments/{{ .external_name }}"),

	// search
	//
	// Search Services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Search/searchServices/service1
	"azurerm_search_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Search/searchServices/{{ .external_name }}"),

	// solutions
	//
	// Managed Application Definition can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applicationDefinitions/appDefinition1
	"azurerm_managed_application_definition": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Solutions/applicationDefinitions/{{ .external_name }}"),

	// signalrservice
	//
	// SignalR services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/signalR/tfex-signalr
	"azurerm_signalr_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.SignalRService/signalR/{{ .external_name }}"),
	// Network ACLs for a SignalR service can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1
	// TODO: Check documentation, it seems there is a bug
	"azurerm_signalr_service_network_acl": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
	"azurerm_web_pubsub": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webPubSub1/hubs/webPubSubhub1
	"azurerm_web_pubsub_hub": config.IdentifierFromProvider,

	// service
	//
	// Resource Groups can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ServiceFabric/managedClusters/clusterName1
	"azurerm_service_fabric_managed_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ServiceFabric/managedClusters/{{ .external_name }}"),

	// servicefabric
	//
	// Service Fabric Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabric/clusters/cluster1
	"azurerm_service_fabric_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ServiceFabric/clusters/{{ .external_name }}"),

	// eventgrid
	//
	// EventGrid Domains can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1
	"azurerm_eventgrid_domain": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventGrid/domains/{{ .external_name }}"),
	// EventGrid Domain Topics can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/domains/domain1/topics/topic1
	"azurerm_eventgrid_domain_topic": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventGrid/domains/{{ .parameters.domain_name }}/topics/{{ .external_name}}"),
	// EventGrid Event Subscription's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscription1
	"azurerm_eventgrid_event_subscription": config.IdentifierFromProvider,
	// Event Grid System Topic can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/systemTopics/systemTopic1
	"azurerm_eventgrid_system_topic": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventGrid/systemTopics/{{ .external_name }}"),
	// EventGrid Topic's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventGrid/topics/topic1
	"azurerm_eventgrid_topic": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventGrid/topics/{{ .external_name }}"),

	// devtest
	//
	// An existing Dev Test Global Shutdown Schedule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-rg/providers/Microsoft.DevTestLab/schedules/shutdown-computevm-SampleVM
	"azurerm_dev_test_global_vm_shutdown_schedule": config.IdentifierFromProvider,
	// Dev Test Labs can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1
	"azurerm_dev_test_lab": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .external_name }}"),
	// Dev Test Linux Virtual Machines can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
	"azurerm_dev_test_linux_virtual_machine": config.IdentifierFromProvider,
	// Dev Test Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/policysets/default/policies/policy1
	"azurerm_dev_test_policy": config.IdentifierFromProvider,
	// DevTest Schedule's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DevTestLab/labs/myDevTestLab/schedules/labvmautostart
	"azurerm_dev_test_schedule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .parameters.lab_name }}/schedules/{{ .external_name }}"),
	// DevTest Virtual Networks can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualnetworks/network1
	"azurerm_dev_test_virtual_network": config.IdentifierFromProvider,
	// DevTest Windows Virtual Machines can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
	"azurerm_dev_test_windows_virtual_machine": config.IdentifierFromProvider,

	// eventhub
	//
	// EventHub Namespace Authorization Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/authorizationRules/rule1
	"azurerm_eventhub_namespace_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/authorizationRules/{{ .external_name }}"),
	// EventHubs can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/disasterRecoveryConfigs/config1
	"azurerm_eventhub_namespace_disaster_recovery_config": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/disasterRecoveryConfigs/{{ .external_name }}"),

	// elastic
	//
	// Elasticsearch's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Elastic/monitors/monitor1
	"azurerm_elastic_cloud_elasticsearch": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Elastic/monitors/{{ .external_name }}"),

	// healthcareapis
	//
	// Healthcare Service can be imported using the resourceid
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource_group/providers/Microsoft.HealthcareApis/services/service_name
	"azurerm_healthcare_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.HealthcareApis/services/{{ .external_name }}"),

	// datamigration
	//
	// Database Migration Projects can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.DataMigration/services/example-dms/projects/project1
	"azurerm_database_migration_project": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataMigration/services/{{ .parameters.service_name }}/projects/{{ .external_name }}"),
	// Database Migration Services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.DataMigration/services/database_migration_service1
	"azurerm_database_migration_service": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataMigration/services/{{ .external_name }}"),

	// dataprotection
	//
	// Backup Instance Blob Storages can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
	"azurerm_data_protection_backup_instance_blob_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupInstances/{{ .external_name }}"),
	// Backup Instance Disks can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
	"azurerm_data_protection_backup_instance_disk": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupInstances/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/backupVaults/vault1/backupInstances/backupInstance1
	"azurerm_data_protection_backup_instance_postgresql": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupInstances/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataProtection/resourceGuards/resourceGuard1
	"azurerm_data_protection_resource_guard": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataProtection/resourceGuards/{{ .external_name }}"),

	// attestation
	//
	// Attestation Providers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Attestation/attestationProviders/provider1
	"azurerm_attestation_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Attestation/attestationProviders/{{ .external_name }}"),

	// web
	//
	// DEPRECATED
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/serverfarms/instance1
	"azurerm_app_service_plan": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/serverfarms/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/instance1/slots/staging/config/virtualNetwork
	// DEPRECATED: azurerm_service_plan should be used instead
	// App Service Source Control Token's can be imported using the type
	// "azurerm_app_service_source_control_token": config.IdentifierFromProvider,
	// DEPRECATED
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1
	"azurerm_function_app": config.IdentifierFromProvider,
	// "azurerm_function_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),
	// DEPRECATED
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1/slots/staging
	"azurerm_function_app_slot": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .parameters.function_app_name }}/slots/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_linux_web_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),
	// /providers/Microsoft.Web/sourceControls/GitHub
	"azurerm_source_control_token": config.IdentifierFromProvider,
	// a Function App Active Slot can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_function_app_active_slot": config.IdentifierFromProvider,
	// a Function App Function can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/functions/function1
	"azurerm_function_app_function": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_web_app_active_slot": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/hybridConnectionNamespaces/hybridConnectionNamespace1/relays/relay1
	"azurerm_web_app_hybrid_connection": config.IdentifierFromProvider,
	// a Function App Hybrid Connection can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/hybridConnectionNamespaces/hybridConnectionNamespace1/relays/relay1
	"azurerm_function_app_hybrid_connection": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_windows_function_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_windows_function_app_slot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.function_app_id }}/slots/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_windows_web_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_windows_web_app_slot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.app_service_id }}/slots/{{ .external_name }}"),

	// web_pubsub
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/webpubsub1
	// TODO: Bug in documentation. Normalize external_name while testing
	"azurerm_web_pubsub_network_acl": config.IdentifierFromProvider,

	// logic
	//
	// Logic App Custom Actions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/custom1
	"azurerm_logic_app_action_custom": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logic_app_id }}/actions/{{ .external_name }}"),
	// Logic App HTTP Actions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/webhook1
	"azurerm_logic_app_action_http": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logic_app_id }}/actions/{{ .external_name }}"),
	// Logic App Integration Account Schemas can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/schemas/schema1
	"azurerm_logic_app_integration_account_schema": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/schemas/{{ .external_name }}"),
	// Logic App Integration Account Sessions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/sessions/session1
	"azurerm_logic_app_integration_account_session": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/sessions/{{ .external_name }}"),
	// Logic App Custom Triggers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/custom1
	"azurerm_logic_app_trigger_custom": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logic_app_id }}/triggers/{{ .external_name }}"),
	// Logic App HTTP Request Triggers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/request1
	"azurerm_logic_app_trigger_http_request": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logic_app_id }}/triggers/{{ .external_name }}"),
	// Logic App Recurrence Triggers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/triggers/daily
	"azurerm_logic_app_trigger_recurrence": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logic_app_id }}/triggers/{{ .external_name }}"),
	// Logic App Workflows can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1
	"azurerm_logic_app_workflow": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/workflows/{{ .external_name }}"),

	// iotcentral
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.IoTCentral/iotApps/app1
	"azurerm_iotcentral_application_network_rule_set": config.IdentifierFromProvider,

	// iothub
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DeviceUpdate/accounts/account1
	"azurerm_iothub_device_update_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DeviceUpdate/accounts/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DeviceUpdate/accounts/account1/instances/instance1
	"azurerm_iothub_device_update_instance": config.TemplatedStringAsIdentifier("name", "{{ .parameters.device_update_account_id }}/instances/{{ .external_name }}"),

	// kubernetes_fleet
	//
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}
	"azurerm_kubernetes_fleet_manager": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ContainerService/fleets/{{ .external_name }}"),

	// kusto
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/ManagedPrivateEndpoints/managedPrivateEndpoint1
	"azurerm_kusto_cluster_managed_private_endpoint": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/clusters/{{ .parameters.cluster_name }}/managedPrivateEndpoints/{{ .external_name }}"),

	// managedidentity
	//
	// An existing User Assigned Identity can be imported into Terraform using the resource
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}
	"azurerm_user_assigned_identity": config.IdentifierFromProvider,

	// maps
	//
	// A Maps Account can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Maps/accounts/my-maps-account
	"azurerm_maps_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Maps/accounts/{{ .external_name }}"),
	// An Azure Maps Creators can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Maps/accounts/account1/creators/creator1
	"azurerm_maps_creator": config.TemplatedStringAsIdentifier("name", "{{ .parameters.maps_account_id }}/creators/{{ .external_name }}"),

	// sentinel
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_ms_security_incident": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_iot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Sentinel Automation Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/automationRules/rule1
	"azurerm_sentinel_automation_rule": config.IdentifierFromProvider,
	// Sentinel Watchlists can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1
	"azurerm_sentinel_watchlist": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/watchlists/{{ .external_name }}"),
	// Security Insights Sentinel Onboarding States can be imported using the resource id
	//  /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/onboardingStates/defaults
	"azurerm_sentinel_log_analytics_workspace_onboarding": config.IdentifierFromProvider,

	// lab_service
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.LabServices/labPlans/labPlan1
	"azurerm_lab_service_plan": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.LabServices/labPlans/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.LabServices/labs/lab1
	"azurerm_lab_service_lab": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.LabServices/labs/{{ .external_name }}"),

	// log_analytics
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.OperationalInsights/queryPacks/queryPack1
	"azurerm_log_analytics_query_pack": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/queryPacks/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.OperationalInsights/queryPacks/queryPack1/queries/15b49e87-8555-4d92-8a7b-2014b469a9df
	"azurerm_log_analytics_query_pack_query": config.IdentifierFromProvider,

	// logic
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/partners/partner1
	"azurerm_logic_app_integration_account_partner": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/partners/{{ .external_name }}"),

	// monitor
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
	"azurerm_monitor_alert_processing_rule_action_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AlertsManagement/actionRules/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.AlertsManagement/actionRules/actionRule1
	"azurerm_monitor_alert_processing_rule_suppression": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AlertsManagement/actionRules/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Insights/dataCollectionEndpoints/endpoint1
	"azurerm_monitor_data_collection_endpoint": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/dataCollectionEndpoints/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Insights/dataCollectionRules/rule1
	"azurerm_monitor_data_collection_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/dataCollectionRules/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.Insights/dataCollectionRuleAssociations/dca1
	"azurerm_monitor_data_collection_rule_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.target_resource_id }}/providers/Microsoft.Insights/dataCollectionRuleAssociations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Insights/scheduledQueryRules/rule1
	"azurerm_monitor_scheduled_query_rules_alert_v2": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/scheduledQueryRules/{{ .external_name }}"),

	// spring_cloud
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/applicationAccelerators/default
	"azurerm_spring_cloud_accelerator": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/applicationLiveViews/default
	"azurerm_spring_cloud_application_live_view": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/applicationLiveViews/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/Spring/spring1/applicationAccelerators/default/customizedAccelerators/customizedAccelerator1
	"azurerm_spring_cloud_customized_accelerator": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_accelerator_id }}/customizedAccelerators/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/Spring/service1/DevToolPortals/default
	"azurerm_spring_cloud_dev_tool_portal": config.IdentifierFromProvider,

	// static_site
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1
	"azurerm_static_site": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/staticSites/{{ .external_name }}"),

	// signalr
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1/sharedPrivateLinkResources/resource1
	"azurerm_signalr_shared_private_link_resource": config.IdentifierFromProvider,

	// site_recovery
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group-name/providers/Microsoft.RecoveryServices/vaults/recovery-vault-name/replicationFabrics/primary-fabric-name/replicationNetworks/azureNetwork/replicationNetworkMappings/mapping-name
	"azurerm_site_recovery_network_mapping": config.IdentifierFromProvider,

	// search
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Search/searchServices/service1/sharedPrivateLinkResources/resource1
	"azurerm_search_shared_private_link_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.search_service_id }}/sharedPrivateLinkResources/{{ .external_name }}"),

	// resource_deployment
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Resources/deploymentScripts/script1
	"azurerm_resource_deployment_script_azure_cli": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Resources/deploymentScripts/script1
	"azurerm_resource_deployment_script_azure_power_shell": config.IdentifierFromProvider,

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

	// certificateregistration
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.CertificateRegistration/certificateOrders/certificateorder1
	"azurerm_app_service_certificate_order": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CertificateRegistration/certificateOrders/{{ .external_name }}"),

	// storage
	//
	// https://account1.dfs.core.windows.net/fileSystem1/path
	"azurerm_storage_data_lake_gen2_path": config.IdentifierFromProvider,
	// https://tomdevsa20.file.core.windows.net/share1/directory1
	"azurerm_storage_share_directory": config.IdentifierFromProvider,
	// https://example.table.core.windows.net/table1(PartitionKey='samplepartition',RowKey='samplerow')
	"azurerm_storage_table_entity": config.IdentifierFromProvider,

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

	// automation
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connectionTypes/type1
	"azurerm_automation_connection_type": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/hybridRunbookWorkerGroups/grp1
	"azurerm_automation_hybrid_runbook_worker_group": config.IdentifierFromProvider,

	// fluid_relay
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.FluidRelay/fluidRelayServers/server1
	"azurerm_fluid_relay_server": config.IdentifierFromProvider,

	// federated_identity
	//
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{parentIdentityName}/federatedIdentityCredentials/{resourceName}
	"azurerm_federated_identity_credential": config.TemplatedStringAsIdentifier("name", "{{ .parameters.parent_id }}/federatedIdentityCredentials/{{ .external_name }}"),

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

	// stream_analytics
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/outputs/output1
	"azurerm_stream_analytics_output_eventhub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_reference_input_mssql": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_name }}/inputs/{{ .external_name }}"),

	// logz
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1/accounts/subAccount1
	"azurerm_logz_sub_account": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logz_monitor_id }}/accounts/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1/accounts/subAccount1/tagRules/ruleSet1
	"azurerm_logz_sub_account_tag_rule": config.IdentifierFromProvider,
	// logz Monitors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1
	"azurerm_logz_monitor": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logz/monitors/{{ .external_name }}"),
	// logz Tag Rules can be imported using the resource id
	//  /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logz/monitors/monitor1/tagRules/ruleSet1
	"azurerm_logz_tag_rule": config.TemplatedStringAsIdentifier("", "{{ .parameters.logz_monitor_id }}/tagRules/default"),

	// media_job
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Media/mediaServices/account1/transforms/transform1/jobs/job1
	"azurerm_media_job": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Media/mediaServices/{{ .parameters.media_services_account_name }}/transforms/{{ .parameters.transform_name }}/jobs/{{ .external_name }}"),

	// network_manager
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1
	"azurerm_network_manager": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkManagers/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/networkGroups/networkGroup1
	"azurerm_network_manager_network_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_manager_id }}/networkGroups/{{ .external_name }}"),

	// mssql
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/devOpsAuditingSettings/default
	"azurerm_mssql_server_microsoft_support_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/devOpsAuditingSettings/default"),

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
	"azurerm_spring_cloud_gateway_custom_domain": config.IdentifierFromProvider,

	// securityinsights
	//
	// Sentinel Fusion Alert Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_fusion": config.IdentifierFromProvider,
	// Sentinel Machine Learning Behavior Analytics Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_machine_learning_behavior_analytics": config.IdentifierFromProvider,

	// private_dns
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/dnsResolvers/dnsResolver1
	"azurerm_private_dns_resolver": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/dnsResolvers/{{ .external_name }}"),

	// orbital_contact
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/spacecrafts/spacecraft1
	"azurerm_orbital_spacecraft": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Orbital/spacecrafts/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Orbital/contactProfiles/contactProfile1
	"azurerm_orbital_contact_profile": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Orbital/contactProfiles/{{ .external_name }}"),

	// policy_virtual_machine
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/assignment1
	"azurerm_policy_virtual_machine_configuration_assignment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.virtual_machine_id }}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{ .external_name }}"),

	// customproviders
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.CustomProviders/resourceProviders/example
	"azurerm_custom_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CustomProviders/resourceProviders/{{ .external_name }}"),

	// api_management
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.ApiManagement/service/service1/gateways/gateway1/apis/api1
	"azurerm_api_management_gateway_api": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/products/myproduct/tags/mytag
	"azurerm_api_management_product_tag": config.IdentifierFromProvider,

	// network
	//
	// /providers/Microsoft.Management/managementGroups/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/networkManagerConnections/networkManagerConnection1
	"azurerm_network_manager_management_group_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Network/networkManagerConnections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Network/networkManagers/networkManager1/networkGroups/networkGroup1/staticMembers/staticMember1
	"azurerm_network_manager_static_member": config.TemplatedStringAsIdentifier("name", "{{ .parameters.network_group_id }}/staticMembers/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/networkManagerConnections/networkManagerConnection1
	"azurerm_network_manager_subscription_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Network/networkManagerConnections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/privateEndpoints/endpoints1|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/applicationSecurityGroups/securityGroup1
	"azurerm_private_endpoint_application_security_group_association": config.IdentifierFromProvider,

	// storage
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/localUsers/user1
	"azurerm_storage_account_local_user": config.TemplatedStringAsIdentifier("name", "{{ .parameters.storage_account_id }}/localUsers/{{ .external_name }}"),

	// machinelearningservices
	//
	// Machine Learning Workspace can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.MachineLearningServices/workspaces/workspace1
	"azurerm_machine_learning_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.MachineLearningServices/workspaces/{{ .external_name }}"),
	// Machine Learning Compute Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1
	"azurerm_machine_learning_compute_cluster": config.IdentifierFromProvider,
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
	"azurerm_maintenance_assignment_dedicated_host": config.IdentifierFromProvider,
	// Maintenance Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/vm1/providers/Microsoft.Maintenance/configurationAssignments/assign1
	"azurerm_maintenance_assignment_virtual_machine": config.TemplatedStringAsIdentifier("", "{{ .parameters.virtual_machine_id }}/providers/Microsoft.Maintenance/configurationAssignments/default"),

	// appconfiguration
	//
	// App Configurations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1
	"azurerm_app_configuration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppConfiguration/configurationStores/{{ .external_name }}"),
}

func keyVaultURLIDConf(resourceType string) config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = getResourceNameFromIDURLFn(2)
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		keyVaultID, ok := parameters["key_vault_id"]
		if !ok {
			return "", errors.New("cannot get key_vault_id")
		}
		words := strings.Split(keyVaultID.(string), "/")
		keyVaultName := words[len(words)-1]

		name, ok := parameters["name"]
		if !ok {
			return "", errors.New("cannot get name")
		}

		version, ok := parameters["version"]
		if !ok {
			return "", nil
		}

		return fmt.Sprintf("https://%s.vault.azure.net/%s/%s/%s",
			keyVaultName, resourceType, name, version), nil
	}
	return e
}

func keyVaultURLIDWithoutVersionConfFn(resourceType string) config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = getResourceNameFromIDURLFn(1)
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		keyVaultID, ok := parameters["key_vault_id"]
		if !ok {
			return "", errors.New("cannot get key_vault_id")
		}
		words := strings.Split(keyVaultID.(string), "/")
		keyVaultName := words[len(words)-1]

		name, ok := parameters["name"]
		if !ok {
			return "", errors.New("cannot get name")
		}

		return fmt.Sprintf("https://%s.vault.azure.net/%s/%s",
			keyVaultName, resourceType, name), nil
	}
	return e
}

func getResourceNameFromIDURLFn(pos int) config.GetExternalNameFn {
	return func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("cannot get id")
		}
		words := strings.Split(id.(string), "/")
		return words[len(words)-pos], nil
	}
}

func keyVaultAccessPolicy() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		objectID, ok := tfstate["object_id"]
		if !ok {
			return "", errors.New("cannot get object_id")
		}
		applicationID := tfstate["application_id"]
		if applicationID == nil || applicationID == "" {
			return objectID.(string), nil
		}
		return fmt.Sprintf("%s/%s", objectID, applicationID), nil
	}
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		keyVaultID, ok := parameters["key_vault_id"]
		if !ok {
			return "", errors.New("cannot get key_vault_id")
		}
		objectID, ok := parameters["object_id"]
		if !ok {
			return "", errors.New("cannot get object_id")
		}

		applicationID := parameters["application_id"]
		if applicationID == nil || applicationID == "" {
			return fmt.Sprintf("%s/objectId/%s", keyVaultID, objectID), nil
		}
		return fmt.Sprintf("%s/objectId/%s/applicationId/%s", keyVaultID, objectID, applicationID), nil
	}
	return e
}

// custom function for azurerm_storage_data_lake_gen2_filesystem
func storageDataLakeGen2Filesystem() config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		w := strings.Split(id.(string), "/")
		return w[len(w)-1], nil
	}
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		storageAccountID, ok := parameters["storage_account_id"]
		if !ok {
			return "", errors.New("cannot get storage_account_id")
		}
		w := strings.Split(storageAccountID.(string), "/")
		storageAccountName := w[len(w)-1]
		return fmt.Sprintf("https://%s.dfs.core.windows.net/%s", storageAccountName, externalName), nil
	}
	return e
}

// custom function for azurerm_management_group_subscription_association
// /managementGroup/MyManagementGroup/subscription/12345678-1234-1234-1234-123456789012
func managementGroupSubscriptionAssociation() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		w := strings.Split(id.(string), "/")
		return fmt.Sprintf("%s/%s", w[len(w)-3], w[len(w)-1]), nil
	}
	// if we construct id according to the full path above, the underlying
	// terraform non-deterministically fails with
	//  "could not read properties for Management Group "example-sub""
	// just populate the id with empty string solves it. Same happens in
	// isolated test with terraform cli
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		return "", nil
	}
	return e
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
