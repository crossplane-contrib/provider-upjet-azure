#!/usr/bin/env bash
set -aeuo pipefail

# Delete the OutputMSSQL resource before deleting secret with password
${KUBECTL} delete outputmssql.streamanalytics.azure.upbound.io --all