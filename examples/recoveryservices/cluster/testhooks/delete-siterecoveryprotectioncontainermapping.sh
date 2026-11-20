#!/usr/bin/env bash
set -euo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

kubectl="${KUBECTL:-kubectl}"

# Azure deletes a Recovery Services vault even when it still holds site recovery
# items, and refuses to remove a fabric while its containers are mapped, so the
# chain has to go from the mapping down to the fabrics before the vault.
"$kubectl" delete siterecoveryprotectioncontainermappings.recoveryservices.azure.upbound.io/container-mapping \
    --ignore-not-found --wait=true --timeout=30m
"$kubectl" delete siterecoveryprotectioncontainers.recoveryservices.azure.upbound.io/primary-protection-container \
    siterecoveryprotectioncontainers.recoveryservices.azure.upbound.io/secondary \
    --ignore-not-found --wait=true --timeout=30m
"$kubectl" delete siterecoveryreplicationpolicies.recoveryservices.azure.upbound.io/policy \
    --ignore-not-found --wait=true --timeout=30m
"$kubectl" delete siterecoveryfabrics.recoveryservices.azure.upbound.io/primary \
    siterecoveryfabrics.recoveryservices.azure.upbound.io/secondary \
    --ignore-not-found --wait=true --timeout=30m
