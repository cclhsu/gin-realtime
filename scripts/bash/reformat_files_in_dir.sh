#!/bin/bash

reformat_files_in_dir() {
    if [ "$#" != "2" ]; then
        echo "Usage: ${FUNCNAME[0]} <BASE_DIR> <re/*.sh/*.py>"
    else
        local BASE_DIR="${1}"
        local FILE_PATTERN="${2}"

        echo "Cleaning files in ${BASE_DIR} with pattern ${FILE_PATTERN}"

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
            find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec bash -c 'unexpand -a -t 4 {} > {}.tmp; mv {}.tmp {};' {} \;
        elif [ "${FILE_PATTERN}" == "*.go" ] || [ "${FILE_PATTERN}" == "*.go.j2" ]; then
            # Convert tabs to spaces with a tab width of 4 for specific file types.
            find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec bash -c 'unexpand -a -t 4 {} > {}.tmp; mv {}.tmp {};' {} \;
        elif [ "${FILE_PATTERN}" == "*.bats" ]; then
            # For *.bats files, you can add specific actions if needed.
            :
        else
            # Change tabs to spaces with a tab width of 4 for other file types.
            find "${BASE_DIR}" -iname "${FILE_PATTERN}" -type f -print -exec bash -c 'expand -t 4 {} > {}.tmp; mv {}.tmp {};' {} \;
        fi
    fi
}

# Replace 'PATH_TO_DIRECTORY' with the actual path to your directory
PATH_TO_DIRECTORY="/Users/clark.hsu/src/myProject/project-suite/project-suite-cli/src/template"

# Call the reformat_files_in_dir function with the directory path
# reformat_files_in_dir "/path/to/directory" "*.sh"
# reformat_files_in_dir "/path/to/directory" "*.md"
reformat_files_in_dir "${PATH_TO_DIRECTORY}" "*.j2"
reformat_files_in_dir "${PATH_TO_DIRECTORY}" "*Makefile"
reformat_files_in_dir "${PATH_TO_DIRECTORY}" "*Makefile.j2"
