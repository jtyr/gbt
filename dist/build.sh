#!/bin/bash

# Import GPG key
gpg --import "$TRAVIS_BUILD_DIR/dist/gpg_key.priv"
echo -e '%_gpg_name Jiri Tyr (PKG) <jiri.tyr@gmail.com>\n%dist .el7' > ~/.rpmmacros

# Create RPM infrastructure
mkdir -p ~/rpmbuild/SOURCES
echo -e '#!/usr/bin/expect -f\nspawn rpmsign --key-id CA67951CD2BBE8AAE4210B72FB90C91F64BED28C --addsign {*}$argv\nexpect -exact "Enter pass phrase: "\nsend -- "\\r"\nexpect eof' > ~/rpm-sign.exp
chmod +x ~/rpm-sign.exp

# Archs to build for
declare -a PLATFORMS=(
    'darwin/amd64'
    'linux/amd64'
    'linux/arm:5'
    'linux/arm:6'
    'linux/arm64'
)

NAME='gbt'
VER="${TRAVIS_TAG:1}"
BUILD="${TRAVIS_COMMIT::6}"
TMP="/tmp/$NAME"

rm -fr "$TMP"

# Process build for each arch
for P in "${PLATFORMS[@]}"; do
    echo "### Building $P"

    PTMP="$TMP/$P/$NAME-$VER"
    OS="${P%%/*}"
    ARCH="${P#*/}"
    ARM="${ARCH#*:}"
    PKG="$NAME-$VER-$OS-$ARCH.tar.gz"

    mkdir -p "$PTMP"

    # Compile GBT
    if [ "$ARCH" == "$ARM" ]; then
        GOOS="$OS" GOARCH="$ARCH" CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=$VER -X main.build=$BUILD" -o "$PTMP/$NAME" github.com/jtyr/gbt/cmd/gbt
    else
        ARCH="${ARCH%%:*}"
        PKG="$NAME-$VER-$OS-$ARCH$ARM.tar.gz"

        GOOS="$OS" GOARCH="$ARCH" GOARM="$ARM" CGO_ENABLED=0 go build -ldflags='-s -w' -o "$PTMP/$NAME" github.com/jtyr/gbt/cmd/gbt
    fi

    # Create .tar.gz package
    (
        cp -r "$TRAVIS_BUILD_DIR"/{README.md,LICENSE,themes,sources} "$PTMP"
        tar -C "$PTMP/.." -czf "$TMP/$PKG" ./
        cd "$TMP"
    )

    # Create Linux distro packages
    if [ "$OS" = 'linux' ]; then
        # DEB
        (
            DEBARCH="$ARCH"

            if [ "$ARM" == '5' ]; then
                DEBARCH='armel'
            elif [ "$ARM" == '6' ]; then
                DEBARCH='armhf'
            fi

            cd "$TRAVIS_BUILD_DIR/contrib"
            rm -f "$TRAVIS_BUILD_DIR/contrib/$NAME"
            ln -s "$PTMP" "$TRAVIS_BUILD_DIR/contrib/$NAME"
            m4 -DVER="$VER" -DDATE="$(date '+%a, %d %b %Y %H:%M:%S %z')" debian/changelog.m4 > debian/changelog
            dpkg-buildpackage -a$DEBARCH -tc -b -kCA67951CD2BBE8AAE4210B72FB90C91F64BED28C
        )
        debsigs --sign=origin -k CA67951CD2BBE8AAE4210B72FB90C91F64BED28C "$TRAVIS_BUILD_DIR"/*.deb
        mv "$TRAVIS_BUILD_DIR"/*.deb "$TMP"

        # RPM
        if [ "$ARCH" = 'amd64' ]; then
            ln -sf "$TMP/$PKG" ~/rpmbuild/SOURCES/
            (
                cd "$TRAVIS_BUILD_DIR/contrib/redhat"
                m4 -DVER="$VER" -DDATE="$(date '+%a %b %d %Y')" gbt.spec.m4 > gbt.spec
                rpmbuild -bb gbt.spec
            )
            ~/rpm-sign.exp ~/rpmbuild/RPMS/x86_64/*.rpm
            mv ~/rpmbuild/RPMS/x86_64/*.rpm "$TMP"
        fi
    fi
done

# Create checksums
cd "$TMP"
sha256sum *.tar.gz *.deb *.rpm | sort -k2 > "$NAME-$VER-checksums-sha256.txt"
