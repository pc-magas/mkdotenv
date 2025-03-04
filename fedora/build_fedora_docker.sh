#!/usr/bin/env bash


SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../ 

VERSION=$(grep 'const VERSION' ./mkdotenv/msg/msg.go | sed -E 's/.*"([^"]+)".*/\1/')

SRC_FOLDER=mkdotenv-${VERSION}
CHANGES_FILE=${SCRIPTPATH}/../../mkdotenv_*_source.changes


mkdir -p ${SOURCEPATH}rpmbuild/SOURCES
mkdir -p ${SOURCEPATH}rpmbuild/RPMS/x86_64

cd ${SCRIPTPATH}
ls -l
mkdir -p ${SRC_FOLDER}
cp -r ../mkdotenv ${SRC_FOLDER}/mkdotenv
cp -r ../man ${SRC_FOLDER}/man
cp ../Makefile ${SRC_FOLDER}/Makefile
cp ../LICENCE ${SRC_FOLDER}/LICENCE

tar -czf ${SOURCEPATH}/rpmbuild/SOURCES/mkdotenv-${VERSION}.tar.gz ${SRC_FOLDER}

docker build -f ${SCRIPTPATH}/dockerfiles/DockerfileFedora -t pcmagas/gopkgbuild:fedora-41 ${SCRIPTPATH}

docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "${SOURCEPATH}/rpmbuild/SOURCES:/home/pkgbuild/rpmbuild/SOURCES" \
    -v "${SOURCEPATH}/rpmbuild/SPECS:/home/pkgbuild/rpmbuild/SPECS" \
    -v "${SOURCEPATH}/rpmbuild/RPMS/x86_64:/home/pkgbuild/rpmbuild/RPMS/x86_64" \
    pcmagas/gopkgbuild:fedora-41 rpmbuild -bb /home/pkgbuild/rpmbuild/SPECS/mkdotenv.spec

# echo "Cleanup"
# rm -rf ${SRC_FOLDER}