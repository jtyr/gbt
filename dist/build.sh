#!/bin/bash

declare -a PLATFORMS=(
    'darwin/amd64'
    'linux/amd64'
)

NAME='gbt'
VER="${TRAVIS_TAG:1}"
TMP="/tmp/$NAME"
rm -fr "$TMP"

gpg --import "$TRAVIS_BUILD_DIR/dist/gpg_key.priv"
echo -e '%_gpg_name Jiri Tyr (PKG) <jiri.tyr@gmail.com>\n%dist .el7' > ~/.rpmmacros

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
            dpkg-buildpackage -tc -b -kCA67951CD2BBE8AAE4210B72FB90C91F64BED28C
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
        echo -e '#!/usr/bin/expect -f\nspawn rpmsign --key-id CA67951CD2BBE8AAE4210B72FB90C91F64BED28C --addsign {*}$argv\nexpect -exact "Enter pass phrase: "\nsend -- "\\r"\nexpect eof' > ~/rpm-sign.exp
        chmod +x ~/rpm-sign.exp
        ~/rpm-sign.exp ~/rpmbuild/RPMS/x86_64/*.rpm
        mv ~/rpmbuild/RPMS/x86_64/*.rpm "$TMP"
    fi
done

cd "$TMP"
sha256sum *.tar.gz *.deb *.rpm | sort -k2 > "$NAME-$VER-checksums-sha256.txt"
