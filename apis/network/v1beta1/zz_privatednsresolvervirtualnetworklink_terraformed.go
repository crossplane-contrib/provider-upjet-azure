package v1beta1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// PrivateDNSResolverVirtualNetworkLink is the Schema for the PrivateDNSResolverVirtualNetworkLinks API.
type PrivateDNSResolverVirtualNetworkLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSResolverVirtualNetworkLinkSpec   `json:"spec"`
	Status            PrivateDNSResolverVirtualNetworkLinkStatus `json:"status,omitempty"`
}

// GetTerraformResourceType returns Terraform resource type for this PrivateDNSResolverVirtualNetworkLink
func (mg *PrivateDNSResolverVirtualNetworkLink) GetTerraformResourceType() string {
	return "azurerm_private_dns_resolver_virtual_network_link"
}

// GetObservation of this PrivateDNSResolverVirtualNetworkLink
func (tr *PrivateDNSResolverVirtualNetworkLink) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this PrivateDNSResolverVirtualNetworkLink
func (tr *PrivateDNSResolverVirtualNetworkLink) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this PrivateDNSResolverVirtualNetworkLink
func (tr *PrivateDNSResolverVirtualNetworkLink) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this PrivateDNSResolverVirtualNetworkLink
func (tr *PrivateDNSResolverVirtualNetworkLink) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this PrivateDNSResolverVirtualNetworkLink using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *PrivateDNSResolverVirtualNetworkLink) LateInitialize(attrs []byte) (bool, error) {
	params := &PrivateDNSResolverVirtualNetworkLinkParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}
