#!/bin/bash

declare -a PLATFORMS=(
    'darwin/amd64'
    'linux/amd64'
)

TMP='/tmp/gbt'
rm -fr "$TMP"

for P in "${PLATFORMS[@]}"; do
    echo "Building $P"

    PTMP="$TMP/$P"
    OS="${P%%/*}"
    ARCH="${P#*/}"

    mkdir -p "$PTMP"
    GOOS="$OS" GOARCH="$ARCH" CGO_ENABLED=0 go build -ldflags='-s -w' -o "$PTMP/gbt"

    (
        cp -r "$TRAVIS_BUILD_DIR"/{README.md,LICENSE,themes,sources} "$PTMP"
        tar -C "$PTMP" -czf "$TMP/gbt-$TRAVIS_TAG-$OS-$ARCH.tar.gz" ./
        cd "$TMP"
        sha256sum "gbt-$TRAVIS_TAG-$OS-$ARCH.tar.gz" >> "$TMP/gbt-$TRAVIS_TAG-checksums.txt"
    )
done
