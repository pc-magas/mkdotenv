#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)
TARGZ_NAME=mkdotenv-${VERSION}.tar.gz

SOURCE_FOLDER=${SCRIPT_DIR}/mkdotenv-${VERSION}

rm -rf ${SOURCE_FOLDER}
mkdir -p ${SOURCE_FOLDER}

mkdir ${SOURCE_FOLDER}/mkdotenv
cp -r ${SCRIPT_DIR}/../mkdotenv/* ${SOURCE_FOLDER}/mkdotenv
cp -r ${SCRIPT_DIR}/../man ${SOURCE_FOLDER}/man
cp ${SCRIPT_DIR}/../Makefile ${SOURCE_FOLDER}/Makefile
cp ${SCRIPT_DIR}/../LICENCE ${SOURCE_FOLDER}/LICENCE

(
  cd ${SCRIPT_DIR} && tar -czf ${TARGZ_NAME} -C $(basename ${SOURCE_FOLDER}) .
)

date > TAR_FILES
echo ${TARGZ_NAME} > TAR_FILES
tar -tzf ${SCRIPT_DIR}/${TARGZ_NAME} >> TAR_FILES

echo ${TARGZ_NAME}
