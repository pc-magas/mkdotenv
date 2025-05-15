#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)
PKG_NAME="mkdotenv-${VERSION}-1-x86_64.pkg.tar.zst"
VOLUME_DIR=${SCRIPT_DIR}/volumes

docker run --rm -v "${VOLUME_DIR}/${PKG_NAME}":/root/${PKG_NAME} -w /root -ti archlinux:latest bash -c "pacman -U --noconfirm ./mkdotenv-0.3.0-1-x86_64.pkg.tar.zst && mkdotenv -h"
