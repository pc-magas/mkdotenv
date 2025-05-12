#!/usr/bin/env bash

test -n "$BASH_VERSION" || exec /bin/bash $0 "$@"
if ! command -v dialog &> /dev/null; then
    echo "Error: 'dialog' is not installed. Install it with: sudo apt install dialog"
    exit 1
fi

prompt_and_save() {
    local file=$1
    local message=$2
    local value=""

    # Check if file exists and read value
    if [ -f "$file" ]; then
        value=$(cat "$file")
    fi

    # Prompt user with dialog
    value=$(dialog --inputbox "$message" 8 50 "$value" 3>&1 1>&2 2>&3)

    # Save the value if not empty
    if [ ! -z "$value" ]; then
        echo "$value" > "$file"
    fi

    # Return the value
    echo "$value"
}

dialog --title "Version Bump Confirmation" --yesno "Do you want to bump the version and update changelogs?" 7 50
response=$?
clear
if [[ $response -ne 0 ]]; then
    echo "Aborting version bump process."
    exit 0
fi

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../ 

cd ${SOURCEPATH}

DEBIAN_CHANGELOG="debian/changelog"
SPEC_FILE="rpmbuild/SPECS/mkdotenv.spec"
DATE=$(date +%Y-%m-%d)
CHANGELOG="Changelog.md"

DEBEMAIL_VAL=$(prompt_and_save "EMAIL" "Enter your email:")
NAME_VAL=$(prompt_and_save "NAME" "Enter your name:")
UPSTREAM_VERSION=$(prompt_and_save "VERSION" "Bump the version (or keep it the same)")
clear

sensible-editor RELEASE_NOTES
RELEASE_NOTES=$(cat RELEASE_NOTES)

NEW_ENTRY="# Version $UPSTREAM_VERSION $DATE"

echo "Prepending new version entry to $CHANGELOG."
echo -e "$NEW_ENTRY\n\n$RELEASE_NOTES\n\n$(cat $CHANGELOG)" > "$CHANGELOG"

# Let user edit the changelog
sensible-editor "$CHANGELOG"

bash ./fedora/update_changelog_docker.sh

# Update Debian changelog

echo "Adding new Debian changelog entry for version $UPSTREAM_VERSION."
DEB_VERSION="$UPSTREAM_VERSION-0debian1~unstable1"

export DEBEMAIL=$DEBEMAIL_VAL

if ! head ${DEBIAN_CHANGELOG} | grep -q "${DEB_VERSION}"; then 
    dch -M --distribution unstable --newversion $DEB_VERSION -m ""
else
    echo "Avoiding to create a new version, appending messages only"
fi

while IFS= read -r line; do
    [[ -z "$line" ]] && echo "LINE EMPTY"&& continue  # Skip empty lines
    echo $line;
    dch -M -a "$line"
done < RELEASE_NOTES

unset DEBEMAIL

# Prompt user to edit Debian changelog
sensible-editor "$DEBIAN_CHANGELOG"

echo "Bump Version for Alpine"
sed -i "s|pkgver=".*"|pkgver="${UPSTREAM_VERSION}"|" ${SOURCEPATH}/alpinebuild/APKBUILD-template
sensible-editor "${SOURCEPATH}/alpinebuild/APKBUILD-template"

echo "Bump version for Arch"
sed -i "s|pkgver=".*"|pkgver="${UPSTREAM_VERSION}"|" ${SOURCEPATH}/aur/PKGBUILD
sensible-editor "${SOURCEPATH}/aur/PKGBUILD"


echo "Version updated successfully: $UPSTREAM_VERSION"
git commit -m "[Autotool] Bump version and fix into nessesary files" ./$CHANGELOG ./$DEBIAN_CHANGELOG ./$SPEC_FILE ./Changelog.md ./VERSION ./RELEASE_NOTES ${SOURCEPATH}/alpinebuild/APKBUILD-template

