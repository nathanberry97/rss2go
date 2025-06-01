.DEFAULT_GOAL := explain

.PHONY: explain
explain:
	@echo rss2go
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## Install pre-commit hooks
	@pre-commit install

.PHONY: build
build: ## Build rss2go api
	@rm -rf web/static/css/* || true
	@go build -o bin/app cmd/app/main.go

.PHONY: run
run: build ## Build and run rss2go api
	@./bin/app

.PHONY: test
test: ## Test backend for rss2go app
	@go test internal/utils/*.go -v

.PHONY: clean
clean: ## Clean up build artifacts
	@rm -rf bin

.PHONY: container
container: ## Run a local containised version of the application
	@mkdir -p ~/.config/rss2go
	@touch ~/.config/rss2go/rss.db
	@podman build --tag rss2go_dev .
	@podman run --name rss2go_dev -dit \
  	 -p 8000:8080 \
	 --replace \
	 rss2go_dev
