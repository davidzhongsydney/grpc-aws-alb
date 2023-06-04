GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)

ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	MODEL_PROTO_FILES=$(shell $(Git_Bash) -c "find model -name *.proto")
else
	MODEL_PROTO_FILES=$(shell find model -name *.proto)
endif

.PHONY: init
# download and update dependencies
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/wire/cmd/wire@latest


.PHONY: model
# generate files from model folder by proto file
model:
	protoc --go_out=. \
			--go_opt=paths=source_relative  \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
			$(MODEL_PROTO_FILES)

.PHONY: buildserver
buildserver:
	mkdir -p bin/; \
	go build -o ./bin/ ./server/...

.PHONY: buildclient
buildclient:
	mkdir -p bin/; \
	go build -o ./bin/ ./client/...

.PHONY: all
# generate all
all:
	make model;
	make build;