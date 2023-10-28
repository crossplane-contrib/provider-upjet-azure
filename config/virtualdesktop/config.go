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

package virtualdesktop

import (
	"github.com/upbound/upjet/pkg/config"
)

const group = "virtualdesktop"

// Configure configures virtual desktop
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("azurerm_virtual_desktop_host_pool", func(r *config.Resource) {
		r.Kind = "HostPool"
		r.ShortGroup = group
	})

	p.AddResourceConfigurator("azurerm_virtual_desktop_host_pool_registration_info", func(r *config.Resource) {
		r.Kind = "HostPoolRegistrationInfo"
		r.ShortGroup = group
	})

	p.AddResourceConfigurator("azurerm_virtual_desktop_application_group", func(r *config.Resource) {
		r.Kind = "ApplicationGroup"
		r.ShortGroup = group
	})

	p.AddResourceConfigurator("azurerm_virtual_desktop_workspace", func(r *config.Resource) {
		r.Kind = "Workspace"
		r.ShortGroup = group
	})

	p.AddResourceConfigurator("azurerm_virtual_desktop_workspace_application_group_association", func(r *config.Resource) {
		r.Kind = "WorkspaceApplicationAssociation"
		r.ShortGroup = group
	})

	p.AddResourceConfigurator("azurerm_virtual_desktop_application", func(r *config.Resource) {
		r.Kind = "Application"
		r.ShortGroup = group
	})

	p.AddResourceConfigurator("azurerm_virtual_desktop_scaling_plan", func(r *config.Resource) {
		r.Kind = "ScalingPlan"
		r.ShortGroup = group
	})
}
