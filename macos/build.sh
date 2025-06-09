#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
FORMULA_PATH="${SCRIPTPATH}/mkdotenv.rb"

rm -rf ./mkdotenv-macos.zip
make bin OS=darwin ARCH=arm64 COMPILED_BIN_PATH="/tmp/mkdotenv"

mkdir  -p ${SCRIPTPATH}/bin

cp ./bin/mkdotenv-darwin-arm64  ${SCRIPTPATH}/bin/mkdotenv
zip -j -o mkdotenv-macos.zip ${SCRIPTPATH}/bin/mkdotenv

SHA256=$(shasum -a 256 mkdotenv-macos.zip | awk '{ print $1 }')

if [[ -f "${FORMULA_PATH}" ]]; then
  sed -i -E "s|sha256 \".*\"|sha256 \"${SHA256}\"|" "${FORMULA_PATH}"
  echo "Updated SHA256 in ${FORMULA_PATH}"
else
  echo "Error: Formula file not found at ${FORMULA_PATH}"
  exit 1
fi
