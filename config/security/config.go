// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package security

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures security group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_advanced_threat_protection", func(r *config.Resource) {
		r.Kind = "AdvancedThreatProtection"
		delete(r.References, "target_resource_id")
	})

	p.AddResourceConfigurator("azurerm_iot_security_device_group", func(r *config.Resource) {
		r.Kind = "IOTSecurityDeviceGroup"
	})

	p.AddResourceConfigurator("azurerm_iot_security_solution", func(r *config.Resource) {
		r.Kind = "IOTSecuritySolution"

		r.References["iothub_ids"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/devices/v1beta1.IOTHub",
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
