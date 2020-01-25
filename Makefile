default: | help

build: ## Build all files
	go build -v

buildstatic: ## Build all of binaries as static building for Linux AMD64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build "-ldflags=-s -w -buildid=" -trimpath -o ecs-gen main.go version.go commands.go

gogen: ## go generate
	go generate ./...

help:  ## Show all of tasks
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: \
	build \
	buildstatic \
	gogen \
	help \