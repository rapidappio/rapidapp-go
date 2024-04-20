PROTO_SRC_DIR := pkg/proto
PROTO_FILES := $(wildcard $(PROTO_SRC_DIR)/**/**/*.proto)

proto-gen:
	@echo "Generating proto files"
	protoc -I=$(PROTO_SRC_DIR) --go_out=$(PROTO_SRC_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_SRC_DIR) --go-grpc_opt=paths=source_relative $(shell find $(PROTO_SRC_DIR) -name '*.proto')
	@echo "Finished generating proto files"
