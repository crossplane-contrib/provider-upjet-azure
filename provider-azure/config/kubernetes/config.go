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

package kubernetes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
	"github.com/upbound/official-providers/provider-azure/config/common"
)

// Configure configures kubernetes group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_kubernetes_cluster", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2

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
		r.ShortGroup = "containerservice"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"kubelet_identity", "private_link_enabled"},
		}
		r.References = config.References{
			"resource_group_name": config.Reference{
				Type: rconfig.ResourceGroupReferencePath,
			},
			"default_node_pool.pod_subnet_id": config.Reference{
				Type:      rconfig.SubnetReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"default_node_pool.vnet_subnet_id": config.Reference{
				Type:      rconfig.SubnetReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"addon_profile.ingress_application_gateway.subnet_id": config.Reference{
				Type:      rconfig.SubnetReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.ContainerService", "managedClusters", "name")

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
		r.Version = common.VersionV1Alpha2
		r.Kind = "KubernetesClusterNodePool"
		r.ShortGroup = "containerservice"
		r.References = config.References{
			"kubernetes_cluster_id": config.Reference{
				Type:      "KubernetesCluster",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"pod_subnet_id": config.Reference{
				Type:      rconfig.SubnetReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
			"vnet_subnet_id": config.Reference{
				Type:      rconfig.SubnetReferencePath,
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ContainerService/managedClusters/cluster1/agentPools/pool1
		r.ExternalName.GetIDFn = func(_ context.Context, name string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			clusterID, ok := parameters["kubernetes_cluster_id"]
			if !ok {
				return "", errors.Errorf(common.ErrFmtNoAttribute, "kubernetes_cluster_id")
			}
			clusterIDStr, ok := clusterID.(string)
			if !ok {
				return "", errors.Errorf(common.ErrFmtUnexpectedType, "kubernetes_cluster_id")
			}
			return fmt.Sprintf("%s/agentPools/%s", clusterIDStr, name), nil
		}
	})
}
