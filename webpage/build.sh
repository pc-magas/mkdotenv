#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

cd $SCRIPTPATH

npm ci
cp ./node_modules/@pc-magas/asciiart/dist/js/asciiart.min.js ./public/
cp ./node_modules/@pc-magas/asciiart/dist/css/style.css ./public/

npx @11ty/eleventy