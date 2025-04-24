#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=$(cat ${SCRIPT_DIR}/../VERSION)
APKFILE=mkdotenv-${VERSION}-r0.apk
APKFILE_FULL=${SCRIPT_DIR}/release/home/x86_64/${APKFILE}

if [ -z ${APKFILE_FULL} ]; then
    bash ${SCRIPT_DIR}/make_apk.sh
fi

docker run \
  -v ${APKFILE_FULL}:/root/${APKFILE} \
  -w /root \
  alpine \
  sh -c "apk add --allow-untrusted ${APKFILE} && mkdotenv -h"
