#!/usr/bin/env bash
set -aeuo pipefail

# Remove LinuxApp first before removing AppPlan
${KUBECTL} delete linuxwebapp.web.azure.upbound.io --all