#!/usr/bin/env sh

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
echo ${SCRIPT_DIR}

docker build -f ${SCRIPT_DIR}/Dockerfile -t pcmagas/alpinebuild ${SCRIPT_DIR}
docker run -v ${SCRIPT_DIR}:/home/packager -ti -u root pcmagas/alpinebuild bash -c "chown -R packager:packager /home/packager/* && bash"