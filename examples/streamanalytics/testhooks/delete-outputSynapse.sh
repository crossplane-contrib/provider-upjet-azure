#!/usr/bin/env bash
set -aeuo pipefail

# Delete the outputSynapse resource before deleting secret with password
${KUBECTL} delete outputsynapse.streamanalytics.azure.upbound.io --all