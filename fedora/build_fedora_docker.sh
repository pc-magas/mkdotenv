#!/usr/bin/env bash

SCRIPTPATH=$(dirname "$0")
SOURCEPATH=${SCRIPTPATH}/../ 

VERSION=$(grep 'const VERSION' ./mkdotenv/msg/msg.go | sed -E 's/.*"([^"]+)".*/\1/')

SRC_FOLDER=${SCRIPTPATH}/mkdotenv-${VERSION}
CHANGES_FILE=${SCRIPTPATH}/../../mkdotenv_*_source.changes


mkdir -p ${SOURCEPATH}rpmbuild/SOURCES
mkdir -p ${SOURCEPATH}rpmbuild/RPMS/x86_64

ls -l ${SOURCEPATH}mkdotenv

mkdir -p ${SRC_FOLDER}
cp -r ${SCRIPTPATH}/../mkdotenv ${SRC_FOLDER}/
cp -r ${SCRIPTPATH}/../man ${SRC_FOLDER}
cp ${SCRIPTPATH}/../Makefile ${SRC_FOLDER}
cp ${SCRIPTPATH}/../LICENCE ${SRC_FOLDER}
tar --exclude=debian --exclude=alpinebuild -czf ${SOURCEPATH}/rpmbuild/SOURCES/mkdotenv-${VERSION}.tar.gz ${SRC_FOLDER}

docker build -f ${SCRIPTPATH}/dockerfiles/DockerfileFedora -t pcmagas/gopkgbuild:fedora-41 ${SCRIPTPATH}
docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "$(pwd)/rpmbuild/SOURCES:/home/pkgbuild/rpmbuild/SOURCES" \
    -v "$(pwd)/rpmbuild/SPECS:/home/pkgbuild/rpmbuild/SPECS" \
    -v "$(pwd)/rpmbuild/RPMS/x86_64:/home/pkgbuild/rpmbuild/RPMS/x86_64" \
    pcmagas/gopkgbuild:fedora-41 rpmbuild -bb /home/pkgbuild/rpmbuild/SPECS/mkdotenv.spec

# echo "Cleanup"
# rm -rf ${SRC_FOLDER}