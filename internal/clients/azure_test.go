// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/crossplane/upjet/v2/pkg/terraform"

	namespacedv1beta1 "github.com/upbound/provider-azure/v2/apis/namespaced/v1beta1"
)

func ptr(s string) *string { return &s }

func TestOIDCAuth(t *testing.T) {
	subscriptionID := "sub-123"

	cases := map[string]struct {
		pcSpec  *namespacedv1beta1.ProviderConfigSpec
		envVars map[string]string
		want    terraform.ProviderConfiguration
		wantErr string
	}{
		"AllFieldsInSpec": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID: ptr(subscriptionID),
				TenantID:       ptr("tenant-spec"),
				ClientID:       ptr("client-spec"),
			},
			want: terraform.ProviderConfiguration{
				keyOidcTokenFilePath: defaultOidcTokenFilePath,
				keySubscriptionID:    subscriptionID,
				keyTenantID:          "tenant-spec",
				keyClientID:          "client-spec",
				keyUseOIDC:           "true",
			},
		},
		"ClientIDAndTenantIDFromEnvVars": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID: ptr(subscriptionID),
			},
			envVars: map[string]string{
				envAzureClientID: "client-env",
				envAzureTenantID: "tenant-env",
			},
			want: terraform.ProviderConfiguration{
				keyOidcTokenFilePath: defaultOidcTokenFilePath,
				keySubscriptionID:    subscriptionID,
				keyTenantID:          "tenant-env",
				keyClientID:          "client-env",
				keyUseOIDC:           "true",
			},
		},
		"SpecTakesPrecedenceOverEnvVars": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID: ptr(subscriptionID),
				TenantID:       ptr("tenant-spec"),
				ClientID:       ptr("client-spec"),
			},
			envVars: map[string]string{
				envAzureClientID: "client-env",
				envAzureTenantID: "tenant-env",
			},
			want: terraform.ProviderConfiguration{
				keyOidcTokenFilePath: defaultOidcTokenFilePath,
				keySubscriptionID:    subscriptionID,
				keyTenantID:          "tenant-spec",
				keyClientID:          "client-spec",
				keyUseOIDC:           "true",
			},
		},
		"CustomTokenFilePath": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID:    ptr(subscriptionID),
				TenantID:          ptr("tenant-spec"),
				ClientID:          ptr("client-spec"),
				OidcTokenFilePath: ptr("/custom/token/path"),
			},
			want: terraform.ProviderConfiguration{
				keyOidcTokenFilePath: "/custom/token/path",
				keySubscriptionID:    subscriptionID,
				keyTenantID:          "tenant-spec",
				keyClientID:          "client-spec",
				keyUseOIDC:           "true",
			},
		},
		"MissingSubscriptionID": {
			pcSpec:  &namespacedv1beta1.ProviderConfigSpec{},
			wantErr: errSubscriptionIDNotSet,
		},
		"MissingTenantIDNoEnvVar": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID: ptr(subscriptionID),
				ClientID:       ptr("client-spec"),
			},
			wantErr: errTenantIDNotSet,
		},
		"MissingClientIDNoEnvVar": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID: ptr(subscriptionID),
				TenantID:       ptr("tenant-spec"),
			},
			wantErr: errClientIDNotSet,
		},
		"TenantIDFromEnvClientIDMissing": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID: ptr(subscriptionID),
			},
			envVars: map[string]string{
				envAzureTenantID: "tenant-env",
			},
			wantErr: errClientIDNotSet,
		},
		"ClientIDFromEnvTenantIDMissing": {
			pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID: ptr(subscriptionID),
			},
			envVars: map[string]string{
				envAzureClientID: "client-env",
			},
			wantErr: errTenantIDNotSet,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// Set env vars for this test case.
			for k, v := range tc.envVars {
				t.Setenv(k, v)
			}

			ps := &terraform.Setup{
				Configuration: terraform.ProviderConfiguration{},
			}

			err := oidcAuth(tc.pcSpec, ps)

			if tc.wantErr != "" {
				if diff := cmp.Diff(tc.wantErr, err.Error(), test.EquateConditions()); diff != "" {
					t.Errorf("oidcAuth() error mismatch (-want +got):\n%s", diff)
				}
				return
			}

			if err != nil {
				t.Fatalf("oidcAuth() unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.want, ps.Configuration); diff != "" {
				t.Errorf("oidcAuth() configuration mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
