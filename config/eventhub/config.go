/*
Copyright 2022 The Crossplane Authors.

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

package eventhub

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_eventhub_namespace", func(r *config.Resource) {
		r.Kind = "EventHubNamespace"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"network_rulesets"},
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			Type: "EventHubNamespace",
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub_consumer_group", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			Type: "EventHubNamespace",
		}
		r.References["eventhub_name"] = config.Reference{
			Type: "EventHub",
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub_authorization_rule", func(r *config.Resource) {
		r.References["namespace_name"] = config.Reference{
			Type: "EventHubNamespace",
		}
		r.References["eventhub_name"] = config.Reference{
			Type: "EventHub",
		}
	})
}
