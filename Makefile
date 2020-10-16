CURRENT_DIR = $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

GOPATH       = $(shell go env GOPATH)
CGO_ENABLED  = 0
GOOS        ?= linux
GOARCH      ?= amd64
GOFLAGS      = CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH)

GOBUILD_TARGET = atlantserver


# Download dependencies.
.PHONY: gomod
gomod:
	@echo "+@"
	@go mod download

# Check lint, code styling rules. e.g. pylint, phpcs, eslint, style (java) etc ...
.PHONY: style
style:
	@echo "+ $@"
	@golangci-lint run -v

# Format code. e.g Prettier (js), format (golang)
.PHONY: format
format:
	@echo "+ $@"
	@go fmt "$(CURRENT_DIR)/..."

# Shortcut to launch all the test tasks (unit, functional and integration).
.PHONY: test
test: test-unit
	@echo "+ $@"

# Launch unit tests. e.g. pytest, jest (js), phpunit, JUnit (java) etc ...
.PHONY: test-unit
test-unit:
	@echo "+ $@"
	@go test \
		-race \
		-v \
		-cover \
		-coverprofile \
		coverage.out
	@echo "+ $@"

# Locally run the application, e.g. node index.js, python -m myapp, go run myapp etc ...
.PHONY: run
run:
	@echo "+ $@"

# Build the Docker container.
.PHONY: docker-build
docker-build:
	@echo "+ $@"

# Build the binary.
.PHONY: go-build
go-build:
	@echo "+ $@"
	@$(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/$(GOBUILD_TARGET) \
		$(CURRENT_DIR)/cmd/atlantserver/main.go

# Build the application with werf.
.PHONY: werf-build
werf-build:
	@echo "+ $@"

# Remove assets.
.PHONY: clear
clear:
	@echo "+ $@"
	@rm -rf $(CURRENT_DIR)/out