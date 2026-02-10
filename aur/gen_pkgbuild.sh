#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

VERSION=$(cat ${SCRIPT_DIR}/../VERSION)

# ${1:-$SCRIPT_DIR} gets the value
# %/ removes a single trailing slash if it exists
TARGET_DIR="${1:-$SCRIPT_DIR}"
TARGET_DIR="${TARGET_DIR%/}"

# Defaults
LOCAL=0
OUTPUT_DIR="${SCRIPT_DIR}"
PKGNAME="mkdotenv"


# Parse arguments (order-independent)
while [[ $# -gt 0 ]]; do
    case "$1" in
        --src_local)
            LOCAL=1
            shift
            ;;

        *)
            TARGET_DIR="$1"
            shift
            ;;
    esac
done

mkdir -p "${TARGET_DIR}"
PKGBUILD_PATH=${TARGET_DIR}/PKGBUILD

echo "Generating ${PKGBUILD_PATH}"

echo "# Maintainer: Dimitrios Desyllas <pcmagas@disroot.org>" > $PKGBUILD_PATH
echo "pkgname='mkdotenv'">>$PKGBUILD_PATH
echo "pkgver=$VERSION">>$PKGBUILD_PATH
echo "pkgrel=1">>$PKGBUILD_PATH
echo "pkgdesc=\"Populate .env files from secrets.\"">>$PKGBUILD_PATH
echo "arch=('x86_64')">>$PKGBUILD_PATH
echo "url=\"https://github.com/pc-magas/mkdotenv\"">>$PKGBUILD_PATH
echo "license=('GPL-3')">>$PKGBUILD_PATH
echo "depends=()">>$PKGBUILD_PATH
echo "makedepends=()">>$PKGBUILD_PATH


if [[ $LOCAL -eq 0 ]]; then
    SOURCEVAL="source=(\"\$pkgname-\$pkgver.tar.gz::https://github.com/pc-magas/mkdotenv/releases/download/v\$pkgver/mkdotenv-\$pkgver.tar.gz\")"

    # IS_PRERELEASE Is an environmental varialbe that github actions exposes.
    if [[ $IS_PRERELEASE == "true" ]]; then
        SOURCEVAL="source=(\"\$pkgname-\$pkgver.tar.gz::https://github.com/pc-magas/mkdotenv/releases/download/v\$pkgver/mkdotenv-\$pkgver.tar.gz\")"
    fi

    echo ${SOURCEVAL} >> "${PKGBUILD_PATH}"
else
    # TODO check if file exists
    echo "source=(\"\$pkgname-\$pkgver.tar.gz\")" >> "${PKGBUILD_PATH}"
fi

echo "">>$PKGBUILD_PATH
cat << 'EOF' >> $PKGBUILD_PATH
prepare() {
  curl -LO https://go.dev/dl/go1.25.3.linux-amd64.tar.gz
  tar -C "$srcdir" -xzf go1.25.3.linux-amd64.tar.gz
  rm go1.25.3.linux-amd64.tar.gz
}

build() {
  export PATH="$srcdir/go/bin:$PATH"
  make bin VERSION="${pkgver}"
}

package() {
  make install \
    DESTDIR="${pkgdir}" \
    INSTALL_BIN_DIR="/usr/bin" \
    INSTALL_MAN_DIR="/usr/share/man/man1"
}
EOF

# NOTE tar.gz should exist on same folder with 
echo "Fxing Checksums"
docker run --rm -i -v "${TARGET_DIR}":/home/builder pcmagas/arch-pkg-builder run_fixperm updpkgsums

SRC_INFO=${TARGET_DIR}/.SRCINFO
echo "Generating  ${SRC_INFO}:"
docker run --rm -i -v "${TARGET_DIR}":/home/builder pcmagas/arch-pkg-builder run_fixperm makepkg --printsrcinfo > ${TARGET_DIR}/.SRCINFO

