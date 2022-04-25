.PHONY: docker_up goose-status goose_migrate goose_degrade goose_reset goose_create

DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USERNAME ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= postgres
DB_SSL ?= disable

docker_up:
	docker-compose up

goose_status:
	goose postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)" status

goose_migrate:
	goose -dir db-library/migrations postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)" up

goose_degrade:
	goose -dir db-library/migrations postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)" down

goose_reset:
	goose -dir db-library/migrations postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)" reset

goose_create:
	goose -dir db-library/migrations create $(FILENAME) sql
