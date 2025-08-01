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

type ServicePlanInitParameters struct {

	// The ID of the App Service Environment to create this Service Plan in.
	AppServiceEnvironmentID *string `json:"appServiceEnvironmentId,omitempty" tf:"app_service_environment_id,omitempty"`

	// The Azure Region where the Service Plan should exist. Changing this forces a new Service Plan to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The maximum number of workers to use in an Elastic SKU Plan or Premium Plan that have premium_plan_auto_scale_enabled set to true. Cannot be set unless using an Elastic or Premium SKU.
	MaximumElasticWorkerCount *float64 `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count,omitempty"`

	// The O/S type for the App Services to be hosted in this plan. Possible values include Windows, Linux, and WindowsContainer. Changing this forces a new resource to be created.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Should Per Site Scaling be enabled. Defaults to false.
	PerSiteScalingEnabled *bool `json:"perSiteScalingEnabled,omitempty" tf:"per_site_scaling_enabled,omitempty"`

	// Should automatic scaling be enabled for the Premium SKU Plan. Defaults to false. Cannot be set unless using a Premium SKU.
	PremiumPlanAutoScaleEnabled *bool `json:"premiumPlanAutoScaleEnabled,omitempty" tf:"premium_plan_auto_scale_enabled,omitempty"`

	// The SKU for the plan. Possible values include B1, B2, B3, D1, F1, I1, I2, I3, I1v2, I1mv2, I2v2, I2mv2, I3v2, I3mv2, I4v2, I4mv2, I5v2, I5mv2, I6v2, P1v2, P2v2, P3v2, P0v3, P1v3, P2v3, P3v3, P1mv3, P2mv3, P3mv3, P4mv3, P5mv3, S1, S2, S3, SHARED, EP1, EP2, EP3, FC1, WS1, WS2, WS3, and Y1.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the AppService.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The number of Workers (instances) to be allocated.
	WorkerCount *float64 `json:"workerCount,omitempty" tf:"worker_count,omitempty"`

	// Should the Service Plan balance across Availability Zones in the region. Changing this forces a new resource to be created.
	ZoneBalancingEnabled *bool `json:"zoneBalancingEnabled,omitempty" tf:"zone_balancing_enabled,omitempty"`
}

type ServicePlanObservation struct {

	// The ID of the App Service Environment to create this Service Plan in.
	AppServiceEnvironmentID *string `json:"appServiceEnvironmentId,omitempty" tf:"app_service_environment_id,omitempty"`

	// The ID of the Service Plan.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A string representing the Kind of Service Plan.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The Azure Region where the Service Plan should exist. Changing this forces a new Service Plan to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The maximum number of workers to use in an Elastic SKU Plan or Premium Plan that have premium_plan_auto_scale_enabled set to true. Cannot be set unless using an Elastic or Premium SKU.
	MaximumElasticWorkerCount *float64 `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count,omitempty"`

	// The O/S type for the App Services to be hosted in this plan. Possible values include Windows, Linux, and WindowsContainer. Changing this forces a new resource to be created.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Should Per Site Scaling be enabled. Defaults to false.
	PerSiteScalingEnabled *bool `json:"perSiteScalingEnabled,omitempty" tf:"per_site_scaling_enabled,omitempty"`

	// Should automatic scaling be enabled for the Premium SKU Plan. Defaults to false. Cannot be set unless using a Premium SKU.
	PremiumPlanAutoScaleEnabled *bool `json:"premiumPlanAutoScaleEnabled,omitempty" tf:"premium_plan_auto_scale_enabled,omitempty"`

	// Whether this is a reserved Service Plan Type. true if os_type is Linux, otherwise false.
	Reserved *bool `json:"reserved,omitempty" tf:"reserved,omitempty"`

	// The name of the Resource Group where the Service Plan should exist. Changing this forces a new Service Plan to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU for the plan. Possible values include B1, B2, B3, D1, F1, I1, I2, I3, I1v2, I1mv2, I2v2, I2mv2, I3v2, I3mv2, I4v2, I4mv2, I5v2, I5mv2, I6v2, P1v2, P2v2, P3v2, P0v3, P1v3, P2v3, P3v3, P1mv3, P2mv3, P3mv3, P4mv3, P5mv3, S1, S2, S3, SHARED, EP1, EP2, EP3, FC1, WS1, WS2, WS3, and Y1.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the AppService.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The number of Workers (instances) to be allocated.
	WorkerCount *float64 `json:"workerCount,omitempty" tf:"worker_count,omitempty"`

	// Should the Service Plan balance across Availability Zones in the region. Changing this forces a new resource to be created.
	ZoneBalancingEnabled *bool `json:"zoneBalancingEnabled,omitempty" tf:"zone_balancing_enabled,omitempty"`
}

type ServicePlanParameters struct {

	// The ID of the App Service Environment to create this Service Plan in.
	// +kubebuilder:validation:Optional
	AppServiceEnvironmentID *string `json:"appServiceEnvironmentId,omitempty" tf:"app_service_environment_id,omitempty"`

	// The Azure Region where the Service Plan should exist. Changing this forces a new Service Plan to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The maximum number of workers to use in an Elastic SKU Plan or Premium Plan that have premium_plan_auto_scale_enabled set to true. Cannot be set unless using an Elastic or Premium SKU.
	// +kubebuilder:validation:Optional
	MaximumElasticWorkerCount *float64 `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count,omitempty"`

	// The O/S type for the App Services to be hosted in this plan. Possible values include Windows, Linux, and WindowsContainer. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Should Per Site Scaling be enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	PerSiteScalingEnabled *bool `json:"perSiteScalingEnabled,omitempty" tf:"per_site_scaling_enabled,omitempty"`

	// Should automatic scaling be enabled for the Premium SKU Plan. Defaults to false. Cannot be set unless using a Premium SKU.
	// +kubebuilder:validation:Optional
	PremiumPlanAutoScaleEnabled *bool `json:"premiumPlanAutoScaleEnabled,omitempty" tf:"premium_plan_auto_scale_enabled,omitempty"`

	// The name of the Resource Group where the Service Plan should exist. Changing this forces a new Service Plan to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU for the plan. Possible values include B1, B2, B3, D1, F1, I1, I2, I3, I1v2, I1mv2, I2v2, I2mv2, I3v2, I3mv2, I4v2, I4mv2, I5v2, I5mv2, I6v2, P1v2, P2v2, P3v2, P0v3, P1v3, P2v3, P3v3, P1mv3, P2mv3, P3mv3, P4mv3, P5mv3, S1, S2, S3, SHARED, EP1, EP2, EP3, FC1, WS1, WS2, WS3, and Y1.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the AppService.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The number of Workers (instances) to be allocated.
	// +kubebuilder:validation:Optional
	WorkerCount *float64 `json:"workerCount,omitempty" tf:"worker_count,omitempty"`

	// Should the Service Plan balance across Availability Zones in the region. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ZoneBalancingEnabled *bool `json:"zoneBalancingEnabled,omitempty" tf:"zone_balancing_enabled,omitempty"`
}

// ServicePlanSpec defines the desired state of ServicePlan
type ServicePlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicePlanParameters `json:"forProvider"`
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
	InitProvider ServicePlanInitParameters `json:"initProvider,omitempty"`
}

// ServicePlanStatus defines the observed state of ServicePlan.
type ServicePlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServicePlan is the Schema for the ServicePlans API. Manages an App Service: Service Plan.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServicePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.osType) || (has(self.initProvider) && has(self.initProvider.osType))",message="spec.forProvider.osType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   ServicePlanSpec   `json:"spec"`
	Status ServicePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicePlanList contains a list of ServicePlans
type ServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicePlan `json:"items"`
}

// Repository type metadata.
var (
	ServicePlan_Kind             = "ServicePlan"
	ServicePlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServicePlan_Kind}.String()
	ServicePlan_KindAPIVersion   = ServicePlan_Kind + "." + CRDGroupVersion.String()
	ServicePlan_GroupVersionKind = CRDGroupVersion.WithKind(ServicePlan_Kind)
)

func init() {
	SchemeBuilder.Register(&ServicePlan{}, &ServicePlanList{})
}
