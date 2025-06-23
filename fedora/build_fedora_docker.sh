#!/usr/bin/env bash


SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/ 

VERSION=$(cat ${SOURCEPATH}/../VERSION)

RPM_SRC=${SCRIPTPATH}/rpmbuild/SOURCES

rm -rf ${RPM_SRC}
mkdir -p ${RPM_SRC}

GENERATED_TAR=$(bash ${SCRIPTPATH}/make_tar.sh)

echo "TAR COntents"
tar tzf ${GENERATED_TAR} | head -n 1

echo "Recreating RPM Storage folder"
rm -rf ${SOURCEPATH}rpmbuild/RPMS/x86_64
mkdir -p ${SOURCEPATH}rpmbuild/RPMS/x86_64

echo "Recreate SRPM Storage folder"
rm -rf ${SOURCEPATH}rpmbuild/SRPMS
mkdir -p ${SOURCEPATH}rpmbuild/SRPMS

echo "Generating SRPM"

docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "${SOURCEPATH}/rpmbuild/SOURCES:/home/pkgbuild/rpmbuild/SOURCES" \
    -v "${SOURCEPATH}/mkdotenv.spec:/home/pkgbuild/rpmbuild/SPECS/mkdotenv.spec" \
    -v "${SOURCEPATH}/rpmbuild/RPMS/x86_64:/home/pkgbuild/rpmbuild/RPMS/x86_64" \
    -v "${SOURCEPATH}/rpmbuild/SRPMS:/home/pkgbuild/rpmbuild/SRPMS" \
    ghcr.io/pc-magas/fedora_rpm_build_docker rpmbuild -bs /home/pkgbuild/rpmbuild/SPECS/mkdotenv.spec

echo "Build SRPM"

docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "${SOURCEPATH}/rpmbuild/RPMS/x86_64:/home/pkgbuild/rpmbuild/RPMS/x86_64" \
    -v "${SOURCEPATH}/rpmbuild/SRPMS:/home/pkgbuild/rpmbuild/SRPMS" \
    ghcr.io/pc-magas/fedora_rpm_build_docker rpmbuild --rebuild /home/pkgbuild/rpmbuild/SRPMS/mkdotenv-${VERSION}-2.fc41.src.rpm

