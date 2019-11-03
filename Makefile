VERSION=0.0.2
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PATH_BUILD=build/
BINARY_NAME=rmbr
BINARY_NAME_LINUX=rmbr_unix

clean:
	@rm -rf ./build

build: clean
	$(GOBUILD) -o $(PATH_BUILD)$(BINARY_NAME) -v

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)$(BINARY_NAME) '$(HOME)/bin/$(BINARY_NAME)'

build-linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(PATH_BUILD)$(BINARY_NAME_LINUX) -v

build-all: clean build build-linux
