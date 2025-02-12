# Package metadata
PKG_NAME = mkdotenv
BUILD = 1
VERSION := $(shell grep 'const VERSION' ./src/mkdotenv.go | sed -E 's/.*"([^"]+)".*/\1/')
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
	GOOS=linux GOARCH=$(ARCH) $(GO) build -o $(BIN_NAME) ./src/*


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

# POackage as debian image
deb:
	dpkg-buildpackage -b
	mv ../*.deb ./

# Step 1: Create the source folder if the tarball does not exist
create_source_folder:
	mkdir -p mkdotenv_$(VERSION)
	cp -r src mkdotenv_$(VERSION)
	cp -r man mkdotenv_$(VERSION)
	cp Makefile mkdotenv_$(VERSION)
	tar --exclude=debian --exclude=alpinebuild -czf ../mkdotenv_$(VERSION).orig.tar.gz mkdotenv_$(VERSION);

# Step 2: Create the source package
source_package: ../mkdotenv_$(VERSION).orig.tar.gz
	sed -i 's/unstable/$(DIST)/g' debian/changelog
	sed -i 's/debian/$(LINUX_DIST)/g' debian/changelog
	dpkg-buildpackage -S -sa
	sed -i 's/$(DIST)/unstable/g' debian/changelog
	sed -i 's/$(LINUX_DIST)/debian/g' debian/changelog

#create files for PPA
ppa: 
	$(MAKE) create_source_folder
	for dist in $(DISTROS); do \
		$(MAKE) source_package DIST=$$dist; \
	done
	dput ppa:pcmagas/mkdotenv ../mkdotenv_*_source.changes


# Raw binary build
bin: compile
	mv $(BIN_NAME) $(PKG_NAME)

# Build into docker image
docker:
	docker build -t pcmagas/mkdotenv:$(VERSION) -t pcmagas/mkdotenv:latest .

docker-push: docker
	docker push pcmagas/mkdotenv:$(VERSION)
	docker push pcmagas/mkdotenv:latest
