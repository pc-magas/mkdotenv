#!/usr/bin/env bash

# docker build -t arch-pkg-builder .

docker run --rm -it -v "$PWD":/build -w /build arch-pkg-builder