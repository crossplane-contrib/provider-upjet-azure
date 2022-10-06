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

package security

import (
	"github.com/upbound/upjet/pkg/config"

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
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
