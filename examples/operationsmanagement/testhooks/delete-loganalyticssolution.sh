#!/usr/bin/env bash
set -aeuo pipefail

# Delete the LogAnalyticsSolution resource before deleting Workspace
${KUBECTL} delete loganalyticssolution.operationsmanagement.azure.upbound.io --all