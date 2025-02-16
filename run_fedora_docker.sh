#!/usr/bin/env bash

docker build -f ./dockerfiles/DockerfileFedora -t pcmagas/gopkgbuild:fedora-41 .
docker run -v .:/home/pkgbuild/code -v ./rpmbuild:/home/pkgbuild/rpmbuild -ti pcmagas/gopkgbuild:fedora-41 /bin/bash