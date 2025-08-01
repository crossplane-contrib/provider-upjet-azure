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

type BastionHostIPConfigurationInitParameters struct {

	// The name of the IP configuration. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Public IP Address to associate with this Bastion Host. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// Reference to a subnet in which this Bastion Host has been created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type BastionHostIPConfigurationObservation struct {

	// The name of the IP configuration. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Public IP Address to associate with this Bastion Host. Changing this forces a new resource to be created.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a subnet in which this Bastion Host has been created. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type BastionHostIPConfigurationParameters struct {

	// The name of the IP configuration. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Reference to a Public IP Address to associate with this Bastion Host. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// Reference to a subnet in which this Bastion Host has been created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type BastionHostInitParameters struct {

	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to true.
	CopyPasteEnabled *bool `json:"copyPasteEnabled,omitempty" tf:"copy_paste_enabled,omitempty"`

	// Is File Copy feature enabled for the Bastion Host. Defaults to false.
	FileCopyEnabled *bool `json:"fileCopyEnabled,omitempty" tf:"file_copy_enabled,omitempty"`

	// A ip_configuration block as defined below. Changing this forces a new resource to be created.
	IPConfiguration *BastionHostIPConfigurationInitParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Is IP Connect feature enabled for the Bastion Host. Defaults to false.
	IPConnectEnabled *bool `json:"ipConnectEnabled,omitempty" tf:"ip_connect_enabled,omitempty"`

	// Is Kerberos authentication feature enabled for the Bastion Host. Defaults to false.
	KerberosEnabled *bool `json:"kerberosEnabled,omitempty" tf:"kerberos_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Review Azure Bastion Host FAQ for supported locations.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The number of scale units with which to provision the Bastion Host. Possible values are between 2 and 50. Defaults to 2.
	ScaleUnits *float64 `json:"scaleUnits,omitempty" tf:"scale_units,omitempty"`

	// Is Session Recording feature enabled for the Bastion Host. Defaults to false.
	SessionRecordingEnabled *bool `json:"sessionRecordingEnabled,omitempty" tf:"session_recording_enabled,omitempty"`

	// Is Shareable Link feature enabled for the Bastion Host. Defaults to false.
	ShareableLinkEnabled *bool `json:"shareableLinkEnabled,omitempty" tf:"shareable_link_enabled,omitempty"`

	// The SKU of the Bastion Host. Accepted values are Developer, Basic, Standard and Premium. Defaults to Basic.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Is Tunneling feature enabled for the Bastion Host. Defaults to false.
	TunnelingEnabled *bool `json:"tunnelingEnabled,omitempty" tf:"tunneling_enabled,omitempty"`

	// The ID of the Virtual Network for the Developer Bastion Host. Changing this forces a new resource to be created.
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`

	// Specifies a list of Availability Zones in which this Public Bastion Host should be located. Changing this forces a new resource to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type BastionHostObservation struct {

	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to true.
	CopyPasteEnabled *bool `json:"copyPasteEnabled,omitempty" tf:"copy_paste_enabled,omitempty"`

	// The FQDN for the Bastion Host.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// Is File Copy feature enabled for the Bastion Host. Defaults to false.
	FileCopyEnabled *bool `json:"fileCopyEnabled,omitempty" tf:"file_copy_enabled,omitempty"`

	// The ID of the Bastion Host.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A ip_configuration block as defined below. Changing this forces a new resource to be created.
	IPConfiguration *BastionHostIPConfigurationObservation `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Is IP Connect feature enabled for the Bastion Host. Defaults to false.
	IPConnectEnabled *bool `json:"ipConnectEnabled,omitempty" tf:"ip_connect_enabled,omitempty"`

	// Is Kerberos authentication feature enabled for the Bastion Host. Defaults to false.
	KerberosEnabled *bool `json:"kerberosEnabled,omitempty" tf:"kerberos_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Review Azure Bastion Host FAQ for supported locations.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Bastion Host. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The number of scale units with which to provision the Bastion Host. Possible values are between 2 and 50. Defaults to 2.
	ScaleUnits *float64 `json:"scaleUnits,omitempty" tf:"scale_units,omitempty"`

	// Is Session Recording feature enabled for the Bastion Host. Defaults to false.
	SessionRecordingEnabled *bool `json:"sessionRecordingEnabled,omitempty" tf:"session_recording_enabled,omitempty"`

	// Is Shareable Link feature enabled for the Bastion Host. Defaults to false.
	ShareableLinkEnabled *bool `json:"shareableLinkEnabled,omitempty" tf:"shareable_link_enabled,omitempty"`

	// The SKU of the Bastion Host. Accepted values are Developer, Basic, Standard and Premium. Defaults to Basic.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Is Tunneling feature enabled for the Bastion Host. Defaults to false.
	TunnelingEnabled *bool `json:"tunnelingEnabled,omitempty" tf:"tunneling_enabled,omitempty"`

	// The ID of the Virtual Network for the Developer Bastion Host. Changing this forces a new resource to be created.
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`

	// Specifies a list of Availability Zones in which this Public Bastion Host should be located. Changing this forces a new resource to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type BastionHostParameters struct {

	// Is Copy/Paste feature enabled for the Bastion Host. Defaults to true.
	// +kubebuilder:validation:Optional
	CopyPasteEnabled *bool `json:"copyPasteEnabled,omitempty" tf:"copy_paste_enabled,omitempty"`

	// Is File Copy feature enabled for the Bastion Host. Defaults to false.
	// +kubebuilder:validation:Optional
	FileCopyEnabled *bool `json:"fileCopyEnabled,omitempty" tf:"file_copy_enabled,omitempty"`

	// A ip_configuration block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IPConfiguration *BastionHostIPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// Is IP Connect feature enabled for the Bastion Host. Defaults to false.
	// +kubebuilder:validation:Optional
	IPConnectEnabled *bool `json:"ipConnectEnabled,omitempty" tf:"ip_connect_enabled,omitempty"`

	// Is Kerberos authentication feature enabled for the Bastion Host. Defaults to false.
	// +kubebuilder:validation:Optional
	KerberosEnabled *bool `json:"kerberosEnabled,omitempty" tf:"kerberos_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created. Review Azure Bastion Host FAQ for supported locations.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Bastion Host. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The number of scale units with which to provision the Bastion Host. Possible values are between 2 and 50. Defaults to 2.
	// +kubebuilder:validation:Optional
	ScaleUnits *float64 `json:"scaleUnits,omitempty" tf:"scale_units,omitempty"`

	// Is Session Recording feature enabled for the Bastion Host. Defaults to false.
	// +kubebuilder:validation:Optional
	SessionRecordingEnabled *bool `json:"sessionRecordingEnabled,omitempty" tf:"session_recording_enabled,omitempty"`

	// Is Shareable Link feature enabled for the Bastion Host. Defaults to false.
	// +kubebuilder:validation:Optional
	ShareableLinkEnabled *bool `json:"shareableLinkEnabled,omitempty" tf:"shareable_link_enabled,omitempty"`

	// The SKU of the Bastion Host. Accepted values are Developer, Basic, Standard and Premium. Defaults to Basic.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Is Tunneling feature enabled for the Bastion Host. Defaults to false.
	// +kubebuilder:validation:Optional
	TunnelingEnabled *bool `json:"tunnelingEnabled,omitempty" tf:"tunneling_enabled,omitempty"`

	// The ID of the Virtual Network for the Developer Bastion Host. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VirtualNetworkID *string `json:"virtualNetworkId,omitempty" tf:"virtual_network_id,omitempty"`

	// Specifies a list of Availability Zones in which this Public Bastion Host should be located. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// BastionHostSpec defines the desired state of BastionHost
type BastionHostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BastionHostParameters `json:"forProvider"`
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
	InitProvider BastionHostInitParameters `json:"initProvider,omitempty"`
}

// BastionHostStatus defines the observed state of BastionHost.
type BastionHostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BastionHostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BastionHost is the Schema for the BastionHosts API. Manages a Bastion Host.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BastionHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   BastionHostSpec   `json:"spec"`
	Status BastionHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BastionHostList contains a list of BastionHosts
type BastionHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BastionHost `json:"items"`
}

// Repository type metadata.
var (
	BastionHost_Kind             = "BastionHost"
	BastionHost_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BastionHost_Kind}.String()
	BastionHost_KindAPIVersion   = BastionHost_Kind + "." + CRDGroupVersion.String()
	BastionHost_GroupVersionKind = CRDGroupVersion.WithKind(BastionHost_Kind)
)

func init() {
	SchemeBuilder.Register(&BastionHost{}, &BastionHostList{})
}
