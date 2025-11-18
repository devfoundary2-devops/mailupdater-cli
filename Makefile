.PHONY: help db-up db-down db-logs build run clean test

help: ## Show this help message
	@printf "\n\033[1mAvailable Targets:\033[0m\n\n"
	@awk 'BEGIN {FS = ":.*## " } \
		/^[a-zA-Z0-9_-]+:.*## / { \
			printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 \
		}' $(MAKEFILE_LIST)
	@printf "\n"

db-up: ## Start the database
	docker compose up -d
	@echo "Waiting for database to be ready..."
	@sleep 3
	@echo "Database is ready!"

db-down: ## Stop the database
	docker compose down

db-logs: ## Show database logs
	docker compose logs -f postgres

db-clean: ## Stop database and remove volumes
	docker compose down -v

build: ## Build the CLI
	go build -o bin/mailupdater .

run: build ## Build and run the CLI
	./bin/mailupdater

clean: ## Clean build artifacts
	rm -rf bin/

test: ## Run tests
	go test -v ./...

install: build ## Install the CLI locally
	go install ./mailupdater