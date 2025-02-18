#!/usr/bin/env bash

VERSION=$(grep 'const VERSION' ./src/mkdotenv.go | sed -E 's/.*"([^"]+)".*/\1/')
mkdir -p mkdotenv-$VERSION
cp -r src mkdotenv-${VERSION}
cp -r man mkdotenv-${VERSION}
cp Makefile mkdotenv-${VERSION}
cp LICENCE mkdotenv-${VERSION}
cp go.mod mkdotenv-${VERSION}
tar --exclude=debian --exclude=alpinebuild -czf ./rpmbuild/SOURCES/mkdotenv-${VERSION}.tar.gz mkdotenv-${VERSION};

docker build -f ./dockerfiles/DockerfileFedora -t pcmagas/gopkgbuild:fedora-41 .
docker run -v "$(pwd)/rpmbuild:/home/pkgbuild/rpmbuild" pcmagas/gopkgbuild:fedora-41 bash -c "rpmbuild -bb ~/rpmbuild/SPECS/mkdotenv.spec"
