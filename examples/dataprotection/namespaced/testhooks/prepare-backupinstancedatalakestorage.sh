#!/usr/bin/env bash
set -euo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

kubectl="${KUBECTL:-kubectl}"
account=accounts.storage.azure.m.upbound.io/example
instance=backupinstancedatalakestorages.dataprotection.azure.m.upbound.io/example
vault=backupvaults.dataprotection.azure.m.upbound.io/example
assignment=roleassignments.authorization.azure.m.upbound.io/example

# Uptest's random subdomain includes a hyphen, which storage account names forbid.
uid=$("$kubectl" -n upbound-system get "$account" -o jsonpath='{.metadata.uid}')
suffix="${uid//-/}"
if [[ ! "$suffix" =~ ^[a-f0-9]{32}$ ]]; then
    echo "Expected a UUID for the storage account metadata.uid, got: $uid" >&2
    exit 1
fi
"$kubectl" -n upbound-system patch "$account" --type=merge \
    -p "{\"metadata\":{\"annotations\":{\"crossplane.io/external-name\":\"adls${suffix:0:20}\",\"crossplane.io/paused\":null}}}"
"$kubectl" -n upbound-system wait "$vault" "$account" --for=condition=Ready --timeout=15m
principal=$("$kubectl" -n upbound-system get "$vault" -o jsonpath='{.status.atProvider.identity.principalId}')
scope=$("$kubectl" -n upbound-system get "$account" -o jsonpath='{.status.atProvider.id}')
if [[ -z "$principal" || -z "$scope" ]]; then
    echo "BackupVault principal ID or storage account ID is missing" >&2
    exit 1
fi
patch=$(jq -n --arg principal "$principal" --arg scope "$scope" \
    '{metadata: {annotations: {"crossplane.io/paused": null}}, spec: {forProvider: {principalId: $principal, scope: $scope}}}')
"$kubectl" -n upbound-system patch "$assignment" --type=merge -p "$patch"
"$kubectl" -n upbound-system wait "$assignment" --for=condition=Ready --timeout=15m
"$kubectl" -n upbound-system wait containers.storage.azure.m.upbound.io/example \
    containers.storage.azure.m.upbound.io/example2 --for=condition=Ready --timeout=15m
"$kubectl" -n upbound-system annotate "$instance" crossplane.io/paused-
