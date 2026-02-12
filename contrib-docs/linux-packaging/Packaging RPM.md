# Automation Script for this repo

Rpms are built upon fedora Linux and a Docker image is used. In order to build them run:
```
cd ./fedora
bash ./build_fedora_docker.sh
```

Built rpm is in `./rpmbuild/RPMS/x86_64`.