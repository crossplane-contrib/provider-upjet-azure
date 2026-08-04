// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"

	namespacedv1beta1 "github.com/upbound/provider-azure/v2/apis/namespaced/v1beta1"
)

func Test_armServiceFromPath(t *testing.T) {
	cases := map[string]struct {
		path string
		want string
	}{
		"arm_service_namespace_exists": {
			path: "/subscriptions/sub1/resourceGroups/rg/providers/Microsoft.Compute/virtualMachines/vm",
			want: "Microsoft.Compute",
		},
		"no_arm_service_resource_group": {
			path: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-rg",
			want: "resourceGroups",
		},
		"no_arm_service_resource_group_collection": {
			path: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups",
			want: "resourceGroups",
		},
		"no_arm_service_subscription": {
			path: "/subscriptions/00000000-0000-0000-0000-000000000000",
			want: "subscriptions",
		},
		"not_valid_arm_service": {
			path: "/some/random/path",
			want: "unknown",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := armServiceFromPath(tc.path)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("armServiceFromPath(%q) mismatch (-want +got):\n%s", tc.path, diff)
			}
		})
	}
}

func Test_oidcAuth_tokenFilePath(t *testing.T) {
	subID, tenantID, clientID := "sub", "tenant", "client"
	explicitPath := "/explicit/path/azure-identity-token"
	envPath := "/var/run/secrets/azure/wi/token/azure-identity-token"

	cases := map[string]struct {
		oidcTokenFilePath *string
		envValue          string
		envSet            bool
		want              string
	}{
		"explicit_path_wins_over_env": {
			oidcTokenFilePath: &explicitPath,
			envValue:          envPath,
			envSet:            true,
			want:              explicitPath,
		},
		"env_used_when_no_explicit_path": {
			envValue: envPath,
			envSet:   true,
			want:     envPath,
		},
		"falls_back_to_default_when_env_unset": {
			want: defaultOidcTokenFilePath,
		},
		"falls_back_to_default_when_env_empty": {
			envValue: "",
			envSet:   true,
			want:     defaultOidcTokenFilePath,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.envSet {
				t.Setenv(envAzureFederatedTokenFile, tc.envValue)
			}
			pcSpec := &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID:    &subID,
				TenantID:          &tenantID,
				ClientID:          &clientID,
				OidcTokenFilePath: tc.oidcTokenFilePath,
			}
			ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}
			if err := oidcAuth(pcSpec, ps); err != nil {
				t.Fatalf("oidcAuth() returned unexpected error: %v", err)
			}
			got, _ := ps.Configuration[keyOidcTokenFilePath].(string)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("oidc_token_file_path mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
