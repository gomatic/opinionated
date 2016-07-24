export APP_NAME := $(notdir $(shell pwd))
DESC :=
PROJECT_URL := "https://github.com/gomatic/$(APP_NAME)"

SOURCE = $(patsubst %/,%,$(sort $(dir $(wildcard *.go vendor/application/*.go))))

.PHONY : $(DIRS)
.PHONY : all
.PHONY : help
.DEFAULT_GOAL := help

PREFIX ?= usr/local

export COMMIT_ID := $(shell git describe --tags --always --dirty 2>/dev/null)
export COMMIT_TIME := $(shell git show -s --format=%ct 2>/dev/null)

export STARTD := $(shell pwd)
export THIS := $(abspath $(lastword $(MAKEFILE_LIST)))
export THISD := $(dir $(THIS))


build $(APP_NAME): $(SOURCE) ## Build opinionated
	go build -ldflags "-X main.VERSION=$(COMMIT_TIME)-$(COMMIT_ID)" -o $(APP_NAME)


run: $(APP_NAME) ## Run opinionated
	./$(APP_NAME)


help: ## This help.
	@echo Targets:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_ -]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
