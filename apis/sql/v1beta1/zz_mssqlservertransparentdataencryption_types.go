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

type MSSQLServerTransparentDataEncryptionInitParameters struct {

	// When enabled, the server will continuously check the key vault for any new versions of the key being used as the TDE protector. If a new version of the key is detected, the TDE protector on the server will be automatically rotated to the latest key version within 60 minutes.
	AutoRotationEnabled *bool `json:"autoRotationEnabled,omitempty" tf:"auto_rotation_enabled,omitempty"`

	// To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta2.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`

	// To use customer managed keys from a managed HSM, provide the Managed HSM Key ID. To use service managed keys, omit this field.
	ManagedHSMKeyID *string `json:"managedHsmKeyId,omitempty" tf:"managed_hsm_key_id,omitempty"`
}

type MSSQLServerTransparentDataEncryptionObservation struct {

	// When enabled, the server will continuously check the key vault for any new versions of the key being used as the TDE protector. If a new version of the key is detected, the TDE protector on the server will be automatically rotated to the latest key version within 60 minutes.
	AutoRotationEnabled *bool `json:"autoRotationEnabled,omitempty" tf:"auto_rotation_enabled,omitempty"`

	// The ID of the MSSQL encryption protector
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// To use customer managed keys from a managed HSM, provide the Managed HSM Key ID. To use service managed keys, omit this field.
	ManagedHSMKeyID *string `json:"managedHsmKeyId,omitempty" tf:"managed_hsm_key_id,omitempty"`

	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type MSSQLServerTransparentDataEncryptionParameters struct {

	// When enabled, the server will continuously check the key vault for any new versions of the key being used as the TDE protector. If a new version of the key is detected, the TDE protector on the server will be automatically rotated to the latest key version within 60 minutes.
	// +kubebuilder:validation:Optional
	AutoRotationEnabled *bool `json:"autoRotationEnabled,omitempty" tf:"auto_rotation_enabled,omitempty"`

	// To use customer managed keys from Azure Key Vault, provide the AKV Key ID. To use service managed keys, omit this field.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta2.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`

	// To use customer managed keys from a managed HSM, provide the Managed HSM Key ID. To use service managed keys, omit this field.
	// +kubebuilder:validation:Optional
	ManagedHSMKeyID *string `json:"managedHsmKeyId,omitempty" tf:"managed_hsm_key_id,omitempty"`

	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/sql/v1beta2.MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a MSSQLServer in sql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer in sql to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// MSSQLServerTransparentDataEncryptionSpec defines the desired state of MSSQLServerTransparentDataEncryption
type MSSQLServerTransparentDataEncryptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLServerTransparentDataEncryptionParameters `json:"forProvider"`
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
	InitProvider MSSQLServerTransparentDataEncryptionInitParameters `json:"initProvider,omitempty"`
}

// MSSQLServerTransparentDataEncryptionStatus defines the observed state of MSSQLServerTransparentDataEncryption.
type MSSQLServerTransparentDataEncryptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLServerTransparentDataEncryptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MSSQLServerTransparentDataEncryption is the Schema for the MSSQLServerTransparentDataEncryptions API. Manages the transparent data encryption configuration for a MSSQL Server
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLServerTransparentDataEncryption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLServerTransparentDataEncryptionSpec   `json:"spec"`
	Status            MSSQLServerTransparentDataEncryptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerTransparentDataEncryptionList contains a list of MSSQLServerTransparentDataEncryptions
type MSSQLServerTransparentDataEncryptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServerTransparentDataEncryption `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServerTransparentDataEncryption_Kind             = "MSSQLServerTransparentDataEncryption"
	MSSQLServerTransparentDataEncryption_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServerTransparentDataEncryption_Kind}.String()
	MSSQLServerTransparentDataEncryption_KindAPIVersion   = MSSQLServerTransparentDataEncryption_Kind + "." + CRDGroupVersion.String()
	MSSQLServerTransparentDataEncryption_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServerTransparentDataEncryption_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServerTransparentDataEncryption{}, &MSSQLServerTransparentDataEncryptionList{})
}
