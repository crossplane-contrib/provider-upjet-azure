#!/usr/bin/env bash
set -aeuo pipefail

# the InteractiveQueryCluster resource should be deleted before Account
${KUBECTL} delete interactivequerycluster.hdinsight.azure.upbound.io --all