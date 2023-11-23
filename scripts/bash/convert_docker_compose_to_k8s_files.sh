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

# ******************************************************************************
# Functions

function convert_docker_compose_to_helm_and_manifest() {
    if [ "$#" != "1" ]; then
        log_error "Usage: ${FUNCNAME[0]} <PROJECT_DIR>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    # local PROJECT_DIR=$(dirname $(dirname $(dirname "${0}")))
    local PROJECT_DIR="${1}"
    local DOCKER_COMPOSE_FILE=docker-compose.yaml.j2 # docker-compose.yml | docker-compose.yaml.j2
    local FILES=$(find "${PROJECT_DIR}" -type f -name "${DOCKER_COMPOSE_FILE}" -not -path '*/dist*' -not -path '*/node_modules*' -not -path '*/.git*')

    for FILE in ${FILES}; do
        # Exclude the current directory
        if [[ "${FILE}" != "${PROJECT_DIR}" ]]; then
            # Get the directory of the docker-compose.yml file
            DIR=$(dirname $(dirname "${FILE}"))
            log_info "Converting ${FILE} to Kubernetes files in ${DIR}"

            # Change to the directory for kubernetes-manifest
            mkdir -p "${DIR}/kubernetes-manifest"
            cd "${DIR}/kubernetes-manifest" || exit
            # Convert docker-compose.yml to kubernetes-manifest
            find "${DIR}/kubernetes-manifest" \
                ! -name "cmd.sh" \
                ! -name "cmd.sh.j2" \
                ! -name "Makefile" \
                ! -name "Makefile.j2" \
                ! -name "README.md" \
                ! -name "README.md.j2" \
                -type f -exec rm -f {} \;
            kompose --file "${FILE}" convert --with-kompose-annotation=false
            # Rename files to .j2 format if not already in that format
            find_and_add_j2_extension_to_files_in_directory "${DIR}/kubernetes-manifest" ".j2"

            # Change to the directory for kubernetes-helm
            mkdir -p "${DIR}/kubernetes-helm"
            cd "${DIR}/kubernetes-helm" || exit
            # Convert docker-compose.yml to kubernetes-helm
            find "${DIR}/kubernetes-helm" \
                ! -name "cmd.sh" \
                ! -name "cmd.sh.j2" \
                ! -name "Makefile" \
                ! -name "Makefile.j2" \
                ! -name "README.md" \
                ! -name "README.md.j2" \
                -type f -exec rm -f {} \;
            kompose --file "${FILE}" convert --chart --out . --with-kompose-annotation=false
            # Rename files to .j2 format if not already in that format
            find_and_add_j2_extension_to_files_in_directory "${DIR}/kubernetes-helm" ".j2"

            # Change back to the current directory
            cd "${PROJECT_DIR}" || exit
        fi
    done

}

# ******************************************************************************
# Main Program

convert_docker_compose_to_helm_and_manifest "${PROJECT_DIR}"

# ******************************************************************************
#set +e # Exit on error Off
#set +x # Trace Off
#echo "End: $(basename "${0}")"
echo -e "\n================================================================================\n"
exit 0
# ******************************************************************************
