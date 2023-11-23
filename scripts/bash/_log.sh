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
# -[](<https://en.wikipedia.org/wiki/ANSI_escape_code>)
# -[](<https://gist.github.com/JBlond/2fea43a3049b38287e5e9cefc87b2124>)
# -[](<https://github.com/fidian/ansi>)

ENDCOLOR="\033[0m"

BOLD_BLACK="\033[1;30m"   # output
BOLD_RED="\033[1;31m"     # error
BOLD_GREEN="\033[1;32m"   # success
BOLD_YELLOW="\033[1;33m"  # warn
BOLD_BLUE="\033[1;34m"    # info
BOLD_MAGENTA="\033[1;35m" # verbose
BOLD_CYAN="\033[1;36m"    # debug
BOLD_GRAY="\033[1;37m"
BOLD_BRIGHT_GRAY="\033[1;90m"
BOLD_BRIGHT_RED="\033[1;91m" # fatal
BOLD_BRIGHT_GREEN="\033[1;92m"
BOLD_BRIGHT_YELLOW="\033[1;93m"
BOLD_BRIGHT_BLUE="\033[1;94m"
BOLD_BRIGHT_MAGENTA="\033[1;95m"
BOLD_BRIGHT_CYAN="\033[1;96m"
BOLD_WHITE="\033[1;97m"

BOLD_BLACK_BG="\033[1;40m"
BOLD_RED_BG="\033[1;41m"
BOLD_GREEN_BG="\033[1;42m"
BOLD_YELLOW_BG="\033[1;43m"
BOLD_BLUE_BG="\033[1;44m"
BOLD_MAGENTA_BG="\033[1;45m"
BOLD_CYAN_BG="\033[1;46m"
BOLD_GRAY_BG="\033[1;47m"
BOLD_BRIGHT_Gray_BG="\033[1;100m"
BOLD_BRIGHT_RED_BG="\033[1;101m"
BOLD_BRIGHT_GREEN_BG="\033[1;102m"
BOLD_BRIGHT_YELLOW_BG="\033[1;103m"
BOLD_BRIGHT_BLUE_BG="\033[1;104m"
BOLD_BRIGHT_MAGENTA_BG="\033[1;105m"
BOLD_BRIGHT_CYAN_BG="\033[1;106m"
BOLD_WHITE_BG="\033[1;107m"

function section() {
    echo ""
    echo "==============================================================================="
    echo ""
}

function divider() {
    echo ""
    echo "-------------------------------------------------------------------------------"
    echo ""
}

# fatal
function log_fatal() {
    echo -e "${BOLD_BRIGHT_RED}\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [FATAL] ${*} ...\n${ENDCOLOR}" # bright red
    exit 1
}

# error
function log_error() {
    echo -e "${BOLD_RED}\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [ERROR] ${*} ...\n${ENDCOLOR}" # red
    # exit 1
}

# warn | failure
function log_warn() {
    echo -e "${BOLD_YELLOW}\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [WARN] ${*} ...\n${ENDCOLOR}" # yellow
}

# info | success
function log_info() {
    echo -e "${BOLD_BLUE}\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [INFO] ${*} ...\n${ENDCOLOR}" # blue
}

function log_exec() {
    echo -e "${BOLD_GREEN}\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [EXEC] ${*} ...\n${ENDCOLOR}" # green
    eval "${*}"
}

# output | console
function log_output() {
    echo -e "\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [OUTPUT] ${*} ...\n" # black
}

# debug | message
function log_debug() {
    echo -e "${BOLD_CYAN}\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [DEBUG] ${*} ...\n${ENDCOLOR}" # cyan
    # divider
}

# verbose | message
function log_verbose() {
    echo -e "${BOLD_MAGENTA}\n>>> [$(date +'%Y-%m-%d %H:%M:%S')] [VERBOSE] ${*} ...\n${ENDCOLOR}" # magenta
    # section
}

# ******************************************************************************
