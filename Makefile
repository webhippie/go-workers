NAME := workers
IMPORT := github.com/webhippie/$(NAME)

PACKAGES ?= $(shell go list ./... | grep -v /vendor/ | grep -v /_tools/)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./_tools/*")

TAGS ?=

.PHONY: all
all: build

.PHONY: update
update:
	retool do dep ensure -update

.PHONY: sync
sync:
	retool do dep ensure

.PHONY: graph
graph:
	retool do dep status -dot | dot -T png -o docs/deps.png

.PHONY: clean
clean:
	go clean -i ./...

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: megacheck
megacheck:
	retool do megacheck -tags '$(TAGS)' $(PACKAGES)

.PHONY: lint
lint:
	for PKG in $(PACKAGES); do retool do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: test
test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

.PHONY: build
build:
	go build -i -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' $(IMPORT)

HAS_RETOOL := $(shell command -v retool)

.PHONY: retool
retool:
ifndef HAS_RETOOL
	go get -u github.com/twitchtv/retool
endif
	retool sync
	retool build
