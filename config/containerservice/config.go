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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
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
			IgnoredFields: []string{"kubelet_identity", "private_link_enabled"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			if kc, ok := attr["kube_config_raw"].(string); ok {
				return map[string][]byte{
					"kubeconfig": []byte(kc),
				}, nil
			}
			return nil, nil
		}
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
