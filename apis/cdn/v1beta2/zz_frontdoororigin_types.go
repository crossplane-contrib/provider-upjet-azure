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

type FrontdoorOriginInitParameters struct {

	// Specifies whether certificate name checks are enabled for this origin.
	CertificateNameCheckEnabled *bool `json:"certificateNameCheckEnabled,omitempty" tf:"certificate_name_check_enabled,omitempty"`

	// Should the origin be enabled? Possible values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The value of the HTTP port. Must be between 1 and 65535. Defaults to 80.
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// The value of the HTTPS port. Must be between 1 and 65535. Defaults to 443.
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// The IPv4 address, IPv6 address or Domain name of the Origin.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_host",true)
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Reference to a Account in storage to populate hostName.
	// +kubebuilder:validation:Optional
	HostNameRef *v1.Reference `json:"hostNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate hostName.
	// +kubebuilder:validation:Optional
	HostNameSelector *v1.Selector `json:"hostNameSelector,omitempty" tf:"-"`

	// The host header value (an IPv4 address, IPv6 address or Domain name) which is sent to the origin with each request. If unspecified the hostname from the request will be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_host",true)
	OriginHostHeader *string `json:"originHostHeader,omitempty" tf:"origin_host_header,omitempty"`

	// Reference to a Account in storage to populate originHostHeader.
	// +kubebuilder:validation:Optional
	OriginHostHeaderRef *v1.Reference `json:"originHostHeaderRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate originHostHeader.
	// +kubebuilder:validation:Optional
	OriginHostHeaderSelector *v1.Selector `json:"originHostHeaderSelector,omitempty" tf:"-"`

	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy. Must be between 1 and 5 (inclusive). Defaults to 1.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A private_link block as defined below.
	PrivateLink *PrivateLinkInitParameters `json:"privateLink,omitempty" tf:"private_link,omitempty"`

	// The weight of the origin in a given origin group for load balancing. Must be between 1 and 1000. Defaults to 500.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type FrontdoorOriginObservation struct {

	// The ID of the Front Door Origin Group within which this Front Door Origin should exist. Changing this forces a new Front Door Origin to be created.
	CdnFrontdoorOriginGroupID *string `json:"cdnFrontdoorOriginGroupId,omitempty" tf:"cdn_frontdoor_origin_group_id,omitempty"`

	// Specifies whether certificate name checks are enabled for this origin.
	CertificateNameCheckEnabled *bool `json:"certificateNameCheckEnabled,omitempty" tf:"certificate_name_check_enabled,omitempty"`

	// Should the origin be enabled? Possible values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The value of the HTTP port. Must be between 1 and 65535. Defaults to 80.
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// The value of the HTTPS port. Must be between 1 and 65535. Defaults to 443.
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// The IPv4 address, IPv6 address or Domain name of the Origin.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Front Door Origin.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The host header value (an IPv4 address, IPv6 address or Domain name) which is sent to the origin with each request. If unspecified the hostname from the request will be used.
	OriginHostHeader *string `json:"originHostHeader,omitempty" tf:"origin_host_header,omitempty"`

	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy. Must be between 1 and 5 (inclusive). Defaults to 1.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A private_link block as defined below.
	PrivateLink *PrivateLinkObservation `json:"privateLink,omitempty" tf:"private_link,omitempty"`

	// The weight of the origin in a given origin group for load balancing. Must be between 1 and 1000. Defaults to 500.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type FrontdoorOriginParameters struct {

	// The ID of the Front Door Origin Group within which this Front Door Origin should exist. Changing this forces a new Front Door Origin to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta2.FrontdoorOriginGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginGroupID *string `json:"cdnFrontdoorOriginGroupId,omitempty" tf:"cdn_frontdoor_origin_group_id,omitempty"`

	// Reference to a FrontdoorOriginGroup in cdn to populate cdnFrontdoorOriginGroupId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginGroupIDRef *v1.Reference `json:"cdnFrontdoorOriginGroupIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorOriginGroup in cdn to populate cdnFrontdoorOriginGroupId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorOriginGroupIDSelector *v1.Selector `json:"cdnFrontdoorOriginGroupIdSelector,omitempty" tf:"-"`

	// Specifies whether certificate name checks are enabled for this origin.
	// +kubebuilder:validation:Optional
	CertificateNameCheckEnabled *bool `json:"certificateNameCheckEnabled,omitempty" tf:"certificate_name_check_enabled,omitempty"`

	// Should the origin be enabled? Possible values are true or false. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The value of the HTTP port. Must be between 1 and 65535. Defaults to 80.
	// +kubebuilder:validation:Optional
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// The value of the HTTPS port. Must be between 1 and 65535. Defaults to 443.
	// +kubebuilder:validation:Optional
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// The IPv4 address, IPv6 address or Domain name of the Origin.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_host",true)
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Reference to a Account in storage to populate hostName.
	// +kubebuilder:validation:Optional
	HostNameRef *v1.Reference `json:"hostNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate hostName.
	// +kubebuilder:validation:Optional
	HostNameSelector *v1.Selector `json:"hostNameSelector,omitempty" tf:"-"`

	// The host header value (an IPv4 address, IPv6 address or Domain name) which is sent to the origin with each request. If unspecified the hostname from the request will be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_host",true)
	// +kubebuilder:validation:Optional
	OriginHostHeader *string `json:"originHostHeader,omitempty" tf:"origin_host_header,omitempty"`

	// Reference to a Account in storage to populate originHostHeader.
	// +kubebuilder:validation:Optional
	OriginHostHeaderRef *v1.Reference `json:"originHostHeaderRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate originHostHeader.
	// +kubebuilder:validation:Optional
	OriginHostHeaderSelector *v1.Selector `json:"originHostHeaderSelector,omitempty" tf:"-"`

	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy. Must be between 1 and 5 (inclusive). Defaults to 1.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A private_link block as defined below.
	// +kubebuilder:validation:Optional
	PrivateLink *PrivateLinkParameters `json:"privateLink,omitempty" tf:"private_link,omitempty"`

	// The weight of the origin in a given origin group for load balancing. Must be between 1 and 1000. Defaults to 500.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type PrivateLinkInitParameters struct {

	// Specifies the location where the Private Link resource should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("location",false)
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Reference to a Account in storage to populate location.
	// +kubebuilder:validation:Optional
	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate location.
	// +kubebuilder:validation:Optional
	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// The ID of the Azure Resource to connect to via the Private Link.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PrivateLinkTargetID *string `json:"privateLinkTargetId,omitempty" tf:"private_link_target_id,omitempty"`

	// Reference to a Account in storage to populate privateLinkTargetId.
	// +kubebuilder:validation:Optional
	PrivateLinkTargetIDRef *v1.Reference `json:"privateLinkTargetIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate privateLinkTargetId.
	// +kubebuilder:validation:Optional
	PrivateLinkTargetIDSelector *v1.Selector `json:"privateLinkTargetIdSelector,omitempty" tf:"-"`

	// Specifies the request message that will be submitted to the private_link_target_id when requesting the private link endpoint connection. Values must be between 1 and 140 characters in length. Defaults to Access request for CDN FrontDoor Private Link Origin.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// Specifies the type of target for this Private Link Endpoint. Possible values are blob, blob_secondary, Gateway, managedEnvironments, sites, web and web_secondary.
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`
}

type PrivateLinkObservation struct {

	// Specifies the location where the Private Link resource should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Azure Resource to connect to via the Private Link.
	PrivateLinkTargetID *string `json:"privateLinkTargetId,omitempty" tf:"private_link_target_id,omitempty"`

	// Specifies the request message that will be submitted to the private_link_target_id when requesting the private link endpoint connection. Values must be between 1 and 140 characters in length. Defaults to Access request for CDN FrontDoor Private Link Origin.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// Specifies the type of target for this Private Link Endpoint. Possible values are blob, blob_secondary, Gateway, managedEnvironments, sites, web and web_secondary.
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`
}

type PrivateLinkParameters struct {

	// Specifies the location where the Private Link resource should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("location",false)
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Reference to a Account in storage to populate location.
	// +kubebuilder:validation:Optional
	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate location.
	// +kubebuilder:validation:Optional
	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// The ID of the Azure Resource to connect to via the Private Link.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrivateLinkTargetID *string `json:"privateLinkTargetId,omitempty" tf:"private_link_target_id,omitempty"`

	// Reference to a Account in storage to populate privateLinkTargetId.
	// +kubebuilder:validation:Optional
	PrivateLinkTargetIDRef *v1.Reference `json:"privateLinkTargetIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate privateLinkTargetId.
	// +kubebuilder:validation:Optional
	PrivateLinkTargetIDSelector *v1.Selector `json:"privateLinkTargetIdSelector,omitempty" tf:"-"`

	// Specifies the request message that will be submitted to the private_link_target_id when requesting the private link endpoint connection. Values must be between 1 and 140 characters in length. Defaults to Access request for CDN FrontDoor Private Link Origin.
	// +kubebuilder:validation:Optional
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// Specifies the type of target for this Private Link Endpoint. Possible values are blob, blob_secondary, Gateway, managedEnvironments, sites, web and web_secondary.
	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`
}

// FrontdoorOriginSpec defines the desired state of FrontdoorOrigin
type FrontdoorOriginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorOriginParameters `json:"forProvider"`
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
	InitProvider FrontdoorOriginInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorOriginStatus defines the observed state of FrontdoorOrigin.
type FrontdoorOriginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorOriginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FrontdoorOrigin is the Schema for the FrontdoorOrigins API. Manages a Front Door (standard/premium) Origin.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorOrigin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateNameCheckEnabled) || (has(self.initProvider) && has(self.initProvider.certificateNameCheckEnabled))",message="spec.forProvider.certificateNameCheckEnabled is a required parameter"
	Spec   FrontdoorOriginSpec   `json:"spec"`
	Status FrontdoorOriginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorOriginList contains a list of FrontdoorOrigins
type FrontdoorOriginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorOrigin `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorOrigin_Kind             = "FrontdoorOrigin"
	FrontdoorOrigin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorOrigin_Kind}.String()
	FrontdoorOrigin_KindAPIVersion   = FrontdoorOrigin_Kind + "." + CRDGroupVersion.String()
	FrontdoorOrigin_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorOrigin_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorOrigin{}, &FrontdoorOriginList{})
}
