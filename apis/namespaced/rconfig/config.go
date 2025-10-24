// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package rconfig

import (
	"fmt"
	"path"

	"github.com/crossplane/upjet/v2/pkg/resource"

	xpref "github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/upbound/provider-azure/apis/namespaced"
	// ExtractResourceIDFuncPath holds the Azure resource ID extractor func name
	ExtractResourceIDFuncPath = APISPackagePath + "/rconfig.ExtractResourceID()"
	// ExtractResourceLocationFuncPath holds the Azure resource location extractor func name
	ExtractResourceLocationFuncPath = APISPackagePath + "/rconfig.ExtractResourceLocation()"
	// ExtractAccountContainerEndpoint holds the Azure Storage Container endpoint extractor func name
	ExtractAccountContainerEndpointFuncPath = APISPackagePath + "/rconfig.ExtractAccountContainerEndpoint()"
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

// ExtractResourceID extracts the value of `spec.atProvider.location`
// from a Terraformed resource. If mr is not a Terraformed
// resource or status.atProvider.location is not yet populated returns an empty string.
func ExtractResourceLocation() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		ob, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		l, ok := ob["location"]
		if !ok {
			return ""
		}
		location, ok := l.(string)
		if !ok {
			return ""
		}
		return location
	}
}

// ExtractAccountContainerEndpoint extracts the value of `spec.atProvider.storageAccountName` and `spec.atProvider.id`
// and creates blob container endpoint based on well-known format - https://{accountName}.blob.core.windows.net/{containerName}
// from a Terraformed resource. If mr is not a Terraformed
// resource or status.atProvider.storageAccountName status.atProvider.id is not yet populated returns an empty string.
func ExtractAccountContainerEndpoint() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		ob, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		accountNameVal, ok := ob["storageAccountName"]
		if !ok {
			return ""
		}
		accountName, ok := accountNameVal.(string)
		if !ok {
			return ""
		}
		containerIdVal, ok := ob["id"]
		if !ok {
			return ""
		}
		containerId, ok := containerIdVal.(string)
		if !ok {
			return ""
		}
		containerName := path.Base(containerId)

		return fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)
	}
}
