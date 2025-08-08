#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete applicationinsightsanalyticsitem before deleting resource group
${KUBECTL} delete applicationinsightsanalyticsitem.insights.azure.upbound.io --all