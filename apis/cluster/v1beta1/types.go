// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`

	// ClientID is the user-assigned managed identity's ID
	// when Credentials.Source is `InjectedIdentity`. If unset and
	// Credentials.Source is `InjectedIdentity`, then a system-assigned
	// managed identity is used.
	// +optional
	ClientID *string `json:"clientID,omitempty"`

	// SubscriptionID is the Azure subscription ID to be used.
	// If unset, subscription ID from Credentials will be used.
	// Required if Credentials.Source is InjectedIdentity.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionID,omitempty"`

	// TenantID is the Azure AD tenant ID to be used.
	// If unset, tenant ID from Credentials will be used.
	// Required if Credentials.Source is InjectedIdentity.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantID,omitempty"`

	// MSIEndpoint is the optional path to a custom endpoint for
	// Managed Service Identity.
	// +kubebuilder:validation:Optional
	MSIEndpoint *string `json:"msiEndpoint,omitempty"`

	// The Cloud Environment which should be used. Possible values are "public",
	// "usgovernment", "german", and "china". Defaults to "public".
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty"`

	// OIDCTokenFilePath is the optional path to a token file
	// that allows to access a managed identity.
	// +kubebuilder:validation:Optional
	OidcTokenFilePath *string `json:"oidcTokenFilePath,omitempty"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;UserAssignedManagedIdentity;SystemAssignedManagedIdentity;OIDCTokenFile;Upbound;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures the Azure provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,azure}
// +kubebuilder:storageversion
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// Please replace `PROVIDER-NAME` with your actual provider name, like `aws`, `azure`, `gcp`, `alibaba`
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,azure}
// +kubebuilder:storageversion
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
