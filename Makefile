help: # me
	@grep '^[a-z.]*:' Makefile | sed -e 's/^\(.*\): .*# \(.*\)/\1 - \2/'

dev: # Run development server.
	air

test: test-unit test-browser # Run tests.

test-unit: # Run unit tests.
	go test -v ./...

test-browser: # Run browser tests.
	go run ./cmd/playwright/do-tests.go

lint: # Run the linter.
	golangci-lint run ./...
