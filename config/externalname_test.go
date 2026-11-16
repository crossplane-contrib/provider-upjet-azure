// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMSSQLManagedInstanceFailoverGroupGetIDFn(t *testing.T) {
	type args struct {
		externalName string
		parameters   map[string]any
	}
	type want struct {
		id  string
		err string
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"Successful": {
			args: args{
				externalName: "example-fog",
				parameters: map[string]any{
					"managed_instance_id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Sql/managedInstances/example-mi",
					"location":            "westeurope",
				},
			},
			want: want{
				id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Sql/locations/westeurope/instanceFailoverGroups/example-fog",
			},
		},
		"LocationNormalized": {
			args: args{
				externalName: "example-fog",
				parameters: map[string]any{
					"managed_instance_id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Sql/managedInstances/example-mi",
					"location":            "West Europe",
				},
			},
			want: want{
				id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Sql/locations/westeurope/instanceFailoverGroups/example-fog",
			},
		},
		"ManagedInstanceIDMissing": {
			args: args{
				externalName: "example-fog",
				parameters: map[string]any{
					"location": "westeurope",
				},
			},
			want: want{
				err: "managed_instance_id is required to construct the resource ID",
			},
		},
		"ManagedInstanceIDEmpty": {
			args: args{
				externalName: "example-fog",
				parameters: map[string]any{
					"managed_instance_id": "",
					"location":            "westeurope",
				},
			},
			want: want{
				err: "managed_instance_id is required to construct the resource ID",
			},
		},
		"ManagedInstanceIDMalformed": {
			args: args{
				externalName: "example-fog",
				parameters: map[string]any{
					"managed_instance_id": "/subscriptions/00000000-0000-0000-0000-000000000000",
					"location":            "westeurope",
				},
			},
			want: want{
				err: "invalid managed_instance_id format: /subscriptions/00000000-0000-0000-0000-000000000000",
			},
		},
		"LocationMissing": {
			args: args{
				externalName: "example-fog",
				parameters: map[string]any{
					"managed_instance_id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Sql/managedInstances/example-mi",
				},
			},
			want: want{
				err: "location is required to construct the resource ID",
			},
		},
	}

	e := TerraformPluginSDKExternalNameConfigs["azurerm_mssql_managed_instance_failover_group"]
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := e.GetIDFn(context.Background(), tc.args.externalName, tc.args.parameters, nil)
			var gotErr string
			if err != nil {
				gotErr = err.Error()
			}
			if diff := cmp.Diff(tc.want.err, gotErr); diff != "" {
				t.Errorf("GetIDFn(...): -want error, +got error:\n%s", diff)
			}
			if diff := cmp.Diff(tc.want.id, got); diff != "" {
				t.Errorf("GetIDFn(...): -want id, +got id:\n%s", diff)
			}
		})
	}
}

func TestKeyVaultURLIDConfGetIDFn(t *testing.T) {
	const vaultID = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.KeyVault/vaults/example-kv"
	type args struct {
		resourceType string
		versionless  bool
		externalName string
		parameters   map[string]any
	}
	cases := map[string]struct {
		args args
		want string
	}{
		"KeyVersioned": {
			args: args{resourceType: "keys", versionless: true, externalName: "example-key/0123456789abcdef0123456789abcdef"},
			want: "https://example-kv.vault.azure.net/keys/example-key/0123456789abcdef0123456789abcdef",
		},
		"KeyNameOnly": {
			args: args{resourceType: "keys", versionless: true, externalName: "example-key"},
			want: "https://example-kv.vault.azure.net/keys/example-key/latest",
		},
		"KeyNameTrailingSlash": {
			args: args{resourceType: "keys", versionless: true, externalName: "example-key/"},
			want: "https://example-kv.vault.azure.net/keys/example-key/latest",
		},
		"KeyEmptyExternalNameNotYetCreated": {
			args: args{resourceType: "keys", versionless: true, externalName: ""},
			want: "",
		},
		"SecretNameOnly": {
			args: args{resourceType: "secrets", versionless: true, externalName: "example-secret"},
			want: "https://example-kv.vault.azure.net/secrets/example-secret/latest",
		},
		"CertificateNameOnlyUnchanged": {
			args: args{resourceType: "certificates", versionless: false, externalName: "example-cert"},
			want: "",
		},
		"CertificateVersionedUnchanged": {
			args: args{resourceType: "certificates", versionless: false, externalName: "example-cert/0123456789abcdef0123456789abcdef"},
			want: "https://example-kv.vault.azure.net/certificates/example-cert/0123456789abcdef0123456789abcdef",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			params := map[string]any{"key_vault_id": vaultID}
			for k, v := range tc.args.parameters {
				params[k] = v
			}
			e := keyVaultURLIDConf(tc.args.resourceType, tc.args.versionless)
			got, err := e.GetIDFn(context.Background(), tc.args.externalName, params, map[string]any{})
			if err != nil {
				t.Fatalf("GetIDFn(...): unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetIDFn(...): -want, +got:\n%s", diff)
			}
		})
	}
}

func TestKeyVaultURLIDConfGetExternalNameFn(t *testing.T) {
	cases := map[string]struct {
		versionless bool
		id          string
		want        string
	}{
		"KeyVersioned": {
			versionless: true,
			id:          "https://example-kv.vault.azure.net/keys/example-key/0123456789abcdef0123456789abcdef",
			want:        "example-key/0123456789abcdef0123456789abcdef",
		},
		"KeyLatestRoundTripsToNameOnly": {
			versionless: true,
			id:          "https://example-kv.vault.azure.net/keys/example-key/latest",
			want:        "example-key",
		},
		"CertificateLatestIsAnOrdinaryVersion": {
			versionless: false,
			id:          "https://example-kv.vault.azure.net/certificates/example-cert/latest",
			want:        "example-cert/latest",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := keyVaultURLIDConf("keys", tc.versionless)
			got, err := e.GetExternalNameFn(map[string]any{"id": tc.id})
			if err != nil {
				t.Fatalf("GetExternalNameFn(...): unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetExternalNameFn(...): -want, +got:\n%s", diff)
			}
		})
	}
}
