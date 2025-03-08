#/usr/bin/bash

CHANGELOG="Changelog.md"

VERSION=$(cat VERSION)
RELEASE_NOTES=$(cat RELEASE_NOTES)
DATE=$(date +%Y-%m-%d)

# Format new entry
NEW_ENTRY="# Version $VERSION $DATE\n\n$RELEASE_NOTES\n\n"

# Prepend to the changelog
echo -e "$NEW_ENTRY$(cat $CHANGELOG)" > $CHANGELOG

echo "Updated $CHANGELOG with version $VERSION."

