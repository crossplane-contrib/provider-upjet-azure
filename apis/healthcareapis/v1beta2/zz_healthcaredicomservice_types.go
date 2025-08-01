// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthenticationInitParameters struct {
}

type AuthenticationObservation struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://dicom.azurehealthcareapis.azure.com
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`
}

type AuthenticationParameters struct {
}

type CorsInitParameters struct {

	// Whether to allow credentials in CORS. Defaults to false.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// A list of allowed headers for CORS.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// A list of allowed methods for CORS.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A list of allowed origins for CORS.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The maximum age in seconds for the CORS configuration (must be between 0 and 99998 inclusive).
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type CorsObservation struct {

	// Whether to allow credentials in CORS. Defaults to false.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// A list of allowed headers for CORS.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// A list of allowed methods for CORS.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A list of allowed origins for CORS.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The maximum age in seconds for the CORS configuration (must be between 0 and 99998 inclusive).
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type CorsParameters struct {

	// Whether to allow credentials in CORS. Defaults to false.
	// +kubebuilder:validation:Optional
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// A list of allowed headers for CORS.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// A list of allowed methods for CORS.
	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A list of allowed origins for CORS.
	// +kubebuilder:validation:Optional
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The maximum age in seconds for the CORS configuration (must be between 0 and 99998 inclusive).
	// +kubebuilder:validation:Optional
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type HealthcareDICOMServiceInitParameters struct {

	// A cors block as defined below.
	Cors *CorsInitParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// If data partitions are enabled or not. Defaults to false. Changing this forces a new Healthcare DICOM Service to be created.
	DataPartitionsEnabled *bool `json:"dataPartitionsEnabled,omitempty" tf:"data_partitions_enabled,omitempty"`

	// The URL of the key to use for encryption as part of the customer-managed key encryption settings. For more details, refer to the Azure Customer-Managed Keys Overview.
	EncryptionKeyURL *string `json:"encryptionKeyUrl,omitempty" tf:"encryption_key_url,omitempty"`

	// An identity block as defined below.
	Identity *IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Healthcare DICOM Service should be created. Changing this forces a new Healthcare DICOM Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether to enabled public networks when data plane traffic coming from public networks while private endpoint is enabled. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A storage block as defined below.
	Storage *StorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the Healthcare DICOM Service.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthcareDICOMServiceObservation struct {

	// The authentication block as defined below.
	Authentication []AuthenticationObservation `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// A cors block as defined below.
	Cors *CorsObservation `json:"cors,omitempty" tf:"cors,omitempty"`

	// If data partitions are enabled or not. Defaults to false. Changing this forces a new Healthcare DICOM Service to be created.
	DataPartitionsEnabled *bool `json:"dataPartitionsEnabled,omitempty" tf:"data_partitions_enabled,omitempty"`

	// The URL of the key to use for encryption as part of the customer-managed key encryption settings. For more details, refer to the Azure Customer-Managed Keys Overview.
	EncryptionKeyURL *string `json:"encryptionKeyUrl,omitempty" tf:"encryption_key_url,omitempty"`

	// The ID of the Healthcare DICOM Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity *IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Healthcare DICOM Service should be created. Changing this forces a new Healthcare DICOM Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	PrivateEndpoint []PrivateEndpointObservation `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// Whether to enabled public networks when data plane traffic coming from public networks while private endpoint is enabled. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The url of the Healthcare DICOM Services.
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`

	// A storage block as defined below.
	Storage *StorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the Healthcare DICOM Service.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the Healthcare Workspace where the Healthcare DICOM Service should exist. Changing this forces a new Healthcare DICOM Service to be created.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type HealthcareDICOMServiceParameters struct {

	// A cors block as defined below.
	// +kubebuilder:validation:Optional
	Cors *CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// If data partitions are enabled or not. Defaults to false. Changing this forces a new Healthcare DICOM Service to be created.
	// +kubebuilder:validation:Optional
	DataPartitionsEnabled *bool `json:"dataPartitionsEnabled,omitempty" tf:"data_partitions_enabled,omitempty"`

	// The URL of the key to use for encryption as part of the customer-managed key encryption settings. For more details, refer to the Azure Customer-Managed Keys Overview.
	// +kubebuilder:validation:Optional
	EncryptionKeyURL *string `json:"encryptionKeyUrl,omitempty" tf:"encryption_key_url,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Healthcare DICOM Service should be created. Changing this forces a new Healthcare DICOM Service to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether to enabled public networks when data plane traffic coming from public networks while private endpoint is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A storage block as defined below.
	// +kubebuilder:validation:Optional
	Storage *StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the Healthcare DICOM Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the Healthcare Workspace where the Healthcare DICOM Service should exist. Changing this forces a new Healthcare DICOM Service to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/healthcareapis/v1beta1.HealthcareWorkspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a HealthcareWorkspace in healthcareapis to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a HealthcareWorkspace in healthcareapis to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

type IdentityInitParameters struct {

	// A list of User Assigned Identity IDs which should be assigned to this Healthcare DICOM service.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The type of identity used for the Healthcare DICOM service. Possible values are UserAssigned, SystemAssigned and SystemAssigned, UserAssigned. If UserAssigned is set, an identity_ids must be set as well.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// A list of User Assigned Identity IDs which should be assigned to this Healthcare DICOM service.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The ID of the Healthcare DICOM Service.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Healthcare DICOM Service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of identity used for the Healthcare DICOM service. Possible values are UserAssigned, SystemAssigned and SystemAssigned, UserAssigned. If UserAssigned is set, an identity_ids must be set as well.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Identity IDs which should be assigned to this Healthcare DICOM service.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The type of identity used for the Healthcare DICOM service. Possible values are UserAssigned, SystemAssigned and SystemAssigned, UserAssigned. If UserAssigned is set, an identity_ids must be set as well.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type PrivateEndpointInitParameters struct {
}

type PrivateEndpointObservation struct {

	// The ID of the Healthcare DICOM Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Healthcare DICOM Service. Changing this forces a new Healthcare DICOM Service to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PrivateEndpointParameters struct {
}

type StorageInitParameters struct {

	// The filesystem name of connected storage account. Changing this forces a new Healthcare DICOM Service to be created.
	FileSystemName *string `json:"fileSystemName,omitempty" tf:"file_system_name,omitempty"`

	// The resource ID of connected storage account. Changing this forces a new Healthcare DICOM Service to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type StorageObservation struct {

	// The filesystem name of connected storage account. Changing this forces a new Healthcare DICOM Service to be created.
	FileSystemName *string `json:"fileSystemName,omitempty" tf:"file_system_name,omitempty"`

	// The resource ID of connected storage account. Changing this forces a new Healthcare DICOM Service to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type StorageParameters struct {

	// The filesystem name of connected storage account. Changing this forces a new Healthcare DICOM Service to be created.
	// +kubebuilder:validation:Optional
	FileSystemName *string `json:"fileSystemName" tf:"file_system_name,omitempty"`

	// The resource ID of connected storage account. Changing this forces a new Healthcare DICOM Service to be created.
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

// HealthcareDICOMServiceSpec defines the desired state of HealthcareDICOMService
type HealthcareDICOMServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcareDICOMServiceParameters `json:"forProvider"`
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
	InitProvider HealthcareDICOMServiceInitParameters `json:"initProvider,omitempty"`
}

// HealthcareDICOMServiceStatus defines the observed state of HealthcareDICOMService.
type HealthcareDICOMServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcareDICOMServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// HealthcareDICOMService is the Schema for the HealthcareDICOMServices API. Manages a Healthcare DICOM (Digital Imaging and Communications in Medicine) Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HealthcareDICOMService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   HealthcareDICOMServiceSpec   `json:"spec"`
	Status HealthcareDICOMServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareDICOMServiceList contains a list of HealthcareDICOMServices
type HealthcareDICOMServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareDICOMService `json:"items"`
}

// Repository type metadata.
var (
	HealthcareDICOMService_Kind             = "HealthcareDICOMService"
	HealthcareDICOMService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthcareDICOMService_Kind}.String()
	HealthcareDICOMService_KindAPIVersion   = HealthcareDICOMService_Kind + "." + CRDGroupVersion.String()
	HealthcareDICOMService_GroupVersionKind = CRDGroupVersion.WithKind(HealthcareDICOMService_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthcareDICOMService{}, &HealthcareDICOMServiceList{})
}
