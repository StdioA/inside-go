GO ?= go
GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
VETPACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/ | grep -v /examples/)
GOFILES := $(shell find2 . -name "*.go" -type f -not -path "./vendor/*")
TESTFOLDER := $(shell $(GO) list ./... | grep -E 'gin$$|binding$$|render$$' | grep -v examples)

all: build

.PHONY: server
server: build
	main.exe

.PHONY: build
build:
	$(GO) build main.go

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)
