#!/usr/bin/env bash
set -aeuo pipefail

# Remove Linux Function App first before removing AppPlan
${KUBECTL} delete linuxfunctionapp.web.azure.upbound.io --all