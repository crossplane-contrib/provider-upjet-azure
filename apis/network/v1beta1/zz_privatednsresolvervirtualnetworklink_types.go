package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrivateDNSResolverVirtualNetworkLinkInitParameters struct {
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
	Tags     map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSResolverVirtualNetworkLinkObservation struct {
	ID                *string `json:"id,omitempty" tf:"id,omitempty"`
	Location          *string `json:"location,omitempty" tf:"location,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	Tags              map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSResolverVirtualNetworkLinkParameters struct {
	Location          *string `json:"location,omitempty" tf:"location,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	Tags              map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSResolverVirtualNetworkLinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSResolverVirtualNetworkLinkParameters `json:"forProvider"`
	InitProvider    PrivateDNSResolverVirtualNetworkLinkInitParameters `json:"initProvider,omitempty"`
}

type PrivateDNSResolverVirtualNetworkLinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSResolverVirtualNetworkLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

type PrivateDNSResolverVirtualNetworkLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSResolverVirtualNetworkLinkSpec   `json:"spec"`
	Status            PrivateDNSResolverVirtualNetworkLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type PrivateDNSResolverVirtualNetworkLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSResolverVirtualNetworkLink `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSResolverVirtualNetworkLink_Kind             = "PrivateDNSResolverVirtualNetworkLink"
	PrivateDNSResolverVirtualNetworkLink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSResolverVirtualNetworkLink_Kind}.String()
	PrivateDNSResolverVirtualNetworkLink_KindAPIVersion   = PrivateDNSResolverVirtualNetworkLink_Kind + "." + CRDGroupVersion.String()
	PrivateDNSResolverVirtualNetworkLink_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSResolverVirtualNetworkLink_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSResolverVirtualNetworkLink{}, &PrivateDNSResolverVirtualNetworkLinkList{})
}
