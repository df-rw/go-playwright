help: # me
	@grep '^[a-z.]*:' Makefile | sed -e 's/^\(.*\): .*# \(.*\)/\1 - \2/'

dev: # Run development server.
	air

test: # Run playwright tests
	go run ./cmd/playwright/do-tests.go
