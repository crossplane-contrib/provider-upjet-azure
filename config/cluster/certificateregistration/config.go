// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package certificateregistration

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures certificateregistration group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_app_service_certificate_order", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"csr"},
		}
	})
}
