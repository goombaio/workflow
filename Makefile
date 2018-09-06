include Makefile.help.mk

BINARY=workflow
MAIN_PACKAGE=cmd/${BINARY}/main.go
PACKAGES = $(shell go list ./...)
VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
DIST_FOLDER=dist
DIST_INCLUDE_FILES=README.md LICENSE VERSION

# Setup -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Build & Install

install:		## Build and install package on your system
	go install $(LDFLAGS) -v $(PACKAGES)

.PHONY: version
version:		## Show version information
	@echo $(VERSION)-$(BUILD)

# Testing

.PHONY: test
test:			## Execute package tests 
	go test -v $(PACKAGES)

.PHONY: cover-profile
cover-profile:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out

.PHONY: cover
cover: cover-profile	
cover: 			## Generate test coverage data
	go tool cover -func=coverage-all.out

.PHONY: cover-html
cover-html: cover-profile
cover-html:		## Generate coverage report
	go tool cover -html=coverage-all.out

.PHONY: codecov
codecov:
	bash <(curl -s https://codecov.io/bash)

# Lint

lint:			## Lint source code
	gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow

# Dependencies

deps:			## Install build dependencies
	go get -u github.com/google/uuid
	go get -u github.com/goombaio/dag


dev-deps: deps
dev-deps:		## Install dev and build dependencies
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: clean
clean:			## Delete generated development environment
	go clean
	rm -rf ${BINARY}-*-*
	rm -rf ${BINARY}-*-*.exe
	rm -rf ${BINARY}-*-*.zip
	rm -rf coverage-all.out

# Docs

godoc-serve:		## Serve documentation (godoc format) for this package at port HTTP 9090
	godoc -http=":9090"
