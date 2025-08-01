#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Remove Linux Function App first before removing AppPlan
${KUBECTL} delete linuxfunctionapp.web.azure.upbound.io --all