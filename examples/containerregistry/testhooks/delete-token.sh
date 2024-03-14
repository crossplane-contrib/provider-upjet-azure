#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete Tokenpassword resource first before Registry resource
${KUBECTL} delete tokenpassword.containerregistry.azure.upbound.io --all