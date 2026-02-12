#!/usr/bin/env bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color (Reset)


SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../..

cd ${SOURCEPATH}
make bin

cd ${SCRIPTPATH}
chmod +x ${SOURCEPATH}/bin/mkdotenv-linux-amd64
echo
echo "#################################"
echo "Executing"
echo "#################################"

${SOURCEPATH}/bin/mkdotenv-linux-amd64 --output-file=.env

# Extract the value for a specific key
ACTUAL_VALUE=$(grep "^PASSWORD=" .env | cut -d'=' -f2)

echo
echo "#################################"
echo "Validating"
echo "#################################"


if [ "$ACTUAL_VALUE" == "1234" ]; then
    echo -e "${GREEN}PASSWORD is correct.${NC}"
else
    echo "${RED}PASSWORD is missing or incorrect.${NC}"
    exit 1
fi

ACTUAL_VALUE=$(grep "^USERNAME=" .env | cut -d'=' -f2)

if [ "$ACTUAL_VALUE" == "test" ]; then
    echo -e "${GREEN}USERNAME is correct.${NC}"
else
    echo -e "${RED}USERNAME is missing or incorrect.${NC}"
    exit 1
fi