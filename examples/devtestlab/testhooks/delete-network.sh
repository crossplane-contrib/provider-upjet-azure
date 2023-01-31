#!/usr/bin/env bash
set -aeuo pipefail

# Delete VirtualNetwork before removing ResourceGroup
${KUBECTL} delete virtualnetwork.devtestlab.azure.upbound.io --all