/*
Copyright 2022 Upbound Inc.
*/

package cache

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures redis group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_redis_linked_server", func(r *config.Resource) {
		r.References["linked_redis_cache_id"] = config.Reference{
			Type:      "RedisCache",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["target_redis_cache_name"] = config.Reference{
			Type: "RedisCache",
		}
		delete(r.References, "linked_redis_cache_location")
	})

	p.AddResourceConfigurator("azurerm_redis_cache", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "redis_version")
	})
}
