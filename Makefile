export PATH := $(HOME)/.local/bin:$(PATH)

rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

PB_VER = 21.4
PB_URL =  https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VER}
PB_PREFIX =  protoc
PB_FN  = osx-x86_64.zip
PB_OSX = ${PB_PREFIX}-${PB_VER}-${PB_FN}

# gRPC variables
PROTODIR = ./app/interface/rpc/proto
PROTO_FILES = $(wildcard $(PROTODIR)/*.proto)
PB_GO_FILES = $(PROTO_FILES:.proto=.pb.go)
all_pb_go_files: $(PB_GO_FILES)

# GO Files
GO_DIR = ./app
GO_MAIN = ./app/cmd/go-swiss-pairing/main
GO_FILES = $(call rwildcard, $(GO_DIR) , *.go)
all_go_files: $(GO_FILES)

all: $(PB_GO_FILES) $(GO_MAIN)
.PHONY: all

# Build the go files to support gRPC operations. Given a folder with *.proto files, when I run my make rule, it builds
# their *.pb.go files in the same folder.
$(PB_GO_FILES): %.pb.go: %.proto
	protoc --go_out=./ \
		--go_opt=paths=source_relative \
		--go-grpc_out=require_unimplemented_servers=false:. \
		--go-grpc_opt=paths=source_relative \
		$<

# Build Server
$(GO_MAIN): $(GO_FILES)
	go build -o $(GO_MAIN) $(GO_MAIN).go

.PHONY: setup
setup: ## Install dependencies
	go get github.com/google/uuid

	# Protocol Buffer Compiler
	curl -LO ${PB_URL}/${PB_OSX}
	unzip ${PB_OSX} -d $${HOME}/.local
	rm -f ${PB_OSX}

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: run
run:
	./app/cmd/go-swiss-pairing/main

.PHONY: show_go_files
show_go_files:
	echo $(GO_FILES)