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

# Default values
DEFAULT_PEM_PHRASE="changeme"
DEFAULT_PEM_SUBJECT="/C=TW/ST=Taiwan/L=Taipei/O=Example Inc./OU=Example/CN=example.com"
DEFAULT_PEM_VALIDITY_DAYS=3650
DEFAULT_BYTE=2560

# ******************************************************************************
# Parameters

# # Check if the file or directory in the argument exists
# if [ $# -eq 0 ]; then
#     log_error "Usage: ${0} <FILE.ext> or ${0} <DIRECTORY>"
#     exit 1
# fi

# ******************************************************************************
# Functions

# Function to generate key pair
generate_key_pair() {
    if [ "$#" != "3" ] && [ "$#" != "4" ]; then
        log_error "Usage: ${FUNCNAME[0]} <KEY_DIR> <PRIVATE_KEY> <PUBLIC_KEY> [<BYTE>]"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local KEY_DIR="${1}"
    local PRIVATE_KEY="${2}"
    local PUBLIC_KEY="${3}"
    local BYTE="${4:-$DEFAULT_BYTE}"

    mkdir -p "${KEY_DIR}"

    log_info "Generating ${PRIVATE_KEY} and ${PUBLIC_KEY}"
    openssl genpkey -algorithm RSA -out "${KEY_DIR}/${PRIVATE_KEY}" -pkeyopt rsa_keygen_bits:${BYTE}
    openssl rsa -in "${KEY_DIR}/${PRIVATE_KEY}" -pubout -out "${KEY_DIR}/${PUBLIC_KEY}"
}

# Function to initialize the Certificate Authority
initialize_ca() {
    if [ "$#" != "1" ] && "$#" != "2" ]; then
        log_error "Usage: ${FUNCNAME[0]} <CA_DIR> [<PEM_VALIDITY_DAYS>]"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local CA_DIR="${1}"
    local PEM_VALIDITY_DAYS="${2:-$DEFAULT_PEM_VALIDITY_DAYS}"

    mkdir -p "${CA_DIR}"
    export CA_PEM_PHRASE=${CA_PEM_PHRASE:-"${DEFAULT_PEM_PHRASE}"}

    if [ ! -e "${CA_DIR}/ca-key.pem" ]; then
        log_info "Initializing Certificate Authority..."
        openssl genpkey -algorithm RSA -aes256 -out "${CA_DIR}/ca-key.pem" -pass env:CA_PEM_PHRASE
        openssl req -x509 -new -sha256 -days ${PEM_VALIDITY_DAYS} -key "${CA_DIR}/ca-key.pem" -out "${CA_DIR}/ca-certificate.pem" -passin env:CA_PEM_PHRASE -subj "${DEFAULT_PEM_SUBJECT}"
    else
        log_info "Certificate Authority already initialized."
    fi
}

# Function to generate certificate signing requests (CSR)
generate_csr() {
    if [ "$#" != "3" ]; then
        log_error "Usage: ${FUNCNAME[0]} <KEY_DIR> <PRIVATE_KEY> <CSR_NAME>"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local KEY_DIR="${1}"
    local PRIVATE_KEY="${2}"
    local CSR_NAME="${3}"

    log_info "Generating ${CSR_NAME} CSR"
    openssl req -new -sha256 -key "${KEY_DIR}/${PRIVATE_KEY}" -out "${KEY_DIR}/${CSR_NAME}-csr.pem" -subj "${DEFAULT_PEM_SUBJECT}"
}

# Function to sign a certificate
sign_certificate() {
    if [ "$#" != "6" ] && [ "$#" != "7" ]; then
        log_error "Usage: ${FUNCNAME[0]} <CA_DIR> <KEY_DIR> <PRIVATE_KEY> <PUBLIC_KEY> <CSR_NAME> <CERT_NAME> [<PEM_VALIDITY_DAYS>]"
        log_error "[${FUNCNAME[0]}] [$#] ${*}"
        exit 1
    fi

    log_verbose "[${FUNCNAME[0]}] [$#] ${*}"
    # cd "${TOP_DIR:?}" || exit 1

    local CA_DIR="${1}"
    local KEY_DIR="${2}"
    local PRIVATE_KEY="${3}"
    local PUBLIC_KEY="${4}"
    local CSR_NAME="${5}"
    local CERT_NAME="${6}"
    local PEM_VALIDITY_DAYS="${7:-$DEFAULT_PEM_VALIDITY_DAYS}"

    log_info "Signing ${CERT_NAME} Certificate"
    openssl x509 -req -sha256 -days ${PEM_VALIDITY_DAYS} -in "${KEY_DIR}/${CSR_NAME}-csr.pem" -CA "${CA_DIR}/ca-certificate.pem" -passin env:CA_PEM_PHRASE -CAkey "${CA_DIR}/ca-key.pem" -CAcreateserial -out "${KEY_DIR}/${CERT_NAME}-certificate.pem"
}

# ******************************************************************************
# Main Program

# Initialize CA
CA_DIR="${PROJECT_DIR}/pki/ca"
initialize_ca "${CA_DIR}" 3650

# Create server and client key pairs
generate_key_pair "${PROJECT_DIR}/pki/servers" "server-key.pem" "server-public.pem"
generate_key_pair "${PROJECT_DIR}/pki/clients" "client-key.pem" "client-public.pem"

# Generate server and client CSR files
generate_csr "${PROJECT_DIR}/pki/servers" "server-key.pem" "server"
generate_csr "${PROJECT_DIR}/pki/clients" "client-key.pem" "client"

# Sign server and client certificates
sign_certificate "${CA_DIR}" "${PROJECT_DIR}/pki/servers" "server-key" "server-public" "server" "server"
sign_certificate "${CA_DIR}" "${PROJECT_DIR}/pki/clients" "client-key" "client-public" "client" "client" 365

# Clean up intermediate CSR files
log_info "Cleaning up intermediate CSR files"
rm "${PROJECT_DIR}/pki/servers/server-csr.pem" "${PROJECT_DIR}/pki/clients/client-csr.pem"

unset CA_PEM_PHRASE
log_info "PKI Setup Completed"

# ******************************************************************************
#set +e # Exit on error Off
#set +x # Trace Off
#echo "End: $(basename "${0}")"
echo -e "\n================================================================================\n"
exit 0
# ******************************************************************************
