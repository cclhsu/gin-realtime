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

# Replace 'PATH_TO_DIRECTORY' with the actual path to your directory
PATH_TO_DIRECTORY="/Users/clark.hsu/src/myProject/project-suite/project-suite-cli/src/template"

# ******************************************************************************
# Functions

# ******************************************************************************
# Main Program

# Call the reformat_files_in_dir_with_file_pattern function with the directory path
# reformat_files_in_dir_with_file_pattern "/path/to/directory" "*.sh"
# reformat_files_in_dir_with_file_pattern "/path/to/directory" "*.md"
reformat_files_in_dir_with_file_pattern "${PATH_TO_DIRECTORY}" "*.j2"
reformat_files_in_dir_with_file_pattern "${PATH_TO_DIRECTORY}" "*Makefile"
reformat_files_in_dir_with_file_pattern "${PATH_TO_DIRECTORY}" "*Makefile.j2"

# ******************************************************************************
#set +e # Exit on error Off
#set +x # Trace Off
#echo "End: $(basename "${0}")"
echo -e "\n================================================================================\n"
exit 0
# ******************************************************************************
