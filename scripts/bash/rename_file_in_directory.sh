#!/bin/bash

# Get the current directory
CURRENT_DIR="$0"
# CURRENT_DIR=$(dirname $(dirname $(dirname "$0")))

OLD_STRING="health"
NEW_STRING="health"

# Function to rename files
rename_files() {
    PATH_TO_DIRECTORY="$1"
    for FILE in "${PATH_TO_DIRECTORY}"/*; do
        echo "Processing ${FILE} file in ${PATH_TO_DIRECTORY}..."
        if [[ -d "${FILE}" ]]; then
            # Recursive call for subdirectories
            rename_files "${FILE}"
        else
            # Check if the file name contains "metric"
            if [[ "${FILE}" == *"${OLD_STRING}"* ]]; then
                NEW_FILE="${FILE/${OLD_STRING}/${NEW_STRING}}"
                mv "${FILE}" "${NEW_FILE}"
                echo "Renamed: ${FILE} -> ${NEW_FILE}"
            fi
        fi
    done
}

# Replace 'PATH_TO_DIRECTORY' with the actual path to your directory
PATH_TO_DIRECTORY="/Users/clark.hsu/Documents/myProject/CodingTest/src/golang/takehome/solution/gin-realtime"

# Call the rename_files function with the directory path
rename_files "${PATH_TO_DIRECTORY}"
