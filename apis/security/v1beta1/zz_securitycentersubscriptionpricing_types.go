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

type ExtensionInitParameters struct {

	// Key/Value pairs that are required for some extensions.
	// +mapType=granular
	AdditionalExtensionProperties map[string]*string `json:"additionalExtensionProperties,omitempty" tf:"additional_extension_properties,omitempty"`

	// The name of extension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ExtensionObservation struct {

	// Key/Value pairs that are required for some extensions.
	// +mapType=granular
	AdditionalExtensionProperties map[string]*string `json:"additionalExtensionProperties,omitempty" tf:"additional_extension_properties,omitempty"`

	// The name of extension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ExtensionParameters struct {

	// Key/Value pairs that are required for some extensions.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalExtensionProperties map[string]*string `json:"additionalExtensionProperties,omitempty" tf:"additional_extension_properties,omitempty"`

	// The name of extension.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type SecurityCenterSubscriptionPricingInitParameters struct {

	// One or more extension blocks as defined below.
	Extension []ExtensionInitParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// The resource type this setting affects. Possible values are AI, Api, AppServices, ContainerRegistry, KeyVaults, KubernetesService, SqlServers, SqlServerVirtualMachines, StorageAccounts, VirtualMachines, Arm, Dns, OpenSourceRelationalDatabases, Containers, CosmosDbs and CloudPosture. Defaults to VirtualMachines
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Resource type pricing subplan. Contact your MSFT representative for possible values. Changing this forces a new resource to be created.
	Subplan *string `json:"subplan,omitempty" tf:"subplan,omitempty"`

	// The pricing tier to use. Possible values are Free and Standard.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SecurityCenterSubscriptionPricingObservation struct {

	// One or more extension blocks as defined below.
	Extension []ExtensionObservation `json:"extension,omitempty" tf:"extension,omitempty"`

	// The subscription pricing ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource type this setting affects. Possible values are AI, Api, AppServices, ContainerRegistry, KeyVaults, KubernetesService, SqlServers, SqlServerVirtualMachines, StorageAccounts, VirtualMachines, Arm, Dns, OpenSourceRelationalDatabases, Containers, CosmosDbs and CloudPosture. Defaults to VirtualMachines
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Resource type pricing subplan. Contact your MSFT representative for possible values. Changing this forces a new resource to be created.
	Subplan *string `json:"subplan,omitempty" tf:"subplan,omitempty"`

	// The pricing tier to use. Possible values are Free and Standard.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SecurityCenterSubscriptionPricingParameters struct {

	// One or more extension blocks as defined below.
	// +kubebuilder:validation:Optional
	Extension []ExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// The resource type this setting affects. Possible values are AI, Api, AppServices, ContainerRegistry, KeyVaults, KubernetesService, SqlServers, SqlServerVirtualMachines, StorageAccounts, VirtualMachines, Arm, Dns, OpenSourceRelationalDatabases, Containers, CosmosDbs and CloudPosture. Defaults to VirtualMachines
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Resource type pricing subplan. Contact your MSFT representative for possible values. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Subplan *string `json:"subplan,omitempty" tf:"subplan,omitempty"`

	// The pricing tier to use. Possible values are Free and Standard.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

// SecurityCenterSubscriptionPricingSpec defines the desired state of SecurityCenterSubscriptionPricing
type SecurityCenterSubscriptionPricingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterSubscriptionPricingParameters `json:"forProvider"`
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
	InitProvider SecurityCenterSubscriptionPricingInitParameters `json:"initProvider,omitempty"`
}

// SecurityCenterSubscriptionPricingStatus defines the observed state of SecurityCenterSubscriptionPricing.
type SecurityCenterSubscriptionPricingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterSubscriptionPricingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityCenterSubscriptionPricing is the Schema for the SecurityCenterSubscriptionPricings API. Manages the Pricing Tier for Azure Security Center in the current subscription.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterSubscriptionPricing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tier) || (has(self.initProvider) && has(self.initProvider.tier))",message="spec.forProvider.tier is a required parameter"
	Spec   SecurityCenterSubscriptionPricingSpec   `json:"spec"`
	Status SecurityCenterSubscriptionPricingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterSubscriptionPricingList contains a list of SecurityCenterSubscriptionPricings
type SecurityCenterSubscriptionPricingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterSubscriptionPricing `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterSubscriptionPricing_Kind             = "SecurityCenterSubscriptionPricing"
	SecurityCenterSubscriptionPricing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterSubscriptionPricing_Kind}.String()
	SecurityCenterSubscriptionPricing_KindAPIVersion   = SecurityCenterSubscriptionPricing_Kind + "." + CRDGroupVersion.String()
	SecurityCenterSubscriptionPricing_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterSubscriptionPricing_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterSubscriptionPricing{}, &SecurityCenterSubscriptionPricingList{})
}
