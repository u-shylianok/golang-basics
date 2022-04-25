.PHONY: start goose_create goose_create_fixture goose_status goose_migrate goose_degrade goose_reset goose_fixtures_up goose_fixtures_down goose_fixtures_reset

DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USERNAME ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= postgres
DB_SSL ?= disable
GOOSE_DBSTRING ?= "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)"

DIR_MIGRATIONS ?= db-library/migrations
DIR_FIXTURES ?= db-library/fixtures

start:
	docker-compose up


goose_create:
	goose -dir $(DIR_MIGRATIONS) create $(FILENAME) sql

goose_create_fixture:
	goose -dir $(DIR_FIXTURES) create $(FILENAME) sql


goose_status:
	goose postgres $(GOOSE_DBSTRING) status

goose_migrate:
	goose -dir $(DIR_MIGRATIONS) postgres $(GOOSE_DBSTRING) up

goose_degrade:
	goose -dir $(DIR_MIGRATIONS) postgres $(GOOSE_DBSTRING) down

goose_reset:
	goose -dir $(DIR_MIGRATIONS) postgres $(GOOSE_DBSTRING) reset


goose_fixtures_up:
	goose -dir $(DIR_FIXTURES) postgres $(GOOSE_DBSTRING) up

goose_fixtures_down:
	goose -dir $(DIR_FIXTURES) postgres $(GOOSE_DBSTRING) down

goose_fixtures_reset:
	goose -dir $(DIR_FIXTURES) postgres $(GOOSE_DBSTRING) reset
