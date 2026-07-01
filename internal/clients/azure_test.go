// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
