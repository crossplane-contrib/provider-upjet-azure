// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/apis/configuration/v1alpha1"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ReconciliationPolicy(ctx context.Context, client client.Client, mg xpresource.Managed) (*v1alpha1.ReconciliationPolicy, error) {
	spec, err := resolveProviderConfig(ctx, client, mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve the referenced ProviderConfig")
	}
	return spec.ReconciliationPolicy, nil
}
