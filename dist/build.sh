#!/bin/bash

declare -a PLATFORMS=(
    'darwin/amd64'
    'linux/amd64'
)

NAME='gbt'
VER="${TRAVIS_TAG:1}"
TMP="/tmp/$NAME"
rm -fr "$TMP"

for P in "${PLATFORMS[@]}"; do
    echo "Building $P"

    PTMP="$TMP/$P/$NAME-$VER"
    OS="${P%%/*}"
    ARCH="${P#*/}"
    PKG="$NAME-$VER-$OS-$ARCH.tar.gz"

    mkdir -p "$PTMP"
    GOOS="$OS" GOARCH="$ARCH" CGO_ENABLED=0 go build -ldflags='-s -w' -o "$PTMP/$NAME"

    (
        cp -r "$TRAVIS_BUILD_DIR"/{README.md,LICENSE,themes,sources} "$PTMP"
        tar -C "$PTMP/.." -czf "$TMP/$PKG" ./
        cd "$TMP"
    )

    if [ "$OS" = 'linux' ]; then
        # DEB
        (
            cd "$TRAVIS_BUILD_DIR/contrib"
            ln -s "$PTMP" "$TRAVIS_BUILD_DIR/contrib/$NAME"
            m4 -DVER="$VER" -DDATE="$(date '+%a, %d %b %Y %H:%M:%S %z')" debian/changelog.m4 > debian/changelog
            dpkg-buildpackage -us -uc -tc -b
        )
        mv "$TRAVIS_BUILD_DIR"/*.deb $TMP

        # RPM
        mkdir -p ~/rpmbuild/SOURCES
        ln -s "$TMP/$PKG" ~/rpmbuild/SOURCES/
        (
            cd "$TRAVIS_BUILD_DIR/contrib/redhat"
            m4 -DVER="$VER" -DDATE="$(date '+%a %b %d %Y')" gbt.spec.m4 > gbt.spec
            rpmbuild -bb gbt.spec
        )
        mv ~/rpmbuild/RPMS/x86_64/*.rpm "$TMP"
    fi
done

cd "$TMP"
sha256sum *.tar.gz *.deb *.rpm | sort -k2 > "$NAME-$VER-checksums-sha256.txt"
