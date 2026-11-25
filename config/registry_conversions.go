// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	_ "embed"
	"slices"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

// crdSchemaChanges holds the schema changes between the consecutive API
// versions of each CRD. upjet's schemadiff generates it from the CRDs, see
// generate/generate.go.
//
//go:embed crd-schema-changes.json
var crdSchemaChanges []byte

// clusterFieldRenames holds the renamed fields of the cluster-scoped
// resources.
//
//go:embed cluster/field-rename.yaml
var clusterFieldRenames []byte

// namespacedFieldRenames holds the renamed fields of the namespaced resources.
//
//go:embed namespaced/field-rename.yaml
var namespacedFieldRenames []byte

// fieldRenames lists, per Terraform resource, the fields that have another
// name in the target API version than in the source API version. Both the
// keys and the values of a resource's map are CRD field paths.
type fieldRenames struct {
	SourceVersion string                       `json:"sourceVersion"`
	TargetVersion string                       `json:"targetVersion"`
	Data          map[string]map[string]string `json:"data"`
}

// registerAPIConversions registers the conversions between the API versions
// of the resources. A renamed field is copied under its other name. A field
// that only one version has is kept in an annotation while the object is
// served at the other version, and a field whose type differs is converted,
// both as reported in crdSchemaChanges. It must run after
// pc.ConfigureResources(), which assigns the API versions, groups and kinds.
func registerAPIConversions(pc *ujconfig.Provider, fieldRenameData []byte, generationProvider bool) error {
	if err := ujconfig.ExcludeTypeChangesFromIdentity(pc, crdSchemaChanges); err != nil {
		return errors.Wrap(err, "cannot exclude the fields with a type change from the identity conversion")
	}
	for _, r := range pc.Resources {
		if paths := r.AutoConversionRegistrationOptions.IdentityConversionExcludePaths; len(paths) != 0 {
			// the first conversion is the identity conversion of the default resource
			r.Conversions[0] = conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), paths...)
		}
	}
	if err := registerFieldRenameConversions(pc, fieldRenameData); err != nil {
		return err
	}
	// the generator does not convert objects
	if generationProvider {
		return nil
	}
	return errors.Wrap(ujconfig.RegisterAutoConversions(pc, crdSchemaChanges), "cannot register the conversions for the CRD schema changes")
}

// registerFieldRenameConversions registers a conversion in both directions
// for each renamed field. It also keeps both names of the field out of the
// conversions registered from crdSchemaChanges, which reports a renamed field
// as a removed and an added one.
func registerFieldRenameConversions(pc *ujconfig.Provider, fieldRenameData []byte) error {
	var renames []fieldRenames
	if err := yaml.Unmarshal(fieldRenameData, &renames); err != nil {
		return errors.Wrap(err, "cannot parse the field renames")
	}
	for _, fr := range renames {
		for name, fields := range fr.Data {
			r, ok := pc.Resources[name]
			if !ok {
				return errors.Errorf("resource %q in the field renames is not configured", name)
			}
			if r.Version != fr.TargetVersion || !slices.Contains(r.PreviousVersions, fr.SourceVersion) {
				return errors.Errorf("resource %q in the field renames has no API versions %s and %s", name, fr.SourceVersion, fr.TargetVersion)
			}
			for source, target := range fields {
				r.Conversions = append(r.Conversions,
					conversion.NewFieldRenameConversion(fr.SourceVersion, source, fr.TargetVersion, target),
					conversion.NewFieldRenameConversion(fr.TargetVersion, target, fr.SourceVersion, source))
				r.AutoConversionRegistrationOptions.AutoRegisterExcludePaths = append(r.AutoConversionRegistrationOptions.AutoRegisterExcludePaths, source, target)
			}
		}
	}
	return nil
}
