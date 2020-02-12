GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_DIR=bin
BINARY_NAME=join_tsv
TARGET_FILE=./cmd/join_tsv/main.go

all: test build
install: build system_install

.PHONY: test
test:
	$(GOTEST) -v ./test/

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) $(TARGET_FILE)

.PHONY: run
run:
	$(GORUN) $(TARGET_FILE)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -r $(BINARY_DIR)

.PHONY: system_install
system_install:
	cd cmd/join_tsv  && go install

