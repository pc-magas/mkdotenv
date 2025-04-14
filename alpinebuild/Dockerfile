FROM alpine

ENV UID=1000\
GID=1000

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
    addgroup packager abuild

RUN echo "#!/usr/bin/env bash" > /usr/bin/entrypoint &&\
    echo "chown -R packager:packager /home/packager/*" >> /usr/bin/entrypoint &&\
    echo "su packager -c \"\$*\"" >> /usr/bin/entrypoint &&\
    chown root:root /usr/bin/entrypoint &&\
    chmod 0775 /usr/bin/entrypoint

USER packager
WORKDIR /home/packager
VOLUME /home/packager

USER root

ENTRYPOINT ["/usr/bin/entrypoint"]