// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package notificationhubs

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures notificationhubs group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_notification_hub", func(r *config.Resource) {
		r.Kind = "NotificationHub"
	})

	p.AddResourceConfigurator("azurerm_notification_hub_namespace", func(r *config.Resource) {
		r.Kind = "NotificationHubNamespace"
	})
}
