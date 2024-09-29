NAME=search-keyword-service

SRC_PATH:= ${PWD}

.PHONY: build
## build: Compile the packages.
build:
	@CGO_ENABLED=0 go build -o ./bin/$(NAME)

.PHONY: run
## run: Run the application
run:
	go run main.go

.PHONY: run-dev

.PHONY: swag
## swag: Generate swagger docs
swag:
	@swag init

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f ./bin

.PHONY: deps
## deps: Download modules
deps:
	@go mod download

.PHONY: test
## test: Run tests with verbose mode
test:
	@go test -cover -coverprofile=coverage.out ./...

.PHONY:lint
lint:
	golangci-lint run

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo