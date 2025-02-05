# Packaging app as a Debian package.

## Commands:

It is assumed that you cloned the repository.

```
export DEB_SIGN_KEYID=^key_id^
make deb
```

Where the `^key_id^` is rthe signature of your key id. In order to found the key_id run:

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
