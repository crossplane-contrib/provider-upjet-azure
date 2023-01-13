#!/usr/bin/env bash
set -aeuo pipefail

# Delete the SpringCloudAppMySQLAssociation resource before deleting secret with password
${KUBECTL} delete springcloudappmysqlassociation.appplatform.azure.upbound.io --all