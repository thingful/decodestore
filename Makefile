
BIN := dcs
VERSION := $(shell git describe --tags --always --dirty --abbrev=8)
BUILD_DATE := $(shell date "+%FT%H:%M:%S %Z")
BASE_PACKAGE := github.com/thingful/decodestore
PACKAGE := $(BASEPACKAGE)/cmd/$(BIN)
BUILD_DIR := build
BUILD_FLAGS := -v -ldflags "-extldflags -static -X $(BASE_PACKAGE)/pkg/version.Version=$(VERSION) -X \"$(BASE_PACKAGE)/pkg/version.BuildDate=$(BUILD_DATE)\""

BUILD_IMAGE := thingful/decodestorebuilder

OS ?= linux
ARCH ?= amd64

SRC_DIRS := cmd pkg

.PHONY: gen
gen: ## Run the protobuf compiler to generate implementation stubs.
	@docker run \
		-ti \
		--rm \
		-v "$$(pwd)/.go:/go" \
		-v "$$(pwd):/go/src/$(BASE_PACKAGE)" \
	  -w /go/src/$(BASE_PACKAGE) \
		$(BUILD_IMAGE) \
		retool do protoc --proto_path=/go/src:. --twirp_out=. --go_out=. ./pkg/rpc/datastore/service.proto

.PHONY: clean
clean:
	rm -rf .go bin .cache

build-darwin:
	@$(MAKE) --no-print-directory OS=darwin build

build: bin/$(OS)/$(ARCH)/$(BIN)

bin/$(OS)/$(ARCH)/$(BIN): build-dirs
	@echo "Building inside container"
	@docker run \
		-ti \
		--rm \
		-v "$$(pwd)/.go:/go" \
		-v "$$(pwd):/go/src/$(BASE_PACKAGE)" \
		-v "$$(pwd)/bin:/go/bin" \
		-v "$$(pwd)/.go/std/$(OS)/$(ARCH):/usr/local/go/pkg/$(OS)_$(ARCH)_static" \
		-v "$$(pwd)/.cache/go-build:/root/.cache/go-build" \
		-w /go/src/$(BASE_PACKAGE) \
		$(BUILD_IMAGE) \
		/bin/sh -c " \
			ARCH=$(ARCH) \
			OS=$(OS) \
			VERSION=$(VERSION) \
			PKG=$(BASE_PACKAGE) \
			./build/build.sh \
		"

.PHONY: test
test: build-dirs
	@docker run \
		-ti \
		--rm \
		-v "$$(pwd)/.go:/go" \
		-v "$$(pwd):/go/src/$(BASE_PACKAGE)" \
		-v "$$(pwd)/bin:/go/bin" \
		-v "$$(pwd)/.go/std/$(OS)/$(ARCH):/usr/local/go/pkg/$(OS)_$(ARCH)_static" \
		-v "$$(pwd)/.cache/go-build:/root/.cache/go-build" \
		-w /go/src/$(BASE_PACKAGE) \
		$(BUILD_IMAGE) \
		/bin/sh -c " \
			./build/test.sh $(SRC_DIRS) \
		"


.PHONY: shell
shell: build-dirs
	@echo "Launching shell in the containerized build environment"
	@docker run \
		-ti \
		--rm \
		-v "$$(pwd)/.go:/go" \
		-v "$$(pwd):/go/src/$(BASE_PACKAGE)" \
		-v "$$(pwd)/bin:/go/bin" \
		-v "$$(pwd)/.go/std/$(OS)/$(ARCH):/usr/local/go/pkg/$(OS)_$(ARCH)_static" \
		-v "$$(pwd)/.cache/go-build:/root/.cache/go-build" \
	  -w /go/src/$(BASE_PACKAGE) \
		$(BUILD_IMAGE) \
		/bin/sh $(CMD)

.PHONY: version
version: ## returns the current version
	@echo $(VERSION)

.PHONY: build-dirs
build-dirs:
	@mkdir -p bin/$(ARCH)
	@mkdir -p .go/src/$(BASE_PACKAGE) .go/pkg .go/bin .go/std/$(OS)/$(ARCH) .cache/go-build

.PHONY: build-builder
build-builder:
	@docker build -t $(BUILD_IMAGE) -f ./Dockerfile.builder .