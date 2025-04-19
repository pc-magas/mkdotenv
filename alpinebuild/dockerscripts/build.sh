#!/usr/bin/env bash

set -eu

mkdir -p  /home/$USER/packages/

# Append it only once
grep -qxF "PACKAGER_REPODEST=/home/$USER/packages" ~/.abuild/abuild.conf || \
echo "PACKAGER_REPODEST=/home/$USER/packages" >> ~/.abuild/abuild.conf

keyfile=$(grep '^PACKAGER_PRIVKEY=' /home/$USER/.abuild/abuild.conf | cut -d= -f2 | tr -d '"')

echo $keyfile

if ["$keyfile" == ""]; then
   abuild -a
   keyfile=$(grep '^PACKAGER_PRIVKEY=' ~/.abuild/abuild.conf | cut -d= -f2 | tr -d '"')
fi

echo "KEYFILE ${keyfile}"

sudo cp ~/.abuild/*.rsa.pub /etc/apk/keys/

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
sleep 5
abuild -r -K -P "/home/$USER"