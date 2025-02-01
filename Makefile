# Package metadata
PKG_NAME = mkdotenv
VERSION = 0.0.2
BUILD = 1
ARCH = amd64
BUILD_DIR = build
DEB_DIR = $(BUILD_DIR)/deb
RPM_DIR = $(BUILD_DIR)/rpmbuild
RPM_NAME = $(PKG_NAME)-$(VERSION)-$(BUILD).$(ARCH).rpm

# Default target
all: build

# Compile Go binary
build:
	GOOS=linux GOARCH=$(ARCH) go build -o $(PKG_NAME) ./src/*

# Create the .deb package
deb: clean build
	mkdir -p $(DEB_DIR)/DEBIAN
	mkdir -p $(BUILD_DIR)/usr/local/bin
	mkdir -p $(BUILD_DIR)/usr/share/man/man1

	# Create control file
	echo "Package: $(PKG_NAME)" > $(DEB_DIR)/DEBIAN/control
	echo "Version: $(VERSION)" >> $(DEB_DIR)/DEBIAN/control
	echo "Section: utils" >> $(DEB_DIR)/DEBIAN/control
	echo "Priority: optional" >> $(DEB_DIR)/DEBIAN/control
	echo "Architecture: $(ARCH)" >> $(DEB_DIR)/DEBIAN/control
	echo "Maintainer: Your Name <your.email@example.com>" >> $(DEB_DIR)/DEBIAN/control
	echo "Description: A CLI tool for managing .env files" >> $(DEB_DIR)/DEBIAN/control
	echo " Adds or updates variables in .env files with optional input/output file support." >> $(DEB_DIR)/DEBIAN/control

	# Copy binary
	cp $(PKG_NAME) $(DEB_DIR)/usr/local/bin/
	chmod 755 $(DEB_DIR)/usr/local/bin/$(PKG_NAME)

	# Copy man page if exists
	if [ -f man/$(PKG_NAME).1 ]; then \
		cp man/$(PKG_NAME).1 $(DEB_DIR)/usr/share/man/man1/$(PKG_NAME).1; \
		gzip -9 $(DEB_DIR)/usr/share/man/man1/$(PKG_NAME).1; \
	fi

	# Build .deb package
	dpkg-deb --build $(DEB_DIR)
	mv $(DEB_DIR).deb mkdotenv.deb

# Install the programme
install:
	cp $(PKG_NAME) /usr/bin/$(PKG_NAME)
	chmod 755 /usr/bin/$(PKG_NAME)
	cp man/mkdotenv.1 /usr/local/share/man/man1/

# Uninstall the programme
uninstall:
	rm -f /usr/bin/$(PKG_NAME)
	rm -f /usr/local/share/man/man1/mkdotenv.1 


# Clean up build files
clean:
	rm -rf build 

.PHONY: all build deb install uninstall clean
