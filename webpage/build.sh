#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

cd ${SCRIPTPATH}/..

make man

cd $SCRIPTPATH

npm ci
cp ./node_modules/@pc-magas/asciiart/dist/js/asciiart.min.js ./page
cp ./node_modules/@pc-magas/asciiart/dist/css/style.css ./page

npx @11ty/eleventy