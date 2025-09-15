#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# the MSSQLServerMicrosoftSupportAuditingPolicy resource should be deleted before Account
${KUBECTL} delete mssqlservermicrosoftsupportauditingpolicy.sql.azure.upbound.io --all