// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APIInitParameters struct {

	// Type of API. Possible values are graphql, http, soap, and websocket. Defaults to http.
	APIType *string `json:"apiType,omitempty" tf:"api_type,omitempty"`

	// A contact block as documented below.
	Contact []ContactInitParameters `json:"contact,omitempty" tf:"contact,omitempty"`

	// A description of the API Management API, which may include HTML formatting tags.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the API.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A import block as documented below.
	Import []ImportInitParameters `json:"import,omitempty" tf:"import,omitempty"`

	// A license block as documented below.
	License []LicenseInitParameters `json:"license,omitempty" tf:"license,omitempty"`

	// An oauth2_authorization block as documented below.
	Oauth2Authorization []Oauth2AuthorizationInitParameters `json:"oauth2Authorization,omitempty" tf:"oauth2_authorization,omitempty"`

	// An openid_authentication block as documented below.
	OpenIDAuthentication []OpenIDAuthenticationInitParameters `json:"openidAuthentication,omitempty" tf:"openid_authentication,omitempty"`

	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A list of protocols the operations in this API can be invoked. Possible values are http, https, ws, and wss.
	// +listType=set
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// The description of the API Revision of the API Management API.
	RevisionDescription *string `json:"revisionDescription,omitempty" tf:"revision_description,omitempty"`

	// Absolute URL of the backend service implementing this API.
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`

	// The API id of the source API, which could be in format azurerm_api_management_api.example.id or in format azurerm_api_management_api.example.id;rev=1
	SourceAPIID *string `json:"sourceApiId,omitempty" tf:"source_api_id,omitempty"`

	// A subscription_key_parameter_names block as documented below.
	SubscriptionKeyParameterNames []SubscriptionKeyParameterNamesInitParameters `json:"subscriptionKeyParameterNames,omitempty" tf:"subscription_key_parameter_names,omitempty"`

	// Should this API require a subscription key? Defaults to true.
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty" tf:"subscription_required,omitempty"`

	// Absolute URL of the Terms of Service for the API.
	TermsOfServiceURL *string `json:"termsOfServiceUrl,omitempty" tf:"terms_of_service_url,omitempty"`

	// The Version number of this API, if this API is versioned.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The description of the API Version of the API Management API.
	VersionDescription *string `json:"versionDescription,omitempty" tf:"version_description,omitempty"`

	// The ID of the Version Set which this API is associated with.
	VersionSetID *string `json:"versionSetId,omitempty" tf:"version_set_id,omitempty"`
}

type APIObservation struct {

	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Type of API. Possible values are graphql, http, soap, and websocket. Defaults to http.
	APIType *string `json:"apiType,omitempty" tf:"api_type,omitempty"`

	// A contact block as documented below.
	Contact []ContactObservation `json:"contact,omitempty" tf:"contact,omitempty"`

	// A description of the API Management API, which may include HTML formatting tags.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the API.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the API Management API.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A import block as documented below.
	Import []ImportObservation `json:"import,omitempty" tf:"import,omitempty"`

	// Is this the current API Revision?
	IsCurrent *bool `json:"isCurrent,omitempty" tf:"is_current,omitempty"`

	// Is this API Revision online/accessible via the Gateway?
	IsOnline *bool `json:"isOnline,omitempty" tf:"is_online,omitempty"`

	// A license block as documented below.
	License []LicenseObservation `json:"license,omitempty" tf:"license,omitempty"`

	// An oauth2_authorization block as documented below.
	Oauth2Authorization []Oauth2AuthorizationObservation `json:"oauth2Authorization,omitempty" tf:"oauth2_authorization,omitempty"`

	// An openid_authentication block as documented below.
	OpenIDAuthentication []OpenIDAuthenticationObservation `json:"openidAuthentication,omitempty" tf:"openid_authentication,omitempty"`

	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A list of protocols the operations in this API can be invoked. Possible values are http, https, ws, and wss.
	// +listType=set
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Revision which used for this API. Changing this forces a new resource to be created.
	Revision *string `json:"revision,omitempty" tf:"revision,omitempty"`

	// The description of the API Revision of the API Management API.
	RevisionDescription *string `json:"revisionDescription,omitempty" tf:"revision_description,omitempty"`

	// Absolute URL of the backend service implementing this API.
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`

	// The API id of the source API, which could be in format azurerm_api_management_api.example.id or in format azurerm_api_management_api.example.id;rev=1
	SourceAPIID *string `json:"sourceApiId,omitempty" tf:"source_api_id,omitempty"`

	// A subscription_key_parameter_names block as documented below.
	SubscriptionKeyParameterNames []SubscriptionKeyParameterNamesObservation `json:"subscriptionKeyParameterNames,omitempty" tf:"subscription_key_parameter_names,omitempty"`

	// Should this API require a subscription key? Defaults to true.
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty" tf:"subscription_required,omitempty"`

	// Absolute URL of the Terms of Service for the API.
	TermsOfServiceURL *string `json:"termsOfServiceUrl,omitempty" tf:"terms_of_service_url,omitempty"`

	// The Version number of this API, if this API is versioned.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The description of the API Version of the API Management API.
	VersionDescription *string `json:"versionDescription,omitempty" tf:"version_description,omitempty"`

	// The ID of the Version Set which this API is associated with.
	VersionSetID *string `json:"versionSetId,omitempty" tf:"version_set_id,omitempty"`
}

type APIParameters struct {

	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Type of API. Possible values are graphql, http, soap, and websocket. Defaults to http.
	// +kubebuilder:validation:Optional
	APIType *string `json:"apiType,omitempty" tf:"api_type,omitempty"`

	// A contact block as documented below.
	// +kubebuilder:validation:Optional
	Contact []ContactParameters `json:"contact,omitempty" tf:"contact,omitempty"`

	// A description of the API Management API, which may include HTML formatting tags.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the API.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A import block as documented below.
	// +kubebuilder:validation:Optional
	Import []ImportParameters `json:"import,omitempty" tf:"import,omitempty"`

	// A license block as documented below.
	// +kubebuilder:validation:Optional
	License []LicenseParameters `json:"license,omitempty" tf:"license,omitempty"`

	// An oauth2_authorization block as documented below.
	// +kubebuilder:validation:Optional
	Oauth2Authorization []Oauth2AuthorizationParameters `json:"oauth2Authorization,omitempty" tf:"oauth2_authorization,omitempty"`

	// An openid_authentication block as documented below.
	// +kubebuilder:validation:Optional
	OpenIDAuthentication []OpenIDAuthenticationParameters `json:"openidAuthentication,omitempty" tf:"openid_authentication,omitempty"`

	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A list of protocols the operations in this API can be invoked. Possible values are http, https, ws, and wss.
	// +kubebuilder:validation:Optional
	// +listType=set
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Revision which used for this API. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Revision *string `json:"revision" tf:"revision,omitempty"`

	// The description of the API Revision of the API Management API.
	// +kubebuilder:validation:Optional
	RevisionDescription *string `json:"revisionDescription,omitempty" tf:"revision_description,omitempty"`

	// Absolute URL of the backend service implementing this API.
	// +kubebuilder:validation:Optional
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`

	// The API id of the source API, which could be in format azurerm_api_management_api.example.id or in format azurerm_api_management_api.example.id;rev=1
	// +kubebuilder:validation:Optional
	SourceAPIID *string `json:"sourceApiId,omitempty" tf:"source_api_id,omitempty"`

	// A subscription_key_parameter_names block as documented below.
	// +kubebuilder:validation:Optional
	SubscriptionKeyParameterNames []SubscriptionKeyParameterNamesParameters `json:"subscriptionKeyParameterNames,omitempty" tf:"subscription_key_parameter_names,omitempty"`

	// Should this API require a subscription key? Defaults to true.
	// +kubebuilder:validation:Optional
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty" tf:"subscription_required,omitempty"`

	// Absolute URL of the Terms of Service for the API.
	// +kubebuilder:validation:Optional
	TermsOfServiceURL *string `json:"termsOfServiceUrl,omitempty" tf:"terms_of_service_url,omitempty"`

	// The Version number of this API, if this API is versioned.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The description of the API Version of the API Management API.
	// +kubebuilder:validation:Optional
	VersionDescription *string `json:"versionDescription,omitempty" tf:"version_description,omitempty"`

	// The ID of the Version Set which this API is associated with.
	// +kubebuilder:validation:Optional
	VersionSetID *string `json:"versionSetId,omitempty" tf:"version_set_id,omitempty"`
}

type ContactInitParameters struct {

	// The email address of the contact person/organization.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the contact person/organization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Absolute URL of the contact information.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ContactObservation struct {

	// The email address of the contact person/organization.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the contact person/organization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Absolute URL of the contact information.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ContactParameters struct {

	// The email address of the contact person/organization.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the contact person/organization.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Absolute URL of the contact information.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ImportInitParameters struct {

	// The format of the content from which the API Definition should be imported. Possible values are: openapi, openapi+json, openapi+json-link, openapi-link, swagger-json, swagger-link-json, wadl-link-json, wadl-xml, wsdl and wsdl-link.
	ContentFormat *string `json:"contentFormat,omitempty" tf:"content_format,omitempty"`

	// The Content from which the API Definition should be imported. When a content_format of *-link-* is specified this must be a URL, otherwise this must be defined inline. The URL must be accessible and return a valid document; otherwise, deployment may fail.
	ContentValue *string `json:"contentValue,omitempty" tf:"content_value,omitempty"`

	// A wsdl_selector block as defined below, which allows you to limit the import of a WSDL to only a subset of the document. This can only be specified when content_format is wsdl or wsdl-link.
	WsdlSelector []WsdlSelectorInitParameters `json:"wsdlSelector,omitempty" tf:"wsdl_selector,omitempty"`
}

type ImportObservation struct {

	// The format of the content from which the API Definition should be imported. Possible values are: openapi, openapi+json, openapi+json-link, openapi-link, swagger-json, swagger-link-json, wadl-link-json, wadl-xml, wsdl and wsdl-link.
	ContentFormat *string `json:"contentFormat,omitempty" tf:"content_format,omitempty"`

	// The Content from which the API Definition should be imported. When a content_format of *-link-* is specified this must be a URL, otherwise this must be defined inline. The URL must be accessible and return a valid document; otherwise, deployment may fail.
	ContentValue *string `json:"contentValue,omitempty" tf:"content_value,omitempty"`

	// A wsdl_selector block as defined below, which allows you to limit the import of a WSDL to only a subset of the document. This can only be specified when content_format is wsdl or wsdl-link.
	WsdlSelector []WsdlSelectorObservation `json:"wsdlSelector,omitempty" tf:"wsdl_selector,omitempty"`
}

type ImportParameters struct {

	// The format of the content from which the API Definition should be imported. Possible values are: openapi, openapi+json, openapi+json-link, openapi-link, swagger-json, swagger-link-json, wadl-link-json, wadl-xml, wsdl and wsdl-link.
	// +kubebuilder:validation:Optional
	ContentFormat *string `json:"contentFormat" tf:"content_format,omitempty"`

	// The Content from which the API Definition should be imported. When a content_format of *-link-* is specified this must be a URL, otherwise this must be defined inline. The URL must be accessible and return a valid document; otherwise, deployment may fail.
	// +kubebuilder:validation:Optional
	ContentValue *string `json:"contentValue" tf:"content_value,omitempty"`

	// A wsdl_selector block as defined below, which allows you to limit the import of a WSDL to only a subset of the document. This can only be specified when content_format is wsdl or wsdl-link.
	// +kubebuilder:validation:Optional
	WsdlSelector []WsdlSelectorParameters `json:"wsdlSelector,omitempty" tf:"wsdl_selector,omitempty"`
}

type LicenseInitParameters struct {

	// The name of the license .
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Absolute URL of the license.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LicenseObservation struct {

	// The name of the license .
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Absolute URL of the license.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LicenseParameters struct {

	// The name of the license .
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Absolute URL of the license.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type Oauth2AuthorizationInitParameters struct {

	// OAuth authorization server identifier. The name of an OAuth2 Authorization Server.
	AuthorizationServerName *string `json:"authorizationServerName,omitempty" tf:"authorization_server_name,omitempty"`

	// Operations scope.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type Oauth2AuthorizationObservation struct {

	// OAuth authorization server identifier. The name of an OAuth2 Authorization Server.
	AuthorizationServerName *string `json:"authorizationServerName,omitempty" tf:"authorization_server_name,omitempty"`

	// Operations scope.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type Oauth2AuthorizationParameters struct {

	// OAuth authorization server identifier. The name of an OAuth2 Authorization Server.
	// +kubebuilder:validation:Optional
	AuthorizationServerName *string `json:"authorizationServerName" tf:"authorization_server_name,omitempty"`

	// Operations scope.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type OpenIDAuthenticationInitParameters struct {

	// How to send token to the server. A list of zero or more methods. Valid values are authorizationHeader and query.
	// +listType=set
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// OpenID Connect provider identifier. The name of an OpenID Connect Provider.
	OpenIDProviderName *string `json:"openidProviderName,omitempty" tf:"openid_provider_name,omitempty"`
}

type OpenIDAuthenticationObservation struct {

	// How to send token to the server. A list of zero or more methods. Valid values are authorizationHeader and query.
	// +listType=set
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// OpenID Connect provider identifier. The name of an OpenID Connect Provider.
	OpenIDProviderName *string `json:"openidProviderName,omitempty" tf:"openid_provider_name,omitempty"`
}

type OpenIDAuthenticationParameters struct {

	// How to send token to the server. A list of zero or more methods. Valid values are authorizationHeader and query.
	// +kubebuilder:validation:Optional
	// +listType=set
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// OpenID Connect provider identifier. The name of an OpenID Connect Provider.
	// +kubebuilder:validation:Optional
	OpenIDProviderName *string `json:"openidProviderName" tf:"openid_provider_name,omitempty"`
}

type SubscriptionKeyParameterNamesInitParameters struct {

	// The name of the HTTP Header which should be used for the Subscription Key.
	Header *string `json:"header,omitempty" tf:"header,omitempty"`

	// The name of the QueryString parameter which should be used for the Subscription Key.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type SubscriptionKeyParameterNamesObservation struct {

	// The name of the HTTP Header which should be used for the Subscription Key.
	Header *string `json:"header,omitempty" tf:"header,omitempty"`

	// The name of the QueryString parameter which should be used for the Subscription Key.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type SubscriptionKeyParameterNamesParameters struct {

	// The name of the HTTP Header which should be used for the Subscription Key.
	// +kubebuilder:validation:Optional
	Header *string `json:"header" tf:"header,omitempty"`

	// The name of the QueryString parameter which should be used for the Subscription Key.
	// +kubebuilder:validation:Optional
	Query *string `json:"query" tf:"query,omitempty"`
}

type WsdlSelectorInitParameters struct {

	// The name of endpoint (port) to import from WSDL.
	EndpointName *string `json:"endpointName,omitempty" tf:"endpoint_name,omitempty"`

	// The name of service to import from WSDL.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type WsdlSelectorObservation struct {

	// The name of endpoint (port) to import from WSDL.
	EndpointName *string `json:"endpointName,omitempty" tf:"endpoint_name,omitempty"`

	// The name of service to import from WSDL.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type WsdlSelectorParameters struct {

	// The name of endpoint (port) to import from WSDL.
	// +kubebuilder:validation:Optional
	EndpointName *string `json:"endpointName" tf:"endpoint_name,omitempty"`

	// The name of service to import from WSDL.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// APISpec defines the desired state of API
type APISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider APIInitParameters `json:"initProvider,omitempty"`
}

// APIStatus defines the observed state of API.
type APIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// API is the Schema for the APIs API. Manages an API within an API Management Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APISpec   `json:"spec"`
	Status            APIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIList contains a list of APIs
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

// Repository type metadata.
var (
	API_Kind             = "API"
	API_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: API_Kind}.String()
	API_KindAPIVersion   = API_Kind + "." + CRDGroupVersion.String()
	API_GroupVersionKind = CRDGroupVersion.WithKind(API_Kind)
)

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
