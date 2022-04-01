HASH := $(shell git show -s --format=%h)

.PHONY: test
test: ## runs all tests
	go test -v -vet=all -failfast ./...

.PHONY: terminal
terminal: # Run the terminal adapter
	go run ./term-snake

.PHONY: web
web: # Run the web adapter
	go run ./web

.PHONY: validator
validator: # Run the validator adapter
	go run ./validator

.PHONY: push
push: ## Push the image to my personal docker registry with the git hash as the tag
	docker build -t tomarrell/personal:snake-$(HASH) -f Dockerfile.validator .
	docker push tomarrell/personal:snake-$(HASH)
	echo "snake-$(HASH)" | tr -d '\n' | tee /dev/tty | pbcopy

.DEFAULT_GOAL := help

help: ## Prints this help.
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
