#!/usr/bin/env bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

cd ${SCRIPTPATH}/..

VERSION=$(cat VERSION)

docker build -t pcmagas/mkdotenv:${VERSION} --build-arg VERSION=${VERSION} -t pcmagas/mkdotenv:latest .


BRANCH=${GITHUB_REF##*/}

if [[ $BRANCH == 'master' ]]; then
    echo "Pushing"
    docker push pcmagas/mkdotenv:${VERSION}
	docker push pcmagas/mkdotenv:latest
fi
