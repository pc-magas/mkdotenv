#!/usr/bin/env bash

test -n "$BASH_VERSION" || exec /bin/bash $0 "$@"

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../ 

cd ${SOURCEPATH}

CHANGELOG="Changelog.md"
DEBIAN_CHANGELOG="debian/changelog"
SPEC_FILE="rpmbuild/SPECS/mkdotenv.spec"
DATE=$(date +%Y-%m-%d)

sensible-editor VERSION
UPSTREAM_VERSION=$(cat VERSION)

sensible-editor RELEASE_NOTES
RELEASE_NOTES=$(cat RELEASE_NOTES)

NEW_ENTRY="# Version $UPSTREAM_VERSION $DATE"

# Check if the first line contains the same version
if head -n 1 "$CHANGELOG" | grep -q "$UPSTREAM_VERSION"; then
    echo "Version $UPSTREAM_VERSION already exists at the top of $CHANGELOG, overwriting."
    # Overwrite existing top entry
    sed -i "1,/^$/c\\$NEW_ENTRY\n\n$RELEASE_NOTES\n" "$CHANGELOG"
else
    echo "Prepending new version entry to $CHANGELOG."
    # Prepend new entry at the top
    echo -e "$NEW_ENTRY\n\n$RELEASE_NOTES\n\n$(cat $CHANGELOG)" > "$CHANGELOG"
fi

# Let user edit the changelog
sensible-editor "$CHANGELOG"

# Extract version and release from RPM spec file
CURRENT_RPM_VERSION=$(grep -E '^Version:' "$SPEC_FILE" | awk '{print $2}')
CURRENT_RPM_RELEASE=$(grep -E '^Release:' "$SPEC_FILE" | awk '{print $2}' | sed 's/%{?dist}//')

# Determine if version needs bumping
if [ "$CURRENT_RPM_VERSION" == "$UPSTREAM_VERSION" ]; then
    echo "RPM spec file already has version $UPSTREAM_VERSION."
    echo "What would you like to do?"
    echo "1) Increase Release number (current: $CURRENT_RPM_RELEASE)"
    echo "2) Keep the same Release number"
    read -rp "Choose an option [1/2]: " choice

    if [ "$choice" == "1" ]; then
        NEW_RPM_RELEASE=$((CURRENT_RPM_RELEASE + 1))
        echo "Increasing release number to $NEW_RPM_RELEASE."
    else
        NEW_RPM_RELEASE=$CURRENT_RPM_RELEASE
        echo "Keeping the release number at $NEW_RPM_RELEASE."
    fi
else
    NEW_RPM_RELEASE="1"  # Reset release for new version
    echo "Updating RPM version to $UPSTREAM_VERSION with release 1."
fi

# Update Version and Release in spec file
sed -i "s/^Version:.*/Version:        $UPSTREAM_VERSION/" "$SPEC_FILE"
sed -i "s/^Release:.*/Release:        $NEW_RPM_RELEASE%{?dist}/" "$SPEC_FILE"

# Add changelog entry to RPM spec
rpmdev-bumpspec -c "$RELEASE_NOTES" -u "$(whoami)" "$SPEC_FILE"

# Prompt user to edit spec file
sensible-editor "$SPEC_FILE"

# Update Debian changelog

echo "Adding new Debian changelog entry for version $UPSTREAM_VERSION."
DEB_VERSION="$UPSTREAM_VERSION-0debian1-unstable1"
dch --newversion "$DEB_VERSION"
while IFS= read -r line; do
    echo $line;
    dch --newversion "$DEB_VERSION" -a "$line"
done < RELEASE_NOTES
dch --newversion "$DEB_VERSION" --distribution unstable ignored

# Prompt user to edit Debian changelog
sensible-editor "$DEBIAN_CHANGELOG"

echo "Version updated successfully: $UPSTREAM_VERSION"

cd ${SCRIPTPATH}