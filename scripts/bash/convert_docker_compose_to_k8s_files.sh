#!/bin/bash

# Function to log messages with timestamps
function log() {
    echo -e "\033[1;32m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ${*} ...\n\033[0m" # green
    # echo -e "\033[1;34m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ${*} ...\n\033[0m" # blue
    # echo -e "\033[1;92m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ${*} ...\n\033[0m" # bright green
    # echo -e "\033[1;94m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ${*} ...\n\033[0m" # bright blue
    # eval "${*}"
}

# Function to log error messages with timestamps
function error() {
    # echo -e "\033[1;31m$1\033[0m" # red
    # echo -e "\033[1;91m$1\033[0m" # bright red
    echo -e "\033[1;31m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ERROR: ${1} ...\n\033[0m" # red
    # echo -e "\033[1;91m\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] ERROR: ${1} ...\n\033[0m" # bright red
    # eval "${*}"
}

# Function to convert files to .j2 format
function convert_to_j2() {
    local dir="$1"
    local extension="$2"
    find "${dir}" -type f ! -name "Makefile*" ! -name "cmd.sh*" ! -name "*.j2" -exec mv {} {}.j2 \;
}

# Get the current directory
CURRENT_DIR=$(dirname $(dirname $(dirname "$0")))
cd "${CURRENT_DIR}" || exit
CURRENT_DIR="$(pwd)"
log "Using current directory: ${CURRENT_DIR}"

# Find all files named docker-compose.yml in the current directory
DOCKER_COMPOSE_FILE=docker-compose.yaml.j2 # docker-compose.yml | docker-compose.yaml.j2
FILES=$(find "${CURRENT_DIR}" -type f -name "${DOCKER_COMPOSE_FILE}" -not -path '*/dist*' -not -path '*/.git*')

# Loop through subdirectories
for FILE in ${FILES}; do
    # Exclude the current directory
    if [[ "${FILE}" != "${CURRENT_DIR}" ]]; then
        # Get the directory of the docker-compose.yml file
        DIR=$(dirname $(dirname "${FILE}"))
        log "Converting ${FILE} to Kubernetes files in ${DIR}"

        # Change to the directory for kubernetes-manifest
        cd "${DIR}/kubernetes-manifest" || exit
        # Convert docker-compose.yml to kubernetes-manifest
        find "${DIR}/kubernetes-manifest" \
            ! -name "Makefile" \
            ! -name "Makefile.j2" \
            ! -name "cmd.sh" \
            ! -name "cmd.sh.j2" \
            -type f -exec rm -f {} \;
        kompose --file "${FILE}" convert --with-kompose-annotation=false
        # Rename files to .j2 format if not already in that format
        convert_to_j2 "${DIR}/kubernetes-manifest" ".j2"

        # Change to the directory for kubernetes-helm
        cd "${DIR}/kubernetes-helm" || exit
        # Convert docker-compose.yml to kubernetes-helm
        find "${DIR}/kubernetes-helm" \
            ! -name "Makefile" \
            ! -name "Makefile.j2" \
            ! -name "cmd.sh" \
            ! -name "cmd.sh.j2" \
            -type f -exec rm -f {} \;
        kompose --file "${FILE}" convert --chart --out . --with-kompose-annotation=false
        # Rename files to .j2 format if not already in that format
        convert_to_j2 "${DIR}/kubernetes-helm" ".j2"

        # Change back to the current directory
        cd "${CURRENT_DIR}" || exit
    fi
done
