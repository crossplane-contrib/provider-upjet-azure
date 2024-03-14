// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package media

import (
	"github.com/crossplane/upjet/pkg/config"

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

	p.AddResourceConfigurator("azurerm_media_services_account_filter", func(r *config.Resource) {
		r.References["media_services_account_name"] = config.Reference{
			Type: "ServicesAccount",
		}
	})
}
