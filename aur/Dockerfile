FROM archlinux:latest

RUN pacman -Syu --noconfirm \
    && pacman -S --noconfirm base-devel git sudo pacman-contrib\ 
    && useradd -m builder && echo "builder ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers


COPY --chown=root:root --chmod=0755 ./build_n_run.sh /bin/build_n_run

USER builder
WORKDIR /home/builder

# Entrypoint for custom build scripts
CMD ["/bin/bash"]