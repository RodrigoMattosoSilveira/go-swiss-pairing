export PATH := $(HOME)/.local/bin:$(PATH)

rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

# Certificate variables
CERT_DIR=./cert
CERT_SCRIPT=$(CERT_DIR)/gencert.zsh
CERT_CONF=$(CERT_DIR)/server-ext.cnf
CERT_SERVER_PEM=$(CERT_DIR)/server-cert.pem
GENERATED_CERT_FILES=$(shell find $(CERT_DIR) -type f \( -name "*.pem" -o -name "*.srl" \))

# gRPC variables
PB_VER = 21.4
PB_URL =  https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VER}
PB_PREFIX =  protoc
PB_FN  = osx-x86_64.zip
PB_OSX = ${PB_PREFIX}-${PB_VER}-${PB_FN}
GRPC_DIR=./app/interface/rpc
GENERATED_GRPC_FILES=$(shell find $(GRPC_DIR) -type f \( -name "*.pb.go" -o -name "*.js" \))

# gRPC variables
PROTO_DIR = ./app/interface/rpc/proto
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)
PB_GO_FILES = $(PROTO_FILES:.proto=.pb.go)
all_pb_go_files: $(PB_GO_FILES)

# GO Server Files
GO_DIR = ./app
GO_SERVER = $(GO_DIR)/server/main
GO_CLIENT = $(GO_DIR)/client/main
GO_FILES = $(call rwildcard,,*.pb.go)

# GO client Files
GO_CLIENT = $(GO_DIR)/client/main

# GO UI Files
# Not yet

# Keep all PHONY tasks definitions together
.PHONY: setup \
		build \
		clean \
		clean_cert \
		clean_grpc \
		run_server \
		run_client

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

# Build the GO and JS files to support gRPC operations, in the same folder as their proto files
$(PB_GO_FILES): %.pb.go: %.proto
	protoc --js\_out=import\_style=commonjs,binary:./ \
		--grpc-web\_out=import\_style=commonjs,mode=grpcwebtext:./ \
		--go_out=./ \
		--go_opt=paths=source_relative \
		--go-grpc_out=require_unimplemented_servers=false:. \
		--go-grpc_opt=paths=source_relative \
		$<

# Build Server
$(GO_SERVER): $(GO_FILES) $(CERT_SERVER_PEM)
	go build -o $(GO_SERVER) $(GO_SERVER).go

# Build Client
$(GO_CLIENT): $(GO_FILES) $(CERT_SERVER_PEM)
	go build -o $(GO_CLIENT) $(GO_CLIENT).go

# Build UI
# not yet

# Build all files
build: $(CERT_SERVER_PEM) $(PB_GO_FILES) $(GO_SERVER) $(GO_CLIENT)

# Run the server
run_server:
	$(GO_SERVER)

# Run the client
run_client:
	$(GO_CLIENT)

# Clean the generated gRPC files
clean_grpc:
	rm $(GENERATED_GRPC_FILES)

# Clean the generated certificates
clean_cert:
	rm $(GENERATED_CERT_FILES)

# Clean all generated files
clean: clean_cert clean_grpc

# Test targets
# TBD server, client, ui, all

# Install targets
# TBD install, uninstall

# The default goal
.DEFAULT_GOAL := build
