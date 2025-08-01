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

type IdentityInitParameters struct {

	// Specifies a list of user managed identity ids to be assigned. Required if type is UserAssigned.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the identity type of the Log Analytics Workspace. Possible values are SystemAssigned (where Azure will generate a Service Principal for you) and UserAssigned where you can specify the Service Principal IDs in the identity_ids field.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of user managed identity ids to be assigned. Required if type is UserAssigned.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Log Analytics Workspace ID.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Log Analytics Workspace ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the identity type of the Log Analytics Workspace. Possible values are SystemAssigned (where Azure will generate a Service Principal for you) and UserAssigned where you can specify the Service Principal IDs in the identity_ids field.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of user managed identity ids to be assigned. Required if type is UserAssigned.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the identity type of the Log Analytics Workspace. Possible values are SystemAssigned (where Azure will generate a Service Principal for you) and UserAssigned where you can specify the Service Principal IDs in the identity_ids field.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type WorkspaceInitParameters struct {

	// Specifies if the log Analytics Workspace allows users accessing to data associated with the resources they have permission to view, without permission to workspace. Defaults to true.
	AllowResourceOnlyPermissions *bool `json:"allowResourceOnlyPermissions,omitempty" tf:"allow_resource_only_permissions,omitempty"`

	// Is Customer Managed Storage mandatory for query management?
	CmkForQueryForced *bool `json:"cmkForQueryForced,omitempty" tf:"cmk_for_query_forced,omitempty"`

	// The workspace daily quota for ingestion in GB. Defaults to -1 (unlimited) if omitted.
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty" tf:"daily_quota_gb,omitempty"`

	// The ID of the Data Collection Rule to use for this workspace.
	DataCollectionRuleID *string `json:"dataCollectionRuleId,omitempty" tf:"data_collection_rule_id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether to remove the data in the Log Analytics Workspace immediately after 30 days.
	ImmediateDataPurgeOn30DaysEnabled *bool `json:"immediateDataPurgeOn30DaysEnabled,omitempty" tf:"immediate_data_purge_on_30_days_enabled,omitempty"`

	// Should the Log Analytics Workspace support ingestion over the Public Internet? Defaults to true.
	InternetIngestionEnabled *bool `json:"internetIngestionEnabled,omitempty" tf:"internet_ingestion_enabled,omitempty"`

	// Should the Log Analytics Workspace support querying over the Public Internet? Defaults to true.
	InternetQueryEnabled *bool `json:"internetQueryEnabled,omitempty" tf:"internet_query_enabled,omitempty"`

	LocalAuthenticationDisabled *bool `json:"localAuthenticationDisabled,omitempty" tf:"local_authentication_disabled,omitempty"`

	// Specifies if the log Analytics workspace should allow local authentication methods in addition to Microsoft Entra (Azure AD). Defaults to true.
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The capacity reservation level in GB for this workspace. Possible values are 100, 200, 300, 400, 500, 1000, 2000 and 5000.
	ReservationCapacityInGbPerDay *float64 `json:"reservationCapacityInGbPerDay,omitempty" tf:"reservation_capacity_in_gb_per_day,omitempty"`

	// The workspace data retention in days. Possible values are between 30 and 730.
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Specifies the SKU of the Log Analytics Workspace. Possible values are PerNode, Premium, Standard, Standalone, Unlimited, CapacityReservation, PerGB2018, and LACluster. Defaults to PerGB2018.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WorkspaceObservation struct {

	// Specifies if the log Analytics Workspace allows users accessing to data associated with the resources they have permission to view, without permission to workspace. Defaults to true.
	AllowResourceOnlyPermissions *bool `json:"allowResourceOnlyPermissions,omitempty" tf:"allow_resource_only_permissions,omitempty"`

	// Is Customer Managed Storage mandatory for query management?
	CmkForQueryForced *bool `json:"cmkForQueryForced,omitempty" tf:"cmk_for_query_forced,omitempty"`

	// The workspace daily quota for ingestion in GB. Defaults to -1 (unlimited) if omitted.
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty" tf:"daily_quota_gb,omitempty"`

	// The ID of the Data Collection Rule to use for this workspace.
	DataCollectionRuleID *string `json:"dataCollectionRuleId,omitempty" tf:"data_collection_rule_id,omitempty"`

	// The Log Analytics Workspace ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether to remove the data in the Log Analytics Workspace immediately after 30 days.
	ImmediateDataPurgeOn30DaysEnabled *bool `json:"immediateDataPurgeOn30DaysEnabled,omitempty" tf:"immediate_data_purge_on_30_days_enabled,omitempty"`

	// Should the Log Analytics Workspace support ingestion over the Public Internet? Defaults to true.
	InternetIngestionEnabled *bool `json:"internetIngestionEnabled,omitempty" tf:"internet_ingestion_enabled,omitempty"`

	// Should the Log Analytics Workspace support querying over the Public Internet? Defaults to true.
	InternetQueryEnabled *bool `json:"internetQueryEnabled,omitempty" tf:"internet_query_enabled,omitempty"`

	LocalAuthenticationDisabled *bool `json:"localAuthenticationDisabled,omitempty" tf:"local_authentication_disabled,omitempty"`

	// Specifies if the log Analytics workspace should allow local authentication methods in addition to Microsoft Entra (Azure AD). Defaults to true.
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The capacity reservation level in GB for this workspace. Possible values are 100, 200, 300, 400, 500, 1000, 2000 and 5000.
	ReservationCapacityInGbPerDay *float64 `json:"reservationCapacityInGbPerDay,omitempty" tf:"reservation_capacity_in_gb_per_day,omitempty"`

	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The workspace data retention in days. Possible values are between 30 and 730.
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Specifies the SKU of the Log Analytics Workspace. Possible values are PerNode, Premium, Standard, Standalone, Unlimited, CapacityReservation, PerGB2018, and LACluster. Defaults to PerGB2018.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type WorkspaceParameters struct {

	// Specifies if the log Analytics Workspace allows users accessing to data associated with the resources they have permission to view, without permission to workspace. Defaults to true.
	// +kubebuilder:validation:Optional
	AllowResourceOnlyPermissions *bool `json:"allowResourceOnlyPermissions,omitempty" tf:"allow_resource_only_permissions,omitempty"`

	// Is Customer Managed Storage mandatory for query management?
	// +kubebuilder:validation:Optional
	CmkForQueryForced *bool `json:"cmkForQueryForced,omitempty" tf:"cmk_for_query_forced,omitempty"`

	// The workspace daily quota for ingestion in GB. Defaults to -1 (unlimited) if omitted.
	// +kubebuilder:validation:Optional
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty" tf:"daily_quota_gb,omitempty"`

	// The ID of the Data Collection Rule to use for this workspace.
	// +kubebuilder:validation:Optional
	DataCollectionRuleID *string `json:"dataCollectionRuleId,omitempty" tf:"data_collection_rule_id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether to remove the data in the Log Analytics Workspace immediately after 30 days.
	// +kubebuilder:validation:Optional
	ImmediateDataPurgeOn30DaysEnabled *bool `json:"immediateDataPurgeOn30DaysEnabled,omitempty" tf:"immediate_data_purge_on_30_days_enabled,omitempty"`

	// Should the Log Analytics Workspace support ingestion over the Public Internet? Defaults to true.
	// +kubebuilder:validation:Optional
	InternetIngestionEnabled *bool `json:"internetIngestionEnabled,omitempty" tf:"internet_ingestion_enabled,omitempty"`

	// Should the Log Analytics Workspace support querying over the Public Internet? Defaults to true.
	// +kubebuilder:validation:Optional
	InternetQueryEnabled *bool `json:"internetQueryEnabled,omitempty" tf:"internet_query_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LocalAuthenticationDisabled *bool `json:"localAuthenticationDisabled,omitempty" tf:"local_authentication_disabled,omitempty"`

	// Specifies if the log Analytics workspace should allow local authentication methods in addition to Microsoft Entra (Azure AD). Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The capacity reservation level in GB for this workspace. Possible values are 100, 200, 300, 400, 500, 1000, 2000 and 5000.
	// +kubebuilder:validation:Optional
	ReservationCapacityInGbPerDay *float64 `json:"reservationCapacityInGbPerDay,omitempty" tf:"reservation_capacity_in_gb_per_day,omitempty"`

	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The workspace data retention in days. Possible values are between 30 and 730.
	// +kubebuilder:validation:Optional
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Specifies the SKU of the Log Analytics Workspace. Possible values are PerNode, Premium, Standard, Standalone, Unlimited, CapacityReservation, PerGB2018, and LACluster. Defaults to PerGB2018.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
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
	InitProvider WorkspaceInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Workspace is the Schema for the Workspaces API. Manages a Log Analytics (formally Operational Insights) Workspace.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   WorkspaceSpec   `json:"spec"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
