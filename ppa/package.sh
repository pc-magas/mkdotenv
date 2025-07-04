#!/usr/bin/env bash

SCRIPTPATH=$(dirname "$0") 
VERSION=$(cat ${SCRIPTPATH}/../VERSION)
DISTROS=("jammy" "noble")

SRC_FOLDER=mkdotenv_${VERSION}
CHANGES_FILE=${SCRIPTPATH}/../../mkdotenv_*_source.changes

echo "VERSION: ${VERSION}"

# PPA distro config
LINUX_DIST="ubuntu"
DIST=jammy

bash ${SCRIPTPATH}/make_tar.sh

cd ${SCRIPTPATH}/..

for distro in "${DISTROS[@]}"; do
    echo "Create source package for: "${distro}

    sed -i "s/unstable/${distro}/g" ${SCRIPTPATH}/../debian/changelog
	sed -i 's/debian/ubuntu/g' ${SCRIPTPATH}/../debian/changelog
	dpkg-buildpackage -S -sa
	sed -i "s/${distro}/unstable/g" ${SCRIPTPATH}/debian/changelog
	sed -i 's/ubuntu/debian/g' ${SCRIPTPATH}/debian/changelog
done