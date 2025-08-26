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

rm -rf ${SCRIPTPATH}/../../mkdotenv_*.ppa.upload 
dput ppa:pcmagas/ppa-test5 ${SCRIPTPATH}/../../mkdotenv_${VERSION}-0ubuntu1~*1_source.changes
