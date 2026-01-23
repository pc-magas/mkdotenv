#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../..

cd ${SOURCEPATH}
make bin

cd ${SCRIPTPATH}
chmod +x ${SOURCEPATH}/bin/mkdotenv-linux-amd64
echo
echo "#################################"
echo "Executing"
echo "#################################"

${SOURCEPATH}/bin/mkdotenv-linux-amd64 --output-file=.env