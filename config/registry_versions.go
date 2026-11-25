// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	_ "embed"
	"fmt"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// v5BreakingAPIs is a newline-delimited list of the Terraform resources
// whose CRD schema changed incompatibly with the Terraform AzureRM v5
// bump. Each of them is served at a new API version, and the version it
// had before stays served and deprecated. The list is immutable: a later
// breaking change needs its own list and its own version step.
//
//go:embed v5-breaking-apis.txt
var v5BreakingAPIs string

// v5BreakingDeprecationRelease is the provider release that deprecated the
// previous API versions of the resources in v5BreakingAPIs.
const v5BreakingDeprecationRelease = "v3.0.0"

// bumpBreakingAPIVersions moves the resources in v5BreakingAPIs to their
// next API version. It must run after the functions that assign the
// current versions.
func bumpBreakingAPIVersions(pc *ujconfig.Provider) {
	next := map[string]string{"v1beta1": "v1beta2", "v1beta2": "v1beta3"}
	for _, name := range strings.Split(strings.TrimSpace(v5BreakingAPIs), "\n") {
		r, ok := pc.Resources[name]
		if !ok {
			panic(fmt.Sprintf("resource %q in v5-breaking-apis.txt is not configured", name))
		}
		prev := r.Version
		nv, ok := next[prev]
		if !ok {
			panic(fmt.Sprintf("resource %q has unexpected API version %q", name, prev))
		}
		r.PreviousVersions = []string{prev}
		r.Version = nv
		r.SetCRDStorageVersion(nv)
		r.ControllerReconcileVersion = nv //nolint:staticcheck // still read by the generator
		if err := r.SetDeprecatedVersion(prev, ujconfig.VersionDeprecation{
			Warning:            fmt.Sprintf("%s is deprecated since the Terraform AzureRM v5 bump, please migrate to %s.", prev, nv),
			DeprecationRelease: v5BreakingDeprecationRelease,
		}); err != nil {
			panic(err)
		}
		pc.Resources[name] = r
	}
}
