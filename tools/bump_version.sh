#!/usr/bin/env bash

test -n "$BASH_VERSION" || exec /bin/bash $0 "$@"

if ! command -v dialog &> /dev/null; then
    echo "Error: 'dialog' is not installed. Install it with: sudo apt install dialog"
    exit 1
fi

# Use dialog to show an ncurses-style prompt
dialog --title "Version Bump Confirmation" --yesno "Do you want to bump the version and update changelogs?" 7 50

# Capture the exit status of dialog (0 = Yes, 1 = No)
response=$?

clear  # Clear the screen after dialog closes

if [[ $response -ne 0 ]]; then
    echo "Aborting version bump process."
    exit 0
fi

SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../ 

cd ${SOURCEPATH}

CHANGELOG="Changelog.md"
DEBIAN_CHANGELOG="debian/changelog"
SPEC_FILE="rpmbuild/SPECS/mkdotenv.spec"
DATE=$(date +%Y-%m-%d)

if [ -f EMAIL ];then
    DEBEMAIL_VAL=$(cat EMAIL)
fi

DEBEMAIL_VAL=$(dialog --inputbox "Enter your email:" 8 50 "$DEBEMAIL_VAL" 3>&1 1>&2 2>&3)

if [ ! -z "$DEBEMAIL_VAL" ]; then
    echo $DEBEMAIL_VAL > EMAIL
fi

clear

if [ -f NAME ];then
    NAME_VAL=$(cat NAME)
fi

NAME_VAL=$(dialog --inputbox "Enter your name:" 8 50 "$NAME_VAL" 3>&1 1>&2 2>&3)

if [ ! -z "$NAME_VAL" ]; then
    echo $NAME_VAL > NAME
fi
clear

sensible-editor VERSION
UPSTREAM_VERSION=$(cat VERSION)

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

dch -M --distribution unstable --newversion $DEB_VERSION -m ""
while IFS= read -r line; do
    [[ -z "$line" ]] && echo "LINE EMPTY"&& continue  # Skip empty lines
    echo $line;
    dch -a "$line"
done < RELEASE_NOTES

# Prompt user to edit Debian changelog
sensible-editor "$DEBIAN_CHANGELOG"

echo "Bump Version for Alpine"
sed -i "s|pkgver=".*"|pkgver="${UPSTREAM_VERSION}"|" ${SOURCEPATH}/alpinebuild/APKBUILD-template
sensible-editor "${SOURCEPATH}/alpinebuild/APKBUILD-template"

echo "Version updated successfully: $UPSTREAM_VERSION"
git commit -m "[Autotool] Bump version and fix into nessesary files" ./$CHANGELOG ./$DEBIAN_CHANGELOG ./$SPEC_FILE ./Changelog.md ./VERSION ./RELEASE_NOTES ${SOURCEPATH}/alpinebuild/APKBUILD-template

unset DEBEMAIL
