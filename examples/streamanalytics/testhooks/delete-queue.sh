#!/usr/bin/env bash
set -aeuo pipefail

# Delete the OutputServiceBusQueue resource before deleting secret with password
${KUBECTL} delete outputservicebusqueue.streamanalytics.azure.upbound.io --all