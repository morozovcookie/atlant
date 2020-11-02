CURRENT_DIR = $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

GOPATH       = $(shell go env GOPATH)
CGO_ENABLED  = 0
GOOS        ?= linux
GOARCH      ?= amd64
GOFLAGS      = GOOS=$(GOOS) GOARCH=$(GOARCH)

MONGODB_HOST            ?= 127.0.0.1
MONGODB_PORT            ?= 27017
MONGODB_DATABASE        ?= atlant
MONGODB_DATABASE_URL     = mongodb://$(MONGODB_HOST):$(MONGODB_PORT)/$(MONGODB_DATABASE)
MONGODB_MIGRATIONS_PATH ?= $(CURRENT_DIR)/migrations


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

# Build the Docker container.
.PHONY: docker-build
docker-build: atlantserver-docker-build processor-docker-build
	@echo "+ $@"

# Build atlantserver docker image.
.PHONY: atlantserver-docker-build
atlantserver-docker-build:
	@echo "+ $@"
	@docker build \
		--rm \
		-f ./scripts/docker/Dockerfile.atlantserver \
		-t atlantserver:latest \
		.

# Build processor docker image.
.PHONY: processor-docker-build
processor-docker-build:
	@echo "+ $@"
	@docker build \
		--rm \
		-f ./scripts/docker/Dockerfile.processor \
		-t processor:latest \
		.

# Build the binaries.
.PHONY: go-build
go-build: atlantserver-build atlantclient-build processor-build
	@echo "+ $@"

# Build the atlantserver binary.
.PHONY: atlantserver-build
atlantserver-build:
	@echo "+ $@"
	@$(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/atlantserver \
		$(CURRENT_DIR)/cmd/atlantserver/main.go

# Build the atlantclient binary.
.PHONY: atlantclient-build
atlantclient-build:
	@echo "+ $@"
	@CGO_ENABLED = 0 $(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/atlantclient \
		$(CURRENT_DIR)/cmd/atlantclient/main.go

# Build the processor binary.
.PHONY: processor-build
processor-build:
	@echo "+ $@"
	@$(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/atlantclient \
		$(CURRENT_DIR)/cmd/atlantclient/main.go

# Build the application with werf.
.PHONY: werf-build
werf-build:
	@echo "+ $@"

# Remove assets.
.PHONY: clear
clear:
	@echo "+ $@"
	@rm -rf $(CURRENT_DIR)/out

# Generate go-files from proto
.PHONY: protoc
protoc:
	@echo "+ $@"
	@protoc \
		--proto_path=api/proto/v1 \
		--gogofaster_out=plugins=grpc:grpc/v1 \
		atlant.proto

# Up migrations
.PHONY: migrate-up
migrate-up:
	@echo "+ $@"
	@migrate \
		-database $(MONGODB_DATABASE_URL) \
		-path $(MONGODB_MIGRATIONS_PATH) \
		up

# Down migrations
.PHONY: migrate-down
migrate-down:
	@echo "+ $@"
	@migrate \
		-database $(MONGODB_DATABASE_URL) \
		-path $(MONGODB_MIGRATIONS_PATH) \
		down

# Drop migrations
.PHONY: migrate-drop
migrate-drop:
	@echo "+ $@"
	@migrate \
		-database $(MONGODB_DATABASE_URL) \
		-path $(MONGODB_MIGRATIONS_PATH) \
		drop
