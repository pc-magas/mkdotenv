#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
VERSION=$(cat  ${SCRIPTPATH}/../VERSION)
PKGPATH=${SCRIPTPATH}/pkg

cd ${SCRIPTPATH}/..

if [ -f keyid ]; then
    echo "Export Keyid from file"
    export DEB_SIGN_KEYID=$(cat keyid)
fi


mkdir -p ${PKGPATH}
rm -rv ${PKGPATH}/*

dpkg-buildpackage -b
mv ../*.deb ${PKGPATH}/mkdotenv.deb

rm -rf mkdotenv_${VERSION}
rm -rf ../*.changes
rm -rf ../*.buildinfo
rm -rf ../*.dsc	
rm -rf ../*.orig.tar.gz
rm -rf ../*.debian.tar.xz