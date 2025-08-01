// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package notificationhubs

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures notificationhubs group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_notification_hub", func(r *config.Resource) {
		r.Kind = "NotificationHub"
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil {
				return diff, nil
			}
			remove := func(prefix string) bool {
				for k := range diff.Attributes {
					if strings.HasPrefix(k, prefix+".") && k != prefix+".#" {
						return false
					}
				}
				return true
			}
			// the underlying Terraform resource implementation currently fails
			// if empty credentials are provided. And in order not to block
			// MR deletions when the referenced secret is deleted before
			// the MR, upjet passes empty credentials if the referenced
			// secret it not found. These are workarounds to prevent the
			// provider from panicking if the referenced credentials
			// secret is not found.
			if remove("gcm_credential") {
				delete(diff.Attributes, "gcm_credential.#")
			}
			if remove("apns_credential") {
				delete(diff.Attributes, "apns_credential.#")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("azurerm_notification_hub_namespace", func(r *config.Resource) {
		r.Kind = "NotificationHubNamespace"
	})
}
