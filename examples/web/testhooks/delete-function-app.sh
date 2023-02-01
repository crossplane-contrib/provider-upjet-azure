#!/usr/bin/env bash
set -aeuo pipefail

# Remove WebApp first before removing AppPlan
${KUBECTL} delete functionapp.web.azure.upbound.io --all