// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Package provider contains the initialization tasks shared by the generated
// provider binaries.
package provider

import (
	"context"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	authv1 "k8s.io/api/authorization/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const rootGroup = "azure.upbound.io"

// RunStorageVersionMigration re-stores the objects of the CRDs in
// "<shortGroup>.azure.upbound.io" at their current storage version and removes
// the other versions from the CRDs' status.storedVersions. An empty shortGroup
// covers every CRD under the root group, which is what the monolith binary
// needs. The family package, built from the config group, has no init command
// and never calls it: its CRDs are in the root group itself, which a short
// group cannot select, and none of them has a version to migrate. The same
// goes for the azure group, which is generated but not built.
func RunStorageVersionMigration(ctx context.Context, logr logging.Logger, cfg *rest.Config, shortGroup string) error {
	// The migrator lists CRDs, checks its own permissions with a
	// SelfSubjectAccessReview and handles the objects as unstructured, so
	// it needs none of the provider's own types.
	scheme := runtime.NewScheme()
	if err := extv1.AddToScheme(scheme); err != nil {
		return errors.Wrap(err, "cannot add apiextensions/v1 to the scheme")
	}
	if err := authv1.AddToScheme(scheme); err != nil {
		return errors.Wrap(err, "cannot add authorization/v1 to the scheme")
	}

	kube, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		return errors.Wrap(err, "cannot create a Kubernetes client for the storage version migration")
	}

	var opts []ujconfig.CRDMigratorOption
	if shortGroup != "" {
		opts = append(opts, ujconfig.WithShortGroup(shortGroup))
	}
	return errors.Wrap(ujconfig.NewCRDMigrator(rootGroup, opts...).Run(ctx, logr, kube), "cannot migrate the CRD storage versions")
}
