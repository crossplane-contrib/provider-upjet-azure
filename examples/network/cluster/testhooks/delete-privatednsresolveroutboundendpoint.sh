#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(jonasz-lasut): We are getting (permorming Delete: unexpected status 409...) exception if PrivateDNSResolver for
# the PrivateDNSResolverOutboundEndpoint got deleted before the PrivateDNSResolverOutboundEndpoint resource.
# This is a workaround for this problem to make automated tests to work.
${KUBECTL} delete privatednsresolveroutboundendpoint.network.azure.upbound.io -l testing.upbound.io/example-name=example