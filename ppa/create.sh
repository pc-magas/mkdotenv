#!/usr/bin/env bash

SCRIPTPATH=$(dirname "$0") 
VERSION=$(grep 'const VERSION' ./mkdotenv/msg/msg.go | sed -E 's/.*"([^"]+)".*/\1/')

SRC_FOLDER=${SCRIPTPATH}/mkdotenv_${VERSION}

TAR_PATH=${SCRIPTPATH}/../mkdotenv_${VERSION}.orig.tar.gz

mkdir -p ${SRC_FOLDER}
cp -r ${SCRIPTPATH}/../mkdotenv ${SRC_FOLDER}
cp -r ${SCRIPTPATH}/../man ${SRC_FOLDER}
cp ${SCRIPTPATH}/../Makefile ${SRC_FOLDER}
cp ${SCRIPTPATH}/../LICENCE ${SRC_FOLDER}
tar --exclude=debian --exclude=alpinebuild -czf ${TAR_PATH} ${SRC_FOLDER}