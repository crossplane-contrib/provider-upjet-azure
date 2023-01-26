#!/usr/bin/env bash
set -aeuo pipefail

# Delete the OutputServiceBusTopic resource before deleting secret with password
${KUBECTL} delete outputservicebustopic.streamanalytics.azure.upbound.io --all