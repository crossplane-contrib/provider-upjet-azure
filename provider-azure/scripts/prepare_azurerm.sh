#!/usr/bin/env sh

# Copyright 2021 The Crossplane Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eu

REPO_DIR="${WORK_DIR}/.azurerm"
mkdir -p "${REPO_DIR}"

if [ ! -d "${REPO_DIR}/.git" ]; then
  git -C "${REPO_DIR}" init --initial-branch=trunk
  git -C "${REPO_DIR}" remote add origin https://github.com/hashicorp/terraform-provider-azurerm
fi

git -C "${REPO_DIR}" fetch origin v"${TERRAFORM_PROVIDER_VERSION}"
git -C "${REPO_DIR}" reset --hard FETCH_HEAD

mkdir -p "${REPO_DIR}/xpprovider"
cp hack/provider.go.txt "${REPO_DIR}/xpprovider/provider.go"
