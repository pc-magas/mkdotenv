#!/usr/bin/env bash
set -e

SCRIPTPATH=$(dirname "$0")
UBUNTU_VERSION="jammy"
PPA_REPO="ppa:pcmagas/mkdotenv"

if [ -f "${SCRIPTPATH}/../PPA_OVERRIDE" ]; then
    PPA_REPO=$(cat "${SCRIPTPATH}/../PPA_OVERRIDE")
fi

COMMAND="mkdotenv -h"

# Run everything inside a temporary Docker container
docker run --rm -it ubuntu:$UBUNTU_VERSION bash -c "
    set -e
    echo '[*] Updating package lists...'
    apt-get update

    echo '[*] Installing software-properties-common...'
    DEBIAN_FRONTEND=noninteractive apt-get install -y software-properties-common

    echo '[*] Adding PPA: $PPA_REPO...'
    add-apt-repository -y $PPA_REPO

    echo '[*] Updating after adding PPA...'
    apt-get update

    echo '[*] Installing mkdotenv from PPA...'
    DEBIAN_FRONTEND=noninteractive apt-get install -y mkdotenv

    echo '[*] Running command: $COMMAND'
    $COMMAND || echo 'Command failed or not installed'
"
