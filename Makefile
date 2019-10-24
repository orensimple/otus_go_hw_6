# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SOURCE_NAME=main.go
BINARY_NAME=main

all: deps build test
deps:
	$(GOGET) github.com/cheggaaa/pb/v3
	$(GOGET) github.com/spf13/pflag
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(SOURCE_NAME) -v ./...
	./$(BINARY_NAME) 
