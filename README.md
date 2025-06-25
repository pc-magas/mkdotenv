
```text
 __  __ _    _____        _                  
|  \/  | |  |  __ \      | |                 
| \  / | | _| |  | | ___ | |_ ___ _ ____   __
| |\/| | |/ / |  | |/ _ \| __/ _ \ '_ \ \ / /
| |  | |   <| |__| | (_) | ||  __/ | | \ V / 
|_|  |_|_|\_\_____/ \___/ \__\___|_| |_|\_/  
```
                                              
**Simplify Your .env Files – One Variable at a Time!**

MkDotenv is a lightweight and efficient tool for managing your `.env` files. Whether you're adding, updating, or replacing environment variables, MkDotenv makes it easy and hassle-free.

[![Copr build status](https://copr.fedorainfracloud.org/coprs/pc-magas/mkdotenv/package/mkdotenv/status_image/last_build.png)](https://copr.fedorainfracloud.org/coprs/pc-magas/mkdotenv/package/mkdotenv/) 

# Build from source (Supported in Linux & macOS)

Compilation required to have make and golang installed.
Both Go 1.23 and 1.24 supported.

### Step 0: Install golang and make:

#### On Ubuntu Linux and Linux mint 

On Linux Mint and Ubuntu, you can run:

```bash
sudo apt-get install make golang-1.23*
```

#### On macOS and other distros

For other linux distros and macOS follow official instructions https://go.dev/doc/install 


#### Install make upon macOs

You can use either xcode or homebrew. Homebrew method is in https://formulae.brew.sh/formula/make

### Step 1: Clone repo:

```bash
git clone https://github.com/pc-magas/mkdotenv.git
```

For a specific version (e.g `v0.3.2`) you can also do:

```bash
export VERSION=v0.3.2
git clone --depth 1 --branch ${VERSION} https://github.com/pc-magas/mkdotenv.git
```

### Step 2: Compile

```bash
make
```

#### Note
In case you use the `golang-1.23` package shipped with ubuntu and linux mint, and unable to run `go` command line into the shell you can also run this command:

```bash
make GO=/usr/lib/go-1.23/bin/go
```

### Step 3: Install

```bash
sudo make install
```

(If run as root omit `sudo`)

Once `make install` is successfully run golang can be uninstalled if desired, it is a build-only requirement.

#### Uninstall

If cloned this repo and built the tool you can do:

```bash
sudo make uninstall
```

Otherwise you can do it manually:

```bash
sudo rm -f /usr/bin/mkdotenv
sudo rm -f /usr/local/share/man/man1/mkdotenv.1 

sudo rm -f /usr/local/bin/mkdotenv
sudo rm -f /usr/local/share/man/man1/mkdotenv.1 
```

# Installation

## Install In linux 

### From Executable Binaries

```shell
export VERSION=v0.3.2
wget -o mkdotenv https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv-linux-amd64 
sudo cp mkdotenv /bin/mkdotenv
sudo chmod 755 /bin/mkdotenv

mkdotenv --version
```

For environments that do not provide root access use these commands:

```bash
export VERSION=v0.3.2
wget -o mkdotenv https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv-linux-amd64 
sudo cp mkdotenv ~/.local/bin/mkdotenv
sudo chmod 755 ~/.local/bin/mkdotenv

mkdotenv --version
```

#### Detect latest Version

In order to see the latest version check the https://github.com/pc-magas/mkdotenv/releases page once you found the desired release instead of:

```bash
# Replace with actual version number, e.g., v0.3.2
export VERSION=v0.1.0
```

Do (replace `^found_version^` with the version you found upon releases page):

```bash
# Replace with actual version number, e.g., v0.3.2
export VERSION=^found_version^
```

And execute the commands:

```bash
wget -O mkdotenv https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv-linux-amd64 

sudo cp mkdotenv /bin/mkdotenv
sudo chmod 755 /bin/mkdotenv

mkdotenv --version
```

### Uninstall

```
rm -rf /bin/mkdotenv
```

### Via PPA for Ubuntu & Linux Mint

If running ubuntu or Linux mint you can use our PPA repository:

```
sudo add-apt-repository ppa:pcmagas/mkdotenv
sudo apt-get update
sudo apt-get install mkdotenv
```

### From debian package

Works in Debian, Mint and Ubuntu (or any other Debian-compatible distros)

```shell
# Replace with actual version number, e.g., v0.3.2
export VERSION=^found_version^
wget https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv_${VERSION}_amd64.deb
sudo dpkg -i mkdotenv_${VERSION}_amd64.deb
```

At code above replace `^found_version^` with the version shown at [Detect Latest Version](#detect-latest-version).


#### From RPM package

Tested on Fedora

```shell
# Replace with actual version number, e.g., v0.3.2
export VERSION=^found_version^
wget https://github.com/pc-magas/mkdotenv/releases/download/v${VERSION}/mkdotenv.rpm
sudo rpm -i mkdotenv-${VERSION}-1.fc41.x86_64.rpm
```

At code above replace `^found_version^` with the version shown at [Detect Latest Version](#detect-latest-version).

### In Fedora Linux

App is delivered via corpr you can install it like this:

```bash
sudo dnf install dnf-plugins-core
dnf copr enable pc-magas/mkdotenv 
sudo dnf install mkdotenv
```

### In Alpine Linux

```shell
# Replace with actual version number, e.g., v0.3.2
export VERSION=^found_version^
wget https://github.com/pc-magas/mkdotenv/releases/download/v${VERSION}/mkdotenv-${VERSION}-r0.apk
```

> There's a pending release for alpine linux on official repositories.

Then as root user run:

```shell
apk add --allow-untrusted mkdotenv-${VERSION}-r0.apk
```

At code above replace `^found_version^` with the version shown at [Detect Latest Version](#detect-latest-version).

### In arch Linux

Mkdotenv is shipped via [AUR](https://aur.archlinux.org/packages/mkdotenv), use `yay` to install:

```
yay mkdotenv
```

## Install upon Windows

Windows builds are provided as standalone binaries without an installer.

Just download the `.exe` from releases page ( https://github.com/pc-magas/mkdotenv/releases ) and run it through cmd/powershell:

```
mkdotenv-windows-amd64.exe 
```

Once downloaded it can be renamed as `mkdotenv.exe`.

## Install upon MacOS

### Using Homebrew

macOS binaries for M Series are shipped via homebrew.
For intel based releases compile it from source.

```bash
brew tap pc-magas/mkdotenv
brew install pc-magas/mkdotenv/mkdotenv
mkdotenv --help
```

> For intel macs you need to [compile from source](#build-from-source-supported-in-linux--mac)

### From Statically Built Binaries

Statically-built binaries that can be converted into executable for M series macOS are also shipped as well.

```shell
export VERSION=v0.3.2
wget -o mkdotenv https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv-darwin-arm64 
sudo cp mkdotenv /usr/local/bin/mkdotenv
sudo chmod 755 /usr/local/mkdotenv

mkdotenv --version
```

# Usage

## Basic

```
mkdotenv <variable_name> <variable_value>
```

This will output  to *stdout* the contents of a `.env` file with the variable `<variable_name>` having `<variable_value>` instead of the original one.
If no `.env` file exists it will just output the `<variable_name>` having the `<variable_value>`.

### Example:

```
mkdotenv DB_HOST 127.0.0.1
```

This will output:

```
DB_HOST=127.0.0.1
```

If a .env file exists with values:

```
DB_HOST=example.com
DB_USER=xxx
```

The final output would be:

```
DB_HOST=127.0.0.1
DB_USER=xxx
```

## Selecting file to read and write upon

Instead of outputing the .env value you can use the `--output-file` argument in order to write the contents upon a file.
Also you can use the parameter `--input-file` in order to select which file to read upon, if omited `.env` file is used.

### Example 1 Read a specified file and output its contents to *stdout*:

Assuming we run the command

```
mkdotenv DB_HOST 127.0.0.1 --input-file=.env.example
```

This will read the `.env.example` and output:

```
DB_HOST=127.0.0.1
```


### Example 2 Write file upon a .env file:

```
mkdotenv DB_HOST 127.0.0.1 --output-file=.env.production
```

This would **create** a file named `.env.production` containing:

```
DB_HOST=127.0.0.1
```

### Example 3 Read a specified .env file and output its contents to a separate .env file:

Assuming we have a file named `.env.template` containing:

```
DB_HOST=example.com
DB_USER=xxx
DB_PASSWORD=zzz
```

And we want to create a file named `.env.production` containing 

```
DB_HOST=127.0.0.1
DB_USER=xxx
DB_PASSWORD=zzz
```

We have to run:

```
mkdotenv DB_HOST 127.0.0.1 --input-file .env.template --output-file .env.production
```

## Piping outputs

You can provide a .env via a pipe. A common use is to replace multiple variables:

```
mkdotenv DB_HOST 127.0.0.1 | mkdotenv DB_USER maiuser | mkdotenv DB_PASSWORD XXXX --output_file .env.production
```

# Docker

## Upon Image building
Mkdotenv is also shipped via docker image. Its intention is to use it as a stage for your Dockerfile for example:

```Dockerfile

FROM pcmagas/mkdotenv AS mkdotenv

FROM debian 

COPY --from=mkdotenv /usr/bin/mkdotenv /bin/mkdotenv

```

Or alpine based images:

```Dockerfile
FROM pcmagas/mkdotenv AS mkdotenv

FROM alpine 

COPY --from=mkdotenv /usr/bin/mkdotenv /bin/mkdotenv

```

Or temporarily mounting it on a run command:

```Dockerfile
RUN --mount=type=bind,from=pcmagas/mkdotenv:latest,source=/usr/bin/mkdotenv,target=/bin/mkdotenv
```

## Run image into standalone container.

You can also run it as standalone image as well:

```shell
docker run pcmagas/mkdotenv mkdotenv --version
```

If you want to manipulate a `.env` file using the docker image. You can use it like this:

```shell
cat .env | docker run -i pcmagas/mkdotenv mkdotenv Hello BAKA > .env.new
```

Or if you want multiple variables:

```shell
cat .env | docker run -i pcmagas/mkdotenv mkdotenv Hello BAKA | docker run -i pcmagas/mkdotenv mkdotenv BIG BROTHER > .env.new
```
Keep in mind to use the `-i` argument upon docker command that enables to read the input via the pipes. If omited the `mkdotenv` command residing inside the container will not be able to read the contents of .env file piped to it.

### <ins>**Note**</ins>

If running the `pcmagas/mkdotenv` image **as is** the arguments `--env-file`,`--input-file` and `--input-file` will result an unsuccessful execution of `mkdotenv`. 

If a `.env` file needs to be manipulated either pipe the outputs as shown upon examples above or extend the `pcmagas/mkdotenv` using a your own Dockerfile providing a necessary volume:

```Dockerfile
FROM `pcmagas/mkdotenv`

RUN mkdir app

VOLUME app

```

These do not apply if following the instructions shown into [Upon Image building](#upon-image-building) section.

### Ports and volumes

No volumes are provided with this image, also no ports are exposed with docker image as well.
