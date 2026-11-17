// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
)

// v5BreakingAPIs is a newline-delimited list of the Terraform resources
// whose CRD schema changed incompatibly with the Terraform AzureRM v5
// bump. Each of them is served at a new API version, and the version it
// had before stays served and deprecated. The list is immutable: a later
// breaking change needs its own list and its own version step.
//
//go:embed v5-breaking-apis.txt
var v5BreakingAPIs string

// v5APIConversions holds, per Terraform resource in v5BreakingAPIs, how the
// fields of the previous and the new API version relate. It is derived from
// the CRD schemas of the two versions by hack/api-conversions/generate.py
// and is as immutable as the list.
//
//go:embed v5-api-conversions.json
var v5APIConversions []byte

// apiConversion describes the fields that the identity conversion cannot
// copy between the previous and the new API version of a resource. All
// paths are CRD field paths and may contain "[*]" for the elements of a list.
type apiConversion struct {
	// Renames are [previous path, new path] pairs of fields that have the
	// same schema under a different name.
	Renames [][2]string `json:"renames"`
	// PreviousOnly are the top-most fields that only the previous version
	// has.
	PreviousOnly []string `json:"previousOnly"`
	// CurrentOnly are the top-most fields that only the new version has.
	CurrentOnly []string `json:"currentOnly"`
}

// v5BreakingDeprecationRelease is the provider release that deprecated the
// previous API versions of the resources in v5BreakingAPIs.
const v5BreakingDeprecationRelease = "v3.0.0"

// bumpBreakingAPIVersions moves the resources in v5BreakingAPIs to their
// next API version. It must run after the functions that assign the
// current versions.
func bumpBreakingAPIVersions(pc *ujconfig.Provider) {
	next := map[string]string{"v1beta1": "v1beta2", "v1beta2": "v1beta3"}
	conversions := map[string]apiConversion{}
	if err := json.Unmarshal(v5APIConversions, &conversions); err != nil {
		panic(fmt.Sprintf("cannot parse v5-api-conversions.json: %v", err))
	}
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
		c, ok := conversions[name]
		if !ok {
			panic(fmt.Sprintf("resource %q has no entry in v5-api-conversions.json", name))
		}
		r.Conversions = append(r.Conversions, c.conversions(prev, nv)...)
		if err := r.SetDeprecatedVersion(prev, ujconfig.VersionDeprecation{
			Warning:            fmt.Sprintf("%s is deprecated since the Terraform AzureRM v5 bump, please migrate to %s.", prev, nv),
			DeprecationRelease: v5BreakingDeprecationRelease,
		}); err != nil {
			panic(err)
		}
		pc.Resources[name] = r
	}
}

// conversions returns the conversions between the prev and cur API versions.
// The identity conversion copies every field the two versions share. A
// renamed field is copied under its other name. A field that only one
// version has is kept in an annotation while the object is served at the
// other version, so converting an object there and back loses nothing.
func (c apiConversion) conversions(prev, cur string) []conversion.Conversion {
	cs := []conversion.Conversion{
		conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes()),
	}
	for _, r := range c.Renames {
		cs = append(cs,
			conversion.NewFieldRenameConversion(prev, r[0], cur, r[1]),
			conversion.NewFieldRenameConversion(cur, r[1], prev, r[0]))
	}
	for _, p := range c.PreviousOnly {
		cs = append(cs,
			conversion.NewNewlyIntroducedFieldConversion(prev, cur, p, conversion.ToAnnotation),
			conversion.NewNewlyIntroducedFieldConversion(cur, prev, p, conversion.FromAnnotation))
	}
	for _, p := range c.CurrentOnly {
		cs = append(cs,
			conversion.NewNewlyIntroducedFieldConversion(cur, prev, p, conversion.ToAnnotation),
			conversion.NewNewlyIntroducedFieldConversion(prev, cur, p, conversion.FromAnnotation))
	}
	return cs
}
