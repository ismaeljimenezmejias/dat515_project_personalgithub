proto_src	:= proto/kv.proto
proto_go	:= $(proto_src:%.proto=%.pb.go)
grpc_go		:= $(proto_src:%.proto=%_grpc.pb.go)
kvstore		:= ./cmd/kvs

TAGS ?= ""

.PHONY: kvstore test clean gitignore

kvstore: proto gitignore
	@go build -tags $(TAGS) -o ./kvs $(kvstore)

test: proto
	@go test -v -tags $(TAGS) ./kvstore ./proto

proto: $(proto_go) $(grpc_go)

%.pb.go %_grpc.pb.go: %.proto
	@protoc --go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			--go_opt=default_api_level=API_OPAQUE $^

clean:
	@rm -vf proto/*.pb.go ./kvs

gitignore:
	@grep -q "3net/grpc/kvs" ../../.gitignore || echo "3net/grpc/kvs" >> ../../.gitignore
