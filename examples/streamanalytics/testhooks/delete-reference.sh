#!/usr/bin/env bash
set -aeuo pipefail

# Delete the ReferenceInputBlob resource before deleting secret with password
${KUBECTL} delete referenceinputblob.streamanalytics.azure.upbound.io --all