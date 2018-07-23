# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=servo
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	$(GOBUILD) -o dist/$(BINARY_NAME) -v
test:
	$(GOTEST) ./...
cover:
	$(GOTEST) ./... -cover
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o dist/$(BINARY_NAME) -v
	./dist/$(BINARY_NAME)
deps:
	echo "No dependencies..."


# Cross Compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_UNIX) -v

# Docker
docker-build:
	docker build -t pgb/servo .
docker-run:
	docker run --rm -v "$(GOPATH)":/go -w /go/src/github.com/jumballaya/servo pgb/servo go build -o "/dist/$(BINARY_UNIX)" -v
