#!/usr/bin/env bash
set -aeuo pipefail

# the MSSQLServerMicrosoftSupportAuditingPolicy resource should be deleted before Account
${KUBECTL} delete mssqlservermicrosoftsupportauditingpolicy.sql.azure.upbound.io --all