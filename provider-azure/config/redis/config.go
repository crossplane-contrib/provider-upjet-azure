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

package redis

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/apis/rconfig"
	"github.com/upbound/official-providers/provider-azure/config/common"
)

// Configure configures redis group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_redis_firewall_rule", func(r *config.Resource) {
		r.References["redis_cache_name"] = config.Reference{
			Type: "RedisCache",
		}
	})
	p.AddResourceConfigurator("azurerm_redis_linked_server", func(r *config.Resource) {
		r.References["linked_redis_cache_id"] = config.Reference{
			Type:      "RedisCache",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["target_redis_cache_name"] = config.Reference{
			Type: "RedisCache",
		}
		delete(r.References, "resource_group_name")
		delete(r.References, "linked_redis_cache_location")
	})

	p.AddResourceConfigurator("azurerm_redis_enterprise_database", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type:      "RedisEnterpriseCluster",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			id, ok := tfstate["id"]
			if !ok {
				return "", errors.New("cannot find id in tfstate")
			}
			l := strings.Split(id.(string), "/")
			return l[len(l)-1], nil
		}
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1/databases/database1
		r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			subIDStr, err := getStr(providerConfig, "subscription_id")
			if err != nil {
				return "", err
			}
			rgStr, err := getStr(parameters, "resource_group_name")
			if err != nil {
				return "", err
			}
			cidStr, err := getStr(parameters, "cluster_id")
			if err != nil {
				return "", err
			}
			cidParts := strings.Split(cidStr, "/")
			cStr := cidParts[len(cidParts)-1]
			nStr, err := getStr(parameters, "name")
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Cache/redisEnterprise/%s/databases/%s",
				subIDStr, rgStr, cStr, nStr), nil
		}
	})
}

func getStr(from map[string]interface{}, key string) (string, error) {
	out, ok := from[key]
	if !ok {
		return "", errors.Errorf(common.ErrFmtNoAttribute, key)
	}
	outStr, ok := out.(string)
	if !ok {
		return "", errors.Errorf(common.ErrFmtUnexpectedType, key)
	}
	return outStr, nil
}
