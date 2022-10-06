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

package media

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures virtual group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_media_services_account", func(r *config.Resource) {
		r.References["storage_account.id"] = config.Reference{
			Type:      rconfig.StorageAccountReferencePath,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_media_asset", func(r *config.Resource) {
		r.References["media_services_account_name"] = config.Reference{
			Type: "ServicesAccount",
		}
	})

	p.AddResourceConfigurator("azurerm_media_live_event", func(r *config.Resource) {
		r.References["media_services_account_name"] = config.Reference{
			Type: "ServicesAccount",
		}
	})

	p.AddResourceConfigurator("azurerm_media_live_event_output", func(r *config.Resource) {
		r.References["live_event_id"] = config.Reference{
			Type:      "LiveEvent",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["asset_name"] = config.Reference{
			Type: "Asset",
		}
	})

	p.AddResourceConfigurator("azurerm_media_streaming_endpoint", func(r *config.Resource) {
		r.References["media_services_account_name"] = config.Reference{
			Type: "ServicesAccount",
		}
	})

	p.AddResourceConfigurator("azurerm_media_streaming_locator", func(r *config.Resource) {
		r.References["media_services_account_name"] = config.Reference{
			Type: "ServicesAccount",
		}
		r.References["asset_name"] = config.Reference{
			Type: "Asset",
		}
	})

	p.AddResourceConfigurator("azurerm_media_streaming_policy", func(r *config.Resource) {
		r.References["media_services_account_name"] = config.Reference{
			Type: "ServicesAccount",
		}
	})

	p.AddResourceConfigurator("azurerm_media_transform", func(r *config.Resource) {
		r.References["media_services_account_name"] = config.Reference{
			Type: "ServicesAccount",
		}
	})
}
