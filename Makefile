BUILDPATH    =$(CURDIR)
MAIN_NAME    =${BUILDPATH}/cmd/ova-person-api/main.go
GOBIN        =$(CURDIR)/bin
PROJECT_NAME = ova-person-api

.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps:
	@echo "Installing dependencies..."
	@go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u google.golang.org/grpc
	@go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go mod download

.PHONY: generate
generate:
	@echo "Installing dependencies..."
	@protoc -I ${CURDIR}/api \
		--go_out=${CURDIR}/pkg --go_opt=paths=source_relative \
		--go-grpc_out=${CURDIR}/pkg --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=${CURDIR}/pkg --grpc-gateway_opt paths=source_relative \
		${CURDIR}/api/*.proto

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go mod tidy
	@go mod download && go build -o ${GOBIN}/${PROJECT_NAME} ${MAIN_NAME}

.PHONY: run
## run: runs go run main.go
run:
	@echo "Running..."
	@go run ${MAIN_NAME}

.PHONY: format
## format: format the code of project
format:
	@echo "Formatting..."
	@go fmt ./...

.PHONY: lint
## lint: make the code of project is cleaner
lint:
	@echo "Running of linter..."
	@golangci-lint run --fix

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning..."
	@go clean -modcache
	@go clean -i

.PHONY: test
## test: tests
test:
	@echo "Test"
	@go test ./internal/models
	@go test ./internal/utils
	@go test ./internal/flusher
