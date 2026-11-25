#!/usr/bin/env bash
set -euo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

kubectl="${KUBECTL:-kubectl}"

# Azure deletes a Recovery Services vault even when it still holds site recovery
# items, and refuses to remove a fabric while its containers are mapped, so the
# chain has to go from the mapping down to the fabrics before the vault.
"$kubectl" -n upbound-system delete siterecoveryprotectioncontainermappings.recoveryservices.azure.m.upbound.io/container-mapping \
    --ignore-not-found --wait=true --timeout=30m
"$kubectl" -n upbound-system delete siterecoveryprotectioncontainers.recoveryservices.azure.m.upbound.io/primary-protection-container \
    siterecoveryprotectioncontainers.recoveryservices.azure.m.upbound.io/secondary \
    --ignore-not-found --wait=true --timeout=30m
"$kubectl" -n upbound-system delete siterecoveryreplicationpolicies.recoveryservices.azure.m.upbound.io/policy \
    --ignore-not-found --wait=true --timeout=30m
"$kubectl" -n upbound-system delete siterecoveryfabrics.recoveryservices.azure.m.upbound.io/primary \
    siterecoveryfabrics.recoveryservices.azure.m.upbound.io/secondary \
    --ignore-not-found --wait=true --timeout=30m
