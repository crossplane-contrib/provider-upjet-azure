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

type FrontdoorCustomDomainInitParameters struct {

	// The ID of the Azure DNS Zone which should be used for this Front Door Custom Domain. If you are using Azure to host your DNS domains, you must delegate the domain provider's domain name system (DNS) to an Azure DNS Zone. For more information, see Delegate a domain to Azure DNS. Otherwise, if you're using your own domain provider to handle your DNS, you must validate the Front Door Custom Domain by creating the DNS TXT records manually.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.DNSZone
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// Reference to a DNSZone in network to populate dnsZoneId.
	// +kubebuilder:validation:Optional
	DNSZoneIDRef *v1.Reference `json:"dnsZoneIdRef,omitempty" tf:"-"`

	// Selector for a DNSZone in network to populate dnsZoneId.
	// +kubebuilder:validation:Optional
	DNSZoneIDSelector *v1.Selector `json:"dnsZoneIdSelector,omitempty" tf:"-"`

	// The host name of the domain. The host_name field must be the FQDN of your domain(e.g. contoso.fabrikam.com). Changing this forces a new Front Door Custom Domain to be created.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// A tls block as defined below.
	TLS *TLSInitParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type FrontdoorCustomDomainObservation struct {

	// The ID of the Front Door Profile. Changing this forces a new Front Door Custom Domain to be created.
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// The ID of the Azure DNS Zone which should be used for this Front Door Custom Domain. If you are using Azure to host your DNS domains, you must delegate the domain provider's domain name system (DNS) to an Azure DNS Zone. For more information, see Delegate a domain to Azure DNS. Otherwise, if you're using your own domain provider to handle your DNS, you must validate the Front Door Custom Domain by creating the DNS TXT records manually.
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// The date time that the token expires.
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The host name of the domain. The host_name field must be the FQDN of your domain(e.g. contoso.fabrikam.com). Changing this forces a new Front Door Custom Domain to be created.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Front Door Custom Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A tls block as defined below.
	TLS *TLSObservation `json:"tls,omitempty" tf:"tls,omitempty"`

	// Challenge used for DNS TXT record or file based validation.
	ValidationToken *string `json:"validationToken,omitempty" tf:"validation_token,omitempty"`
}

type FrontdoorCustomDomainParameters struct {

	// The ID of the Front Door Profile. Changing this forces a new Front Door Custom Domain to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorProfile
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// Reference to a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDRef *v1.Reference `json:"cdnFrontdoorProfileIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDSelector *v1.Selector `json:"cdnFrontdoorProfileIdSelector,omitempty" tf:"-"`

	// The ID of the Azure DNS Zone which should be used for this Front Door Custom Domain. If you are using Azure to host your DNS domains, you must delegate the domain provider's domain name system (DNS) to an Azure DNS Zone. For more information, see Delegate a domain to Azure DNS. Otherwise, if you're using your own domain provider to handle your DNS, you must validate the Front Door Custom Domain by creating the DNS TXT records manually.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.DNSZone
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// Reference to a DNSZone in network to populate dnsZoneId.
	// +kubebuilder:validation:Optional
	DNSZoneIDRef *v1.Reference `json:"dnsZoneIdRef,omitempty" tf:"-"`

	// Selector for a DNSZone in network to populate dnsZoneId.
	// +kubebuilder:validation:Optional
	DNSZoneIDSelector *v1.Selector `json:"dnsZoneIdSelector,omitempty" tf:"-"`

	// The host name of the domain. The host_name field must be the FQDN of your domain(e.g. contoso.fabrikam.com). Changing this forces a new Front Door Custom Domain to be created.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// A tls block as defined below.
	// +kubebuilder:validation:Optional
	TLS *TLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type TLSInitParameters struct {

	// Resource ID of the Front Door Secret.
	CdnFrontdoorSecretID *string `json:"cdnFrontdoorSecretId,omitempty" tf:"cdn_frontdoor_secret_id,omitempty"`

	// Defines the source of the SSL certificate. Possible values include CustomerCertificate and ManagedCertificate. Defaults to ManagedCertificate.
	CertificateType *string `json:"certificateType,omitempty" tf:"certificate_type,omitempty"`

	// TLS protocol version that will be used for Https. Possible values are TLS12. Defaults to TLS12.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`
}

type TLSObservation struct {

	// Resource ID of the Front Door Secret.
	CdnFrontdoorSecretID *string `json:"cdnFrontdoorSecretId,omitempty" tf:"cdn_frontdoor_secret_id,omitempty"`

	// Defines the source of the SSL certificate. Possible values include CustomerCertificate and ManagedCertificate. Defaults to ManagedCertificate.
	CertificateType *string `json:"certificateType,omitempty" tf:"certificate_type,omitempty"`

	// TLS protocol version that will be used for Https. Possible values are TLS12. Defaults to TLS12.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`
}

type TLSParameters struct {

	// Resource ID of the Front Door Secret.
	// +kubebuilder:validation:Optional
	CdnFrontdoorSecretID *string `json:"cdnFrontdoorSecretId,omitempty" tf:"cdn_frontdoor_secret_id,omitempty"`

	// Defines the source of the SSL certificate. Possible values include CustomerCertificate and ManagedCertificate. Defaults to ManagedCertificate.
	// +kubebuilder:validation:Optional
	CertificateType *string `json:"certificateType,omitempty" tf:"certificate_type,omitempty"`

	// TLS protocol version that will be used for Https. Possible values are TLS12. Defaults to TLS12.
	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`
}

// FrontdoorCustomDomainSpec defines the desired state of FrontdoorCustomDomain
type FrontdoorCustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorCustomDomainParameters `json:"forProvider"`
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
	InitProvider FrontdoorCustomDomainInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorCustomDomainStatus defines the observed state of FrontdoorCustomDomain.
type FrontdoorCustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorCustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FrontdoorCustomDomain is the Schema for the FrontdoorCustomDomains API. Manages a Front Door (standard/premium) Custom Domain.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorCustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostName) || (has(self.initProvider) && has(self.initProvider.hostName))",message="spec.forProvider.hostName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tls) || (has(self.initProvider) && has(self.initProvider.tls))",message="spec.forProvider.tls is a required parameter"
	Spec   FrontdoorCustomDomainSpec   `json:"spec"`
	Status FrontdoorCustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorCustomDomainList contains a list of FrontdoorCustomDomains
type FrontdoorCustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorCustomDomain `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorCustomDomain_Kind             = "FrontdoorCustomDomain"
	FrontdoorCustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorCustomDomain_Kind}.String()
	FrontdoorCustomDomain_KindAPIVersion   = FrontdoorCustomDomain_Kind + "." + CRDGroupVersion.String()
	FrontdoorCustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorCustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorCustomDomain{}, &FrontdoorCustomDomainList{})
}
