#!/usr/bin/env bash


SCRIPTPATH="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"
SOURCEPATH=${SCRIPTPATH}/../ 

UPSTREAM_VERSION=$(cat ${SOURCEPATH}/VERSION)
RELEASE_NOTES_FILE=${SOURCEPATH}/RELEASE_NOTES

SPEC_FILE=$SOURCEPATH/rpmbuild/SPECS/mkdotenv.spec

CURRENT_RPM_VERSION=$(grep -E '^Version:' "$SPEC_FILE" | awk '{print $2}')
CURRENT_RPM_RELEASE=$(grep -E '^Release:' "$SPEC_FILE" | awk '{print $2}' | sed 's/%{?dist}//')

EMAIL_VAL=$(cat EMAIL)
NAME_VAL=$(cat NAME)


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

RPM_DATE=$(date +"%a %b %d %Y")
CHANGELOG_LINE="* $RPM_DATE ${NAME_VAL} <${EMAIL_VAL}> - %{version}-$NEW_RPM_RELEASE"

echo "" >> $SPEC_FILE
echo "$CHANGELOG_LINE" >> $SPEC_FILE

while IFS= read -r line; do
    [[ -z "$line" ]] && echo "LINE EMPTY"&& continue  # Skip empty lines
    echo "- $line" >> $SPEC_FILE
done < "${RELEASE_NOTES_FILE}"

sensible-editor "$SPEC_FILE"
