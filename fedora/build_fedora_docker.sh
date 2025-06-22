#!/usr/bin/env bash


SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/ 

VERSION=$(cat ${SOURCEPATH}/../VERSION)

SRC_FOLDER=mkdotenv-${VERSION}
RPM_SRC=${SCRIPTPATH}/rpmbuild/SOURCES

GENERATED_TAR=$(bash ${SCRIPTPATH}/make_tar.sh)

tar tzf ${GENERATED_TAR} | head -n 1

mkdir -p ${SOURCEPATH}rpmbuild/RPMS/x86_64

docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "${SOURCEPATH}/rpmbuild/SOURCES:/home/pkgbuild/rpmbuild/SOURCES" \
    -v "${SOURCEPATH}/mkdotenv.spec:/home/pkgbuild/rpmbuild/SPECS/mkdotenv.spec" \
    -v "${SOURCEPATH}/rpmbuild/RPMS/x86_64:/home/pkgbuild/rpmbuild/RPMS/x86_64" \
    ghcr.io/pc-magas/fedora_rpm_build_docker rpmbuild -bb /home/pkgbuild/rpmbuild/SPECS/mkdotenv.spec

echo "Cleanup"
rm -rf ${SRC_FOLDER}
