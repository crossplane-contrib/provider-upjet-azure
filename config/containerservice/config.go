/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package containerservice

import (
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
	"github.com/upbound/provider-azure/config/common"
)

// Configure configures kubernetes group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_kubernetes_cluster", func(r *config.Resource) {
		// Note(ezgidemirel): Following fields are not marked as "sensitive" in Terraform cli schema output.
		// We need to configure them explicitly to store in connectionDetails secret.
		r.TerraformResource.Schema["kube_admin_config"].Elem.(*schema.Resource).
			Schema["client_certificate"].Sensitive = true
		r.TerraformResource.Schema["kube_admin_config"].Elem.(*schema.Resource).
			Schema["client_key"].Sensitive = true
		r.TerraformResource.Schema["kube_admin_config"].Elem.(*schema.Resource).
			Schema["cluster_ca_certificate"].Sensitive = true
		r.TerraformResource.Schema["kube_admin_config"].Elem.(*schema.Resource).
			Schema["password"].Sensitive = true
		r.TerraformResource.Schema["kube_config"].Elem.(*schema.Resource).
			Schema["client_certificate"].Sensitive = true
		r.TerraformResource.Schema["kube_config"].Elem.(*schema.Resource).
			Schema["client_key"].Sensitive = true
		r.TerraformResource.Schema["kube_config"].Elem.(*schema.Resource).
			Schema["cluster_ca_certificate"].Sensitive = true
		r.TerraformResource.Schema["kube_config"].Elem.(*schema.Resource).
			Schema["password"].Sensitive = true

		r.Kind = "KubernetesCluster"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"kubelet_identity", "private_link_enabled",
				"api_server_authorized_ip_ranges", "microsoft_defender",
				"oms_agent"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			caData, err := common.GetField(attr, "kube_config[0].cluster_ca_certificate")
			if err != nil {
				return nil, err
			}
			caDataBytes, err := base64.StdEncoding.DecodeString(caData)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot serialize cluster ca data")
			}
			clientCertData, err := common.GetField(attr, "kube_config[0].client_certificate")
			if err != nil {
				return nil, err
			}
			clientCertDataBytes, err := base64.StdEncoding.DecodeString(clientCertData)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot serialize cluster client cert data")
			}
			clientKeyData, err := common.GetField(attr, "kube_config[0].client_key")
			if err != nil {
				return nil, err
			}
			clientKeyDataBytes, err := base64.StdEncoding.DecodeString(clientKeyData)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot serialize cluster client key data")
			}

			if kc, ok := attr["kube_config_raw"].(string); ok {
				return map[string][]byte{
					"kubeconfig":                      []byte(kc),
					"kubeconfig.clustercacertificate": caDataBytes,
					"kubeconfig.clientcertificate":    clientCertDataBytes,
					"kubeconfig.clientkey":            clientKeyDataBytes,
				}, nil
			}
			return nil, nil
		}
		r.MetaResource.ArgumentDocs["api_server_authorized_ip_ranges"] = "Deprecated in favor of `spec.forProvider.apiServerAccessProfile[0].authorizedIpRanges`"
	})

	p.AddResourceConfigurator("azurerm_kubernetes_cluster_node_pool", func(r *config.Resource) {
		r.Kind = "KubernetesClusterNodePool"
		r.ShortGroup = "containerservice"
		r.References["kubernetes_cluster_id"] = config.Reference{
			Type:      "KubernetesCluster",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
