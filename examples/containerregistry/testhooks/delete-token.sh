#!/usr/bin/env bash
set -aeuo pipefail

# Delete Tokenpassword resource first before Registry resource
${KUBECTL} delete tokenpassword.containerregistry.azure.upbound.io --all