#!/usr/bin/env bash

SCRIPTPATH=$(realpath "$(dirname "$0")")

VERSION=$(cat ${SCRIPTPATH}/../VERSION)

TAR_PATH=${SCRIPTPATH}/../../mkdotenv_${VERSION}.orig.tar.gz
CHANGES_FILE=${SCRIPTPATH}/../../mkdotenv_*_source.changes
SRC_FOLDER=mkdotenv_${VERSION}

echo "VERSION: ${VERSION}"

echo "Vendoring Go dependencies..."
(
  cd "${SCRIPTPATH}/../mkdotenv" || exit 1
  go clean -modcache
  go mod tidy
  go mod vendor
  go mod verify
)


cd ${SCRIPTPATH}
mkdir -p ${SRC_FOLDER}
cp -r ${SCRIPTPATH}/../mkdotenv ${SRC_FOLDER}
cp -r ${SCRIPTPATH}/../man ${SRC_FOLDER}
cp ${SCRIPTPATH}/../Makefile ${SRC_FOLDER}
cp ${SCRIPTPATH}/../LICENCE ${SRC_FOLDER}
cp ${SCRIPTPATH}/../VERSION ${SRC_FOLDER}
tar --exclude=debian --exclude=alpinebuild -czf ${TAR_PATH} ${SRC_FOLDER}

if [ ! -f ${TAR_PATH} ]; then
	echo "Tarball does not exist" >&2
fi

tar tzf ${TAR_PATH} >&1

echo ${TAR_PATH}
