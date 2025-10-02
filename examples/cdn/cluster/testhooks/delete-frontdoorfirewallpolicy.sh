#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(jonasz-lasut): We are getting observe failed: failed to observe the resource: [{0 retrieving Front Door Web Application Firewall Policy (Subscription: "0a7d218a-1c62-4113-9660-9fbf629fe279"
# Resource Group Name: "frontdoorfirewallpolicy" Front Door Web Application Firewall Policy Name: "frontdoorfirewallpolicy"):
# webapplicationfirewallpolicies.WebApplicationFirewallPoliciesClient#PoliciesGet: Failure responding to request: StatusCode=404 -- Original Error: autorest/azure: Service returned an error. Status=404 Code="ResourceGroupNotFound" Message="Resource group 'frontdoorfirewallpolicy' could not be found."  []}].
# This is a workaround for this problem to make automated tests to work.
${KUBECTL} delete frontdoorfirewallpolicy.cdn.azure.upbound.io -l testing.upbound.io/example-name=frontdoorfirewallpolicy