# Packaging In Alpine

## Build workflow

Run these commands in order tot est whether apok is built sucessfully and able to run:

```shell
cd ./alpinebuild
bash make_apk.sh
bash test_apk.sh
```

## Directory and scripts for CI/CD building

At `alpinebuild` directory scripts and all nessesary files reside for alpine package generation. There the most basic script is the `make_apk.sh` one.

### The make_apk.sh script

This script does the following:

1. Creates nessesary folders for volume moiunting
2. Aggregates a source code in order to be built into a apk (using the `make_tar.sh` script).
3. Generates an `APKBUILD` file that allows you to build from the generated tar.
4. Releases a production-readt `APKBUILD` upon `alpinebuild/release` directory.

Is goal is to provide a common utility for generating a final apk. It executes `make_tar.sh` and `make_apkbuild.sh` scripts, described bellow.

Upon folder `release` it releases:
* `APKBUILD` ready for aports
* `APKBUILD.local` a APKBUILD variatey that has the source into the generated tar instead or release
* source code with *go* dependencies in a `.tar.gz` named as `mkdotenv-^version^.tar.gz` (replace `^version^` with current version).
* generated alpine packahes upon `release/home/*` folders each built architecture has its owen dedicated folder in it.

### The make_tar.sh script.

It is executed inside the `make_apk.sh` one and packages and it is responsible to aggregate any nessesart source code and dependency into a single tar. 
The source code includes the source code of go dependencies as well.

### The make_apkbuild.sh script

Script that creates the nessesary APKBUILD file. It accepts these arguments:

* `--src_local` it uses the source code tar generated via `make_apk.sh` script instead of the apk uploaded into release
* `--checksum` Source code tarball checksum

Also you can place the path of the directory where APKFILE should be released if not value provided it is assumed the path where `make_apkbuild.sh` is stored upon.

## Docker Image

It uses the `ghcr.io/pc-magas/alpinebuild` image for any build upon alpine.

### Volumes

2 folders are used as volume storage:

* `alpinebuild/volumes` in which any docker volume reside upon.
* `alpinebuild/release` in which all built files (APKBUILD, source code tar and *.apk files) reside upon.
