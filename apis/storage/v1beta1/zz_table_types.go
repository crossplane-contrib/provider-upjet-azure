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

type ACLAccessPolicyInitParameters struct {

	// The ISO8061 UTC time at which this Access Policy should be valid until.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The permissions which should associated with this Shared Identifier.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The ISO8061 UTC time at which this Access Policy should be valid from.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type ACLAccessPolicyObservation struct {

	// The ISO8061 UTC time at which this Access Policy should be valid until.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The permissions which should associated with this Shared Identifier.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The ISO8061 UTC time at which this Access Policy should be valid from.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type ACLAccessPolicyParameters struct {

	// The ISO8061 UTC time at which this Access Policy should be valid until.
	// +kubebuilder:validation:Optional
	Expiry *string `json:"expiry" tf:"expiry,omitempty"`

	// The permissions which should associated with this Shared Identifier.
	// +kubebuilder:validation:Optional
	Permissions *string `json:"permissions" tf:"permissions,omitempty"`

	// The ISO8061 UTC time at which this Access Policy should be valid from.
	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type TableACLInitParameters struct {

	// An access_policy block as defined below.
	AccessPolicy []ACLAccessPolicyInitParameters `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ID which should be used for this Shared Identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TableACLObservation struct {

	// An access_policy block as defined below.
	AccessPolicy []ACLAccessPolicyObservation `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ID which should be used for this Shared Identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TableACLParameters struct {

	// An access_policy block as defined below.
	// +kubebuilder:validation:Optional
	AccessPolicy []ACLAccessPolicyParameters `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The ID which should be used for this Shared Identifier.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`
}

type TableInitParameters struct {

	// One or more acl blocks as defined below.
	ACL []TableACLInitParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the storage table. Only Alphanumeric characters allowed, starting with a letter. Must be unique within the storage account the table is located. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the storage account in which to create the storage table. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

type TableObservation struct {

	// One or more acl blocks as defined below.
	ACL []TableACLObservation `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ID of the Table within the Storage Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the storage table. Only Alphanumeric characters allowed, starting with a letter. Must be unique within the storage account the table is located. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Resource Manager ID of this Storage Table.
	// The Resource Manager ID of this Storage Table.
	ResourceManagerID *string `json:"resourceManagerId,omitempty" tf:"resource_manager_id,omitempty"`

	// Specifies the storage account in which to create the storage table. Changing this forces a new resource to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`
}

type TableParameters struct {

	// One or more acl blocks as defined below.
	// +kubebuilder:validation:Optional
	ACL []TableACLParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the storage table. Only Alphanumeric characters allowed, starting with a letter. Must be unique within the storage account the table is located. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the storage account in which to create the storage table. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
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
	InitProvider TableInitParameters `json:"initProvider,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Table is the Schema for the Tables API. Manages a Table within an Azure Storage Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TableSpec   `json:"spec"`
	Status TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
