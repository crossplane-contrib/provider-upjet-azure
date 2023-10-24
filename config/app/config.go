/*
Copyright 2023 The Crossplane Authors.

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
	p.AddResourceConfigurator("azurerm_container_app", func(r *config.Resource) {
		r.TerraformResource.Schema["secret"].Elem.(*schema.Resource).Schema["value"].Sensitive = true
		r.Kind = "ContainerApp"
		r.References["container_app_environment_id"] = config.Reference{
			Type:      "ContainerAppEnvironment",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_container_app_environment", func(r *config.Resource) {
		r.Kind = "ContainerAppEnvironment"
	})
}
