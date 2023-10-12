#!/bin/bash

# Function to log messages with timestamps
function log() {
    echo -e "\033[1;32m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ${*} ...\n\033[0m" # green
}

# Function to log error messages with timestamps
function error() {
    echo -e "\033[1;31m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ERROR: ${1} ...\n\033[0m" # red
}

# Function to replace placeholders in files
function replace_placeholders() {
    local dir="$1"
    local file="$2"
    local placeholder="$3"
    local value="$4"
    sed -i "s#{{ ${placeholder} }}#${value}#g" "${dir}/${file}"
}

# Get the current directory
CURRENT_DIR=$(dirname $(dirname $(dirname "$0")))
cd "${CURRENT_DIR}" || exit
CURRENT_DIR="$(pwd)"
log "Using current directory: ${CURRENT_DIR}"

# Find all files named *.j2 in the current directory
J2_FILE=*.j2
FILES=$(find "${CURRENT_DIR}" -type f -name "${J2_FILE}" -not -path '*/dist*' -not -path '*/.git*')

# Define the values to replace placeholders with
# GIT_PROVIDER="YourGitProvider"
# PROJECT_USER="YourProjectUser"
# PROJECT_NAME="YourProjectName"
# PROJECT_PORT="YourProjectPort"

# Loop through subdirectories
for FILE in ${FILES}; do
    # Copy the file to a new file without the .j2 extension
    DIR=$(dirname "${FILE}")
    BASENAME=$(basename "${FILE}")
    NEW_FILE="${DIR}/${BASENAME%.*}"
    if [[ -f "${NEW_FILE}" ]]; then
        # error "File ${NEW_FILE} already exists"
        # exit 1
        rm "${NEW_FILE}"
    fi
    cp "${FILE}" "${NEW_FILE}"
    log "Copied ${FILE} to ${NEW_FILE}"
    if [[ ! "${FILE}" =~ "src/template" ]]; then
        rm "${FILE}"
    fi

    # Replace placeholders in the file
    # replace_placeholders "${DIR}" "${BASENAME%.*}" "gitProvider" "${GIT_PROVIDER}"
    # replace_placeholders "${DIR}" "${BASENAME%.*}" "projectUser" "${PROJECT_USER}"
    # replace_placeholders "${DIR}" "${BASENAME%.*}" "projectName" "${PROJECT_NAME}"
    # replace_placeholders "${DIR}" "${BASENAME%.*}" "projectPort" "${PROJECT_PORT}"

    PROJECT_NAME=$(basename $(dirname $(dirname "${NEW_FILE}")))
    replace_placeholders "${DIR}" "${BASENAME%.*}" "projectName" "${PROJECT_NAME}"
done
