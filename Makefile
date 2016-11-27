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

cert: data/server.csr

data/server.csr: data/server.key
	cd $(dir $@); openssl req -new -sha256 -key server.key -out server.csr -config example.conf
	cd $(dir $@); openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650

ecdsa:
	cd data; openssl req -x509 -nodes -newkey ec:secp384r1 -keyout server.ecdsa.key -out server.ecdsa.crt -days 3650
	cd data; ln -sf server.ecdsa.key server.key; ln -sf server.ecdsa.crt server.crt

data/server.key:
	cd $(dir $@); openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650 -config example.conf
	cd $(dir $@); ln -sf server.rsa.key server.key; ln -sf server.rsa.crt server.crt


help: ## This help.
	@echo Targets:
	@awk 'BEGIN {FS = ":.*?## "} / [#][#] / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
