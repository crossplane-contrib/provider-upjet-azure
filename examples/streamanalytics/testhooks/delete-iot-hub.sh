#!/usr/bin/env bash
set -aeuo pipefail

# Delete the StreamInputIOTHub resource before deleting secret with password
${KUBECTL} delete streaminputiothub.streamanalytics.azure.upbound.io --all