// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"testing"

	rtfake "github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

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

// TestResolveProviderConfigModern covers which namespace a modern managed
// resource's ProviderConfig is looked up in.
//
// The ClusterProviderConfig case is the one that matters. It is the default
// when a managed resource sets no providerConfigRef, and it is cluster scoped,
// so it must resolve from a namespaced managed resource in any namespace. That
// is not automatic for the diff server: its in-memory client keys its store by
// an *exact* (GVK, namespace, name) match and holds the cluster-scoped object
// under an empty namespace. A real API server's RESTMapper-aware client hides
// this, which is why it needs a test here.
//
// rtfake.ModernManaged stands in for a generated managed resource on purpose:
// internal/clients is compiled into every one of the provider's per-family
// binaries, so importing a specific resource family's API package here would
// point the dependency the wrong way. For the same reason the scheme is built
// from the shared ProviderConfig API group alone.
func TestResolveProviderConfigModern(t *testing.T) {
	const (
		mrNamespace     = "team-a"
		clusterSubID    = "cluster-sub"
		namespacedSubID = "namespaced-sub"
	)

	scheme := runtime.NewScheme()
	if err := namespacedv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot build the test scheme: %v", err)
	}

	cases := map[string]struct {
		kind    string
		wantSub string
		wantErr bool
	}{
		"cluster_scoped_resolves_from_a_namespaced_resource": {
			kind:    namespacedv1beta1.ClusterProviderConfigKind,
			wantSub: clusterSubID,
		},
		"namespaced_resolves_from_the_resource_own_namespace": {
			kind:    namespacedv1beta1.ProviderConfigKind,
			wantSub: namespacedSubID,
		},
		"unknown_kind_is_rejected": {
			kind:    "NoSuchProviderConfig",
			wantErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			clusterSub, namespacedSub := clusterSubID, namespacedSubID
			crClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(
				&namespacedv1beta1.ClusterProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "default"},
					Spec:       namespacedv1beta1.ProviderConfigSpec{SubscriptionID: &clusterSub},
				},
				&namespacedv1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: mrNamespace},
					Spec:       namespacedv1beta1.ProviderConfigSpec{SubscriptionID: &namespacedSub},
				},
			).Build()

			mg := &rtfake.ModernManaged{
				ObjectMeta: metav1.ObjectMeta{Name: "example", Namespace: mrNamespace, UID: "test-uid"},
			}
			mg.SetProviderConfigReference(&xpv2.ProviderConfigReference{Kind: tc.kind, Name: "default"})

			got, err := resolveProviderConfigModern(context.Background(), crClient, mg)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected resolving provider config kind %q to fail", tc.kind)
				}
				return
			}
			if err != nil {
				t.Fatalf("resolveProviderConfigModern() returned unexpected error: %v", err)
			}
			if got.SubscriptionID == nil {
				t.Fatal("resolved provider config has no subscription ID, the wrong object was fetched")
			}
			if diff := cmp.Diff(tc.wantSub, *got.SubscriptionID); diff != "" {
				t.Errorf("resolved the wrong provider config (-want +got):\n%s", diff)
			}
		})
	}
}
