#!/usr/bin/env bash

CHANGELOG="Changelog.md"
DEBIAN_CHANGELOG="debian/changelog"
SPEC_FILE="rpmbuild/SPECS/mkdotenv.spec"
UPSTREAM_VERSION=$(cat VERSION)  # Read the new upstream version
DATE=$(date +%Y-%m-%d)

# Function to select an editor
select_editor() {
  local editor="$VISUAL"
  [ -z "$editor" ] && editor="$EDITOR"
  [ -z "$editor" ] && command -v nano >/dev/null && editor=nano
  [ -z "$editor" ] && command -v vim >/dev/null && editor=vim
  [ -z "$editor" ] && echo "No editor found" >&2 && return 1
  echo "$editor"
}

EDITOR_CHOICE=$(select_editor) || exit 1

$EDITOR_CHOICE RELEASE_NOTES
RELEASE_NOTES=$(cat RELEASE_NOTES)

# Format new entry
NEW_ENTRY="# Version $UPSTREAM_VERSION $DATE"

# Insert release notes into the changelog
echo -e "$NEW_ENTRY\n\n$RELEASE_NOTES\n\n$(cat $CHANGELOG)" > "$CHANGELOG"
$EDITOR_CHOICE "$CHANGELOG"

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
$EDITOR_CHOICE "$SPEC_FILE"

# Update Debian changelog
echo "Adding new Debian changelog entry for version $UPSTREAM_VERSION."
dch -D unstable -m "$RELEASE_NOTES" --newversion "$UPSTREAM_VERSION-0debian1"

# Prompt user to edit Debian changelog
$EDITOR_CHOICE "$DEBIAN_CHANGELOG"

echo "Version updated successfully: $UPSTREAM_VERSION"
