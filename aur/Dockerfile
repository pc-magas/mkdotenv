FROM archlinux:latest

RUN pacman -Syu --noconfirm \
    && pacman -S --noconfirm base-devel git sudo pacman-contrib\ 
    && useradd -m builder && echo "builder ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers


COPY --chown=root:root --chmod=0755 ./build_n_run.sh /bin/build_n_run

RUN echo '#!/usr/bin/env bash' > /usr/local/bin/entrypoint.sh && \
    echo 'set -e' >> /usr/local/bin/entrypoint.sh && \
    echo '' >> /usr/local/bin/entrypoint.sh && \
    echo 'sudo chown builder:builder /home/builder' &&\
    echo 'exec "$@"' >> /usr/local/bin/entrypoint.sh && \
    chmod +x /usr/local/bin/entrypoint.sh

USER builder
WORKDIR /home/builder
VOLUME /home/builder

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["/bin/bash"]