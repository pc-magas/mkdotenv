# Package metadata
PKG_NAME = mkdotenv
BUILD = 1
VERSION := $(shell grep 'const VERSION' ./src/mkdotenv.go | sed -E 's/.*"([^"]+)".*/\1/')
ARCH = amd64
BIN_NAME = mkdotenv_$(VERSION)

.PHONY: all compile ci

# Default target
all: compile

# Compile Go binary
compile:
	GOOS=linux GOARCH=$(ARCH) go build -o $(BIN_NAME) ./src/*


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

# POackage as debian image
deb:
	dpkg-buildpackage -b
	mv ../*.deb ./

# Raw binary build
bin: build
	mv $(BIN_NAME) $(PKG_NAME)

# Build into docker image
docker:
	docker build -t pcmagas/mkdotenv:$(VERSION) -t pcmagas/mkdotenv:latest .

docker-push: docker
	docker push pcmagas/mkdotenv:$(VERSION)
	docker push pcmagas/mkdotenv:latest
