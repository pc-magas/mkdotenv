#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

cd $SCRIPTPATH

npm ci
cp webpage/node_modules/@pc-magas/asciiart/dist/js/asciiart.min.js ./public/
cp webpage/node_modules/@pc-magas/asciiart/dist/css/style.css ./public/