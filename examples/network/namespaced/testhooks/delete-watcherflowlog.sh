#!/usr/bin/env bash
set -euo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

kubectl="${KUBECTL:-kubectl}"

# Azure cannot delete a flow log once its storage account is gone, and the
# network watcher cannot be deleted while it still holds the flow log.
"$kubectl" -n upbound-system delete watcherflowlogs.network.azure.m.upbound.io/example \
    --ignore-not-found --wait=true --timeout=15m
