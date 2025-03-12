#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)

echo ${VERSION}
# Path that APKBUILD is Overriden 
OVERLAY=${SCRIPT_DIR}/apkbuild-overlay

# TARBALL name and path  
TARGZ_NAME=mkdotenv-${VERSION}.tar.gz
TARGZ=${OVERLAY}/${TARGZ_NAME}

APKBUILD_OVERLAY=${OVERLAY}/APKBUILD

SOURCE_FOLDER=${SCRIPT_DIR}/mkdotenv-${VERSION}

rm -rf ${SOURCE_FOLDER}
mkdir -p ${SOURCE_FOLDER}

cp -r ${SCRIPT_DIR}/../mkdotenv ${SOURCE_FOLDER}/
cp -r ${SCRIPT_DIR}/../man ${SOURCE_FOLDER}/man
cp ${SCRIPT_DIR}/../Makefile ${SOURCE_FOLDER}/Makefile
cp ${SCRIPT_DIR}/../LICENCE ${SOURCE_FOLDER}/LICENCE


rm -rf ${OVERLAY}
mkdir -p ${OVERLAY}

tar -czf ${TARGZ} ${SOURCE_FOLDER}/*
cp ${SCRIPT_DIR}/APKBUILD ${APKBUILD_OVERLAY}


sed -i "s|source=".*"|source=\"mkdotenv-\$\{pkgver\}\"|" ${APKBUILD_OVERLAY}
sed -i "s|pkgver=".*"|pkgver="${VERSION}"|" ${APKBUILD_OVERLAY}


docker build -f ${SCRIPT_DIR}/Dockerfile -t pcmagas/alpinebuild ${SCRIPT_DIR}
docker run  \
    -v ${OVERLAY}:/home/packager \
    -v ${SCRIPT_DIR}/.build:/home/packager/.abuild \
    -ti -u root pcmagas/alpinebuild bash -c "chown -R packager:packager /home/packager/* && bash"