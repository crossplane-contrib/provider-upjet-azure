#!/usr/bin/env bash
set -aeuo pipefail

# Delete applicationinsightsanalyticsitem before deleting resource group
${KUBECTL} delete applicationinsightsanalyticsitem.insights.azure.upbound.io --all