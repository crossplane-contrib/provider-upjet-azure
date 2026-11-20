// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package roundtrip

import (
	"maps"
	"reflect"

	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// equateFieldConversionAnnotation ignores the upjet field-conversions
// annotation when comparing object annotations. upjet's default option only
// drops the map entry, so an object that started with nil annotations would
// still differ from its round-tripped copy, which has an annotations map
// with just that key. A field that only one API version of a resource has
// is kept in this annotation while the object is converted to the other
// version.
func equateFieldConversionAnnotation() cmp.Option {
	objectMetaType := reflect.TypeFor[metav1.ObjectMeta]()
	return cmp.FilterPath(func(p cmp.Path) bool {
		sf, ok := p.Last().(cmp.StructField)
		return ok && sf.Name() == "Annotations" && p.Index(-2).Type() == objectMetaType
	}, cmp.FilterValues(func(x, y map[string]string) bool {
		// both empty is already handled by cmpopts.EquateEmpty, and
		// applying here too would make the options ambiguous.
		return len(x) != 0 || len(y) != 0
	}, cmp.Comparer(func(x, y map[string]string) bool {
		x, y = maps.Clone(x), maps.Clone(y)
		delete(x, conversion.AnnotationKey)
		delete(y, conversion.AnnotationKey)
		return maps.Equal(x, y)
	})))
}
