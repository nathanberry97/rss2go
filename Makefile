.DEFAULT_GOAL := explain

.PHONY: explain
explain:
	@echo personalWebsite
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: setup
setup: ## Install pre-commit hooks
	@pre-commit install

.PHONY: start-db
start-db: ## Run PostgreSQL database
	@cd database && podman build -t postgresql .
	@podman run --name rss2go-db -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgresql

.PHONY: connect-db
connect-db: ## Connect to PostgreSQL database
	@podman exec -it rss2go-db psql -U postgres -d postgres

.PHONY: build
build: ## Build backend for RSS2GO API
	@cd api && go build -o bin/api src/*.go

.PHONY: run
run: build ## Build and run backend API
	@cd api && ./bin/api

.PHONY: test
test: ## Test backend for RSS2GO API
	@cd api && go test src/routes/*.go -v
	@cd api && go test src/utils/*.go -v

.PHONY: clean
clean: ## Clean up build artifacts
	@rm -rf api/bin
	@podman rm rss2go-db
