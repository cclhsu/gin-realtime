#!/bin/bash

set -e # Exit immediately on error

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

# Get the current directory
CURRENT_DIR=$(dirname $(dirname $(dirname "$0")))
log "Using current directory: ${CURRENT_DIR}"

# Load environment variables from .env file
if [ -f "${CURRENT_DIR}/.env" ]; then
    log "Loading environment variables from .env file"
    source "${CURRENT_DIR}/.env"
fi

# Default values
DEFAULT_PEM_PHRASE="changeme"
DEFAULT_PEM_SUBJECT="/C=TW/ST=Taiwan/L=Taipei/O=Example Inc./OU=Example/CN=example.com"
DEFAULT_PEM_VALIDITY_DAYS=3650
DEFAULT_BYTE=2560

# Function to generate key pair
generate_key_pair() {
    local KEY_DIR="$1"
    local PRIVATE_KEY="$2"
    local PUBLIC_KEY="$3"
    local BYTE="${4:-$DEFAULT_BYTE}"

    mkdir -p "${KEY_DIR}"

    log "Generating ${PRIVATE_KEY} and ${PUBLIC_KEY}"
    openssl genpkey -algorithm RSA -out "${KEY_DIR}/${PRIVATE_KEY}" -pkeyopt rsa_keygen_bits:${BYTE}
    openssl rsa -in "${KEY_DIR}/${PRIVATE_KEY}" -pubout -out "${KEY_DIR}/${PUBLIC_KEY}"
}

# Function to initialize the Certificate Authority
initialize_ca() {
    local CA_DIR="$1"
    local PEM_VALIDITY_DAYS="${2:-$DEFAULT_PEM_VALIDITY_DAYS}"

    mkdir -p "${CA_DIR}"
    export CA_PEM_PHRASE=${CA_PEM_PHRASE:-"${DEFAULT_PEM_PHRASE}"}

    if [ ! -e "${CA_DIR}/ca-key.pem" ]; then
        log "Initializing Certificate Authority..."
        openssl genpkey -algorithm RSA -aes256 -out "${CA_DIR}/ca-key.pem" -pass env:CA_PEM_PHRASE
        openssl req -x509 -new -sha256 -days ${PEM_VALIDITY_DAYS} -key "${CA_DIR}/ca-key.pem" -out "${CA_DIR}/ca-certificate.pem" -passin env:CA_PEM_PHRASE -subj "${DEFAULT_PEM_SUBJECT}"
    else
        log "Certificate Authority already initialized."
    fi
}

# Function to generate certificate signing requests (CSR)
generate_csr() {
    local KEY_DIR="$1"
    local PRIVATE_KEY="$2"
    local CSR_NAME="$3"

    log "Generating ${CSR_NAME} CSR"
    openssl req -new -sha256 -key "${KEY_DIR}/${PRIVATE_KEY}" -out "${KEY_DIR}/${CSR_NAME}-csr.pem" -subj "${DEFAULT_PEM_SUBJECT}"
}

# Function to sign a certificate
sign_certificate() {
    local CA_DIR="$1"
    local KEY_DIR="$2"
    local PRIVATE_KEY="$3"
    local PUBLIC_KEY="$4"
    local CSR_NAME="$5"
    local CERT_NAME="$6"
    local PEM_VALIDITY_DAYS="${7:-$DEFAULT_PEM_VALIDITY_DAYS}"

    log "Signing ${CERT_NAME} Certificate"
    openssl x509 -req -sha256 -days ${PEM_VALIDITY_DAYS} -in "${KEY_DIR}/${CSR_NAME}-csr.pem" -CA "${CA_DIR}/ca-certificate.pem" -passin env:CA_PEM_PHRASE -CAkey "${CA_DIR}/ca-key.pem" -CAcreateserial -out "${KEY_DIR}/${CERT_NAME}-certificate.pem"
}

# Initialize CA
CA_DIR="${CURRENT_DIR}/pki/ca"
initialize_ca "${CA_DIR}" 3650

# Create server and client key pairs
generate_key_pair "${CURRENT_DIR}/pki/servers" "server-key.pem" "server-public.pem"
generate_key_pair "${CURRENT_DIR}/pki/clients" "client-key.pem" "client-public.pem"

# Generate server and client CSR files
generate_csr "${CURRENT_DIR}/pki/servers" "server-key.pem" "server"
generate_csr "${CURRENT_DIR}/pki/clients" "client-key.pem" "client"

# Sign server and client certificates
sign_certificate "${CA_DIR}" "${CURRENT_DIR}/pki/servers" "server-key" "server-public" "server" "server"
sign_certificate "${CA_DIR}" "${CURRENT_DIR}/pki/clients" "client-key" "client-public" "client" "client" 365

# Clean up intermediate CSR files
log "Cleaning up intermediate CSR files"
rm "${CURRENT_DIR}/pki/servers/server-csr.pem" "${CURRENT_DIR}/pki/clients/client-csr.pem"

unset CA_PEM_PHRASE
log "PKI Setup Completed"
