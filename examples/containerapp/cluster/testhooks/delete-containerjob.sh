#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(jakubramut): We are getting exception - Delete: unexpected status 409 (409 Conflict) with error: 
# ManagedEnvironmentHasContainerAppsJobs: The specified environment example cannot be deleted because it still contains 1 Container Apps Jobs.
# Environment cannot be deleted in expected time - it's hanging but finally deleted.
# This is a workaround for this problem to make autmated tests to work.
${KUBECTL} delete containerjob.containerapp.azure.upbound.io -l testing.upbound.io/example-name=example