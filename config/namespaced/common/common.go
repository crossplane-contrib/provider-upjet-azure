// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package common

import (
	"context"
	"regexp"
	"strings"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/password"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tjconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-azure/apis/namespaced/rconfig"
)

const (
	// ErrGetPasswordSecret is an error string for failing to get password secret
	ErrGetPasswordSecret = "cannot get password secret"
	// errManagedNotNamespaced is an error string for non-namespaced MRs
	errManagedNotNamespaced = "managed resource is not namespaced"
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
	tfName         string
}

// In the regular expressions used when determining the references, we can observe many rules that have common strings.
// Some of these may be more special than others. For example, "ex_subnet$" is a more special case than "subnet$".
// By putting the more specific rules before the general rules in array, processing and capturing the specific
// rules first will be possible. Since there is no fixed index for key-value pairs in maps, it is not possible to place
// rules from specific to general. Therefore, array is used here.
var referenceRules = []refSpecification{
	{wildcardGroup, "resource_group$", "azurerm_resource_group"},
	{wildcardGroup, "subnet$", "azurerm_subnet"},
	{"cosmosdb", "account$", "azurerm_cosmosdb_account"},
	{"dbformariadb", "server$", "azurerm_mariadb_server"},
	{"network", "loadbalancer$", "azurerm_lb"},
}

// AddCommonReferences adds some common reference fields.
// This is a part of resource generation pipeline.
func AddCommonReferences(r *tjconfig.Resource) error {
	delete(r.References, "location")
	return addCommonReferences(r.References, r.TerraformResource, r.ShortGroup, []string{})
}

func addCommonReferences(references tjconfig.References, resource *schema.Resource, shortGroup string, nestedFieldNames []string) error {
	for fieldName, s := range resource.Schema {
		if s.Elem != nil {
			e, ok := s.Elem.(*schema.Resource)
			if ok {
				if err := addCommonReferences(references, e, shortGroup, append(nestedFieldNames, fieldName)); err != nil {
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
		refTFName, err := searchReference(shortGroup, referenceNameWithoutKind)
		if err != nil {
			return err
		}
		if refTFName == "" {
			continue
		}
		if references == nil {
			references = make(map[string]tjconfig.Reference)
		}
		if _, ok := references[referenceName]; !ok {
			addReference(references, referenceKind, referenceName, refTFName)
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
			return rule.tfName, nil
		}
	}
	return "", nil
}

func addReference(references tjconfig.References, referenceKind, referenceName, refTFName string) {
	switch referenceKind {
	case nameReferenceKind:
		references[referenceName] = tjconfig.Reference{
			TerraformName: refTFName,
		}
	case idReferenceKind:
		references[referenceName] = tjconfig.Reference{
			TerraformName: refTFName,
			Extractor:     rconfig.ExtractResourceIDFuncPath,
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
			// this would be a programming/configuration error
			// should be used only for namespaced MRs
			if mg.GetNamespace() == "" {
				return errors.New(errManagedNotNamespaced)
			}
			paved, err := fieldpath.PaveObject(mg)
			if err != nil {
				return errors.Wrap(err, "cannot pave object")
			}
			sel := &v1.LocalSecretKeySelector{}
			if err := paved.GetValueInto(secretRefFieldPath, sel); err != nil {
				return errors.Wrapf(resource.Ignore(fieldpath.IsNotFound, err), "cannot unmarshal %s into a secret key selector", secretRefFieldPath)
			}
			s := &corev1.Secret{}
			if err := client.Get(ctx, types.NamespacedName{Namespace: mg.GetNamespace(), Name: sel.Name}, s); resource.IgnoreNotFound(err) != nil {
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
			s.SetNamespace(mg.GetNamespace())
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
