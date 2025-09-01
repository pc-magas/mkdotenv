#!/usr/bin/env bash

set -e  # Exit on any error

SCRIPTPATH=$(dirname "$0") 
bash ${SCRIPTPATH}/package.sh
VERSION=$(cat ${SCRIPTPATH}/../VERSION)

ls -l ${SCRIPTPATH}/../../mkdotenv_${VERSION}.orig.tar.gz
ls -l ${SCRIPTPATH}/../../mkdotenv_${VERSION}-*.dsc
ls -l ${SCRIPTPATH}/../../mkdotenv_*_source.changes

lintian ${SCRIPTPATH}/../../mkdotenv_${VERSION}-*_source.changes

echo "Testbuild source packages"

DISTROS=("jammy" "noble")
MIRROR="http://archive.ubuntu.com/ubuntu"

PPA_REPO="ppa:pcmagas/mkdotenv"

if [ -f ${SCRIPTPATH}/../PPA_OVERRIDE ]; then
    PPA_REPO=$(cat ${SCRIPTPATH}/../PPA_OVERRIDE)
fi

echo "UPLOAD into ${PPA_REPO}"
sleep 5
rm -rf ${SCRIPTPATH}/../../mkdotenv_*.ppa.upload 
dput ${PPA_REPO} ${SCRIPTPATH}/../../mkdotenv_${VERSION}-0ubuntu1~*1_source.changes
