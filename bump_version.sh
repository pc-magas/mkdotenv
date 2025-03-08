#!/usr/bin/env bash

CHANGELOG="Changelog.md"

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

# Edit version file
$EDITOR_CHOICE VERSION
VERSION=$(cat VERSION)

# Edit release notes
$EDITOR_CHOICE RELEASE_NOTES
RELEASE_NOTES=$(cat RELEASE_NOTES)
DATE=$(date +%Y-%m-%d)

# Format new entry
NEW_ENTRY="# Version $VERSION $DATE"

# Check if entry already exists
if grep -q "$NEW_ENTRY" "$CHANGELOG"; then
  echo "Entry for version $VERSION already exists in $CHANGELOG."
  # Prompt user whether to prepend or skip
  read -p "Do you want to prepend this entry to the changelog? (y/n): " user_input
  if [[ "$user_input" =~ ^[Yy]$ ]]; then
    echo -e "$NEW_ENTRY\n\n$RELEASE_NOTES\n\n$(cat $CHANGELOG)" > "$CHANGELOG"
    echo "Prepended new entry to $CHANGELOG."
  else
    echo "Skipped updating $CHANGELOG."
  fi
else
  # If entry doesn't exist, prepend it
  echo -e "$NEW_ENTRY\n\n$RELEASE_NOTES\n\n$(cat $CHANGELOG)" > "$CHANGELOG"
  echo "Updated $CHANGELOG with version $VERSION."
fi

# Edit changelog
echo "Editing changelog"
$EDITOR_CHOICE "$CHANGELOG"
