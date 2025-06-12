#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
FORMULA_PATH="${SCRIPTPATH}/mkdotenv.rb"

ZIP_PATH="${SCRIPTPATH}/mkdotenv-macos.zip"

echo ${ZIP_PATH}

rm -rf ${ZIP_PATH}

OS_TYPE="$(uname)"
if [[ "$OS_TYPE" == "Darwin" ]]; then
  echo "Building Natively on macOS"
  make bin COMPILED_BIN_PATH="/tmp/mkdotenv"
else
  echo "Buolding for MacOs from ${OS_TYPE}"
  make bin OS=darwin ARCH=arm64 COMPILED_BIN_PATH="/tmp/mkdotenv"
fi

mkdir  -p ${SCRIPTPATH}/bin

cp ./bin/mkdotenv-darwin-arm64  ${SCRIPTPATH}/bin/mkdotenv
zip -j -o ${ZIP_PATH} ${SCRIPTPATH}/bin/mkdotenv

SHA256=$(shasum -a 256 ${ZIP_PATH} | awk '{ print $1 }')

echo "Fixing Checksum upon ${FORMULA_PATH}"
sed -i -E "s|sha256 \".*\"|sha256 \"${SHA256}\"|" "${FORMULA_PATH}"
echo "Updated SHA256 in ${FORMULA_PATH}"


if [[ "$OS_TYPE" == "Darwin" ]]; then

  echo "[MAC] Modifying Formula for local installation"

  FORMULA_TEST_PATH="${SCRIPTPATH}/bin/mkdotenv.rb"
  cp ${FORMULA_PATH} ${FORMULA_TEST_PATH}
  ls -l  ${FORMULA_TEST_PATH}
  sed -i -E "s|url.*|url \"file://${ZIP_PATH}\"|" ${FORMULA_TEST_PATH}

fi