#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# the SentinelLogAnalyticsWorkspaceOnboarding should be deleted before the Workspace resource.
${KUBECTL} delete sentinelloganalyticsworkspaceonboarding.securityinsights.azure.upbound.io --all