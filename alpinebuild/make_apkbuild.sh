#!/usr/bin/env bash

set -e

echo "Script Started"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat "${SCRIPT_DIR}/../VERSION")

VOLUME_DIR=${SCRIPT_DIR}/volumes

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
        --checksum)
            CHECKSUM="$2"
            shift 2
            ;;
        *)
            OUTPUT_DIR="$1"
            shift
            ;;
    esac
done

echo "HERE"

# Set default directory if not provided
if [[ -z "$OUTPUT_DIR" ]]; then
    OUTPUT_DIR="${SCRIPT_DIR}"
fi


# Ensure output directory exists
mkdir -p "${OUTPUT_DIR}"
APKBUILD_PATH="${OUTPUT_DIR}/APKBUILD"

echo "Version: ${VERSION}"
echo "Provided SHA512: ${CHECKSUM}"
echo "APKBUILD path ${APKBUILD_PATH}"

# Write APKBUILD
echo "# Maintainer: Dimitrios Desyllas <pcmagas@disroot.org>" > "${APKBUILD_PATH}"
echo "pkgname=${PKGNAME}" >> "${APKBUILD_PATH}"
echo "pkgver=${VERSION}" >> "${APKBUILD_PATH}"
echo "pkgrel=0" >> "${APKBUILD_PATH}"
echo "pkgdesc=\"Lightweight and efficient tool for managing your .env files.\"" >> "${APKBUILD_PATH}"
echo "url=\"https://github.com/pc-magas/mkdotenv\"" >> "${APKBUILD_PATH}"
echo "arch=\"all\"" >> "${APKBUILD_PATH}"
echo "license=\"GPL-3.0-only\"" >> "${APKBUILD_PATH}"
echo "makedepends=\"go\"" >> "${APKBUILD_PATH}"

if [[ $LOCAL -eq 0 ]]; then
    echo "source=\"\$pkgname-\$pkgver.tar.gz::https://github.com/pc-magas/mkdotenv/releases/download/v\$pkgver/mkdotenv-\$pkgver.tar.gz\"" >> "${APKBUILD_PATH}"
else
    echo "source=\"\$pkgname-\$pkgver.tar.gz\"" >> "${APKBUILD_PATH}"
fi

echo "options=\"!check\" # No tests" >> "${APKBUILD_PATH}"

echo "" >> "${APKBUILD_PATH}"

if [[ -d "${SCRIPT_DIR}/APKBUILD.d" ]]; then
    for step_file in "${SCRIPT_DIR}/APKBUILD.d/"*; do
        base_file=$(basename "$step_file")
        [[ "$base_file" == ".gitignore" ]] && continue
        if [[ -f "$step_file" ]]; then
            echo "source file: $base_file"
            echo "$base_file(){" >> "${APKBUILD_PATH}"
            sed '/./!d' "$step_file" | sed 's/^[[:space:]]*//; s/[[:space:]]*$//' | tr -d "\n" | sed 's/^/    /' >> "${APKBUILD_PATH}"
            echo $value >> "${APKBUILD_PATH}"
            echo "}" >> "${APKBUILD_PATH}"
            echo "" >> "${APKBUILD_PATH}"
        fi
    done
fi

if [[ -n "$CHECKSUM" ]]; then

    echo "Write checksum"
    echo "sha512sums=\"" >> "${APKBUILD_PATH}"
    echo "${CHECKSUM} ${PKGNAME}-${VERSION}.tar.gz"  >> "${APKBUILD_PATH}"
    echo "\"" >> "${APKBUILD_PATH}"

fi


echo "APKBUILD written to ${APKBUILD_PATH}"
