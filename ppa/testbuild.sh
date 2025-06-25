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
    CHROOT_NAME="${distro}-amd64"
    CHROOT_PATH="/srv/chroot/${CHROOT_NAME}"

    echo "Checking if sbuild chroot exists for ${distro}..."

    if ! schroot -l | grep -q "$CHROOT_NAME"; then
        echo "Chroot for ${distro} not found. Creating..."
        sudo sbuild-createchroot ${distro} ${CHROOT_PATH} ${MIRROR}
    else
        echo "Chroot for ${distro} already exists."
    fi

    echo "Building for ${distro}"
    sudo sbuild -d ${distro} ${SCRIPTPATH}/../../mkdotenv_${VERSION}-${distro}1.dsc
    sleep 10
done
