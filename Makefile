# cna-installer makefile
PROJECTNAME := $(shell basename "$(PWD)")

# Detect OS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OS = linux
endif
ifeq ($(UNAME_S),Darwin)
	OS = darwin
endif


# Go related variables.
GOBASE := $(shell pwd)
#GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/build/bin
GOFILES := $(wildcard *.go)
TFBIN := $(GOBIN)/terraform
TFVERSION := 0.12.9

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

## compile: Compile the binary.
build:
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR)| sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECTNAME) 2> /dev/null
	@-$(MAKE) go-clean
	@-rm $(GOBIN)/terraform 2> /dev/null

go-compile: go-get go-build get-terraform

go-build:
	@echo "  >  Building binary..."
	@go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)
	#GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-generate:
	@echo "  >  Generating dependency files..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOBIN=$(GOBIN) go get $(get)

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

get-terraform:
	test -s $(TFBIN) && echo "  >  Found terraform binary at $(TFBIN). Skipping download."|| $(MAKE) download-terraform

download-terraform:
	@echo "  >  Downloading dependent binaries..."
	@curl -sS https://releases.hashicorp.com/terraform/$(TFVERSION)/terraform_$(TFVERSION)_$(OS)_amd64.zip --output $(GOBIN)/terraform.zip 2> $(STDERR)
	@unzip -q -o $(GOBIN)/terraform.zip -d $(GOBIN) 2> $(STDERR)
	@rm -f $(GOBIN)/terraform.zip

.PHONY: build clean help

all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
