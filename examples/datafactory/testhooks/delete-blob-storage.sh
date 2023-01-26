#!/usr/bin/env bash
set -aeuo pipefail

# Remove LinkedServiceAzureBlobStorage before removing secret
${KUBECTL} delete linkedserviceazureblobstorage.datafactory.azure.upbound.io/lsabsexample --all