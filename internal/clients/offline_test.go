// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfazureclient "github.com/hashicorp/terraform-provider-azurerm/xpprovider"
)

// azurermResourceGroup is the Terraform resource type the offline tests use.
// It has no Secret-backed or sensitive attributes and no CustomizeDiff of its
// own, which keeps these tests about Configure and basic diffing rather than
// about any one resource's behaviour.
const azurermResourceGroup = "azurerm_resource_group"

// TestOfflineConfigure asserts that the AzureRM Terraform provider can be
// configured with no credentials and no access to Azure. The assertion is
// meaningful because the offline configuration's client secret is a
// placeholder: had the token request left the process, Azure would have
// rejected it and Configure would have failed.
func TestOfflineConfigure(t *testing.T) {
	EnableOfflineAuthentication()
	if _, ok := auth.Client.(offlineTokenClient); !ok {
		t.Fatalf("EnableOfflineAuthentication did not install the offline token client, auth.Client is %T", auth.Client)
	}

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureRM provider schema: %v", err)
	}

	ps, err := configureOffline(context.Background(), p, nil)
	if err != nil {
		t.Fatalf("cannot configure the AzureRM provider offline: %v", err)
	}
	if ps.Meta == nil {
		t.Fatal("offline configuration left terraform.Setup.Meta unset, the provider's CustomizeDiff functions need it")
	}
}

// TestOfflineDiff asserts that an offline configured provider can produce a
// diff, by making the call the upjet Terraform plugin SDK external client makes
// on the diff server's behalf.
func TestOfflineDiff(t *testing.T) {
	EnableOfflineAuthentication()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureRM provider schema: %v", err)
	}
	ps, err := configureOffline(context.Background(), p, nil)
	if err != nil {
		t.Fatalf("cannot configure the AzureRM provider offline: %v", err)
	}

	r := p.ResourcesMap[azurermResourceGroup]
	state := &tfsdk.InstanceState{
		ID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo-rg",
		Attributes: map[string]string{
			"id":         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo-rg",
			"name":       "demo-rg",
			"location":   "westeurope",
			"managed_by": "before",
		},
	}
	config := &tfsdk.ResourceConfig{
		Config: map[string]any{
			"name":       "demo-rg",
			"location":   "westeurope",
			"managed_by": "after",
		},
	}

	diff, err := schema.InternalMap(r.Schema).Diff(context.Background(), state, config, r.CustomizeDiff, ps.Meta, false)
	if err != nil {
		t.Fatalf("cannot diff %s offline: %v", azurermResourceGroup, err)
	}

	got := map[string][2]string{}
	for k, a := range diff.Attributes {
		if a.Old != a.New {
			got[k] = [2]string{a.Old, a.New}
		}
	}
	want := map[string][2]string{
		"managed_by": {"before", "after"},
	}
	if d := cmp.Diff(want, got); d != "" {
		t.Errorf("unexpected diff for %s -want, +got:\n%s", azurermResourceGroup, d)
	}
}
