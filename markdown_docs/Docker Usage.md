
# Docker

## Upon Image building
Mkdotenv is also shipped via docker image. Its original intention is to use it as a stage for your Dockerfile for example:

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

## Run image into a standalone container.

You can also run it as standalone image as well:

```shell
docker run pcmagas/mkdotenv mkdotenv --version
```

### Ports and volumes

No volumes are provided with this image, also no ports are exposed with docker image as well. 
Extend Image in case you want to use volumes or access any template files.