#!/usr/bin/env bash
set -aeuo pipefail

# Delete the StreamInputEventHub resource before deleting secret with password
${KUBECTL} delete streaminputeventhub.streamanalytics.azure.upbound.io --all