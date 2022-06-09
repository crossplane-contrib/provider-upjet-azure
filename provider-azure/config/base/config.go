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

package base

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-azure/config/common"
)

const (
	errFmtNoAttribute    = `"attribute not found: %s`
	errFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// Configure configures the base group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_subscription", func(r *config.Resource) {
		r.ShortGroup = ""
	})

	p.AddResourceConfigurator("azurerm_resource_provider_registration", func(r *config.Resource) {
		r.ShortGroup = ""
	})

	p.AddResourceConfigurator("azurerm_resource_group", func(r *config.Resource) {
		r.Kind = "ResourceGroup"
		r.ShortGroup = ""
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = common.GetNameFromFullyQualifiedID
		// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example
		r.ExternalName.GetIDFn = func(ctx context.Context, name string, _ map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			subID, ok := providerConfig["subscription_id"]
			if !ok {
				return "", errors.Errorf(errFmtNoAttribute, "subscription_id")
			}
			subIDStr, ok := subID.(string)
			if !ok {
				return "", errors.Errorf(errFmtUnexpectedType, "subscription_id")
			}
			return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", subIDStr, name), nil
		}
	})
}
