# Package metadata
PKG_NAME = mkdotenv
BUILD = 1
VERSION ?= $(shell [ -f VERSION ] && cat VERSION || echo dev)
GO := go

INSTALL_BIN_DIR ?= /usr/local/bin
INSTALL_MAN_DIR ?= /usr/local/share/man/man1

OS   ?= $(GOOS)
ARCH ?= $(GOARCH)

# Fallback: only if not set
ifeq ($(OS),)
  UNAME_S := $(shell uname -s 2>/dev/null || echo Unknown)
  ifeq ($(UNAME_S),Darwin)
    OS := darwin
  else ifeq ($(UNAME_S),Linux)
    OS := linux
  else ifneq (,$(findstring MINGW,$(UNAME_S)))
    OS := windows
  else ifneq (,$(findstring MSYS,$(UNAME_S)))
    OS := windows
  else ifneq (,$(findstring CYGWIN,$(UNAME_S)))
    OS := windows
  else
    OS := unknown
  endif
endif

ifeq ($(ARCH),)
  UNAME_M := $(shell uname -m)
  ifeq ($(UNAME_M),x86_64)
    ARCH := amd64
  else ifeq ($(UNAME_M),arm64)
    ARCH := arm64
  else
    ARCH := unknown
  endif
endif

EXT :=
CGO := 0

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
	@echo "Building on OS=$(OS), ARCH=$(ARCH)"
	cd ./mkdotenv && \
	mkdir -p /tmp/go-mod-cache &&\
	GOCACHE=/tmp/go-build-cache \
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=$(CGO) \
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
	
