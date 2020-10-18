CURRENT_DIR = $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

GOPATH       = $(shell go env GOPATH)
CGO_ENABLED  = 0
GOOS        ?= linux
GOARCH      ?= amd64
GOFLAGS      = CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH)


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

# Build the binaries.
.PHONY: go-build
go-build: atlantserver-build atlantclient-build csvgen-build fileserver-build
	@echo "+ $@"

# Build the binaries.
.PHONY: atlantserver-build
atlantserver-build:
	@echo "+ $@"
	@$(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/atlantserver \
		$(CURRENT_DIR)/cmd/atlantserver/main.go

# Build the binaries.
.PHONY: atlantclient-build
atlantclient-build:
	@echo "+ $@"
	@$(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/atlantclient \
		$(CURRENT_DIR)/cmd/atlantclient/main.go

# Build the csvgen binary.
.PHONY: csvgen-build
csvgen-build:
	@echo "+ $@"
	@$(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/csvgen \
		$(CURRENT_DIR)/cmd/csvgen/main.go

# Build the fileserver binary.
.PHONY: fileserver-build
fileserver-build:
	@echo "+ $@"
	@$(GOFLAGS) go build \
		-ldflags "-s -w" \
		-o $(CURRENT_DIR)/out/fileserver \
		$(CURRENT_DIR)/cmd/fileserver/main.go

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
		--proto_path=third_party \
		--gogofaster_out=plugins=grpc:grpc/v1 \
		atlant.proto
