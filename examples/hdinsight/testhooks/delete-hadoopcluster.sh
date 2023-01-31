#!/usr/bin/env bash
set -aeuo pipefail

# the HadoopCluster resource should be deleted before Account
${KUBECTL} delete hadoopcluster.hdinsight.azure.upbound.io --all