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

type SentinelLogAnalyticsWorkspaceOnboardingInitParameters struct {

	// Specifies if the Workspace is using Customer managed key. Defaults to false. Changing this forces a new resource to be created.
	CustomerManagedKeyEnabled *bool `json:"customerManagedKeyEnabled,omitempty" tf:"customer_managed_key_enabled,omitempty"`

	// Specifies the Workspace Id. Changing this forces the Log Analytics Workspace off the board and onboard again. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

type SentinelLogAnalyticsWorkspaceOnboardingObservation struct {

	// Specifies if the Workspace is using Customer managed key. Defaults to false. Changing this forces a new resource to be created.
	CustomerManagedKeyEnabled *bool `json:"customerManagedKeyEnabled,omitempty" tf:"customer_managed_key_enabled,omitempty"`

	// The ID of the Security Insights Sentinel Onboarding States.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Workspace Id. Changing this forces the Log Analytics Workspace off the board and onboard again. Changing this forces a new resource to be created.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type SentinelLogAnalyticsWorkspaceOnboardingParameters struct {

	// Specifies if the Workspace is using Customer managed key. Defaults to false. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CustomerManagedKeyEnabled *bool `json:"customerManagedKeyEnabled,omitempty" tf:"customer_managed_key_enabled,omitempty"`

	// Specifies the Workspace Id. Changing this forces the Log Analytics Workspace off the board and onboard again. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// SentinelLogAnalyticsWorkspaceOnboardingSpec defines the desired state of SentinelLogAnalyticsWorkspaceOnboarding
type SentinelLogAnalyticsWorkspaceOnboardingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelLogAnalyticsWorkspaceOnboardingParameters `json:"forProvider"`
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
	InitProvider SentinelLogAnalyticsWorkspaceOnboardingInitParameters `json:"initProvider,omitempty"`
}

// SentinelLogAnalyticsWorkspaceOnboardingStatus defines the observed state of SentinelLogAnalyticsWorkspaceOnboarding.
type SentinelLogAnalyticsWorkspaceOnboardingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelLogAnalyticsWorkspaceOnboardingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SentinelLogAnalyticsWorkspaceOnboarding is the Schema for the SentinelLogAnalyticsWorkspaceOnboardings API. Manages a Security Insights Sentinel Onboarding States.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelLogAnalyticsWorkspaceOnboarding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelLogAnalyticsWorkspaceOnboardingSpec   `json:"spec"`
	Status            SentinelLogAnalyticsWorkspaceOnboardingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelLogAnalyticsWorkspaceOnboardingList contains a list of SentinelLogAnalyticsWorkspaceOnboardings
type SentinelLogAnalyticsWorkspaceOnboardingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelLogAnalyticsWorkspaceOnboarding `json:"items"`
}

// Repository type metadata.
var (
	SentinelLogAnalyticsWorkspaceOnboarding_Kind             = "SentinelLogAnalyticsWorkspaceOnboarding"
	SentinelLogAnalyticsWorkspaceOnboarding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelLogAnalyticsWorkspaceOnboarding_Kind}.String()
	SentinelLogAnalyticsWorkspaceOnboarding_KindAPIVersion   = SentinelLogAnalyticsWorkspaceOnboarding_Kind + "." + CRDGroupVersion.String()
	SentinelLogAnalyticsWorkspaceOnboarding_GroupVersionKind = CRDGroupVersion.WithKind(SentinelLogAnalyticsWorkspaceOnboarding_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelLogAnalyticsWorkspaceOnboarding{}, &SentinelLogAnalyticsWorkspaceOnboardingList{})
}
