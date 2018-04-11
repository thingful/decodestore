
.PHONY: gen
gen: ## Run the protobuf compiler to generate implementation stubs.
	retool do protoc --proto_path=$$GOPATH/src:. --twirp_out=. --go_out=. ./pkg/rpc/datastore/service.proto
