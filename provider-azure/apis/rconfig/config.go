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
	"github.com/upbound/upjet/pkg/resource"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/upbound/official-providers/provider-azure/apis"
	// ExtractResourceIDFuncPath holds the Azure resource ID extractor func name
	ExtractResourceIDFuncPath = APISPackagePath + "/rconfig.ExtractResourceID()"

	// ResourceGroupPath is used as subpackage path for ResourceGroup
	ResourceGroupPath = "/azure/v1alpha2.ResourceGroup"

	// SubnetPath is used as subpackage path for network.Subnet
	SubnetPath = "/network/v1alpha2.Subnet"

	// ResourceGroupReferencePath is used as import path for ResourceGroup
	ResourceGroupReferencePath = APISPackagePath + ResourceGroupPath

	// SubnetReferencePath is used as import path for network.Subnet
	SubnetReferencePath = APISPackagePath + SubnetPath

	// StorageAccountReferencePath is used as import path for StorageAccount
	StorageAccountReferencePath = APISPackagePath + "/storage/v1alpha2.Account"

	// VaultKeyReferencePath is used as import path for VaultKey
	VaultKeyReferencePath = APISPackagePath + "/keyvault/v1alpha2.Key"
)

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
