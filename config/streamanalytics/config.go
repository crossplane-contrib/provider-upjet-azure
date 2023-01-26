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

package streamanalytics

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures storagesync group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_stream_analytics_output_synapse", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_function_javascript_uda", func(r *config.Resource) {
		r.References["stream_analytics_job_id"] = config.Reference{
			Type:      "Job",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_output_blob", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_output_mssql", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_output_mssql", func(r *config.Resource) {
		r.References["server"] = config.Reference{
			Type: "github.com/upbound/provider-azure/apis/sql/v1beta1.MSSQLServer",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_output_mssql", func(r *config.Resource) {
		r.References["table"] = config.Reference{
			Type: "github.com/upbound/provider-azure/apis/storage/v1beta1.Table",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_reference_input_blob", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_stream_input_blob", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_stream_input_eventhub", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_stream_input_iothub", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_stream_input_iothub", func(r *config.Resource) {
		r.References["eventhub_consumer_group_name"] = config.Reference{
			Type: "github.com/upbound/provider-azure/apis/eventhub/v1beta1.ConsumerGroup",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_output_servicebus_queue", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
	p.AddResourceConfigurator("azurerm_stream_analytics_output_servicebus_topic", func(r *config.Resource) {
		r.References["stream_analytics_job_name"] = config.Reference{
			Type: "Job",
		}
	})
}
