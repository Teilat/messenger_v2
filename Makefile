_remote_origin := $(shell git config --get remote.origin.url)
BINARY_NAME :=  $(shell echo $(_remote_origin) | awk -F/ '{print $$5}' | awk -F. '{print $$1}')
WIN_BINARY_NAME := $(BINARY_NAME).exe
VERSION_PATH := $(shell echo github.com/$(shell echo $(_remote_origin) | awk -Fgithub.com/ '{print $$2}')/configuration | sed "s/\/\//\//g")
BINARY_VERSION := $(shell git describe --tags)
BINARY_BUILD_DATE := $(shell date +%d.%m.%Y)
BUILD_FOLDER := .build

DOCKERHUB_PROFILE := "teilat"

PRINTF_FORMAT := "\033[35m%-18s\033[33m %s\033[0m\n"

ABS_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
ifeq ($(shell go env GOHOSTOS), windows)
	ABS_PATH = $(PWD)
endif

.PHONY: all build windows linux vendor gen-webapi clean docker-build docker-publish docker-up docker-tag gen-swagger

all: build

build: windows linux ## Default: build for windows and linux

gen-swagger:
	cd ./webapi && swag init --parseDependency --parseInternal -g webapi.go

windows: vendor ## Build artifacts for windows
	@printf $(PRINTF_FORMAT) BINARY_NAME: $(WIN_BINARY_NAME)
	@printf $(PRINTF_FORMAT) BINARY_BUILD_DATE: $(BINARY_BUILD_DATE)
	mkdir -p $(BUILD_FOLDER)/windows
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc  go build -ldflags "-s -w -X $(VERSION_PATH).Version=$(BINARY_VERSION) -X $(VERSION_PATH).BuildDate=$(BINARY_BUILD_DATE)" -o $(BUILD_FOLDER)/windows/$(WIN_BINARY_NAME) .

linux: vendor ## Build artifacts for linux
	@printf $(PRINTF_FORMAT) BINARY_NAME: $(BINARY_NAME)
	@printf $(PRINTF_FORMAT) BINARY_BUILD_DATE: $(BINARY_BUILD_DATE)
	mkdir -p $(BUILD_FOLDER)/linux
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X $(VERSION_PATH).Version=$(BINARY_VERSION) -X $(VERSION_PATH).BuildDate=$(BINARY_BUILD_DATE)" -o $(BUILD_FOLDER)/linux/$(BINARY_NAME) .

docker-build: linux ## Build artifacts for linux
	docker build -t $(BINARY_NAME) .

docker-tag: docker-build ## Generate container `{version}` tag
	@echo 'create tag latest'
	docker tag $(BINARY_NAME) $(DOCKERHUB_PROFILE)/$(BINARY_NAME):latest

docker-publish: docker-tag ## Build the container without caching
	@echo 'publish latest to $(DOCKERHUB_PROFILE)'
	docker push $(DOCKERHUB_PROFILE)/$(BINARY_NAME):latest

docker-up: ## Build the container without caching
	docker compose up -d

vendor: ## Get dependencies according to go.sum
	env GO111MODULE=auto go mod tidy
	env GO111MODULE=auto go mod vendor

clean: ## Remove vendor and artifacts
	rm -rf vendor
	rm -rf $(BUILD_FOLDER)