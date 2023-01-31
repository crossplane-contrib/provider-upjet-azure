#!/usr/bin/env bash
set -aeuo pipefail

# the HBaseCluster resource should be deleted before Account
${KUBECTL} delete hbasecluster.hdinsight.azure.upbound.io --all