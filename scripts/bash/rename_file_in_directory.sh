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
PATH_TO_DIRECTORY="/Users/clark.hsu/Documents/myProject/CodingTest/src/golang/takehome/solution/gin-realtime"
OLD_PATTERN=""
NEW_PATTERN=""

# ******************************************************************************
# Functions

# Function to rename files
function rename_files() {
    if [ "$#" != "3" ]; then
        log_error "Usage: ${FUNCNAME[0]} <PATH_TO_DIRECTORY> <OLD_PATTERN> <NEW_PATTERN>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local PATH_TO_DIRECTORY="${1}"
    local OLD_PATTERN="${2}"
    local NEW_PATTERN="${3}"

    if [ -z "${PATH_TO_DIRECTORY}" ]; then
        log_error "Path to directory must be set."
        return
    fi

    # Ensure that both old and new strings are set
    if [ -z "${OLD_PATTERN}" ] || [ -z "${NEW_PATTERN}" ]; then
        log_error "Both OLD_PATTERN and NEW_PATTERN must be set."
        return
    fi

    # find directory with old string in name and replace with new string
    local DIRS=$(find ${PATH_TO_DIRECTORY} -type d -name "*${OLD_PATTERN}*" | sort)
    for DIR in ${DIRS[*]}; do
        log_info "Processing ${DIR} in ${PATH_TO_DIRECTORY}"
        if [ -d "${DIR}" ]; then
            # Create the new name by replacing OLD_PATTERN with NEW_PATTERN
            NEW_NAME="${DIR//${OLD_PATTERN}/${NEW_PATTERN}}"
            mv "${DIR}" "${NEW_NAME}"
            log_info "Renamed: ${DIR} -> ${NEW_NAME}"
        fi
    done

    # Loop through all files and directories in the current directory recursively and rename files
    for ENTRY in "${PATH_TO_DIRECTORY}"/*; do
        log_info "Processing ${ENTRY} in ${PATH_TO_DIRECTORY}"
        if [ -d "${ENTRY}" ]; then
            # Recursive call for subdirectories
            rename_files "${ENTRY}"
        else
            # Check if the file name contains OLD_PATTERN
            if [[ "${ENTRY}" == *"${OLD_PATTERN}"* ]]; then
                # Create the new name by replacing OLD_PATTERN with NEW_PATTERN
                NEW_NAME="${ENTRY//${OLD_PATTERN}/${NEW_PATTERN}}"
                mv "${ENTRY}" "${NEW_NAME}"
                log_info "Renamed: ${ENTRY} -> ${NEW_NAME}"
            fi
        fi
    done

    log_info "Done for ${PATH_TO_DIRECTORY}."
}

# ******************************************************************************
# Main Program

# Call the rename_files function with the directory path
rename_files "${PATH_TO_DIRECTORY}" "${OLD_PATTERN}" "${NEW_PATTERN}"

# ******************************************************************************
#set +e # Exit on error Off
#set +x # Trace Off
#echo "End: $(basename "${0}")"
echo -e "\n================================================================================\n"
exit 0
# ******************************************************************************
