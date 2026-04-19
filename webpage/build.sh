#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

cd $SCRIPTPATH

npm ci
cp ./node_modules/@pc-magas/asciiart/dist/js/asciiart.min.js ../docs/
cp ./node_modules/@pc-magas/asciiart/dist/css/style.css ../docs/

npx @11ty/eleventy