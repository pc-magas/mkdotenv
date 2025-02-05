
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

# Install

## From source code:

## Step 0 Install golang:

Upon Linux Mint you can run:

```
sudo apt-get install golang-go golang-1.23*
```

For other linux distros look upon: https://go.dev/doc/install


### Step1 Clone repo:

```shell
git clone https://github.com/pc-magas/mkdotenv.git
```

### Step 2 build source code

```shell
make
```

### Step 3 Install

```shell
sudo make install
```

(If run as root ommit `sudo`)


# Uninstall

If cloned this repo and built the tool you can do:

```
sudo make uninstall
```

Otherwize you can do it manually:

```
rm -f /usr/bin/mkdotenv
rm -f /usr/local/share/man/man1/mkdotenv.1 
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
Also you can use the parameter `--input-file` in order to select which file to read upon, if ommited `.env` file is used.

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

### Example 3 Read a specified .env file and output its contents to a seperate .env file:

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

Or temporaly mounting it on a run command:

```Dockerfile
RUN --mount=type=bind,from=pcmagas/mkdotenv:latest,source=/usr/bin/mkdotenv,target=/bin/mkdotenv
```

## Run image into standalone container.

You can also run it as stanalone image as well:

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
Keep in mind to use the `-i` argument upon docker command that enables to read the input via the pipes. If ommited the `mkdotenv` command residing inside the container will not be able to read the contents of .env file piped to it.

### <ins>**Note**</ins>

If running the `pcmagas/mkdotenv` image **as is** the arguments `--env-file`,`--input-file` and `--input-file` will result an unsucessfull execution of `mkdotenv`. 

If a `.env` file needs to be manipulated either pipe the ouitputs as shown upon examples above or extend the `pcmagas/mkdotenv` using a your own Dockerfile providing a nessesary volume:

```Dockerfile
FROM `pcmagas/mkdotenv`

RUN mkdir app

VOLUME app

```

These do not apply if following the instructions shown into [Upon Image building](#upon-image-building) section.

### Ports and volumes

No volumes are provided with this image, also no ports are exposed with docker image as well.
