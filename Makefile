PKGS := $(shell go list ./... | grep -v /vendor)
GO_FILES := $(shell find . -iname '*.go' -type f | grep -v /vendor/) # All the .go files, excluding vendor/

BINARY := dumb-http
PLATFORMS := windows linux darwin
VERSION ?= latest
os = $(word 1, $@)

bootstrap:
	@mkdir -p release
	go get golang.org/x/lint                  # Linter
	go get honnef.co/go/tools/cmd/staticcheck # Badass static analyzer/linter
	go get github.com/fzipp/gocyclo           # Cyclomatic complexity check

test:
	go test -v -race $(PKGS)        # Normal Test
	go vet ./...                    # go vet is the official Go static analyzer
	staticcheck ./...                 # "go vet on steroids" + linter
	gocyclo -over 19 $(GO_FILES)    # forbid code with huge functions
	golint -set_exit_status $(PKGS) # one last linter


$(PLATFORMS):
	GOOS=$(os) GOARCH=amd64 go build -ldflags '-X main.version=$(VERSION)' -o release/$(BINARY)
	tar -czf release/$(BINARY)-$(VERSION)-$(os)-amd64.tar.gz README.md -C release/ $(BINARY)
	rm release/$(BINARY)

.PHONY: release
release: windows linux darwin

clean:
	rm -rf release/*
