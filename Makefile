# Define Go command
GO ?= go
SRC_DIR := src

# Run all unit tests
test:
	cd $(SRC_DIR) && GOMAXPROCS=2 $(GO) test ./... -v

# Run tests with race condition detection
test-race:
	cd $(SRC_DIR) && GOMAXPROCS=2 $(GO) test -race ./... -v

# Run tests with code coverage
test-cover:
	cd $(SRC_DIR) && $(GO) test -cover ./... -v

# Format Go code
fmt:
	cd $(SRC_DIR) && $(GO) fmt ./...

# Run static analysis with vet
lint:
	cd $(SRC_DIR) && $(GO) vet ./...

# Run all checks (formatting, linting, tests)
check: fmt lint test

# Clean up temporary files
clean:
	rm -rf $(SRC_DIR)/*.out

.PHONY: test test-race test-cover fmt lint check clean
