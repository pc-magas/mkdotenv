
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
# Replace with your own release on https://github.com/pc-magas/mkdotenv/releases
export VERSION=v0.3.2
wget -o mkdotenv https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv-linux-amd64 
sudo cp mkdotenv-linux-amd64  /bin/mkdotenv
sudo chmod 755 /bin/mkdotenv

mkdotenv --version
```

For environments that do not provide root access use these commands:

```bash
# Replace with your own release on https://github.com/pc-magas/mkdotenv/releases
export VERSION=v0.3.2
wget -o mkdotenv https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv-linux-amd64 
cp mkdotenv-linux-amd64  ~/.local/bin/mkdotenv
chmod 755 ~/.local/bin/mkdotenv

mkdotenv --version
```

### Uninstall

```
rm -rf /bin/mkdotenv
```

### Using PPA for Ubuntu & Linux Mint

If running ubuntu or Linux mint you can use our PPA repository:

```
sudo add-apt-repository ppa:pcmagas/mkdotenv
sudo apt-get update
sudo apt-get install mkdotenv
```

### From debian package

Works in Debian, Mint and Ubuntu (or any other Debian-compatible distros)

```shell
# Replace with your own release on https://github.com/pc-magas/mkdotenv/releases
export VERSION=v0.3.2
wget https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv_${VERSION}_amd64.deb
sudo dpkg -i mkdotenv_${VERSION}_amd64.deb
```

At code above replace `^found_version^` with the version shown at [Detect Latest Version](#detect-latest-version).

Uninstalling the package is easy as:

```bash
sudo apt-get remove mkdotenv
```

#### Using RPM package

Tested on Fedora

```shell
# Replace with your own release on https://github.com/pc-magas/mkdotenv/releases
export VERSION=v0.3.2
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
### In amazon Linux 2023

Packages available through corpr

```
sudo dnf install -y 'dnf-command(copr)'
sudo dnf copr enable pc-magas/mkdotenv 
sudo dnf install -y mkdotenv
```

### In OpenSuse

Fedora Corpr repositories also are available for OpenSuse (both leap and tumbleweed)

```bash
sudo zypper install opi
opi copr pc-magas/mkdotenv
sudo zypper install mkdotenv
```

### In Alpine Linux

```shell
# Replace with your own release on https://github.com/pc-magas/mkdotenv/releases
export VERSION=v0.3.2
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
# Replace with your own release on https://github.com/pc-magas/mkdotenv/releases
export VERSION=v0.3.2
wget -o mkdotenv https://github.com/pc-magas/mkdotenv/releases/download/${VERSION}/mkdotenv-darwin-arm64 
sudo cp mkdotenv /usr/local/bin/mkdotenv
sudo chmod 755 /usr/local/mkdotenv

mkdotenv --version
```

# Usage

For usage consult `./docs` folder.
