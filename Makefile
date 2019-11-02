VERSION=0.0.2
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PATH_BUILD=build/
FILE_COMMAND=rmbr
FILE_ARCH=darwin_amd64

clean:
	@rm -rf ./build

builds: clean
	$(GOCMD) \
	  -bc="darwin,amd64" \
	  -pv=$(VERSION) \
	  -d=$(PATH_BUILD) \
	  -build-ldflags "-X main.VERSION=$(VERSION)"


build: clean
	$(GOBUILD) -o $(PATH_BUILD)$(FILE_COMMAND) -v

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)$(FILE_COMMAND) '$(HOME)/bin/$(FILE_COMMAND)'