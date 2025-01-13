.DEFAULT_GOAL := explain

.PHONY: explain
explain:
	@echo personalWebsite
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## Install pre-commit hooks
	@pre-commit install

.PHONY: build
build: ## Build rss2go app
	@cd app && go build -o bin/app src/*.go

.PHONY: run
run: build ## Build and run rss2go app
	@cd app && ./bin/app

.PHONY: test
test: ## Test backend for rss2go app
	@cd app && go test src/routes/*.go -v
	@cd app && go test src/utils/*.go -v

.PHONY: clean
clean: ## Clean up build artifacts
	@rm -rf api/bin
