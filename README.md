<!--
SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>

SPDX-License-Identifier: CC-BY-4.0
-->

# Upjet-based Crossplane provider for Azure

<div style="text-align: center;">

![CI](https://github.com/crossplane-contrib/provider-upjet-azure/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/crossplane-contrib/provider-upjet-azure/all.svg)](https://github.com/crossplane-contrib/provider-upjet-azure/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/crossplane-contrib/provider-upjet-azure)](https://goreportcard.com/report/github.com/crossplane-contrib/provider-upjet-azure)
[![Contributors](https://img.shields.io/github/contributors/crossplane-contrib/provider-upjet-azure)](https://github.com/crossplane-contrib/provider-upjet-azure/graphs/contributors)
[![Slack](https://img.shields.io/badge/Slack-4A154B?logo=slack)](https://crossplane.slack.com/archives/C05E4LDNNG5)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/crossplane_io)](https://twitter.com/crossplane_io)

</div>

Provider Upjet-Azure is a [Crossplane](https://crossplane.io/) provider that is
built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for
[Microsoft Azure](https://azure.microsoft.com/).

## Getting Started

Follow the quick start
guide [here](https://github.com/crossplane-contrib/provider-upjet-azure/blob/main/docs/family/Quickstart.md).

You can find a detailed API reference for all the managed resources with examples in
the [Upbound Marketplace](https://marketplace.upbound.io/providers/upbound/provider-family-azure#managedResources).

For more information about monitoring the Upjet runtime, please
see [Monitoring Guide](https://github.com/crossplane/upjet/blob/main/docs/monitoring.md)

## Contributing

For the general contribution guide,
see [Upjet Contribution Guide](https://github.com/crossplane/upjet/blob/main/CONTRIBUTING.md)

If you'd like to learn how to use Upjet, see [Usage Guide](https://github.com/crossplane/upjet/tree/main/docs).

### Add a New Resource

Follow the Upjet guide
for [adding new resources](https://github.com/crossplane/upjet/blob/main/docs/adding-new-resource.md).

### Build instructions

This repository contains the whole provider-family-azure. It can be built monolithic or seperated providers for each subpackage (e.g. provider-azure-containerservice) can be created:
```
export SUBPACKAGES=containerservice
```

Please note: Binaries built can't be installed directly. They must be packaged according to crossplane's requirements.
Cross-compilation (e.g. build on Apple Silicon for Linux) is possible through setting the BUILD_PLATFORMS accordingly:

```
export BUILD_PLATFORMS=linux_amd64
```

Binaries are created via the build target in Makefile:

```
SUBPACKAGES=containerservice PLATFORMS=linux_amd64 make build
```

### Packaging

Once you've created the binaries you can package them in crossplane's xpkg format, upload them to a container registry and update the provider reference in your crossplane configuration to use the updated version.

The xpkg is created and uploaded through the crossplane-cli.

As first step we create a crossplane.yaml file with the package definition and copy the crds in scope for the subpackage over.

```
STAGING=$(mktemp -d)
mkdir -p $STAGING/package

cat > $STAGING/package/crossplane.yaml << 'EOF'
  apiVersion: meta.pkg.crossplane.io/v1
  kind: Provider
  metadata:
    name: provider-azure-containerservice
    labels:
      pkg.crossplane.io/provider-family: provider-family-azure
  spec:
    capabilities:
    - SafeStart
    crossplane:
      version: ">=v1.12.1-0"
    dependsOn:
      - provider: xpkg.upbound.io/upbound/provider-family-azure
        version: ">=v2.0.0"
EOF

cp _output/package/crds/containerservice.azure.* $STAGING/package/
```

As next step we build a docker runtime image with the provider binary from build step embedded (Note: we use next as version tag):

```
docker build --platform linux/amd64 \
    -t provider-azure-containerservice:next -f - . << 'DOCKERFILE'
  FROM build-30968d6a/provider-azure-amd64:latest
  COPY _output/bin/linux_amd64/containerservice /usr/local/bin/provider
DOCKERFILE
```

Afterwards we create the xpkg that combines the runtime image with the package definition (Note: we use next as version tag):

```
crossplane xpkg build \
    --package-root=$STAGING/package \
    --embed-runtime-image=provider-azure-containerservice:next \
    -o provider-azure-containerservice-next.xpkg
```

Finally we can push the xpkg to a registry:

```
crossplane xpkg push \
<MYCR>/provider-azure-containerservice:next \
-f provider-azure-containerservice-next.xpkg
```

The new version with tag next can then be referenced from the provider definition in crossplane (Note: if you're re-using a tag, you should set packagePullPolicy: Always, so always a fresh version is pulled from the upstream registry):

```
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: upbound-provider-azure-containerservice
spec:
  package: <MYCR>/provider-azure-containerservice:next
  packagePullPolicy: Always
```

## Getting help

For filing bugs, suggesting improvements, or requesting new resources or features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-azure/issues/new/choose).

For general help on using the provider consider asking the Crossplane community in the
[#upjet-provider-azure](https://crossplane.slack.com/archives/C05E4LDNNG5) channel in
[Crossplane Slack](https://slack.crossplane.io)

## License

The provider is released under the [the Apache 2.0 license](LICENSE) with [notice](NOTICE).
