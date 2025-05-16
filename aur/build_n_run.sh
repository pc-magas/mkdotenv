#!/usr/bin/env bash

# Script intented to run *INSIDE* a docker container ti builds and runs mkdotenv inside an arch linux container
set -euo pipefail

sudo chown -R $(whoami) /build

updpkgsums 
makepkg -si --noconfirm 
mkdotenv -h