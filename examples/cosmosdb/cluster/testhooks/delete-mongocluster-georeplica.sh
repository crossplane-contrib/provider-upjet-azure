#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Remove MongoCluster with spec.forProvider.createMode==GeoReplica first before removing source MongoCluster
# If deletion executed in parallel the source MongoCluster will end with status Failed at provider.
${KUBECTL} delete mongoclusters.cosmosdb.azure.upbound.io examplemcreplica