#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Remove FlexibleServer with spec.forProvider.createMode==Replica first before removing source FlexibleServer
# If deletion executed in parallel the source FlexibleServer will end with status Failed at provider.
${KUBECTL} delete flexibleservers.dbforpostgresql.azure.upbound.io psqlreplica