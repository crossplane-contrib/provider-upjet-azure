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

package management

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures management group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_management_group", func(r *config.Resource) {
		r.UseAsync = false
		r.Kind = "ManagementGroup"
	})
	p.AddResourceConfigurator("azurerm_management_group_subscription_association", func(r *config.Resource) {
		r.Kind = "ManagementGroupSubscriptionAssociation"
		r.References["management_group_id"] = config.Reference{
			Type:      "ManagementGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["subscription_id"] = config.Reference{
			Type:      rconfig.APISPackagePath + "/azure/v1beta1.Subscription",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
