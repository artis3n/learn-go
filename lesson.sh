#!/usr/bin/env bash

NAME=$1

if [[ "$(basename -- "$0")" = "lesson.sh\n" ]]; then
    # Script is run directly, not sourced
    echo "source lesson.sh <lesson-title>"
    exit 0
else
    # Script is sourced
    if [[ $# -eq 0 ]]; then
        # No parameters passed in
        echo "source lesson.sh <lesson-title>"
        # Only run `exit` if script is not being sourced
        (return 0 2>/dev/null) && return || exit 0
    fi

    cd ~/Dropbox/Development/learn-go
    mkdir "${NAME}"
    cd "${NAME}"
    touch "${NAME}".go
    code "${NAME}".go
fi
