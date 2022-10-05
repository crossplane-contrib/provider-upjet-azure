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

package dataprotection

import (
	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures dataprotection group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_blob_storage", func(r *config.Resource) {
		r.References["vault_id"] = config.Reference{
			Type:      "BackupVault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_disk", func(r *config.Resource) {
		r.References["vault_id"] = config.Reference{
			Type:      "BackupVault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_data_protection_backup_policy_postgresql", func(r *config.Resource) {
		r.References["vault_name"] = config.Reference{
			Type: "BackupVault",
		}
	})
}
