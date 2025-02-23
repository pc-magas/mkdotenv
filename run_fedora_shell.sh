#!/usr/bin/env bash

VERSION=$(grep 'const VERSION' ./src/mkdotenv.go | sed -E 's/.*"([^"]+)".*/\1/')


docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "$(pwd)/rpmbuild/RPMS/x86_64/:/root/rpmbuild/RPMS/x86_64" \
    -ti -u root \
    fedora:41 bash
