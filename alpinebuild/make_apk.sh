#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)

VOLUME_DIR=${SCRIPT_DIR}/volumes

sudo rm -rf ${VOLUME_DIR}
mkdir -p ${VOLUME_DIR}

echo ${VERSION}
# Path that APKBUILD is Overriden 
OVERLAY=${VOLUME_DIR}/apkbuild-overlay
ABUILD_VOLUME=${VOLUME_DIR}/abuild
RELEASE_DIR=${SCRIPT_DIR}/release

mkdir -p ${OVERLAY}
mkdir -p ${ABUILD_VOLUME}

# Release dir may contain unwanted structure therefore it is re-created
rm -rf ${RELEASE_DIR}
mkdir -p ${RELEASE_DIR}

# TARBALL name and path  
TARGZ_NAME=mkdotenv-${VERSION}.tar.gz
TARGZ=${OVERLAY}/${TARGZ_NAME}

ORIG_TAR=$(bash ${SCRIPT_DIR}/make_tar.sh)
echo "GENERATED TAR ${ORIG_TAR}"
cp ${ORIG_TAR} ${TARGZ}

echo "Generate APKBUILD"
echo ${SCRIPT_DIR}
echo ${TARGZ}
CHECKSUM=$(sha512sum ${TARGZ} | awk '{print $1}')""
bash ${SCRIPT_DIR}/make_apkbuild.sh ${SCRIPT_DIR} --src_local --checksum "${CHECKSUM}"

cp ${SCRIPT_DIR}/APKBUILD ${OVERLAY}/

echo "TAR contents"
tar -tzf ${TARGZ}
cp ${TARGZ} ${RELEASE_DIR}/

docker run \
    -v ${OVERLAY}:/usr/src/apkbuild  \
    -v ${ABUILD_VOLUME}:/home/packager/.abuild \
    -v ${VOLUME_DIR}/keys:/etc/apk/keys \
    -v ${RELEASE_DIR}:/home/packager/release \
    ghcr.io/pc-magas/alpinebuild build --no-checksum

echo "Releasing source file"
cp ${TARGZ} ${RELEASE_DIR}/
echo "Tar.gz released upon ${RELEASE_DIR}"

echo "Fixing APKBUILD for remote"
cp ${RELEASE_DIR}/APKBUILD ${RELEASE_DIR}/APKBUILD.local
rm -rf ${RELEASE_DIR}/APKBUILD
bash ${SCRIPT_DIR}/make_apkbuild.sh ${RELEASE_DIR} --checksum "${CHECKSUM}"

echo "RELEASED FILES"

ls -l ${RELEASE_DIR}