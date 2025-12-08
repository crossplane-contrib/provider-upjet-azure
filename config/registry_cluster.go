// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	_ "embed"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/upbound/provider-azure/v2/config/cluster"
	"github.com/upbound/provider-azure/v2/config/cluster/common"
	"github.com/upbound/provider-azure/v2/hack"
)

// GetProvider returns provider configuration
func GetProvider(_ context.Context, sdkProvider *schema.Provider, generationProvider bool) (*ujconfig.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		ujconfig.WithShortName("azure"),
		ujconfig.WithRootGroup("azure.upbound.io"),
		ujconfig.WithIncludeList(CLIReconciledResourceList()),
		ujconfig.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		ujconfig.WithSkipList(skipList),
		ujconfig.WithDefaultResourceOptions(ResourceConfigurator()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	bumpVersionsWithEmbeddedLists(pc)
	// "azure" group contains resources that actually do not have a specific
	// group, e.g. ResourceGroup with APIVersion "azure.upbound.io/v1beta1".
	// We need to include the controllers for this group into the base packages
	// list to get their controllers packaged together with the config package
	// controllers (provider family config package).
	for _, c := range []string{"azure/resourcegroup", "azure/resourceproviderregistration", "azure/subscription"} {
		pc.BasePackages.ControllerMap[c] = "config"
	}

	// API group overrides from Terraform import statements
	for _, r := range pc.Resources {
		groupKindOverride(r)
	}

	// add custom config functions
	for _, configure := range cluster.ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()

	// This function runs after the special configurations were applied. However, if some references were configured in
	// the configuration files, the reference generator does not override them.
	for _, r := range pc.Resources {
		delete(r.References, "resource_group_name")
		if err := common.AddCommonReferences(r); err != nil {
			panic(err)
		}
	}
	return pc, nil
}

func bumpVersionsWithEmbeddedLists(pc *ujconfig.Provider) {
	l := strings.Split(strings.TrimSpace(oldSingletonListAPIs), "\n")
	oldSLAPIs := make(map[string]struct{}, len(l))
	for _, n := range l {
		oldSLAPIs[n] = struct{}{}
	}

	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		if _, ok := oldSLAPIs[name]; ok {
			r.Version = "v1beta2"
			r.PreviousVersions = []string{"v1beta1"}
			// we would like to set the storage version to v1beta1 to facilitate
			// downgrades.
			r.SetCRDStorageVersion("v1beta1")
			// because the controller reconciles on the API version with the singleton list API,
			// no need for a Terraform conversion.
			r.ControllerReconcileVersion = "v1beta1"
			r.Conversions = []conversion.Conversion{
				conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths()...),
				conversion.NewSingletonListConversion("v1beta1", "v1beta2", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToEmbeddedObject),
				conversion.NewSingletonListConversion("v1beta2", "v1beta1", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToSingletonList)}
		} else {
			// the controller will be reconciling on the CRD API version
			// with the converted API (with embedded objects in place of
			// singleton lists), so we need the appropriate Terraform
			// converter in this case.
			r.TerraformConversions = []ujconfig.TerraformConversion{
				ujconfig.NewTFSingletonConversion(),
			}
		}
		pc.Resources[name] = r
	}
}
