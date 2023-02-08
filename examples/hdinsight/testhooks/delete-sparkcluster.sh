#!/usr/bin/env bash
set -aeuo pipefail

# the SparkCluster resource should be deleted before Account
${KUBECTL} delete sparkcluster.hdinsight.azure.upbound.io --all