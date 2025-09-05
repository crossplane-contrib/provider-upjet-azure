#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete the LogAnalyticsSolution resource before deleting Workspace
${KUBECTL} delete loganalyticssolution.operationsmanagement.azure.upbound.io --all