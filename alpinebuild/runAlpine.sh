#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)

echo ${VERSION}
# Path that APKBUILD is Overriden 
OVERLAY=${SCRIPT_DIR}/apkbuild-overlay

# TARBALL name and path  
TARGZ_NAME=mkdotenv-${VERSION}.tar.gz
TARGZ=${OVERLAY}/${TARGZ_NAME}

sed -i "s|pkgver=".*"|pkgver="${VERSION}"|" ${SCRIPT_DIR}/APKBUILD-template

APKBUILD_OVERLAY=${OVERLAY}/APKBUILD

ORIG_TAR=$(bash ${SCRIPT_DIR}/make_tar.sh)

cp ${ORIG_TAR} ${TARGZ}

cp ${SCRIPT_DIR}/APKBUILD-template ${APKBUILD_OVERLAY}
cp ${SCRIPT_DIR}/build.sh ${OVERLAY}/

tar -tzf ${TARGZ}

sed -i '/^source="\$pkgname-\$pkgver.tar.gz::https:\/\/github.com\/pc-magas\/mkdotenv\/archive\/refs\/tags\/v\$pkgver.tar.gz"/d' ${APKBUILD_OVERLAY}


docker build -f ${SCRIPT_DIR}/Dockerfile -t pcmagas/alpinebuild ${SCRIPT_DIR}
docker run  -v ${OVERLAY}:/home/packager -ti pcmagas/alpinebuild bash