#!/usr/bin/env bash
set -aeuo pipefail

# Delete the Active Deployment resource before deleting the Cloud App istself
${KUBECTL} delete springcloudactivedeployment.appplatform.azure.upbound.io --all