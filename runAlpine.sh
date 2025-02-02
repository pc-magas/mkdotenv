#!/usr/bin/env sh

docker run -v ./APKBUILD:/home/packager/APKBUILD -ti -u root pcmagas/alpinebuild bash -c "chown -R packager:packager /home/packager/APKBUILD && bash"