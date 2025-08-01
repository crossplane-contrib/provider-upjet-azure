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

type IdentifierInitParameters struct {

	// The Offer Name for this Shared Image. Changing this forces a new resource to be created.
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The Publisher Name for this Gallery Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// The Name of the SKU for this Gallery Image. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`
}

type IdentifierObservation struct {

	// The Offer Name for this Shared Image. Changing this forces a new resource to be created.
	Offer *string `json:"offer,omitempty" tf:"offer,omitempty"`

	// The Publisher Name for this Gallery Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// The Name of the SKU for this Gallery Image. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`
}

type IdentifierParameters struct {

	// The Offer Name for this Shared Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// The Publisher Name for this Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// The Name of the SKU for this Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku" tf:"sku,omitempty"`
}

type PurchasePlanInitParameters struct {

	// The Purchase Plan Name for this Shared Image. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Purchase Plan Product for this Gallery Image. Changing this forces a new resource to be created.
	Product *string `json:"product,omitempty" tf:"product,omitempty"`

	// The Purchase Plan Publisher for this Gallery Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type PurchasePlanObservation struct {

	// The Purchase Plan Name for this Shared Image. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Purchase Plan Product for this Gallery Image. Changing this forces a new resource to be created.
	Product *string `json:"product,omitempty" tf:"product,omitempty"`

	// The Purchase Plan Publisher for this Gallery Image. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type PurchasePlanParameters struct {

	// The Purchase Plan Name for this Shared Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Purchase Plan Product for this Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Product *string `json:"product,omitempty" tf:"product,omitempty"`

	// The Purchase Plan Publisher for this Gallery Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type SharedImageInitParameters struct {

	// Specifies if the Shared Image supports Accelerated Network. Changing this forces a new resource to be created.
	AcceleratedNetworkSupportEnabled *bool `json:"acceleratedNetworkSupportEnabled,omitempty" tf:"accelerated_network_support_enabled,omitempty"`

	// CPU architecture supported by an OS. Possible values are x64 and Arm64. Defaults to x64. Changing this forces a new resource to be created.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// Specifies if Confidential Virtual Machines enabled. It will enable all the features of trusted, with higher confidentiality features for isolate machines or encrypted data. Available for Gen2 machines. Changing this forces a new resource to be created.
	ConfidentialVMEnabled *bool `json:"confidentialVmEnabled,omitempty" tf:"confidential_vm_enabled,omitempty"`

	// Specifies if supports creation of both Confidential virtual machines and Gen2 virtual machines with standard security from a compatible Gen2 OS disk VHD or Gen2 Managed image. Changing this forces a new resource to be created.
	ConfidentialVMSupported *bool `json:"confidentialVmSupported,omitempty" tf:"confidential_vm_supported,omitempty"`

	// A description of this Shared Image.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Shared Image supports NVMe disks. Changing this forces a new resource to be created.
	DiskControllerTypeNvmeEnabled *bool `json:"diskControllerTypeNvmeEnabled,omitempty" tf:"disk_controller_type_nvme_enabled,omitempty"`

	// One or more Disk Types not allowed for the Image. Possible values include Standard_LRS and Premium_LRS.
	// +listType=set
	DiskTypesNotAllowed []*string `json:"diskTypesNotAllowed,omitempty" tf:"disk_types_not_allowed,omitempty"`

	// The end of life date in RFC3339 format of the Image.
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty" tf:"end_of_life_date,omitempty"`

	// The End User Licence Agreement for the Shared Image. Changing this forces a new resource to be created.
	Eula *string `json:"eula,omitempty" tf:"eula,omitempty"`

	// Specifies if the Shared Image supports hibernation. Changing this forces a new resource to be created.
	HibernationEnabled *bool `json:"hibernationEnabled,omitempty" tf:"hibernation_enabled,omitempty"`

	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are V1 and V2. Defaults to V1. Changing this forces a new resource to be created.
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// An identifier block as defined below.
	Identifier *IdentifierInitParameters `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Maximum memory in GB recommended for the Image.
	MaxRecommendedMemoryInGb *float64 `json:"maxRecommendedMemoryInGb,omitempty" tf:"max_recommended_memory_in_gb,omitempty"`

	// Maximum count of vCPUs recommended for the Image.
	MaxRecommendedVcpuCount *float64 `json:"maxRecommendedVcpuCount,omitempty" tf:"max_recommended_vcpu_count,omitempty"`

	// Minimum memory in GB recommended for the Image.
	MinRecommendedMemoryInGb *float64 `json:"minRecommendedMemoryInGb,omitempty" tf:"min_recommended_memory_in_gb,omitempty"`

	// Minimum count of vCPUs recommended for the Image.
	MinRecommendedVcpuCount *float64 `json:"minRecommendedVcpuCount,omitempty" tf:"min_recommended_vcpu_count,omitempty"`

	// The type of Operating System present in this Shared Image. Possible values are Linux and Windows. Changing this forces a new resource to be created.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The URI containing the Privacy Statement associated with this Shared Image. Changing this forces a new resource to be created.
	PrivacyStatementURI *string `json:"privacyStatementUri,omitempty" tf:"privacy_statement_uri,omitempty"`

	// A purchase_plan block as defined below.
	PurchasePlan *PurchasePlanInitParameters `json:"purchasePlan,omitempty" tf:"purchase_plan,omitempty"`

	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteURI *string `json:"releaseNoteUri,omitempty" tf:"release_note_uri,omitempty"`

	// Specifies that the Operating System used inside this Image has not been Generalized (for example, sysprep on Windows has not been run). Changing this forces a new resource to be created.
	Specialized *bool `json:"specialized,omitempty" tf:"specialized,omitempty"`

	// A mapping of tags to assign to the Shared Image.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies if Trusted Launch has to be enabled for the Virtual Machine created from the Shared Image. Changing this forces a new resource to be created.
	TrustedLaunchEnabled *bool `json:"trustedLaunchEnabled,omitempty" tf:"trusted_launch_enabled,omitempty"`

	// Specifies if supports creation of both Trusted Launch virtual machines and Gen2 virtual machines with standard security created from the Shared Image. Changing this forces a new resource to be created.
	TrustedLaunchSupported *bool `json:"trustedLaunchSupported,omitempty" tf:"trusted_launch_supported,omitempty"`
}

type SharedImageObservation struct {

	// Specifies if the Shared Image supports Accelerated Network. Changing this forces a new resource to be created.
	AcceleratedNetworkSupportEnabled *bool `json:"acceleratedNetworkSupportEnabled,omitempty" tf:"accelerated_network_support_enabled,omitempty"`

	// CPU architecture supported by an OS. Possible values are x64 and Arm64. Defaults to x64. Changing this forces a new resource to be created.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// Specifies if Confidential Virtual Machines enabled. It will enable all the features of trusted, with higher confidentiality features for isolate machines or encrypted data. Available for Gen2 machines. Changing this forces a new resource to be created.
	ConfidentialVMEnabled *bool `json:"confidentialVmEnabled,omitempty" tf:"confidential_vm_enabled,omitempty"`

	// Specifies if supports creation of both Confidential virtual machines and Gen2 virtual machines with standard security from a compatible Gen2 OS disk VHD or Gen2 Managed image. Changing this forces a new resource to be created.
	ConfidentialVMSupported *bool `json:"confidentialVmSupported,omitempty" tf:"confidential_vm_supported,omitempty"`

	// A description of this Shared Image.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Shared Image supports NVMe disks. Changing this forces a new resource to be created.
	DiskControllerTypeNvmeEnabled *bool `json:"diskControllerTypeNvmeEnabled,omitempty" tf:"disk_controller_type_nvme_enabled,omitempty"`

	// One or more Disk Types not allowed for the Image. Possible values include Standard_LRS and Premium_LRS.
	// +listType=set
	DiskTypesNotAllowed []*string `json:"diskTypesNotAllowed,omitempty" tf:"disk_types_not_allowed,omitempty"`

	// The end of life date in RFC3339 format of the Image.
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty" tf:"end_of_life_date,omitempty"`

	// The End User Licence Agreement for the Shared Image. Changing this forces a new resource to be created.
	Eula *string `json:"eula,omitempty" tf:"eula,omitempty"`

	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName *string `json:"galleryName,omitempty" tf:"gallery_name,omitempty"`

	// Specifies if the Shared Image supports hibernation. Changing this forces a new resource to be created.
	HibernationEnabled *bool `json:"hibernationEnabled,omitempty" tf:"hibernation_enabled,omitempty"`

	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are V1 and V2. Defaults to V1. Changing this forces a new resource to be created.
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// The ID of the Shared Image.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identifier block as defined below.
	Identifier *IdentifierObservation `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Maximum memory in GB recommended for the Image.
	MaxRecommendedMemoryInGb *float64 `json:"maxRecommendedMemoryInGb,omitempty" tf:"max_recommended_memory_in_gb,omitempty"`

	// Maximum count of vCPUs recommended for the Image.
	MaxRecommendedVcpuCount *float64 `json:"maxRecommendedVcpuCount,omitempty" tf:"max_recommended_vcpu_count,omitempty"`

	// Minimum memory in GB recommended for the Image.
	MinRecommendedMemoryInGb *float64 `json:"minRecommendedMemoryInGb,omitempty" tf:"min_recommended_memory_in_gb,omitempty"`

	// Minimum count of vCPUs recommended for the Image.
	MinRecommendedVcpuCount *float64 `json:"minRecommendedVcpuCount,omitempty" tf:"min_recommended_vcpu_count,omitempty"`

	// The type of Operating System present in this Shared Image. Possible values are Linux and Windows. Changing this forces a new resource to be created.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The URI containing the Privacy Statement associated with this Shared Image. Changing this forces a new resource to be created.
	PrivacyStatementURI *string `json:"privacyStatementUri,omitempty" tf:"privacy_statement_uri,omitempty"`

	// A purchase_plan block as defined below.
	PurchasePlan *PurchasePlanObservation `json:"purchasePlan,omitempty" tf:"purchase_plan,omitempty"`

	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteURI *string `json:"releaseNoteUri,omitempty" tf:"release_note_uri,omitempty"`

	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies that the Operating System used inside this Image has not been Generalized (for example, sysprep on Windows has not been run). Changing this forces a new resource to be created.
	Specialized *bool `json:"specialized,omitempty" tf:"specialized,omitempty"`

	// A mapping of tags to assign to the Shared Image.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies if Trusted Launch has to be enabled for the Virtual Machine created from the Shared Image. Changing this forces a new resource to be created.
	TrustedLaunchEnabled *bool `json:"trustedLaunchEnabled,omitempty" tf:"trusted_launch_enabled,omitempty"`

	// Specifies if supports creation of both Trusted Launch virtual machines and Gen2 virtual machines with standard security created from the Shared Image. Changing this forces a new resource to be created.
	TrustedLaunchSupported *bool `json:"trustedLaunchSupported,omitempty" tf:"trusted_launch_supported,omitempty"`
}

type SharedImageParameters struct {

	// Specifies if the Shared Image supports Accelerated Network. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AcceleratedNetworkSupportEnabled *bool `json:"acceleratedNetworkSupportEnabled,omitempty" tf:"accelerated_network_support_enabled,omitempty"`

	// CPU architecture supported by an OS. Possible values are x64 and Arm64. Defaults to x64. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// Specifies if Confidential Virtual Machines enabled. It will enable all the features of trusted, with higher confidentiality features for isolate machines or encrypted data. Available for Gen2 machines. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ConfidentialVMEnabled *bool `json:"confidentialVmEnabled,omitempty" tf:"confidential_vm_enabled,omitempty"`

	// Specifies if supports creation of both Confidential virtual machines and Gen2 virtual machines with standard security from a compatible Gen2 OS disk VHD or Gen2 Managed image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ConfidentialVMSupported *bool `json:"confidentialVmSupported,omitempty" tf:"confidential_vm_supported,omitempty"`

	// A description of this Shared Image.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Shared Image supports NVMe disks. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DiskControllerTypeNvmeEnabled *bool `json:"diskControllerTypeNvmeEnabled,omitempty" tf:"disk_controller_type_nvme_enabled,omitempty"`

	// One or more Disk Types not allowed for the Image. Possible values include Standard_LRS and Premium_LRS.
	// +kubebuilder:validation:Optional
	// +listType=set
	DiskTypesNotAllowed []*string `json:"diskTypesNotAllowed,omitempty" tf:"disk_types_not_allowed,omitempty"`

	// The end of life date in RFC3339 format of the Image.
	// +kubebuilder:validation:Optional
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty" tf:"end_of_life_date,omitempty"`

	// The End User Licence Agreement for the Shared Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Eula *string `json:"eula,omitempty" tf:"eula,omitempty"`

	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta2.SharedImageGallery
	// +kubebuilder:validation:Optional
	GalleryName *string `json:"galleryName,omitempty" tf:"gallery_name,omitempty"`

	// Reference to a SharedImageGallery in compute to populate galleryName.
	// +kubebuilder:validation:Optional
	GalleryNameRef *v1.Reference `json:"galleryNameRef,omitempty" tf:"-"`

	// Selector for a SharedImageGallery in compute to populate galleryName.
	// +kubebuilder:validation:Optional
	GalleryNameSelector *v1.Selector `json:"galleryNameSelector,omitempty" tf:"-"`

	// Specifies if the Shared Image supports hibernation. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	HibernationEnabled *bool `json:"hibernationEnabled,omitempty" tf:"hibernation_enabled,omitempty"`

	// The generation of HyperV that the Virtual Machine used to create the Shared Image is based on. Possible values are V1 and V2. Defaults to V1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	HyperVGeneration *string `json:"hyperVGeneration,omitempty" tf:"hyper_v_generation,omitempty"`

	// An identifier block as defined below.
	// +kubebuilder:validation:Optional
	Identifier *IdentifierParameters `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Maximum memory in GB recommended for the Image.
	// +kubebuilder:validation:Optional
	MaxRecommendedMemoryInGb *float64 `json:"maxRecommendedMemoryInGb,omitempty" tf:"max_recommended_memory_in_gb,omitempty"`

	// Maximum count of vCPUs recommended for the Image.
	// +kubebuilder:validation:Optional
	MaxRecommendedVcpuCount *float64 `json:"maxRecommendedVcpuCount,omitempty" tf:"max_recommended_vcpu_count,omitempty"`

	// Minimum memory in GB recommended for the Image.
	// +kubebuilder:validation:Optional
	MinRecommendedMemoryInGb *float64 `json:"minRecommendedMemoryInGb,omitempty" tf:"min_recommended_memory_in_gb,omitempty"`

	// Minimum count of vCPUs recommended for the Image.
	// +kubebuilder:validation:Optional
	MinRecommendedVcpuCount *float64 `json:"minRecommendedVcpuCount,omitempty" tf:"min_recommended_vcpu_count,omitempty"`

	// The type of Operating System present in this Shared Image. Possible values are Linux and Windows. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The URI containing the Privacy Statement associated with this Shared Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrivacyStatementURI *string `json:"privacyStatementUri,omitempty" tf:"privacy_statement_uri,omitempty"`

	// A purchase_plan block as defined below.
	// +kubebuilder:validation:Optional
	PurchasePlan *PurchasePlanParameters `json:"purchasePlan,omitempty" tf:"purchase_plan,omitempty"`

	// The URI containing the Release Notes associated with this Shared Image.
	// +kubebuilder:validation:Optional
	ReleaseNoteURI *string `json:"releaseNoteUri,omitempty" tf:"release_note_uri,omitempty"`

	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies that the Operating System used inside this Image has not been Generalized (for example, sysprep on Windows has not been run). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Specialized *bool `json:"specialized,omitempty" tf:"specialized,omitempty"`

	// A mapping of tags to assign to the Shared Image.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies if Trusted Launch has to be enabled for the Virtual Machine created from the Shared Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TrustedLaunchEnabled *bool `json:"trustedLaunchEnabled,omitempty" tf:"trusted_launch_enabled,omitempty"`

	// Specifies if supports creation of both Trusted Launch virtual machines and Gen2 virtual machines with standard security created from the Shared Image. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TrustedLaunchSupported *bool `json:"trustedLaunchSupported,omitempty" tf:"trusted_launch_supported,omitempty"`
}

// SharedImageSpec defines the desired state of SharedImage
type SharedImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedImageParameters `json:"forProvider"`
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
	InitProvider SharedImageInitParameters `json:"initProvider,omitempty"`
}

// SharedImageStatus defines the observed state of SharedImage.
type SharedImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SharedImage is the Schema for the SharedImages API. Manages a Shared Image within a Shared Image Gallery.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SharedImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identifier) || (has(self.initProvider) && has(self.initProvider.identifier))",message="spec.forProvider.identifier is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.osType) || (has(self.initProvider) && has(self.initProvider.osType))",message="spec.forProvider.osType is a required parameter"
	Spec   SharedImageSpec   `json:"spec"`
	Status SharedImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedImageList contains a list of SharedImages
type SharedImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedImage `json:"items"`
}

// Repository type metadata.
var (
	SharedImage_Kind             = "SharedImage"
	SharedImage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedImage_Kind}.String()
	SharedImage_KindAPIVersion   = SharedImage_Kind + "." + CRDGroupVersion.String()
	SharedImage_GroupVersionKind = CRDGroupVersion.WithKind(SharedImage_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedImage{}, &SharedImageList{})
}
