
.PHONY: protoc
protoc:
	retool do protoc --proto_path=$$GOPATH/src:. --twirp_out=. --go_out=. ./pkg/rpc/service.proto
