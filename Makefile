BUILDPATH=$(CURDIR)
MAINNAME=${BUILDPATH}/cmd/ova-person-api/main.go

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go mod tidy
	@go mod download && go build -o ${BUILDPATH}/bin ${MAINNAME}

.PHONY: run
## run: runs go run main.go
run:
	@echo "Running..."
	@go run ${MAINNAME}

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
	@go clean -i

.PHONY: test
## test: tests
test:
	@echo "Test"
	@go test -v ./...
