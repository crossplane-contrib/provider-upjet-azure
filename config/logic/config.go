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

package logic

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures logic group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_integration_service_environment", func(r *config.Resource) {
		r.Kind = "IntegrationServiceEnvironment"

		r.References["virtual_network_subnet_ids"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/network/v1beta1.Subnet",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
