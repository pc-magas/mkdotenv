# Packaging In Alpine

## Directory and scripts for CI/CD building

At `alpinebuild` directory scripts and all nessesary files reside for alpine package generation. There the most basic script is the `make_apk.sh` one.

### The make_apk.sh script

This script does the following:

1. Creates nessesary folders for volume moiunting
2. Aggregates a source code in order to be built into a apk (using the `make_tar.sh` script).
3. Generates an `APKBUILD` file from `APKBUILD-template` that allows you to build from the generated tar.

Is goal is to provide a common utility for generating a final apk.

## Docker Image

The `Dockerfile` at `alpinebuild` directory is the one that setups an environment for packaging the app as an apk file.

### Volumes

It uses these volumes:

1. `/usr/src/apkbuild` where `APKBUILD` and optionally nessesary tarball with *source code* reside upon.
2. `/home/packager/.abuild` where private package signing keys and confoig for abuild is located upon
3. `/etc/apk/keys/` Where public signing keys reside upon
4. `/home/packager/release` where generated apk are stored upon.


## Github actions

Upon `release.yml` workflow, at action `alpine_source` the following files are generated:

1. A `APKBUILD` file containing a remoter path for an aports build.
2. A tar.gz with the source code that can be used for alpine building, already tested via `make_apk.sh` .

Once generates are releases into a github actions release using the `release`. Then we can use the generated `APKBUILD` and tarball in order to create an aports build. 
