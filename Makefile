COMPOSE := docker compose
MIGRATE := $(COMPOSE) run --rm migrate -path=/migrations \
	-database=postgres://$${POSTGRES_USER:-flowforge}:$${POSTGRES_PASSWORD:-flowforge}@postgres:5432/$${POSTGRES_DB:-flowforge}?sslmode=disable

.PHONY: up down reset logs psql migrate-up migrate-down migrate-version migrate-create sqlc

up: ## Start Postgres and apply migrations
	$(COMPOSE) up -d

down: ## Stop containers (keeps data)
	$(COMPOSE) down

reset: ## Stop containers and delete the database volume
	$(COMPOSE) down -v

logs:
	$(COMPOSE) logs -f postgres

psql:
	$(COMPOSE) exec postgres psql -U $${POSTGRES_USER:-flowforge} -d $${POSTGRES_DB:-flowforge}

migrate-up:
	$(MIGRATE) up

migrate-down: ## Roll back the latest migration
	$(MIGRATE) down 1

migrate-version:
	$(MIGRATE) version

migrate-create: ## make migrate-create name=add_something
	$(COMPOSE) run --rm migrate create -ext sql -dir /migrations -seq $(name)

sqlc: ## Generate Go code from SQL
	$(COMPOSE) run --rm sqlc generate
