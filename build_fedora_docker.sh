#!/usr/bin/env bash

VERSION=$(grep 'const VERSION' ./src/mkdotenv.go | sed -E 's/.*"([^"]+)".*/\1/')
mkdir -p ./rpmbuild/SOURCES
mkdir -p ./rpmbuild/RPMS/x86_64

mkdir -p mkdotenv-$VERSION
cp -r src mkdotenv-${VERSION}
cp -r man mkdotenv-${VERSION}
cp Makefile mkdotenv-${VERSION}
cp LICENCE mkdotenv-${VERSION}
cp go.mod mkdotenv-${VERSION}
tar --exclude=debian --exclude=alpinebuild -czf ./rpmbuild/SOURCES/mkdotenv-${VERSION}.tar.gz mkdotenv-${VERSION};

ls -l ./rpmbuild/SOURCES

docker build -f ./dockerfiles/DockerfileFedora -t pcmagas/gopkgbuild:fedora-41 .
docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "$(pwd)/rpmbuild/SOURCES:/home/pkgbuild/rpmbuild/SOURCES" \
    -v "$(pwd)/rpmbuild/SPECS:/home/pkgbuild/rpmbuild/SPECS" \
    -v "$(pwd)/rpmbuild/RPMS/x86_64:/home/pkgbuild/rpmbuild/RPMS/x86_64" \
    pcmagas/gopkgbuild:fedora-41 rpmbuild -bb /home/pkgbuild/rpmbuild/SPECS/mkdotenv.spec
