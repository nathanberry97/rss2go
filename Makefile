.DEFAULT_GOAL := explain

.PHONY: explain
explain:
	@echo personalWebsite
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## Install pre-commit hooks
	@pre-commit install

.PHONY: compile
compile: ## Compile scss to css
	@sass --no-source-map scss/style.scss static/css/style.css

.PHONY: build
build: compile ## Build rss2go api
	@go build -o bin/api cmd/api/main.go

.PHONY: run
run: build ## Build and run rss2go api
	@./bin/api

.PHONY: seed
seed: ## Seed rss2go database
	@./scripts/seed.sh

.PHONY: test
test: ## Test backend for rss2go app
	@go test pkg/routes/*.go -v
	@go test internal/utils/*.go -v

.PHONY: clean
clean: ## Clean up build artifacts
	@rm -rf bin
