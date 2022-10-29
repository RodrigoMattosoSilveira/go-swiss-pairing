export PATH := $(HOME)/.local/bin:$(PATH)

rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

# Certificate variables
CERT_DIR=./cert
CERT_SCRIPT=$(CERT_DIR)/gencert.zsh
CERT_CONF=$(CERT_DIR)/server-ext.cnf
CERT_SERVER_PEM=$(CERT_DIR)/server-cert.pem

# gRPC variables
PB_VER = 21.4
PB_URL =  https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VER}
PB_PREFIX =  protoc
PB_FN  = osx-x86_64.zip
PB_OSX = ${PB_PREFIX}-${PB_VER}-${PB_FN}

# gRPC variables
PROTO_DIR = ./app/interface/rpc/proto
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)
PB_GO_FILES = $(PROTO_FILES:.proto=.pb.go)
all_pb_go_files: $(PB_GO_FILES)

# GO Server Files
GO_DIR = ./app
GO_SERVER = $(GO_DIR)/server/main
GO_CLIENT = $(GO_DIR)/client/main
GO_FILES = $(call rwildcard, $(GO_DIR) , *.go)
all_go_files: $(GO_FILES)

# Install certificates
$(CERT_SERVER_PEM) : $(CERT_SCRIPT) $(CERT_CONF)
	$(CERT_SCRIPT)

# Build the go files to support gRPC operations. Given a folder with *.proto files, when I run my make rule, it builds
# their *.pb.go files in the same folder.
$(PB_GO_FILES): %.pb.go: %.proto
	protoc --go_out=./ \
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

.PHONY: setup
setup: ## Install dependencies
	go get github.com/google/uuid

	# Protocol Buffer Compiler
	curl -LO ${PB_URL}/${PB_OSX}
	unzip ${PB_OSX} -d $${HOME}/.local
	rm -f ${PB_OSX}

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: run_server
run_server:
	$(GO_SERVER)

.PHONY: run_client
run_client:
	$(GO_CLIENT)

.PHONY: show_go_files
show_go_files:
	echo $(GO_FILES)

.PHONY: build
build: $(CERT_SERVER_PEM) $(PB_GO_FILES) $(GO_SERVER) $(GO_CLIENT) run_server


# The default goal
.DEFAULT_GOAL := build
