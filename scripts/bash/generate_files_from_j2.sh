#!/usr/bin/env bash
# ******************************************************************************
# Copyright 2020 Clark Hsu
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ******************************************************************************
# How To
# - [Company/Project](<https://{{ GITHUB_PROJECT }}.io/>)
# - [Documentation](<https://{{ GITHUB_PROJECT }}.io/doc>)
# - [Github](<https://github.com/{{ GITHUB_USER }}/{{ GITHUB_PROJECT }}>)
# - [Wikipedia](<https://en.wikipedia.org/wiki/{{ TOPIC }}>)
# ******************************************************************************
# Mark Off this section if use as lib
PROGRAM_NAME=$(basename "${0}")
AUTHOR=clark_hsu
VERSION=0.0.1
# ******************************************************************************
echo -e "\n================================================================================\n"
#echo "Begin: $(basename "${0}")"
#set -e # Exit on error On
#set -x # Trace On
# ******************************************************************************
# Load Configuration

echo -e "\n>>> Load Configuration ...\n"
TOP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${TOP_DIR}/_log.sh"
source "${TOP_DIR}/_common_lib.sh"

PROJECT_DIR=$(dirname $(dirname $(dirname "${0}")))
cd "${PROJECT_DIR}" || exit
PROJECT_DIR="$(pwd)"
log_info "Using current directory: ${PROJECT_DIR}"

# # Load environment variables from _config.sh file
# if [ -f "${TOP_DIR}/_config.sh" ]; then
#     log_info "Loading environment variables from _config.sh file"
#     source "${TOP_DIR}/_config.sh"
# fi

# # Load environment variables from .env file
# if [ -f "${PROJECT_DIR}/.env" ]; then
#     log_info "Loading environment variables from .env file"
#     source "${PROJECT_DIR}/.env"
# fi

# ******************************************************************************
# Parameters

# # Check if the file or directory in the argument exists
# if [ $# -eq 0 ]; then
#     log_error "Usage: ${0} <FILE.ext> or ${0} <DIRECTORY>"
#     exit 1
# fi

TEMPLATE_DIR="${HOME}/Documents/myProject/Library/bash/src/bash_source_lib/j2_templates/common/bash/bash_source_project_app"
NEW_PROJECT_DIR="${HOME}/Documents/myProject/CodingTest/src/typescript/others/test"
GITHUB_PROJECT="programming-language-test-01"

# Define the values to replace placeholders with
# GIT_PROVIDER="YourGitProvider"
# GITHUB_USER="YourProjectUser"
# GITHUB_PROJECT="YourProjectName"
# PROJECT_PORT="YourProjectPort"

# ******************************************************************************
# Functions

# ******************************************************************************
# Main Program

generate_directories_and_files_from_j2_directory "TEMPLATE_DIR:${PROJECT_TEMPLATE_DIR}" "PROJECT_DIR:${NEW_PROJECT_DIR}/${GITHUB_PROJECT}" "J2_ENABLE:true" "GITHUB_PROJECT:${GITHUB_PROJECT}"

# ******************************************************************************
#set +e # Exit on error Off
#set +x # Trace Off
#echo "End: $(basename "${0}")"
echo -e "\n================================================================================\n"
exit 0
# ******************************************************************************
