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

	"github.com/upbound/official-providers/provider-azure/config/common"
)

// Configure configures resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_eventhub_namespace", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.Kind = "EventHubNamespace"
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Eventhub/namespaces/namespace1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Eventhub",
			"namespaces", "name",
		)
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"network_rulesets"},
		}
	})

	p.AddResourceConfigurator("azurerm_eventhub", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.References = config.References{
			"namespace_name": config.Reference{
				Type: "EventHubNamespace",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Eventhub",
			"namespaces", "namespace_name",
			"eventhubs", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_eventhub_consumer_group", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.References = config.References{
			"namespace_name": config.Reference{
				Type: "EventHubNamespace",
			},
			"eventhub_name": config.Reference{
				Type: "EventHub",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/consumerGroups/consumerGroup1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Eventhub",
			"namespaces", "namespace_name",
			"eventhubs", "eventhub_name",
			"consumergroups", "name",
		)
	})

	p.AddResourceConfigurator("azurerm_eventhub_authorization_rule", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		r.References = config.References{
			"namespace_name": config.Reference{
				Type: "EventHubNamespace",
			},
			"eventhub_name": config.Reference{
				Type: "EventHub",
			},
		}
		r.UseAsync = true
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.EventHub/namespaces/namespace1/eventhubs/eventhub1/authorizationRules/rule1
		r.ExternalName.GetIDFn = common.GetFullyQualifiedIDFn("Microsoft.Eventhub",
			"namespaces", "namespace_name",
			"eventhubs", "eventhub_name",
			"authorizationRules", "name",
		)
	})
}
