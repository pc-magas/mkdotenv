#!/bin/bash

# Get list of fingerprints and associated user names
fingerprints=$(gpg --list-keys --with-colons | awk -F: '
    $1 == "fpr" { fpr=$10 }
    $1 == "uid" { print fpr " " $10 }
')

# Check if we found any keys
if [[ -z "$fingerprints" ]]; then
    echo "No GPG keys found."
    exit 1
fi

# Display available keys for selection
echo "Available GPG keys:"
echo "$fingerprints" | nl

# Ask user to select a key
read -p "Enter the number of the key to export: " selection

selected_fpr=$(echo "$fingerprints" | sed -n "${selection}p" | awk '{print $1}')

if [[ -z "$selected_fpr" ]]; then
    echo "Invalid selection."
    exit 1
fi

SCRIPTPATH=$(dirname "$0") 
FILE=${SCRIPTPATH}/../keyid

echo ${selected_fpr} > ${FILE}
