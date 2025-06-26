#!/usr/bin/env bash

SCRIPTPATH=$(dirname "$0")
CHANGES_FILE=${SCRIPTPATH}/../../mkdotenv_*_source.changes

bash ${SCRIPTPATH}/package.sh

dput ppa:pcmagas/mkdotenv ${CHANGES_FILE}

