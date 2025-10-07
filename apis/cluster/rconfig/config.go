// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package rconfig

import (
	"github.com/crossplane/upjet/v2/pkg/resource"

	xpref "github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/upbound/provider-azure/apis/cluster"
	// ExtractResourceIDFuncPath holds the Azure resource ID extractor func name
	ExtractResourceIDFuncPath = APISPackagePath + "/rconfig.ExtractResourceID()"
	// ExtractResourceLocationFuncPath holds the Azure resource location extractor func name
	ExtractResourceLocationFuncPath = APISPackagePath + "/rconfig.ExtractResourceLocation()"
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
