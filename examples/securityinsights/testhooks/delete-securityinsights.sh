#!/usr/bin/env bash
set -aeuo pipefail

# the SentinelLogAnalyticsWorkspaceOnboarding should be deleted before the Workspace resource.
${KUBECTL} delete sentinelloganalyticsworkspaceonboarding.securityinsights.azure.upbound.io --all