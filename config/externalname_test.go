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
