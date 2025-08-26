# Packaging app as a Debian package.

## Required tools and dependencies (for building only)

```shell
sudo apt-get install golang-1.23-go debhelper make build-essential dput
```

## Generate a key and export key_id

### Generate key

If not generating deb files at all as first step you need to generate a signature key. That is used to sign the debian packages. 
For this run:

```shell
sudo apt-get install gnupg debian-keyring
gpg --full-generate-key
```

### Export Key Id

Either uploading to ppa or building a binary deb you need to export the signature at the command.

```
export DEB_SIGN_KEYID=^key_id^
```

Where the `^key_id^` is the signature of your key. 
There is a simple utility script that lists and allows to export the appropriate value:

```
bash ./tools/export_deb_keyid.sh
```

That generated a file named `keyid` containing the selected signature. In order to export the keyid you need to run:

```
export DEB_SIGN_KEYID=$(keyid)
```

Many tools located both in `./deb` or `./ppa` folder do read the file and export the appropriate value like this:

```bash
if [ -f keyid ]; then
    echo "Export Keyid from file"
    export DEB_SIGN_KEYID=$(cat keyid)
fi
```

Meaning that this step is optional.

## Supporting files

* `keyid` is used for debian building scripts to export the keyid in ordeer to sign the packages.
* `EMAIL` contains the email of package maintainer. It is used for changelog syncronization.
* `NAME` contains the name of debian package maintainer. It is used for changelog syncronization.

All files are in `.gitignore` and never should be commited upon git.

## Build Binary Deb:

Just run

```
bash ./deb/build_debian.sh
```

## Upload To PPA

Run: 

```
sudo apt-get update
sudo apt-get install dput
bash ./ppa/create.sh
```
