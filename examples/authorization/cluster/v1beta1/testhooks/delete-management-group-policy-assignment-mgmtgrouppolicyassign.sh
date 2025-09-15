#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Remove ManagementGroupPolicyAssignment first before removing ManagementGroup
# If deletion executed in parallel the ManagementGroupPolicyAssignment
# will stuck with
# ...unexpected status 403 (403 Forbidden) with error: AuthorizationFailed: The client '$uuid' with object id '$uuid' does not have authorization to perform action 'Microsoft.Authorization/policyAssignments/read' over scope '/providers/Microsoft.Management/managementGroups/<name>/providers/Microsoft.Authorization/policyAssignments/<name>' or the scope is invalid. If access was recently granted, please refresh your credentials.
${KUBECTL} delete managementgrouppolicyassignment.authorization.azure.upbound.io mgmtgrouppolicyassign
