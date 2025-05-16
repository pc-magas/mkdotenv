#!/usr/bin/env bash

echo "MAKE PKGBUILD and .SRCINFO for aur"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)

VOLUME_DIR=${SCRIPT_DIR}/staging
mkdir -p ${VOLUME_DIR}

PKGBUILD_STAGING=${VOLUME_DIR}/PKGBUILD
SOURCEVAL="\$pkgname-\$pkgver.tar.gz::https://github.com/pc-magas/mkdotenv/releases/download/v\$pkgver/mkdotenv-\$pkgver.tar.gz"
LANG=C sed "s|source=.*|source=(\"${SOURCEVAL}\")|" ${SCRIPT_DIR}/PKGBUILD > ${PKGBUILD_STAGING}

docker run --rm -i -v "${VOLUME_DIR}":/home/builder pcmagas/arch-pkg-builder makepkg --printsrcinfo > ${VOLUME_DIR}/.SRCINFO

