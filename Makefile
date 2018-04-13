
EXECUTABLE := dcs
VERSION := $(shell git describe --tags --always --dirty --abbrev=8)
BUILDDATE := $(shell date "+%FT%H:%M:%S %Z")
BASEPACKAGE := github.com/thingful/decodestore
PACKAGE := $(BASEPACKAGE)/cmd/$(EXECUTABLE)
BUILD_DIR := build
BUILDFLAGS := -v -ldflags "-extldflags -static -X $(BASEPACKAGE)/pkg/version.Version=$(VERSION) -X \"$(BASEPACKAGE)/pkg/version.BuildDate=$(BUILDDATE)\""

.PHONY: gen
gen: ## Run the protobuf compiler to generate implementation stubs.
	retool do protoc --proto_path=$$GOPATH/src:. --twirp_out=. --go_out=. ./pkg/rpc/datastore/service.proto

.PHONY: build
build:
	export CGO_ENABLED=0
	go build $(BUILDFLAGS) -o $(BUILD_DIR)/$(EXECUTABLE) $(PACKAGE)

.PHONY: install
install:
	export CGO_ENABLED=0
	go install $(BUILDFLAGS) $(PACKAGE)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)