#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
FORMULA_PATH="${SCRIPTPATH}/mkdotenv.rb"

ZIP_PATH="${SCRIPTPATH}/mkdotenv-macos.zip"

echo ${ZIP_PATH}

rm -rf ${ZIP_PATH}
make bin OS=darwin ARCH=arm64 COMPILED_BIN_PATH="/tmp/mkdotenv"

mkdir  -p ${SCRIPTPATH}/bin

cp ./bin/mkdotenv-darwin-arm64  ${SCRIPTPATH}/bin/mkdotenv
zip -j -o ${ZIP_PATH} ${SCRIPTPATH}/bin/mkdotenv

SHA256=$(shasum -a 256 ${ZIP_PATH} | awk '{ print $1 }')

sed -i -E "s|sha256 \".*\"|sha256 \"${SHA256}\"|" "${FORMULA_PATH}"
echo "Updated SHA256 in ${FORMULA_PATH}"


FORMULA_TEST_PATH="${SCRIPTPATH}/bin/mkdotenv.rb"
cp ${FORMULA_PATH} ${FORMULA_TEST_PATH}

ls -l  ${FORMULA_TEST_PATH}

sed -i -E "s|url.*|url \"file://${ZIP_PATH}\"|" ${FORMULA_TEST_PATH}
