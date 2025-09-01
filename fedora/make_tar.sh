#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

VERSION=$(cat ${SCRIPTPATH}/../VERSION)

RPM_SRC=${SCRIPTPATH}/rpmbuild/SOURCES

SRC_FOLDER=mkdotenv-${VERSION}
TAR_NAME=${SRC_FOLDER}-rpm.tar.gz
SRC_DEST=/tmp/${SRC_FOLDER}

cd ${SCRIPTPATH}

echo "Vendoring Go dependencies..."
(
  cd "${SCRIPTPATH}/../mkdotenv" || exit 1
  go clean -modcache
  go mod tidy
  go mod vendor
  go mod verify
)

mkdir -p ${RPM_SRC}
mkdir -p ${SRC_DEST}

cp -r ../mkdotenv ${SRC_DEST}/mkdotenv
cp -r ../man ${SRC_DEST}/man
cp ../Makefile ${SRC_DEST}/Makefile
cp ../LICENCE ${SRC_DEST}/LICENCE

FINAL_TAR_DEST=${RPM_SRC}/${TAR_NAME}
echo "TAR PATH: ${FINAL_TAR_DEST}" >&2
echo "SRC FOLDER: ${SRC_DEST}" >&2

ls -l ${SRC_DEST} >&2

rm -rf ${FINAL_TAR_DEST}
cd /tmp
tar -czf ${FINAL_TAR_DEST} ${SRC_FOLDER}

echo ${FINAL_TAR_DEST}
