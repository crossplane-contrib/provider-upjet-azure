#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(turkenf): We are getting ("application_insights.0.instrumentation_key\\\" to not be an empty string) exception if logger for
# the applicationinsights got deleted before the logger resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete logger.apimanagement.azure.upbound.io -l testing.upbound.io/example-name=example