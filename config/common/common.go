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
	"context"
	"regexp"
	"strings"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/password"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tjconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
	// ErrGetPasswordSecret is an error string for failing to get password secret
	ErrGetPasswordSecret = "cannot get password secret"
	// VersionV1Beta1 is used for resources that meet the v1beta1 criteria
	// here: https://github.com/upbound/arch/pull/33
	VersionV1Beta1 = "v1beta1"

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

// PasswordGenerator returns an InitializerFn that will generate a password
// for a resource if the toggle field is set to true and the secret referenced
// by the secretRefFieldPath is not found or does not have content corresponding
// to the password key.
func PasswordGenerator(secretRefFieldPath, toggleFieldPath string) tjconfig.NewInitializerFn { //nolint:gocyclo
	// NOTE(muvaf): This function is just 1 point over the cyclo limit but there
	// is no easy way to reduce it without making it harder to read.
	return func(client client.Client) managed.Initializer {
		return managed.InitializerFn(func(ctx context.Context, mg resource.Managed) error {
			paved, err := fieldpath.PaveObject(mg)
			if err != nil {
				return errors.Wrap(err, "cannot pave object")
			}
			sel := &v1.SecretKeySelector{}
			if err := paved.GetValueInto(secretRefFieldPath, sel); err != nil {
				return errors.Wrapf(resource.Ignore(fieldpath.IsNotFound, err), "cannot unmarshal %s into a secret key selector", secretRefFieldPath)
			}
			s := &corev1.Secret{}
			if err := client.Get(ctx, types.NamespacedName{Namespace: sel.Namespace, Name: sel.Name}, s); resource.IgnoreNotFound(err) != nil {
				return errors.Wrap(err, ErrGetPasswordSecret)
			}
			if err == nil && len(s.Data[sel.Key]) != 0 {
				// Password is already set.
				return nil
			}
			// At this point, either the secret doesn't exist, or it doesn't
			// have the password filled.
			if gen, err := paved.GetBool(toggleFieldPath); err != nil || !gen {
				// If there is error, then we return that.
				// If the toggle field is not set to true, then we return nil.
				// Because we don't want to generate a password if the user
				// doesn't want to.
				return errors.Wrapf(resource.Ignore(fieldpath.IsNotFound, err), "cannot get the value of %s", toggleFieldPath)
			}
			pw, err := password.Generate()
			if err != nil {
				return errors.Wrap(err, "cannot generate password")
			}
			s.SetName(sel.Name)
			s.SetNamespace(sel.Namespace)
			if !meta.WasCreated(s) {
				// We don't want to own the Secret if it is created by someone
				// else, otherwise the deletion of the managed resource will
				// delete the Secret that we didn't create in the first place.
				meta.AddOwnerReference(s, meta.AsOwner(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind())))
			}
			if s.Data == nil {
				s.Data = make(map[string][]byte, 1)
			}
			s.Data[sel.Key] = []byte(pw)
			return errors.Wrap(resource.NewAPIPatchingApplicator(client).Apply(ctx, s), "cannot apply password secret")
		})
	}
}

func RemoveIndex(s []string, elementToRemove string) []string {
	var result []string
	for _, value := range s {
		if value != elementToRemove {
			result = append(result, value)
		}
	}
	return result
}
