FROM alpine

RUN apk update && apk add --no-cache \
    alpine-sdk \
    abuild \
    git \
    go \
    bash \
    fakeroot \
    curl \
    sudo \
    make && \
    adduser -D packager &&\
    echo "packager ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers &&\
    chown -R packager:packager /var/cache/distfiles &&\
    addgroup packager abuild

USER packager
WORKDIR /home/packager
VOLUME /home/packager

CMD ["bash"]