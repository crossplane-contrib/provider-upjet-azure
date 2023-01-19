#!/usr/bin/env bash
set -aeuo pipefail

# Delete the StreamInputBlob resource before deleting secret with password
${KUBECTL} delete streaminputblob.streamanalytics.azure.upbound.io --all