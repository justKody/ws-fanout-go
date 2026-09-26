.PHONY: all build run test cover fmt vet tidy clean help

APP := ws-fanout
BIN := bin/$(APP)
GO  ?= go

all: build

## build: compile the server to bin/ws-fanout
build:
	$(GO) build -o $(BIN) .

## run: build and start the server
run: build
	./$(BIN)

## test: run all tests
test:
	$(GO) test ./...

## cover: run tests with coverage
cover:
	$(GO) test -cover ./...

## fmt: format Go sources
fmt:
	$(GO) fmt ./...

## vet: static analysis
vet:
	$(GO) vet ./...

## tidy: sync go.mod and go.sum
tidy:
	$(GO) mod tidy

## clean: remove build output
clean:
	rm -rf bin

## help: list targets
help:
	@printf '%s\n' \
		'all     build the binary (default)' \
		'build   compile to bin/ws-fanout' \
		'run     build and start the server' \
		'test    run all tests' \
		'cover   run tests with coverage' \
		'fmt     gofmt the module' \
		'vet     go vet the module' \
		'tidy    go mod tidy' \
		'clean   remove build output'
