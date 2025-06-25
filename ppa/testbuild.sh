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

for distro in "${DISTROS[@]}"; do
    echo "Building for ${distro}"
    dput ppa:pcmagas/mkdotenv-test2 ${SCRIPTPATH}/../../mkdotenv_${VERSION}-0ubuntu1~${distro}1_source.changes
    sleep 10
done
