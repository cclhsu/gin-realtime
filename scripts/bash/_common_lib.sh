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
# PROGRAM_NAME=$(basename "${0}")
# AUTHOR=clark_hsu
# VERSION=0.0.1
# ******************************************************************************
# echo -e "\n================================================================================\n"
#echo "Begin: $(basename "${0}")"
#set -e # Exit on error On
#set -x # Trace On
# ******************************************************************************
# Load Configuration

TOP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${TOP_DIR}/_log.sh"
# source "${TOP_DIR}/_common_lib.sh"

# ******************************************************************************
# Functions

function check_command_exist() {
    # which "${1}" 2>/dev/null
    type "${1}" 2>/dev/null
    # command -v "${1}" >/dev/null 2>&1
    return $?
}

function check_file_exist() {
    test -f "${1}" 2>/dev/null
    return $?
}

function check_directory_exist() {
    test -d "${1}" 2>/dev/null
    return $?
}

function check_run_state() {
    if [ "${1}" != "0" ]; then
        exit 1
    fi
}

function find_program() {
    if [ "$#" != "2" ]; then
        log_error "Usage: ${FUNCNAME[0]} <PROGRAM_VARIABLE> <PROGRAM_NAME>"
        exit 1
    fi

    local PROGRAM RETURN
    PROGRAM=$(which ${2} 2>/dev/null || :)
    RETURN=$?
    if [ ${RETURN} != 0 ]; then
        PROGRAM=
    fi
    eval ${1}=\"${PROGRAM}\"
    return ${RETURN}
}

function check_commands_exist() {
    if [ "$#" == "0" ]; then
        log_error "Usage: ${FUNCNAME[0]} <BINARIES>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local BINARIES=${*}

    for BIN in ${BINARIES[*]}; do
        # if [ -x "$(command -v ${BIN})" ]; then
        if [ -x "$(type ${BIN})" ]; then
            echo "${BIN} is not installed"
        else
            echo "${BIN} is installed"
        fi
    done
}
# check_commands_exist curl grep awk sed jq blkid findmnt lsblk

function check_services_started() {
    if [ "$#" == "0" ]; then
        log_error "Usage: ${FUNCNAME[0]} <SERVICES>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local SERVICES=${*}

    for SERVICE in ${SERVICES[*]}; do
        if [ "$(
            sudo systemctl status ${SERVICE}.service
            echo $?
        )" != 0 ]; then
            echo "${SERVICE} is not functioning properly"
        else
            echo "${SERVICE} is functioning properly"
        fi
    done
}
# check_services_started iscsid

# ******************************************************************************
# Semantic

function semantic_versioning_parse() {
    if [ "$#" != "1" ]; then
        log_error "Usage: ${FUNCNAME[0]} <ARGS>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    # log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    MAJOR="${1%%.*}"
    MINOR="${1#${MAJOR}.}"
    MINOR="${MINOR%%.*}"
    PATCH="${1#${MAJOR}.${MINOR}.}"
    PATCH="${PATCH%%[-.]*}"
    # echo "${MAJOR} ${MINOR} ${PATCH}"
}
# semantic_versioning_parse 1.2.3

# ******************************************************************************

function update_datetime() {
    if [ "$#" != "0" ] && [ "$#" != "1" ]; then
        log_error "Usage: ${FUNCNAME[0]} <ARGS>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    # log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    echo -e "\n>>> Time Sync...\n"
    sudo date -s "$(curl -s --head http://google.com | grep ^Date: | sed 's/Date: //g') ${1}"
}

# ******************************************************************************
# Convert J2 files to files without J2 extension in directory and subdirectories. (Jinja2)

function find_and_remove_j2_extension_from_files_in_directory() {
    if [ "$#" != "1" ]; then
        log_error "Usage: ${FUNCNAME[0]} <BASE_DIR>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_info "[${FUNCNAME[0]}] Remove extension .j2 from files in directory ${1}"

    if [ "$#" == "1" ]; then
        local BASE_DIR=${1}
    fi

    cd "${BASE_DIR:?}" || exit 1

    # find "${BASE_DIR}" -type f -name "*.j2" -exec bash -c 'echo {} $(dirname {})/$(basename {} .j2)' "{}" \;
    find "${BASE_DIR}" -type f -name "*.j2" -exec bash -c 'mv {} $(dirname {})/$(basename {} .j2)' "{}" \;
    # find "${BASE_DIR}" -type f ! -name "Makefile*" ! -name "cmd.sh*" ! -name "*.j2" -exec bash -c 'mv {} $(dirname {})/$(basename {} .j2)' "{}" \;
}

function find_and_add_j2_extension_to_files_in_directory() {
    if [ "$#" != "1" ]; then
        log_error "Usage: ${FUNCNAME[0]} <BASE_DIR>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_info "[${FUNCNAME[0]}] Add extsion j2 to files in directory ${1}"

    if [ "$#" == "1" ]; then
        local BASE_DIR=${1}
    fi

    cd "${BASE_DIR:?}" || exit 1

    # find "${BASE_DIR}" -type f -not -name "*.j2" -exec bash -c 'echo {} {}.j2' "{}" \;
    find "${BASE_DIR}" -type f -not -name "*.j2" -exec bash -c 'mv {} {}.j2' "{}" \;
    # find "${BASE_DIR}" -type f ! -name "Makefile*" ! -name "cmd.sh*" ! -name "*.j2" -exec bash -c 'mv {} {}.j2' "{}" \;
}

# Usage:
# create_file_from_j2_file \
#     "INFILE:${TOP_DIR:?}/data/template/app.conf.j2" \
#     "OUTFILE:${TOP_DIR:?}/data/rootfs/etc/app.conf" \
#     "PROJECT_TYPE:${PROJECT_TYPE}" \
#     "GIT_PROVIDER:${GIT_PROVIDER}" \
#     "GITHUB_USER:${GITHUB_USER}" \
#     "GITHUB_PROJECT:${GITHUB_PROJECT}" \
#     "SERVICE_PORT:${SERVICE_PORT}"
function create_file_from_j2_file() {
    if [[ $# < 2 ]]; then
        log_error "Usage: ${FUNCNAME[0]} \"INFILE:${TOP_DIR:?}/conf/${SERVICE_HOME}.conf\" [\"OUTFILE:${SERVICE_HOME}/etc/${SERVICE_HOME}.config\" ] \"<VARIABLE>:${VARIABLE}\""
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local INDEX=1
    local PARAMETERS=()
    local J2_ENABLE=true
    for ELEMENT in "$@"; do
        IFS=':', read -ra ARRAY <<<"${ELEMENT}"
        if [[ "${ARRAY[0]}" = "INFILE" ]]; then
            local INPUT_FILE=${ARRAY[1]}
            continue
        elif [[ "${ARRAY[0]}" = "OUTFILE" ]]; then
            local OUTPUT_FILE=${ARRAY[1]}
            continue
        elif [[ "${ARRAY[0]}" = "J2_ENABLE" ]]; then
            J2_ENABLE=${ARRAY[1]}
            continue
        fi
        if [ ${J2_ENABLE} == true ]; then
            PARAMETERS[${INDEX}]="-e 's#{{ ${ARRAY[0]} }}#${ARRAY[1]}#g' "
        else
            PARAMETERS[${INDEX}]="-e 's#${ARRAY[0]}#${ARRAY[1]}#g' "
        fi
        PARAMETERS+=${PARAMETERS[${INDEX}]}
        i=$((${INDEX} + 1))
    done

    if [[ "${OUTPUT_FILE}" ]]; then
        if [ ! -e "${OUTPUT_FILE}.bak" ] && [ -e "${OUTPUT_FILE}" ]; then
            cp "${OUTPUT_FILE}" "${OUTPUT_FILE}.bak"
        fi
        eval sed ${PARAMETERS} "${INPUT_FILE}" >"${OUTPUT_FILE}"
    else
        if [ ! -e "${INPUT_FILE}.bak" ] && [ -e "${INPUT_FILE}" ]; then
            cp "${INPUT_FILE}" "${INPUT_FILE}.bak"
        fi
        eval sed -i ${PARAMETERS} "${INPUT_FILE}"
    fi
}

# Usage:
# generate_directories_and_files_from_j2_directory \
#     "TEMPLATE_DIR:/project/templates" \
#     "PROJECT_DIR:/project" \
#     "J2_ENABLE:true" \
#     "PROJECT_TYPE:${PROJECT_TYPE}" \
#     "GIT_PROVIDER:${GIT_PROVIDER}" \
#     "GITHUB_USER:${GITHUB_USER}" \
#     "GITHUB_PROJECT:${GITHUB_PROJECT}" \
#     "SERVICE_PORT:${SERVICE_PORT}"
function generate_directories_and_files_from_j2_directory() {
    if [ "$#" -lt 3 ]; then
        log_error "Usage: ${FUNCNAME[0]} \"TEMPLATE_DIR:/path/to/template/\" \"PROJECT_DIR:/path/to/project/\" \"J2_ENABLE:true\" \"PARAMETER1:VALUE1\" \"PARAMETER2:VALUE2\" ..."
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local TEMPLATE_DIR=""
    local PROJECT_DIR=""
    local J2_ENABLE=true
    local PARAMETERS=()

    for ELEMENT in "$@"; do
        IFS=':' read -ra ARRAY <<<"${ELEMENT}"
        case "${ARRAY[0]}" in
            "TEMPLATE_DIR")
                TEMPLATE_DIR="${ARRAY[1]}"
                ;;
            "PROJECT_DIR")
                PROJECT_DIR="${ARRAY[1]}"
                ;;
            "J2_ENABLE")
                J2_ENABLE="${ARRAY[1]}"
                ;;
            *)
                if [ "${J2_ENABLE}" == "true" ]; then
                    PARAMETERS+=("-e 's#{{ ${ARRAY[0]} }}#${ARRAY[1]}#g'")
                else
                    PARAMETERS+=("-e 's#${ARRAY[0]}#${ARRAY[1]}#g'")
                fi
                ;;
        esac
    done

    if [ ! -d "${TEMPLATE_DIR}" ]; then
        log_error "Directory ${TEMPLATE_DIR} does not exist"
        exit 1
    fi

    # if [ -d "${PROJECT_DIR}" ]; then
    #     log_error "Directory ${PROJECT_DIR} exists"
    #     exit 1
    # else
    #     mkdir -p "${PROJECT_DIR}"
    # fi

    # log_verbose "TEMPLATE_DIR: ${TEMPLATE_DIR}"
    # log_verbose "PROJECT_DIR: ${PROJECT_DIR}"
    # log_verbose "J2_ENABLE: ${J2_ENABLE}"
    # log_verbose "PARAMETERS: ${PARAMETERS[@]}"

    cd "${TEMPLATE_DIR}" || exit 1
    local FILES=$(find . -type f -not -path '*/dist*' -not -path '*/node_modules*' -not -path '*/.git*')

    cd "${PROJECT_DIR}" || exit 1

    for FILE in ${FILES[@]}; do
        # log_info "${TEMPLATE_DIR}/${FILE}"
        BASENAME=$(basename "${FILE}")
        NEW_FILE="${PROJECT_DIR}/$(dirname ${FILE})/${BASENAME%.*}"
        log_info "Generating: ${NEW_FILE}"

        if [ ! -e "$(dirname ${NEW_FILE})" ]; then
            mkdir -p "$(dirname ${NEW_FILE})"
        fi

        if [ ! -e "${NEW_FILE}.bak" ] && [ -e "${NEW_FILE}" ]; then
            cp "${NEW_FILE}" "${NEW_FILE}.bak"
        fi

        echo eval sed "${PARAMETERS}" "${TEMPLATE_DIR}/${FILE}" "${NEW_FILE}"
        eval sed ${PARAMETERS} "${TEMPLATE_DIR}/${FILE}" >"${NEW_FILE}"
        if [ "${FILE: -3}" == ".sh" ]; then
            chmod +x "${FILE}"
        fi
    done
}

# ******************************************************************************
# Reformat files in directory and subdirectories.

function reformat_files_in_dir_with_file_pattern() {
    if [ "$#" != "2" ]; then
        log_error "Usage: ${FUNCNAME[0]} <BASE_DIR> <re/*.sh/*.py>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    local BASE_DIR="${1}"
    local FILE_PATTERN="${2}"

    log_verbose "Cleaning files in ${BASE_DIR} with pattern ${FILE_PATTERN}"

    # Delete all leading blank lines at the top of each file.
    find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec sed -i '/./,$!d' "{}" \;

    # Delete all trailing blank lines at the end of each file.
    find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec sed -i -e :a -e '/^\n*$/{$d;N;};/\n$/ba' "{}" \;

    # Remove leading white spaces from the beginning of each file.
    find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec sed -i 's/^ *#!/#!/' "{}" \;

    # Remove white spaces from the end of each line.
    find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec sed -i 's/[[:space:]]\{1,\}$//' "{}" \;

    # Add a single blank line to the end of each file.
    find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec sed -i -e '$a\' "{}" \;

    if [ "${FILE_PATTERN}" == "*Makefile" ] || [ "${FILE_PATTERN}" == "*Makefile.j2" ]; then
        # Convert tabs to spaces with a tab width of 4 for specific file types.
        # - [](<https://stackoverflow.com/questions/16931770/makefile4-missing-separator-stop>)
        find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec bash -c 'unexpand -a -t 4 {} > {}.tmp; mv {}.tmp {};' {} \;
    elif [ "${FILE_PATTERN}" == "*.go" ] || [ "${FILE_PATTERN}" == "*.go.j2" ]; then
        # Convert tabs to spaces with a tab width of 4 for specific file types.
        # - [](<https://stackoverflow.com/questions/16931770/makefile4-missing-separator-stop>)
        find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec bash -c 'unexpand -a -t 4 {} > {}.tmp; mv {}.tmp {};' {} \;
    elif [ "${FILE_PATTERN}" == "*.bats" ]; then
        # For *.bats files, you can add specific actions if needed.
        :
    else
        # Change tabs to spaces with a tab width of 4 for other file types.
        # - [](<https://github.com/wilzbach/shell-tools/blob/master/expandall>)
        find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec bash -c 'expand -t 4 {} > {}.tmp; mv {}.tmp {};' {} \;
    fi
}

# ******************************************************************************
#set +e # Exit on error Off
#set +x # Trace Off
#echo "End: $(basename "${0}")"
# echo -e "\n================================================================================\n"
# exit 0
# ******************************************************************************
