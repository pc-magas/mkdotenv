# Package metadata
PKG_NAME = mkdotenv
BUILD = 1
VERSION := $(shell grep 'const VERSION' ./mkdotenv/msg/msg.go | sed -E 's/.*"([^"]+)".*/\1/')
ARCH = amd64
BIN_NAME = mkdotenv_$(VERSION)
LINUX_DIST?=ubuntu
DISTROS ?= jammy noble
DIST=jammy

GO := go

.PHONY: all compile

# Default target
all: compile

# Compile Go binary
compile:
	cd ./mkdotenv &&\
	GOOS=linux GOARCH=$(ARCH) $(GO) build -o ../$(BIN_NAME) mkdotenv.go &&\
	cd ../

# Install the programme
install:
	mkdir -p $(DESTDIR)/usr/bin	
	cp $(BIN_NAME) "$(DESTDIR)/usr/bin/$(PKG_NAME)"
	chmod 755 "$(DESTDIR)/usr/bin/$(PKG_NAME)"
	mkdir -p $(DESTDIR)/usr/share/man/man1
	cp man/$(PKG_NAME).1 $(DESTDIR)/usr/share/man/man1/$(PKG_NAME).1
	chmod 644 $(DESTDIR)/usr/share/man/man1/$(PKG_NAME).1

# Uninstall the programme
uninstall:
	rm -f /usr/bin/$(PKG_NAME)
	rm -f /usr/share/man/man1/mkdotenv.1 

# Clean up build files
clean:
	rm -rf $(BIN_NAME)
	rm -rf *.deb
	rm -rf mkdotenv_$(VERSION)


clean-deb:
	rm -rf mkdotenv_$(VERSION)
	rm -rf ../*.changes
	rm -rf ../*.buildinfo
	rm -rf ../*.dsc	
	rm -rf ../*.orig.tar.gz
	rm -rf ../*.debian.tar.xz

# Package as binary debian image
deb:
	dpkg-buildpackage -b
	mv ../*.deb ./

# Raw binary build
bin: compile
	mv $(BIN_NAME) $(PKG_NAME)

# Build into docker image
docker:
	docker build -t pcmagas/mkdotenv:$(VERSION) -t pcmagas/mkdotenv:latest .

docker-push: docker
	docker push pcmagas/mkdotenv:$(VERSION)
	docker push pcmagas/mkdotenv:latest
