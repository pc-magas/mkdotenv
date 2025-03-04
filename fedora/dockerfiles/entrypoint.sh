#!/usr/bin/env bash

usermod -u ${UID} pkgbuild
groupmod -g ${GID} pkgbuild

chown -R pkgbuild:pkgbuild /home/pkgbuild/rpmbuild/RPMS

su pkgbuild -c "$*"
