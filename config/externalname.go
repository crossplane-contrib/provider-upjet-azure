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

	// cdn
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1/endpoints/myendpoint1
	"azurerm_cdn_endpoint": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cdn/profiles/{{ .parameters.profile_name }}/endpoints/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/myprofile1
	"azurerm_cdn_profile": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cdn/profiles/{{ .external_name }}"),

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
	// "azurerm_data_protection_backup_policy_disk":         config.TemplatedStringAsIdentifier("name", "{{ .parameters.vault_id }}/backupPolicies/{{ .external_name }}"),
	// "azurerm_data_protection_backup_policy_postgresql": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataProtection/backupVaults/{{ .parameters.vault_name }}/backupPolicies/{{ .external_name }}"),

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

	// eventhub
	"azurerm_eventhub_namespace":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .external_name }}"),
	"azurerm_eventhub":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .external_name }}"),
	"azurerm_eventhub_consumer_group":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/consumerGroups/{{ .external_name }}"),
	"azurerm_eventhub_authorization_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.EventHub/namespaces/{{ .parameters.namespace_name }}/eventhubs/{{ .parameters.eventhub_name }}/authorizationRules/{{ .external_name }}"),

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

	// kusto
	"azurerm_kusto_cluster":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .external_name }}"),
	"azurerm_kusto_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/Databases/{{ .external_name }}"),
	// Kusto Attached Database Configurations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/AttachedDatabaseConfigurations/configuration1
	"azurerm_kusto_attached_database_configuration": config.IdentifierFromProvider,
	// Data Explorer Cluster Principal Assignments can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/PrincipalAssignments/assignment1
	"azurerm_kusto_cluster_principal_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/PrincipalAssignments/{{ .external_name }}"),
	// Kusto Database Principal Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/PrincipalAssignments/assignment1
	"azurerm_kusto_database_principal_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/Databases/{{ .parameters.database_name }}/PrincipalAssignments/{{ .external_name }}"),
	// Kusto Event Grid Data Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/dataConnection1
	"azurerm_kusto_eventgrid_data_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/Databases/{{ .parameters.database_name }}/DataConnections/{{ .external_name }}"),
	// Kusto EventHub Data Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/eventHubConnection1
	"azurerm_kusto_eventhub_data_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/Databases/{{ .parameters.database_name }}/DataConnections/{{ .external_name }}"),
	// Kusto IotHub Data Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/DataConnections/dataConnection1
	"azurerm_kusto_iothub_data_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/Databases/{{ .parameters.database_name }}/DataConnections/{{ .external_name }}"),

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

	// operationalinsights
	"azurerm_log_analytics_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/workspaces/{{ .external_name }}"),

	// insights
	"azurerm_application_insights": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/components/{{ .external_name }}"),
	// Application Insights API keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Insights/components/instance1/apiKeys/00000000-0000-0000-0000-000000000000
	"azurerm_application_insights_api_key":        config.IdentifierFromProvider,
	"azurerm_monitor_action_group":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/actionGroups/{{ .external_name }}"),
	"azurerm_monitor_metric_alert":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/metricAlerts/{{ .external_name }}"),
	"azurerm_monitor_private_link_scope":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/privateLinkScopes/{{ .external_name }}"),
	"azurerm_monitor_private_link_scoped_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/privateLinkScopes/{{ .parameters.scope_name }}/scopedResources/{{ .external_name }}"),

	// logic
	"azurerm_integration_service_environment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationServiceEnvironments/{{ .external_name }}"),

	// management
	"azurerm_management_group": config.TemplatedStringAsIdentifier("name", "/providers/Microsoft.Management/managementGroups/{{ .external_name }}"),

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

	// mixedreality
	"azurerm_spatial_anchors_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{{ .external_name }}"),

	// dbformysql
	"azurerm_mysql_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .external_name }}"),
	"azurerm_mysql_configuration":        config.IdentifierFromProvider,
	"azurerm_mysql_firewall_rule":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_mysql_virtual_network_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .external_name }}"),
	// dbformysql flexible variants
	"azurerm_mysql_flexible_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .external_name }}"),
	"azurerm_mysql_flexible_database":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	"azurerm_mysql_flexible_server_configuration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/configurations/{{ .external_name }}"),
	"azurerm_mysql_flexible_server_firewall_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),

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

	// notification
	"azurerm_notification_hub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.NotificationHubs/namespaces/{{ .parameters.namespace_name }}/notificationHubs/{{ .external_name }}"),

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
	"azurerm_redis_cache":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/Redis/{{ .external_name }}"),
	"azurerm_redis_firewall_rule":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/Redis/{{ .parameters.redis_cache_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_redis_enterprise_cluster":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redisEnterprise/{{ .external_name }}"),
	"azurerm_redis_enterprise_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cluster_id }}/databases/{{ .external_name }}"),
	"azurerm_redis_linked_server":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/Redis/{{ .parameters.target_redis_cache_name }}/linkedServers/{{ .external_name }}"),

	// resource
	"azurerm_resource_group_template_deployment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Resources/deployments/{{ .external_name }}"),

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
	// "azurerm_stream_analytics_output_table":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	// "azurerm_stream_analytics_function_javascript_udf": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.job_name }}/outputs/{{ .external_name }}"),
	// "azurerm_stream_analytics_output_cosmosdb":          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_id }}/outputs/{{ .external_name }}"),
	// Does not work, tracked in https://github.com/upbound/official-providers/issues/368
	// "azurerm_stream_analytics_output_powerbi": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_id }}/outputs/{{ .external_name }}"),
	"azurerm_stream_analytics_output_blob": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingJobs/{{ .parameters.stream_analytics_job_id }}/outputs/{{ .external_name }}"),

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
	"azurerm_spring_cloud_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/Spring/{{ .parameters.service_name }}/apps/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
	"azurerm_spring_cloud_app_cosmosdb_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/bindings/bind1
	"azurerm_spring_cloud_app_mysql_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.AppPlatform/spring/myservice/apps/myapp/bindings/bind1
	"azurerm_spring_cloud_app_redis_association": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/bindings/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/spring1/certificates/cert1
	"azurerm_spring_cloud_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/Spring/{{ .parameters.service_name }}/certificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/domains/domain.com
	"azurerm_spring_cloud_custom_domain": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.AppPlatform/spring/service1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_java_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AppPlatform/spring/spring1
	"azurerm_spring_cloud_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppPlatform/Spring/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/storages/storage1
	"azurerm_spring_cloud_storage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/storages/{{ .external_name }}"),

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

	// api
	//
	// API Management API Tags can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/apis/api1/tags/tag1
	"azurerm_api_management_api_tag": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}/tags/{{ .external_name }}"),
	// API Management Notification Recipient Users can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/notifications/notificationName1/recipientUsers/userid1
	"azurerm_api_management_notification_recipient_user": config.IdentifierFromProvider,
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
