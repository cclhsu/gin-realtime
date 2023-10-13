#!/bin/bash

# Get the current directory
CURRENT_DIR="$0"
# CURRENT_DIR=$(dirname $(dirname $(dirname "$0")))

# Function to rename files
rename_files() {
    local PATH_TO_DIRECTORY="$1"
    local OLD_STRING="websocket"
    local NEW_STRING="webpush"

    if [ -z "${PATH_TO_DIRECTORY}" ]; then
        echo "Path to directory must be set."
        return
    fi

    # Ensure that both old and new strings are set
    if [ -z "$OLD_STRING" ] || [ -z "$NEW_STRING" ]; then
        echo "Both OLD_STRING and NEW_STRING must be set."
        return
    fi

    # find directory with old string in name and replace with new string
    local DIRS=$(find ${PATH_TO_DIRECTORY} -type d -name "*$OLD_STRING*")
    for DIR in ${DIRS[*]}; do
        echo "Processing ${DIR} in ${PATH_TO_DIRECTORY}..."
        if [ -d "${DIR}" ]; then
            # Create the new name by replacing OLD_STRING with NEW_STRING
            NEW_NAME="${DIR//$OLD_STRING/$NEW_STRING}"
            mv "${DIR}" "${NEW_NAME}"
            echo "Renamed: ${DIR} -> ${NEW_NAME}"
        fi
    done

    # Loop through all files and directories in the current directory recursively and rename files
    for ENTRY in "${PATH_TO_DIRECTORY}"/*; do
        echo "Processing ${ENTRY} in ${PATH_TO_DIRECTORY}..."
        if [ -d "${ENTRY}" ]; then
            # Recursive call for subdirectories
            rename_files "${ENTRY}"
        else
            # Check if the file name contains OLD_STRING
            if [[ "${ENTRY}" == *"$OLD_STRING"* ]]; then
                # Create the new name by replacing OLD_STRING with NEW_STRING
                NEW_NAME="${ENTRY//$OLD_STRING/$NEW_STRING}"
                mv "${ENTRY}" "${NEW_NAME}"
                echo "Renamed: ${ENTRY} -> ${NEW_NAME}"
            fi
        fi
    done

    echo "Done for ${PATH_TO_DIRECTORY}."
}

# Replace 'PATH_TO_DIRECTORY' with the actual path to your directory
PATH_TO_DIRECTORY="/Users/clark.hsu/Documents/myProject/CodingTest/src/golang/takehome/solution/gin-realtime"

# Call the rename_files function with the directory path
rename_files "${PATH_TO_DIRECTORY}"
