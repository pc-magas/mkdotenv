#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

cd $SCRIPTPATH

bash ./build.sh

npx @11ty/eleventy --serve