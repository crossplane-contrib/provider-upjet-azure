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

package healthcareapis

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure healthcareapis resource group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_healthcare_medtech_service", func(r *config.Resource) {
		r.References["workspace_id"] = config.Reference{
			Type:      "HealthcareWorkspace",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["eventhub_namespace_name"] = config.Reference{
			Type: rconfig.APISPackagePath + "/eventhub/v1beta1.EventHubNamespace",
		}
		r.References["eventhub_name"] = config.Reference{
			Type: rconfig.APISPackagePath + "/eventhub/v1beta1.EventHub",
		}
		r.References["eventhub_consumer_group_name"] = config.Reference{
			Type: rconfig.APISPackagePath + "/eventhub/v1beta1.ConsumerGroup",
		}
	})

	p.AddResourceConfigurator("azurerm_healthcare_medtech_service_fhir_destination", func(r *config.Resource) {
		r.References["medtech_service_id"] = config.Reference{
			Type:      "HealthcareMedtechService",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["destination_fhir_service_id"] = config.Reference{
			Type:      "HealthcareFHIRService",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
