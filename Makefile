SHELL := /bin/bash


test: ## runs all tests
	GO111MODULE=on go test -v -vet=all -failfast ./...

terminal: # Run the terminal adapter
	go run ./term-snake

web: # Run the web adapter
	go run ./web

validator: # Run the validator adapter
	go run ./validator

.DEFAULT_GOAL := help

help: ## Prints this help.
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: terminal web validator test help

