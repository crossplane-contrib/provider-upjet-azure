#!/usr/bin/env bash
set -euo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

kubectl="${KUBECTL:-kubectl}"

# Keep the vault's storage permissions until the backup instance is deleted.
"$kubectl" delete backupinstancedatalakestorages.dataprotection.azure.upbound.io/example \
    --ignore-not-found --wait=true --timeout=15m
"$kubectl" delete roleassignments.authorization.azure.upbound.io/example \
    --ignore-not-found --wait=true --timeout=15m
