#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)
sed -i "s|pkgver=".*"|pkgver="${VERSION}"|" ${SCRIPT_DIR}/APKBUILD-template

cp ${SCRIPT_DIR}/APKBUILD-template ${SCRIPT_DIR}/APKFILE
sed -i '/^source=\"$pkgname-$pkgver.tar.gz\"/d' ${SCRIPT_DIR}/APKFILE
