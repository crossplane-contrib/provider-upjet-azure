/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tjconfig "github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`

	// nameReferenceKind represents name reference kind
	nameReferenceKind = "name"
	// idReferenceKind represents ID reference kind
	idReferenceKind = "id"
	// wildcard group name
	wildcardGroup = "*"
)

type refSpecification struct {
	group          string
	attributeRegex string
	packagePath    string
}

// In the regular expressions used when determining the references, we can observe many rules that have common strings.
// Some of these may be more special than others. For example, "ex_subnet$" is a more special case than "subnet$".
// By putting the more specific rules before the general rules in array, processing and capturing the specific
// rules first will be possible. Since there is no fixed index for key-value pairs in maps, it is not possible to place
// rules from specific to general. Therefore, array is used here.
var referenceRules = []refSpecification{
	{wildcardGroup, "resource_group$", rconfig.GetDefaultVersionedPath("azure", "ResourceGroup")},
	{wildcardGroup, "subnet$", rconfig.GetDefaultVersionedPath("network", "Subnet")},
	{"cosmosdb", "account$", rconfig.GetDefaultVersionedPath("cosmosdb", "Account")},
	{"dbformariadb", "server$", rconfig.GetDefaultVersionedPath("dbformariadb", "Server")},
	{"network", "loadbalancer$", rconfig.GetDefaultVersionedPath("network", "LoadBalancer")},
}

// AddCommonReferences adds some common reference fields.
// This is a part of resource generation pipeline.
func AddCommonReferences(r *tjconfig.Resource) error {
	delete(r.References, "location")
	return addCommonReferences(r.References, r.TerraformResource, r.ShortGroup, r.Version, []string{})
}

func addCommonReferences(references tjconfig.References, resource *schema.Resource, shortGroup, version string, nestedFieldNames []string) error {
	for fieldName, s := range resource.Schema {
		if s.Elem != nil {
			e, ok := s.Elem.(*schema.Resource)
			if ok {
				if err := addCommonReferences(references, e, shortGroup, version, append(nestedFieldNames, fieldName)); err != nil {
					return err
				}
				continue
			}
		}

		referenceName := strings.Join(append(nestedFieldNames, fieldName), ".")
		referenceNameWithoutKind, referenceKind := splitReferenceKindFromReferenceName(referenceName)
		if referenceKind == "" {
			continue
		}
		referenceType, err := searchReference(shortGroup, referenceNameWithoutKind)
		if err != nil {
			return err
		}
		if referenceType == "" {
			continue
		}
		if references == nil {
			references = make(map[string]tjconfig.Reference)
		}
		if _, ok := references[referenceName]; !ok {
			referenceType = prepareReferenceType(shortGroup, version, referenceType)
			addReference(references, referenceKind, referenceName, referenceType)
		}
	}
	return nil
}

func splitReferenceKindFromReferenceName(fieldName string) (string, string) {
	p := strings.Split(fieldName, "_")
	if len(p) < 2 {
		return "", ""
	}
	return strings.Join(p[:len(p)-1], "_"), p[len(p)-1]
}

func searchReference(shortGroup, fieldName string) (string, error) {
	for _, rule := range referenceRules {
		if rule.group != wildcardGroup && rule.group != shortGroup {
			continue
		}
		r, err := regexp.Compile(rule.attributeRegex)
		if err != nil {
			return "", err
		}
		if r.MatchString(fieldName) {
			return rule.packagePath, nil
		}
	}
	return "", nil
}

func prepareReferenceType(shortGroup, version, referenceType string) string {
	p := strings.Split(referenceType, "/")
	if shortGroup == p[1] && strings.Split(p[2], ".")[0] == version {
		referenceType = strings.Split(p[2], ".")[1]
	} else {
		referenceType = rconfig.APISPackagePath + referenceType
	}
	return referenceType
}

func addReference(references tjconfig.References, referenceKind, referenceName, referenceType string) {
	switch referenceKind {
	case nameReferenceKind:
		references[referenceName] = tjconfig.Reference{
			Type: referenceType,
		}
	case idReferenceKind:
		references[referenceName] = tjconfig.Reference{
			Type:      referenceType,
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	}
}

// GetField returns the value of field as a string in a map[string]interface{},
//
//	fails properly otherwise.
func GetField(from map[string]interface{}, path string) (string, error) {
	return fieldpath.Pave(from).GetString(path)
}
