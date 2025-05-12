#!/usr/bin/env bash

docker build -t pcmagas/arch-pkg-builder .

docker run --rm -v "$PWD":/build -w /build pcmagas/arch-pkg-builder makepkg -si --noconfirm
docker run --rm -i -v "$PWD":/build -w /build pcmagas/arch-pkg-builder makepkg --printsrcinfo > .SRCINFO