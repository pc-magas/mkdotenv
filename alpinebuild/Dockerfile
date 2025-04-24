FROM alpine

ENV UID=1000\
GID=1000\
EXTERNAL_KEY=false\
PUBKEYNAME="pubkey.asc"\
PRIVKEYNAME="pubkey"

RUN apk update && apk add --no-cache \
    alpine-sdk \
    abuild \
    git \
    go \
    bash \
    fakeroot \
    curl \
    sudo \
    make\
    shadow && \
    adduser -D packager &&\
    echo "packager ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers &&\
    chown -R packager:packager /var/cache/distfiles &&\
    addgroup packager abuild &&\
    mkdir -p /usr/src/apkbuild

COPY --chown=root:root --chmod=0755 ./dockerscripts/* /usr/bin/

VOLUME /etc/apk/keys/
VOLUME /usr/src/apkbuild 

USER packager
WORKDIR /home/packager
RUN mkdir -p /home/packager/.abuild &&\
    mkdir -p /home/packager/release

VOLUME /home/packager/.abuild
VOLUME /home/packager/release

USER root

ENTRYPOINT ["/usr/bin/entrypoint"]
CMD [ "build" ]