BINARY_NAME	= server
GO_BINARY_PATH	=$(shell go env GOPATH)/bin

all: install

build: $(BINARY_NAME)

$(BINARY_NAME):
	@go build -o $(BINARY_NAME)

install:
	@go mod tidy

start:
	@$(GO_BINARY_PATH)/air -c .air.toml

.PHONY: all \
		build \
		install \
		start \