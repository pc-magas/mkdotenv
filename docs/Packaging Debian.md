# Packaging app as a Debian package.

## Required tools and dependencies (for building only)

```
sudo apt-get install golang-1.23-go golang-go debhelper make
```

## Generate a key and export key_id

Either uploading to ppa or building a binary deb you need to export the signature at the command.

```
export DEB_SIGN_KEYID=^key_id^
```

Where the `^key_id^` is the signature of your key. 


In order to find the `^key_id^` run:

```
gpg --list-keys
```

This will output a value for example:

```
/home/pcmagas/.gnupg/pubring.kbx
--------------------------------
pub   rsa4096 2025-02-03 [SC]
      42F71A9B087D2AF8786DE39442DD352E68415A45
uid           [ultimate] John Doe (Debian signing key) <pcmagas@example.com>
sub   rsa4096 2025-02-03 [E]
```

The `^key_id` value is the line bellow pub for example for the output above is:

```
42F71A9B087D2AF8786DE39442DD352E68415A45
```

## Build Binary Deb:

Just run

```
make deb
```

## Upload To PPA

Run: 

```
make ppa
```
