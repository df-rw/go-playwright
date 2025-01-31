help: # me
	@grep '^[-a-z.]*:' Makefile | sed -e 's/^\(.*\): .*# \(.*\)/\1 - \2/'

dev: # Run development server.
	air

test: test-unit test-browser-standalone test-browser-with-testing # Run tests.

test-unit: # Run unit tests.
	go test -v ./internal/...

test-browser-standalone: # Run standalone browser tests.
	go run ./cmd/playwright/do-tests.go

test-browser-with-testing: # Run browser tests with go testing library.
	go test -v -count=1 ./tests/...

lint: # Run the linter.
	golangci-lint run ./...
