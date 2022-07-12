GO=go
GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)
GO_IMPORTS=goimports
GO_LDFLAGS=-ldflags="-s -w"
TARGET_DIR=bin/
CONTAINER_IMAGE=guni1192/spelunker:dev

.PHONY: build test fmt vet clean

build:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=$(GOOS) GO_ARCH=$(GOARCH) $(GO) build $(GO_LDFLAGS) -o $(TARGET_DIR) ./...

check: fmt vet lint

test:
	$(GO) test -v ./...

fmt:
	$(GO_IMPORTS) -w .

vet:
	$(GO) vet ./...

lint:
	golangci-lint run

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/gopls@latest

clean:
	rm -rf bin

docker-build:
	DOCKER_BUILDKIT=1 docker build -t $(CONTAINER_IMAGE) .

docker-run:
	docker run --rm -v $(pwd):/ $(CONTAINER_IMAGE)
