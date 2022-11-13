export PATH := $(HOME)/.local/bin:$(PATH)

rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

# Certificate variables
CERT_DIR=./cert
CERT_SCRIPT=$(CERT_DIR)/gencert.zsh
CERT_CONF=$(CERT_DIR)/server-ext.cnf
CERT_SERVER_PEM=$(CERT_DIR)/server-cert.pem
GENERATED_CERT_FILES=$(shell find $(CERT_DIR) -type f \( -name "*.pem" -o -name "*.srl" \))

# gRPC variables
# gRPC variables
PROTO_DIR=./server/interface/rpc/proto
PROTO_FILES=$(shell find ./server/interface/rpc/proto -type f -name "*.proto")
GRPC_DIR=./grpc
PB_GO_FILES=$(shell find $(GRPC_DIR) -type f -name "*_grpc.pb.go")

# Server
EXEC_DIR = ./dist
SERVER_SRC_DIR = ./server
SERVER_SRC_MAIN = $(SERVER_SRC_DIR)/main.go
SERVER_EXEC_DIR = $(EXEC_DIR)/server
SERVER_EXEC_MAIN = $(SERVER_EXEC_DIR)/main
SERVER_FILES=$(shell find $(SERVER_SRC_DIR) -type f \( -name "*.go" -o -name "*.go" \))

# GO client Files
CLIENT_SRC_DIR = ./client
CLIENT_SRC_MAIN = $(CLIENT_SRC_DIR)/main.go
CLIENT_EXEC_DIR = $(EXEC_DIR)/client
CLIENT_EXEC_MAIN = $(CLIENT_EXEC_DIR)/main
CLIENT_FILES=$(shell find $(CLIENT_SRC_DIR) -type f \( -name "*.go" -o -name "*.go" \))

# GO UI Files
# Not yet

# Keep all PHONY tasks definitions together
.PHONY: setup \
		gen_grpc \
		build \
		clean \
		clean_cert \
		clean_grpc \
		clean_dist \
		run_server \
		run_client \
		test_server

# Install dependencies
setup:
	go get github.com/google/uuid

	# Protocol Buffer Compiler
	curl -LO ${PB_URL}/${PB_OSX}
	unzip ${PB_OSX} -d $${HOME}/.local
	rm -f ${PB_OSX}

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/google/wire/cmd/wire@latest

# Build certificates
$(CERT_SERVER_PEM) : $(CERT_SCRIPT) $(CERT_CONF)
	$(CERT_SCRIPT)

## Generate the gRPC files
$(PB_GO_FILES): $(PROTO_FILES)
	buf generate --template ./server/interface/rpc/buf.gen.yamlmrs


# Build Server
$(SERVER_EXEC_MAIN): $(SERVER_FILES) $(CERT_SERVER_PEM) $(PB_GO_FILES)
	go build -o $(SERVER_EXEC_MAIN) $(SERVER_SRC_MAIN)

# Build Client
$(CLIENT_EXEC_MAIN): $(CLIENT_FILES) $(CERT_SERVER_PEM) $(PB_GO_FILES)
	go build -o $(CLIENT_EXEC_MAIN) $(CLIENT_SRC_MAIN)

# Build UI
# not yet

# Build all files
build: $(CERT_SERVER_PEM) $(PB_GO_FILES) $(SERVER_EXEC_MAIN) $(CLIENT_EXEC_MAIN)

# Run the server
run_server:
	$(SERVER_EXEC_MAIN)

# Run the client
run_client:
	$(CLIENT_EXEC_MAIN)

# Clean the generated gRPC files
clean_grpc:
	rm $(GENERATED_GRPC_FILES)

# Clean the generated certificates
clean_cert:
	rm $(GENERATED_CERT_FILES)

# Clean the generated certificates
clean_dist:
	rm $(SERVER_EXEC_MAIN)
	rm $(CLIENT_EXEC_MAIN)

# Clean all generated files
clean: clean_cert clean_grpc

# Test targets
test_server:
	ginkgo $(SERVER_SRC_DIR)/...

# Install targets
# TBD install, uninstall

# The default goal
.DEFAULT_GOAL := build
