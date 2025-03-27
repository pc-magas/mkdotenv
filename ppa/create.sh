#!/usr/bin/env bash

# PAth config

SCRIPTPATH=$(dirname "$0") 
VERSION=$(grep 'const VERSION' ./mkdotenv/msg/msg.go | sed -E 's/.*"([^"]+)".*/\1/')
SRC_FOLDER=${SCRIPTPATH}/mkdotenv_${VERSION}
TAR_PATH=${SCRIPTPATH}/../../mkdotenv_${VERSION}.orig.tar.gz
CHANGES_FILE=${SCRIPTPATH}/../../mkdotenv_*_source.changes


# PPA distro config
LINUX_DIST="ubuntu"
DISTROS=("jammy" "noble")
DIST=jammy


mkdir -p ${SRC_FOLDER}
cp -r ${SCRIPTPATH}/../mkdotenv ${SRC_FOLDER}
cp -r ${SCRIPTPATH}/../man ${SRC_FOLDER}
cp ${SCRIPTPATH}/../Makefile ${SRC_FOLDER}
cp ${SCRIPTPATH}/../LICENCE ${SRC_FOLDER}
tar --exclude=debian --exclude=alpinebuild -czf ${TAR_PATH} ${SRC_FOLDER}

echo "Generated tar name ${TAR_PATH}"

if [ ! -f ${TAR_PATH} ]; then
	echo "Tarball does not exist"
fi

for distro in "${DISTROS[@]}"; do
    echo "Create source package for: "${distro}

    sed -i "s/unstable/${distro}/g" debian/changelog
	sed -i 's/debian/ubuntu/g' debian/changelog
	dpkg-buildpackage -S -sa
	sed -i "s/${distro}/unstable/g" debian/changelog
	sed -i 's/ubuntu/debian/g' debian/changelog
done

dput ppa:pcmagas/mkdotenv ${CHANGES_FILE}

