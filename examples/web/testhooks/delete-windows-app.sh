#!/usr/bin/env bash
set -aeuo pipefail

# Remove Windows Function App first before removing AppPlan
${KUBECTL} delete windowswebapp.web.azure.upbound.io --all