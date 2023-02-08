#!/usr/bin/env bash
set -aeuo pipefail

# the KafkaCluster resource should be deleted before Account
${KUBECTL} delete kafkacluster.hdinsight.azure.upbound.io --all