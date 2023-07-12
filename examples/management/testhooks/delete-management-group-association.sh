#!/usr/bin/env bash
set -aeuo pipefail

# Remove ManagementGroupSubscriptionAssociation first before removing ManagementGroup
# If deletion executed in parallel the ManagementGroupSubscriptionAssociation
# will stuck with
#  Original Error: autorest/azure: Service returned an error. Status=403 Code="AuthorizationFailed" Message="The client '$uuid' with object id '$uuid' does not have authorization to perform action 'Microsoft.Management/managementGroups/read' over scope '/providers/Microsoft.Management/managementGroups/example-sub' or the scope is invalid. If access was recently granted, please refresh your credentials.":
${KUBECTL} delete managementgroupsubscriptionassociations.management.azure.upbound.io --all
