FROM archlinux:latest

RUN pacman -Syu --noconfirm \
    && pacman -S --noconfirm base-devel git sudo

# Create a non-root user for building
RUN useradd -m builder && echo "builder ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers
USER builder
WORKDIR /home/builder

# Entrypoint for custom build scripts
CMD ["/bin/bash"]