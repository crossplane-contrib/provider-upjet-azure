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

package rconfig

import (
	"fmt"

	"github.com/upbound/upjet/pkg/resource"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/upbound/provider-azure/apis"
	// ExtractResourceIDFuncPath holds the Azure resource ID extractor func name
	ExtractResourceIDFuncPath = APISPackagePath + "/rconfig.ExtractResourceID()"

	// VersionV1Alpha2 is used as minimum version for all manually configured resources.
	// Deprecated: Please use VersionV1Beta1 as minimum.
	VersionV1Alpha2 = "v1alpha2"
	// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
	VersionV1Beta1 = "v1beta1"

	// StorageAccountReferencePath is used as import path for StorageAccount
	StorageAccountReferencePath = APISPackagePath + "/storage/" + VersionV1Beta1 + ".Account"

	// VaultKeyReferencePath is used as import path for VaultKey
	VaultKeyReferencePath = APISPackagePath + "/keyvault/" + VersionV1Beta1 + ".Key"

	// ContainerReferencePath is used as import path for Container
	ContainerReferencePath = APISPackagePath + "/storage/" + VersionV1Beta1 + ".Container"
)

// GetDefaultVersionedPath gets the package path from repo root
// for the specified group and kind name.
func GetDefaultVersionedPath(group, kindName string) string {
	return fmt.Sprintf("/%s/%s.%s", group, VersionV1Beta1, kindName)
}

// ExtractResourceID extracts the value of `spec.atProvider.id`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractResourceID() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}
