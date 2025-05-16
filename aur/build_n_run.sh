#!/usr/bin/env bash

# Script intented to run *INSIDE* a docker container ti builds and runs mkdotenv inside an arch linux container
set -euo pipefail

updpkgsums 
makepkg -si --noconfirm 

if [[ $# -gt 0 ]]; then
    echo "Running post-install command: $*"
    exec "$@"
fi
