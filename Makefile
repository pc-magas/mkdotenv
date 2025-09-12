# Package metadata
PKG_NAME = mkdotenv
BUILD = 1
VERSION ?= $(shell [ -f VERSION ] && cat VERSION || echo dev)
GO := go

INSTALL_BIN_DIR ?= /usr/local/bin
INSTALL_MAN_DIR ?= /usr/local/share/man/man1

OS ?= $(GOOS)
CGO := 0

ifeq ($(OS),)
  OS := $(shell uname -s 2>/dev/null || echo Unknown)
endif

ifeq ($(OS),Darwin)
    OS := darwin
else ifeq ($(OS),Linux)
    OS := linux
else ifneq (,$(findstring MINGW,$(OS)))
    OS := windows
else ifneq (,$(findstring MSYS,$(OS)))
    OS := windows
else ifneq (,$(findstring CYGWIN,$(OS)))
    OS := windows
else
    OS := unknown
endif

ARCH ?= $(GOARCH)

ifeq ($(ARCH),)
  ARCH := $(shell uname -m)
endif

ifeq ($(ARCH),x86_64)
    ARCH := amd64
else ifeq ($(ARCH),i386)
    ARCH := 386
	CGO := 1
else ifeq ($(ARCH),i686)
    ARCH := 386
	CGO := 1
else ifeq ($(ARCH),x86)
    ARCH := 386
	CGO := 1
else ifeq ($(ARCH),arm64)
    ARCH := arm64
else ifeq ($(ARCH),aarch64)
    ARCH := arm64
else ifeq ($(ARCH),armv7l)
    ARCH := arm
else ifeq ($(ARCH),armv7)
    ARCH := arm
else ifeq ($(ARCH),armv6l)
    ARCH := arm
else ifeq ($(ARCH),armhf)
    ARCH := arm
    GOARM := 6
else ifeq ($(ARCH),ppc64le)
    ARCH := ppc64le
else ifeq ($(ARCH),s390x)
    ARCH := s390x
	CGO := 1
else ifeq ($(ARCH),riscv64)
    ARCH := riscv64
else ifeq ($(ARCH),loongarch64)
    ARCH := loong64
else
    ARCH := unknown
endif

EXT :=

ifeq ($(OS),windows)
    EXT := .exe
endif


BIN_NAME ?= $(PKG_NAME)-$(OS)-$(ARCH)$(EXT)
COMPILED_BIN_PATH ?= /tmp/$(BIN_NAME)

VENDOR ?= 0
MODFLAG :=
ifeq ($(VENDOR),1)
    MODFLAG := -mod=vendor
endif


.PHONY: all,compile,install

# Default target
all: bin

make_bin_folder:
	mkdir -p bin

# Compile Go binary
compile:
	@echo "Building on OS=$(OS), ARCH=$(ARCH), GOARM=$(GOARM)"
	
	@if [ "$(OS)" = "windows" ] && [ "$(ARCH)" != "amd64" ]; then \
		echo "Error: Windows builds are only supported on x86_64 (amd64)."; \
		exit 1; \
	fi

	cd ./mkdotenv && \
	mkdir -p /tmp/go-mod-cache &&\
	GOCACHE=/tmp/go-build-cache \
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=$(CGO) GOARM=$(GOARM) \
	$(GO) build $(MODFLAG) -ldflags "-X github.com/pc-magas/mkdotenv/msg.version=$(VERSION)" -o $(COMPILED_BIN_PATH) . &&\
	cd ../

test_run:
	cd ./mkdotenv &&\
	$(GO) run mkdotenv.go $(ARGS)

test:
	cd ./mkdotenv &&\
	mkdir -p /tmp/go-mod-cache &&\
	GOCACHE=/tmp/go-build-cache \
    $(GO) test $(MODFLAG) ./... &&\
    cd ../

# Raw binary build
bin: compile make_bin_folder
	mv $(COMPILED_BIN_PATH) ./bin/$(BIN_NAME)

install_bin:
	mkdir -p $(DESTDIR)$(INSTALL_BIN_DIR)
	install -m 755 ./bin/$(BIN_NAME) "$(DESTDIR)$(INSTALL_BIN_DIR)/$(PKG_NAME)"

install_man:
	mkdir -p $(DESTDIR)$(INSTALL_MAN_DIR)
	install -m 644 man/$(PKG_NAME).1 "$(DESTDIR)$(INSTALL_MAN_DIR)/$(PKG_NAME).1"

uninstall:
	rm -f $(INSTALL_BIN_DIR)/$(PKG_NAME)
	rm -f $(INSTALL_MAN_DIR)/$(PKG_NAME).1

# Install the programme
install: bin install_bin install_man

# Clean up build files
clean:
	rm -rf $(COMPILED_BIN_PATH)
	rm -rf *.deb
	rm -rf mkdotenv_$(VERSION)

vendor-clean:
	cd mkdotenv && \
	rm -rf vendor && \
	go clean -modcache && \
	go mod tidy && \
	go mod vendor && \
	go mod verify
	
