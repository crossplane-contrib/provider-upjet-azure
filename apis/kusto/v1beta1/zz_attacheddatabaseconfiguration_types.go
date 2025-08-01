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

type AttachedDatabaseConfigurationInitParameters struct {

	// The resource id of the cluster where the databases you would like to attach reside. Changing this forces a new resource to be created.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// The Kusto Attached Database Configuration ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ClusterResourceID *string `json:"clusterResourceId,omitempty" tf:"cluster_resource_id,omitempty"`

	// Reference to a Cluster in kusto to populate clusterResourceId.
	// +kubebuilder:validation:Optional
	ClusterResourceIDRef *v1.Reference `json:"clusterResourceIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate clusterResourceId.
	// +kubebuilder:validation:Optional
	ClusterResourceIDSelector *v1.Selector `json:"clusterResourceIdSelector,omitempty" tf:"-"`

	// The name of the database which you would like to attach, use * if you want to follow all current and future databases. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Database
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The default principals modification kind. Valid values are: None (default), Replace and Union. Defaults to None.
	DefaultPrincipalModificationKind *string `json:"defaultPrincipalModificationKind,omitempty" tf:"default_principal_modification_kind,omitempty"`

	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A sharing block as defined below.
	Sharing []SharingInitParameters `json:"sharing,omitempty" tf:"sharing,omitempty"`
}

type AttachedDatabaseConfigurationObservation struct {

	// The list of databases from the cluster_resource_id which are currently attached to the cluster.
	AttachedDatabaseNames []*string `json:"attachedDatabaseNames,omitempty" tf:"attached_database_names,omitempty"`

	// The resource id of the cluster where the databases you would like to attach reside. Changing this forces a new resource to be created.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// The Kusto Attached Database Configuration ID.
	ClusterResourceID *string `json:"clusterResourceId,omitempty" tf:"cluster_resource_id,omitempty"`

	// The name of the database which you would like to attach, use * if you want to follow all current and future databases. Changing this forces a new resource to be created.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The default principals modification kind. Valid values are: None (default), Replace and Union. Defaults to None.
	DefaultPrincipalModificationKind *string `json:"defaultPrincipalModificationKind,omitempty" tf:"default_principal_modification_kind,omitempty"`

	// The Kusto Attached Database Configuration ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A sharing block as defined below.
	Sharing []SharingObservation `json:"sharing,omitempty" tf:"sharing,omitempty"`
}

type AttachedDatabaseConfigurationParameters struct {

	// The resource id of the cluster where the databases you would like to attach reside. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// The Kusto Attached Database Configuration ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterResourceID *string `json:"clusterResourceId,omitempty" tf:"cluster_resource_id,omitempty"`

	// Reference to a Cluster in kusto to populate clusterResourceId.
	// +kubebuilder:validation:Optional
	ClusterResourceIDRef *v1.Reference `json:"clusterResourceIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate clusterResourceId.
	// +kubebuilder:validation:Optional
	ClusterResourceIDSelector *v1.Selector `json:"clusterResourceIdSelector,omitempty" tf:"-"`

	// The name of the database which you would like to attach, use * if you want to follow all current and future databases. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Database
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The default principals modification kind. Valid values are: None (default), Replace and Union. Defaults to None.
	// +kubebuilder:validation:Optional
	DefaultPrincipalModificationKind *string `json:"defaultPrincipalModificationKind,omitempty" tf:"default_principal_modification_kind,omitempty"`

	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A sharing block as defined below.
	// +kubebuilder:validation:Optional
	Sharing []SharingParameters `json:"sharing,omitempty" tf:"sharing,omitempty"`
}

type SharingInitParameters struct {

	// List of external tables exclude from the follower database.
	// +listType=set
	ExternalTablesToExclude []*string `json:"externalTablesToExclude,omitempty" tf:"external_tables_to_exclude,omitempty"`

	// List of external tables to include in the follower database.
	// +listType=set
	ExternalTablesToInclude []*string `json:"externalTablesToInclude,omitempty" tf:"external_tables_to_include,omitempty"`

	// List of materialized views exclude from the follower database.
	// +listType=set
	MaterializedViewsToExclude []*string `json:"materializedViewsToExclude,omitempty" tf:"materialized_views_to_exclude,omitempty"`

	// List of materialized views to include in the follower database.
	// +listType=set
	MaterializedViewsToInclude []*string `json:"materializedViewsToInclude,omitempty" tf:"materialized_views_to_include,omitempty"`

	// List of tables to exclude from the follower database.
	// +listType=set
	TablesToExclude []*string `json:"tablesToExclude,omitempty" tf:"tables_to_exclude,omitempty"`

	// List of tables to include in the follower database.
	// +listType=set
	TablesToInclude []*string `json:"tablesToInclude,omitempty" tf:"tables_to_include,omitempty"`
}

type SharingObservation struct {

	// List of external tables exclude from the follower database.
	// +listType=set
	ExternalTablesToExclude []*string `json:"externalTablesToExclude,omitempty" tf:"external_tables_to_exclude,omitempty"`

	// List of external tables to include in the follower database.
	// +listType=set
	ExternalTablesToInclude []*string `json:"externalTablesToInclude,omitempty" tf:"external_tables_to_include,omitempty"`

	// List of materialized views exclude from the follower database.
	// +listType=set
	MaterializedViewsToExclude []*string `json:"materializedViewsToExclude,omitempty" tf:"materialized_views_to_exclude,omitempty"`

	// List of materialized views to include in the follower database.
	// +listType=set
	MaterializedViewsToInclude []*string `json:"materializedViewsToInclude,omitempty" tf:"materialized_views_to_include,omitempty"`

	// List of tables to exclude from the follower database.
	// +listType=set
	TablesToExclude []*string `json:"tablesToExclude,omitempty" tf:"tables_to_exclude,omitempty"`

	// List of tables to include in the follower database.
	// +listType=set
	TablesToInclude []*string `json:"tablesToInclude,omitempty" tf:"tables_to_include,omitempty"`
}

type SharingParameters struct {

	// List of external tables exclude from the follower database.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExternalTablesToExclude []*string `json:"externalTablesToExclude,omitempty" tf:"external_tables_to_exclude,omitempty"`

	// List of external tables to include in the follower database.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExternalTablesToInclude []*string `json:"externalTablesToInclude,omitempty" tf:"external_tables_to_include,omitempty"`

	// List of materialized views exclude from the follower database.
	// +kubebuilder:validation:Optional
	// +listType=set
	MaterializedViewsToExclude []*string `json:"materializedViewsToExclude,omitempty" tf:"materialized_views_to_exclude,omitempty"`

	// List of materialized views to include in the follower database.
	// +kubebuilder:validation:Optional
	// +listType=set
	MaterializedViewsToInclude []*string `json:"materializedViewsToInclude,omitempty" tf:"materialized_views_to_include,omitempty"`

	// List of tables to exclude from the follower database.
	// +kubebuilder:validation:Optional
	// +listType=set
	TablesToExclude []*string `json:"tablesToExclude,omitempty" tf:"tables_to_exclude,omitempty"`

	// List of tables to include in the follower database.
	// +kubebuilder:validation:Optional
	// +listType=set
	TablesToInclude []*string `json:"tablesToInclude,omitempty" tf:"tables_to_include,omitempty"`
}

// AttachedDatabaseConfigurationSpec defines the desired state of AttachedDatabaseConfiguration
type AttachedDatabaseConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachedDatabaseConfigurationParameters `json:"forProvider"`
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
	InitProvider AttachedDatabaseConfigurationInitParameters `json:"initProvider,omitempty"`
}

// AttachedDatabaseConfigurationStatus defines the observed state of AttachedDatabaseConfiguration.
type AttachedDatabaseConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachedDatabaseConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AttachedDatabaseConfiguration is the Schema for the AttachedDatabaseConfigurations API. Manages Kusto / Data Explorer Attached Database Configuration
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AttachedDatabaseConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AttachedDatabaseConfigurationSpec   `json:"spec"`
	Status AttachedDatabaseConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachedDatabaseConfigurationList contains a list of AttachedDatabaseConfigurations
type AttachedDatabaseConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AttachedDatabaseConfiguration `json:"items"`
}

// Repository type metadata.
var (
	AttachedDatabaseConfiguration_Kind             = "AttachedDatabaseConfiguration"
	AttachedDatabaseConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AttachedDatabaseConfiguration_Kind}.String()
	AttachedDatabaseConfiguration_KindAPIVersion   = AttachedDatabaseConfiguration_Kind + "." + CRDGroupVersion.String()
	AttachedDatabaseConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(AttachedDatabaseConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&AttachedDatabaseConfiguration{}, &AttachedDatabaseConfigurationList{})
}
