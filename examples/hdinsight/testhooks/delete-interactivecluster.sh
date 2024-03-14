#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# the InteractiveQueryCluster resource should be deleted before Account
${KUBECTL} delete interactivequerycluster.hdinsight.azure.upbound.io --all