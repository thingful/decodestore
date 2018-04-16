
BIN := dcs
VERSION := $(shell git describe --tags --always --dirty --abbrev=8)
BUILD_DATE := $(shell date "+%FT%H:%M:%S %Z")
BASE_PACKAGE := github.com/thingful/decodestore
PACKAGE := $(BASEPACKAGE)/cmd/$(BIN)
BUILD_DIR := build
BUILD_FLAGS := -v -ldflags "-extldflags -static -X $(BASE_PACKAGE)/pkg/version.Version=$(VERSION) -X \"$(BASE_PACKAGE)/pkg/version.BuildDate=$(BUILD_DATE)\""

BUILD_IMAGE := thingful/dcs-builder

OS ?= linux
ARCH ?= amd64

SRC_DIRS := cmd pkg

REGISTRY ?= thingful

IMAGE := $(REGISTRY)/$(BIN)-$(ARCH)

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

DOTFILE_IMAGE = $(subst :,_,$(subst /,_,$(IMAGE))-$(VERSION))

container: .container-$(DOTFILE_IMAGE) container-name
.container-$(DOTFILE_IMAGE): bin/$(OS)/$(ARCH)/$(BIN) Dockerfile.in
ifeq ($(OS),darwin)
	@echo "Unable to build darwin container"
	exit 1
endif

	@sed \
		-e 's|ARG_BIN|$(BIN)|g' \
		-e 's|ARG_ARCH|$(ARCH)|g' \
		Dockerfile.in > .dockerfile-$(ARCH)
	@docker build -t $(IMAGE):$(VERSION) -f .dockerfile-$(ARCH) .
	@docker images -q $(IMAGE):$(VERSION) > $@

container-name:
	@echo "container: $(IMAGE):$(VERSION)"

push: .push-$(DOTFILE_IMAGE) push-name
.push-$(DOTFILE_IMAGE): .container-$(DOTFILE_IMAGE)
	@docker push $(IMAGE):$(VERSION)
	@docker images -q $(IMAGE):$(VERSION) > $@

.PHONY: push-name
push-name: ## display the name of the just pushed image
	@echo "pushed: $(IMAGE):$(VERSION)"

.PHONY: version
version: ## returns the current version
	@echo $(VERSION)

.PHONY: build-dirs
build-dirs: ## creates build directories
	@mkdir -p bin/$(ARCH)
	@mkdir -p .go/src/$(BASE_PACKAGE) .go/pkg .go/bin .go/std/$(OS)/$(ARCH) .cache/go-build

.PHONY: build-builder
build-builder: ## builds our builder base container locally
	@docker build -t $(BUILD_IMAGE) -f ./Dockerfile.builder .

.PHONY: clean
clean: container-clean bin-clean ## clean everything

.PHONY: container-clean
container-clean: ## clean container artefacts
	rm -rf .container-* .dockerfile-* .push-*

.PHONY: bin-clean
bin-clean: ## clean build artefacts
	rm -rf .go bin .cache