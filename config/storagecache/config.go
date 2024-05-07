// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package storagecache

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures storagesync group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_hpc_cache_nfs_target", func(r *config.Resource) {
		delete(r.References, "target_host_name")
	})

	p.AddResourceConfigurator("azurerm_hpc_cache_access_policy", func(r *config.Resource) {
		r.OverrideFieldNames["AccessRuleParameters"] = "HPCCacheAccessPolicyAccessRuleParameters"
		r.OverrideFieldNames["AccessRuleInitParameters"] = "HPCCacheAccessPolicyAccessRuleInitParameters"
		r.OverrideFieldNames["AccessRuleObservation"] = "HPCCacheAccessPolicyAccessRuleObservation"
	})
}
