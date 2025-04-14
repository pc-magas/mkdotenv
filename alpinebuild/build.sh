#!/usr/bin/env bash

set -eu

mkdir -p  /home/packager/packages//home

# Append it only once
grep -qxF 'PACKAGER_REPODEST=/home/packager/packages' ~/.abuild/abuild.conf || \
echo 'PACKAGER_REPODEST=/home/packager/packages' >> ~/.abuild/abuild.conf

keyfile=$(grep '^PACKAGER_PRIVKEY=' ~/.abuild/abuild.conf | cut -d= -f2 | tr -d '"')

# Make sure it ends with .rsa and the pub file exists
pubkey="${keyfile}.pub"

if [ -f "$pubkey" ]; then
    echo "Copying $pubkey to /etc/apk/keys/"
    sudo cp "$pubkey" /etc/apk/keys/
else
    echo "ERROR: Public key not found: $pubkey"
    exit 1
fi

echo "Cleanup"
abuild clear
abuild clean
sleep 5
echo "Fix Checksum"
abuild checksum
echo "Installing Dependencies"
sleep 5
abuild deps
echo "Final Build"
sleep 10
abuild -r -K -P "/home/$USER"