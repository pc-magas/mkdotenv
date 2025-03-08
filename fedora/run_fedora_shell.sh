#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../ 

VERSION=$(cat ${SOURCEPATH}/VERSION)


docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "$(pwd)/rpmbuild/RPMS/x86_64/:/root/rpmbuild/RPMS/x86_64" \
    -ti -u root \
    fedora:41 bash
