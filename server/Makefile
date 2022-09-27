BINARY_NAME	= server
GO_BINARY_PATH	=$(shell go env GOPATH)/bin

all: install

$(BINARY_NAME):
	@go build main.go

install:
	@go mod tidy

start:
	@$(GO_BINARY_PATH)/air -c .air.toml
