#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)

VOLUME_DIR=${SCRIPT_DIR}/volumes

rm -rf ${VOLUME_DIR}
mkdir -p ${VOLUME_DIR}

echo ${VERSION}

TARGZ_NAME=mkdotenv-${VERSION}.tar.gz
TARGZ=${VOLUME_DIR}/${TARGZ_NAME}
PKGBUILD_LOCAL=${VOLUME_DIR}/PKGBUILD

ORIG_TAR=$(bash ${SCRIPT_DIR}/../alpinebuild/make_tar.sh)

cp ${SCRIPT_DIR}/../alpinebuild/${ORIG_TAR} ${TARGZ}

ls -l ${VOLUME_DIR}

LANG=C sed "s/source=.*/source=(\"${TARGZ_NAME}\")/" ${SCRIPT_DIR}/PKGBUILD > ${PKGBUILD_LOCAL}

docker build -f ${SCRIPT_DIR}/Dockerfile -t pcmagas/arch-pkg-builder ${SCRIPT_DIR}

docker run --rm -v "${VOLUME_DIR}":/build -w /build pcmagas/arch-pkg-builder build_n_run

cp ${PKGBUILD_LOCAL} ${SCRIPT_DIR}/PKGBUILD


