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
SOURCE_FOLDER=${SCRIPT_DIR}/mkdotenv-${VERSION}

rm -rf ${SOURCE_FOLDER}
mkdir -p ${SOURCE_FOLDER}

mkdir ${SOURCE_FOLDER}/mkdotenv
cp -r ${SCRIPT_DIR}/../mkdotenv/* ${SOURCE_FOLDER}/mkdotenv
cp -r ${SCRIPT_DIR}/../man ${SOURCE_FOLDER}/man
cp ${SCRIPT_DIR}/../Makefile ${SOURCE_FOLDER}/Makefile
cp ${SCRIPT_DIR}/../LICENCE ${SOURCE_FOLDER}/LICENCE

echo ${SOURCE_FOLDER}
ls -l
rm -rf ${OVERLAY}
mkdir -p ${OVERLAY}

(
  cd ${SCRIPT_DIR} && tar -czf ${TARGZ} -C $(basename ${SOURCE_FOLDER}) .
)

cp ${SCRIPT_DIR}/APKBUILD-template ${APKBUILD_OVERLAY}
cp ${SCRIPT_DIR}/build.sh ${OVERLAY}/

tar -tzf ${TARGZ}

sed -i '/^source="\$pkgname-\$pkgver.tar.gz::https:\/\/github.com\/pc-magas\/mkdotenv\/archive\/refs\/tags\/v\$pkgver.tar.gz"/d' ${APKBUILD_OVERLAY}


docker build -f ${SCRIPT_DIR}/Dockerfile -t pcmagas/alpinebuild ${SCRIPT_DIR}
docker run  \
    -v ${OVERLAY}:/home/packager \
    -v ${SCRIPT_DIR}/.build:/home/packager/.abuild \
    -ti -u root pcmagas/alpinebuild bash -c "chown -R packager:packager /home/packager/* && bash"