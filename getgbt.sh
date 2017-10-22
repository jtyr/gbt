#!/bin/bash

set -e
trap "if [[ \"\$?\" != '0' ]]; then echo ERROR; fi" EXIT

if [ -z "$VERSION" ]; then
    VERSION="v1.0.11"
fi
if [ -z "$DOWNLOAD_URL" ]; then
    DOWNLOAD_URL="https://github.com/jtyr/gbt/archive/$VERSION.tar.gz"
fi
if [ -z "$INSTALL_DEST" ]; then
    INSTALL_DEST="/usr/bin/gbp"
fi

msg() {
    echo "$1: $2"

    if [[ $1 == 'E' ]]; then
        exit 1
    fi
}

command_exists() {
    command -v "$1" > /dev/null 2>&1
}

do_install() {
    if command_exists 'gbt'; then
        version="$(gbt -v | cut -d ' ' -f2)"

        if [[ "$VERSION" == "$version" ]]; then
            msg "E" "Latest version is already installed."
        else
            IFS='.' read -r -a VERSION_ARRAY <<< "${VERSION:1}"
            IFS='.' read -r -a version_array <<< "${version:1}"

            if [  "${version_array[0]}" -le "${VERSION_ARRAY[0]}" ] && \
               [  "${version_array[1]}" -le "${VERSION_ARRAY[1]}" ] && \
               [  "${version_array[2]}" -le "${VERSION_ARRAY[2]}" ]; then
                msg "I" "Newer version available."
            else
                msg "E" "Newer version is installed."
            fi
        fi
    fi

    if command_exists 'curl'; then
        echo curl -L -o "$INSTALL_DEST" "$DOWNLOAD_URL"
        msg "I" "Done."
    else
        msg "E" "In order to download GBT, you must have 'curl' installed."
    fi
}

# TODO
#do_install
