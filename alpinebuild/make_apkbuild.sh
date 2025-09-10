#!/usr/bin/env bash

set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat "${SCRIPT_DIR}/../VERSION")

VOLUME_DIR=${SCRIPT_DIR}/volumes
OVERLAY=${VOLUME_DIR}/apkbuild-overlay

# Defaults
LOCAL=0
OUTPUT_DIR=""

# Parse arguments (order-independent)
for arg in "$@"; do
    case "$arg" in
        --local)
            LOCAL=1
            ;;
        *)
            # If OUTPUT_DIR not set yet, treat as directory
            if [[ -z "$OUTPUT_DIR" ]]; then
                OUTPUT_DIR="$arg"
            fi
            ;;
    esac
done

# Set default directory if not provided
if [[ -z "$OUTPUT_DIR" ]]; then
    if [[ $LOCAL -eq 1 ]]; then
        OUTPUT_DIR="${OVERLAY}"
    else
        OUTPUT_DIR="${SCRIPT_DIR}"
    fi
fi

# Ensure output directory exists
mkdir -p "${OUTPUT_DIR}"
APKBUILD_PATH="${OUTPUT_DIR}/APKBUILD"

# Write APKBUILD
echo "# Maintainer: Dimitrios Desyllas <pcmagas@disroot.org>" > "${APKBUILD_PATH}"
echo "pkgname=mkdotenv" >> "${APKBUILD_PATH}"
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

echo "APKBUILD written to ${APKBUILD_PATH}"
