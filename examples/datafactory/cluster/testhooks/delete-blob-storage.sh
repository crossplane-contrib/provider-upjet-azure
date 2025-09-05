#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Remove LinkedServiceAzureBlobStorage before removing secret
${KUBECTL} delete linkedserviceazureblobstorage.datafactory.azure.upbound.io/lsabsexample --all