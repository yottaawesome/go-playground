# Makefile — common tasks for the go-sample-project.
# Run `make help` to see available targets.

APP_NAME   := bookserver
BUILD_DIR  := bin
CMD_DIR    := ./cmd/bookserver

.PHONY: all build run test lint clean help

## all: build and test (default target)
all: build test

## build: compile the server binary into bin/
build:
	@echo "==> Building $(APP_NAME)…"
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

## run: build and start the server
run: build
	@echo "==> Starting $(APP_NAME) on :8080…"
	$(BUILD_DIR)/$(APP_NAME)

## test: run all tests with verbose output
test:
	@echo "==> Running tests…"
	go test -v -race ./...

## lint: run go vet (install staticcheck separately for deeper analysis)
lint:
	@echo "==> Vetting code…"
	go vet ./...

## clean: remove build artefacts
clean:
	@echo "==> Cleaning…"
	rm -rf $(BUILD_DIR)

## help: print this help message
help:
	@echo "Available targets:"
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':'
