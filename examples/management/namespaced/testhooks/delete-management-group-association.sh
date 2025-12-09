#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Remove ManagementGroupSubscriptionAssociation first before removing ManagementGroup
# If deletion executed in parallel the ManagementGroupSubscriptionAssociation
# will stuck with
#  Original Error: autorest/azure: Service returned an error. Status=403 Code="AuthorizationFailed" Message="The client '$uuid' with object id '$uuid' does not have authorization to perform action 'Microsoft.Management/managementGroups/read' over scope '/providers/Microsoft.Management/managementGroups/example-sub' or the scope is invalid. If access was recently granted, please refresh your credentials.":
${KUBECTL} delete managementgroupsubscriptionassociations.management.azure.m.upbound.io -n upbound-system --all
