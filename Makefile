go := $(shell which go)

setup: install
install:
	$(go) install golang.org/x/tools/cmd/goimports@latest
	$(go) install honnef.co/go/tools/cmd/staticcheck@latest
	$(go) mod tidy

run:
	$(go) run .

test: tests
tests:
	$(go) test -v -failfast ./...

build:
	$(go) build -buildvcs=false

lint:
	@goimports -w .
	@staticcheck ./...

