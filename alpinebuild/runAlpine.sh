#!/usr/bin/env sh

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat VERSION)

# Path that APKBUILD is Overriden 
OVERLAY=${SCRIPT_DIR}/apkbuild-overlay

# TARBALL name and path  
TARGZ_NAME=mkdotenv-$(cat VERSION).tar.gz
TARGZ=${OVERLAY}/${TARGZ_NAME}

APKBUILD_OVERLAY=${OVERLAY}/APKBUILD

SOURCE_FOLDER=${SCRIPT_DIR}/mkdotenv-$(cat VERSION)

rm -rf ${OVERLAY}
cp -rT ${SCRIPT_DIR}/../ ${SOURCE_FOLDER}/

mkdir -p ${OVERLAY}

tar -czf ${TARGZ} ${SOURCE_FOLDER}

cp ${SCRIPT_DIR}/APKBUILD ${APKBUILD_OVERLAY}


SUBSTITUTE="s|source=".*"|source=\"${TARGZ_NAME}\"|"
sed -i "s|source=".*"|source=\"${TARGZ_NAME}\"|" ${APKBUILD_OVERLAY}
sed -i 's|version=".*"|source="${VERSION}"|' ${APKBUILD_OVERLAY}


docker build -f ${SCRIPT_DIR}/Dockerfile -t pcmagas/alpinebuild ${SCRIPT_DIR}
docker run  \
    -v ${OVERLAY}:/home/packager/src \
    -ti -u root pcmagas/alpinebuild bash -c "chown -R packager:packager /home/packager/* && bash"