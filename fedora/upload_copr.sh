#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../ 

# docker run \
#     -e UID=$(id -u) -e GID=$(id -g)\
#     -v "${SOURCEPATH}/rpmbuild/SRPMS:/home/pkgbuild/rpmbuild/SRPMS" \
#     -v "${SOURCEPATH}/.config:/home/pkgbuild/.config" \
#     ghcr.io/pc-magas/fedora_rpm_build_docker copr-cli build mkdotenv /home/pkgbuild/rpmbuild/SRPMS/mkdotenv*.src.rpm

docker run \
    -e UID=$(id -u) -e GID=$(id -g)\
    -v "${SOURCEPATH}/rpmbuild/SRPMS:/home/pkgbuild/rpmbuild/SRPMS" \
    -v "${SOURCEPATH}/.config/copr:/home/pkgbuild/.config/copr:ro" \
    -ti \
    ghcr.io/pc-magas/fedora_rpm_build_docker /bin/bash
