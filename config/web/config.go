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

package web

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures web group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_linux_web_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_function_app", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
	p.AddResourceConfigurator("azurerm_function_app_slot", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"key_vault_reference_identity_id"},
		}
	})
}
