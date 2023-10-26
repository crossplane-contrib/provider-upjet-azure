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

package resources

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_resource_deployment_script_azure_cli", func(r *config.Resource) {
		r.References["identity.identity_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.Path = "resourcedeploymentscriptazureclicli"
	})
	p.AddResourceConfigurator("azurerm_resource_deployment_script_azure_power_shell", func(r *config.Resource) {
		r.References["identity.identity_ids"] = config.Reference{
			Type:      "github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
