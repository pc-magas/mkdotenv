#!/bin/bash

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

cp -r ${SCRIPTPATH}/git-hooks/* ${SCRIPTPATH}/../.git/hooks/

chmod +x ${SCRIPTPATH}/../.git/hooks/*