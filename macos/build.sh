#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

rm -rf ./mkdotenv-macos.zip
make bin OS=darwin ARCH=arm64 COMPILED_BIN_PATH="/tmp/mkdotenv"

mkdir ${SCRIPTPATH}/bin
cp ./bin/mkdotenv-darwin-arm64  ${SCRIPTPATH}/bin/mkdotenv
zip -j -o mkdotenv-macos.zip ${SCRIPTPATH}/bin/mkdotenv
