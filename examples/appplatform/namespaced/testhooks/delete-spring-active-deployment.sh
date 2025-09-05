#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete the Active Deployment resource before deleting the Cloud App istself
${KUBECTL} delete springcloudactivedeployment.appplatform.azure.upbound.io --all