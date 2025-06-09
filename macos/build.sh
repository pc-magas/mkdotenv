#!/ust/bin/env bash

rm -rf ./mkdotenv-macos.zip
make compile OS=darwin ARCH=arm64 COMPILED_BIN_PATH="/tmp/mkdotenv"
zip -j -o mkdotenv-macos.zip /tmp/mkdotenv
