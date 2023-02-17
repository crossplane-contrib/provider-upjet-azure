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

package orbital

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures orbital group
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("azurerm_orbital_contact", func(r *config.Resource) {
		r.References["spacecraft_id"] = config.Reference{
			TerraformName: "azurerm_orbital_spacecraft",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["contact_profile_id"] = config.Reference{
			TerraformName: "azurerm_orbital_contact_profile",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
