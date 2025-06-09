#!/ust/bin/env bash


# THIS SCRIPT IS DESIGNED TO RUN UPON MACOS

make bin OS=darwin ARCH=arm64 COMPILED_BIN_PATH="/tmp/mkdotenv"
zip mkdotenv-macos.zip mkdotenv

