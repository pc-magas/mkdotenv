#!/usr/bin/env bash

CHANGELOG="Changelog.md"
DEBIAN_CHANGELOG="debian/changelog"
UPSTREAM_VERSION=$(cat VERSION)  # Read the new upstream version
DATE=$(date +%Y-%m-%d)

# Function to select an editor
select_editor() {
  local editor="$VISUAL"
  [ -z "$editor" ] && editor="$EDITOR"
  [ -z "$editor" ] && command -v editor >/dev/null && editor=editor
  [ -z "$editor" ] && command -v nano   >/dev/null && editor=nano
  [ -z "$editor" ] && command -v vim    >/dev/null && editor=vim
  [ -z "$editor" ] && editor=""

  if [ -z "$editor" ]; then
    echo "No editor found" >&2
    return 1
  fi

  echo "$editor"
}

EDITOR_CHOICE=$(select_editor) || exit 1

# Edit release notes
$EDITOR_CHOICE RELEASE_NOTES
RELEASE_NOTES=$(cat RELEASE_NOTES)

# Format new entry
NEW_ENTRY="# Version $UPSTREAM_VERSION $DATE"

# Insert release notes into the changelog
echo -e "$NEW_ENTRY\n\n$RELEASE_NOTES\n\n$(cat $CHANGELOG)" > "$CHANGELOG"
$EDITOR_CHOICE "$CHANGELOG"

# Check if the version already exists in debian/changelog
if grep -q "$UPSTREAM_VERSION" "$DEBIAN_CHANGELOG"; then
    echo "Version $UPSTREAM_VERSION already exists, updating changelog entry."

    # Replace the existing entry with the new one
    sed -i "/$UPSTREAM_VERSION/{
        r /dev/stdin
        d
    }" "$DEBIAN_CHANGELOG" <<< "$NEW_ENTRY $RELEASE_NOTES"
else
    # Version doesn't exist, add a new entry
    echo "Version $UPSTREAM_VERSION not found, adding new changelog entry."
    dch -D unstable -m "$RELEASE_NOTES" --newversion "$UPSTREAM_VERSION-0debian1" # Add -0debian1 as the debian revision
fi

# Finally, edit the changelog file
echo "Editing changelog"
$EDITOR_CHOICE "$DEBIAN_CHANGELOG"