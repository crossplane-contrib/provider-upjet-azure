// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package management

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/upbound/provider-azure/config/management"
)

var (
	// PathSubscriptionIDExtractor is the golang path to SubscriptionIDExtractor function
	// in this package.
	PathSubscriptionIDExtractor = SelfPackagePath + ".SubscriptionIDExtractor()"
)

// Configure configures management group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_management_group", func(r *config.Resource) {
		r.Kind = "ManagementGroup"
	})
	p.AddResourceConfigurator("azurerm_management_group_subscription_association", func(r *config.Resource) {
		r.Kind = "ManagementGroupSubscriptionAssociation"
		r.References["management_group_id"] = config.Reference{
			TerraformName: "azurerm_management_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["subscription_id"] = config.Reference{
			TerraformName: "azurerm_subscription",
			Extractor:     PathSubscriptionIDExtractor,
		}
	})
}

// SubscriptionIDExtractor extracts the Subscription ID from "status.atProvider.subscriptionId"
// and returns it in the format "/subscriptions/<subscription-id>".
func SubscriptionIDExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		subID, err := paved.GetString("status.atProvider.subscriptionId")
		if err != nil || subID == "" {
			return ""
		}
		return "/subscriptions/" + subID
	}
}
